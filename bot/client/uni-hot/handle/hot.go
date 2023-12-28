package handle

import (
	"bot/client/conf"
	"bot/client/data"
	"bot/client/tweet"
	"bot/client/uni-hot/db"
	bot "bot/proto"
	"context"
	"encoding/json"
	"fmt"
	"go-micro.dev/v4/logger"
	"gorm.io/gorm"
	"io/ioutil"
	"net/http"
	security "security/proto"
	"strconv"
	"strings"
	"time"
	uni "uni/proto"
)

const LimitTime = 300

type PushType int64

const (
	Tweet  PushType = 1
	Vol    PushType = 2
	Tx     PushType = 3
	Trader PushType = 4
)

var HotPool map[string]PushLog

type Srv struct {
	Bs bot.BotService           `json:"bs"`
	Ss security.SecurityService `json:"ss"`
	Us uni.UniService           `json:"us"`
}

type PushLog struct {
	Time   int64   `json:"time"`
	Times  int64   `json:"times"`
	Vol    float64 `json:"vol"`
	Tx     int     `json:"tx"`
	Trader int     `json:"trader"`
	Tweet  int     `json:"tweet"`
}

func InitHotPoolMap() {
	HotPool = make(map[string]PushLog)
}

func GetHotPool(level string) []data.HotPoolData {
	url := getHotPoolURI(level)
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return nil
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	var hotPool data.HotPool
	json.Unmarshal(body, &hotPool)

	return hotPool.Data
}
func (sr *Srv) HotPoolHandle(pool []data.HotPoolData, database *gorm.DB, level string) {

	var volCard data.Card
	volCard.Header.Title.Tag = "plain_text"
	volCard.Header.Title.Content = "普通消息"
	volCard.Header.Template = "blue"
	var txCard data.Card
	txCard.Header.Title.Tag = "plain_text"
	txCard.Header.Title.Content = "普通消息"
	txCard.Header.Template = "blue"
	var traderCard data.Card
	txCard.Header.Title.Tag = "plain_text"
	txCard.Header.Title.Content = "普通消息"
	txCard.Header.Template = "blue"

	var tmpVolElements []data.CardElement
	var tmpTxElements []data.CardElement
	var tmpTraderElements []data.CardElement

	t := tweet.Tweet{
		Url: conf.Config.Tweet.KeywordURI,
	}

	pushBot := getBotInfo(level)
	params := getLimitParams(level)
	passToken := 0
	unPassToken := 0
	for i := 0; i < len(pool); i++ {

		token := strings.ToLower(pool[i].CoinAddr)

		now := time.Now()
		nowTime := now.Unix()
		twNum := t.GetTweetNum(token) + t.GetTweetNum(pool[i].Pair)

		tweetPushLog := db.GetHotLatest(database, token, int64(Tweet))
		if twNum >= 10 && twNum-tweetPushLog.Tweet > params.TweetIncrease {
			msg := "代币地址 : " + token + " (" + strconv.Itoa(tweetPushLog.Tweet) + "次)" +
				"\\n推特 ：" + strconv.Itoa(twNum) + " 增长:" + strconv.Itoa(twNum-tweetPushLog.Tweet) + "条" +
				"\\n" + level + "内交易量 : " + fmt.Sprintf("%f", pool[i].VolSell+pool[i].VolBuy) +
				"\\n" + level + "内交易笔数 : " + strconv.Itoa(pool[i].TxsBuy+pool[i].TxsSell) +
				"\\n流动性 : " + fmt.Sprintf("%f", pool[i].Liquidity) + " ETH" +
				"\\n名称 ：" + pool[i].CoinSymbol +
				"\\n年龄 : " + pool[i].Created

			for j := 0; j < len(pushBot.Tweet); j++ {
				sr.PushMsg(msg, pushBot.Tweet[j], token)
			}

			db.SaveHot(database, &pool[i], twNum, now, tweetPushLog.Times+1, int64(Tweet), level)
		}
		if pool[i].Liquidity < params.HotPoolMinEth {
			continue
		}

		s, err := sr.checkToken(token)
		if err != nil {
			logger.Log(logger.ErrorLevel, "check security err: ", err, " token: ", token)
			continue
		}

		if s != nil {
			if s.Score <= 85 {
				unPassToken++
				continue
			} else {
				passToken++
			}
		}

		var weight int
		volPushLog := db.GetHotLatest(database, token, int64(Vol))
		volElement, singleLevel, volSuccess := volHandle(volPushLog, pool[i], params, now, twNum, level, s)
		if volSuccess {
			weight += 1
			tmpVolElements = append(tmpVolElements, volElement)
			if singleLevel {
				volCard.Header.Title.Content = "重点消息"
			}
			db.SaveHot(database, &pool[i], twNum, now, volPushLog.Times+1, int64(Vol), level)
		}

		txPushLog := db.GetHotLatest(database, token, int64(Tx))
		txElement, singleLevel, txSuccess := txHandle(txPushLog, pool[i], params, now, twNum, level, s)
		if txSuccess {
			weight += 1
			tmpTxElements = append(tmpTxElements, txElement)
			if singleLevel {
				txCard.Header.Title.Content = "重点消息"
			}
			db.SaveHot(database, &pool[i], twNum, now, txPushLog.Times+1, int64(Tx), level)
		}

		if volSuccess || txSuccess {
			if volPushLog.Times == 0 && txPushLog.Times == 0 {
				if pool[i].Txs > 10 && pool[i].Traders > 10 {
					logger.Log(logger.InfoLevel, "开始自动购买: ", pool[i].CoinAddr)
					sr.AutoTrade(&pool[i], s, twNum)
					for j := 0; j < len(pushBot.Safe); j++ {
						bot := pushBot.Safe[j]
						msg := "代币名称: " + pool[i].CoinSymbol +
							"\\n代币地址: " + token +
							"\\n代币评分: " + strconv.Itoa(int(s.Score)) +
							"\\n流动性: " + fmt.Sprintf("%f", pool[i].Liquidity)
						sr.PushMsg(msg, bot, token)
					}
				}
			}
		}

		traderPushLog := db.GetHotLatest(database, token, int64(Trader))
		traderElement, singleLevel, success := traderHandle(traderPushLog, pool[i], params, now, twNum, level, s)
		if success {
			tmpTraderElements = append(tmpTraderElements, traderElement)
			if singleLevel {
				traderCard.Header.Title.Content = "重点消息"
			}
			db.SaveHot(database, &pool[i], twNum, now, traderPushLog.Times+1, int64(Trader), level)
		}

		if weight == 2 && nowTime-HotPool[token].Time > LimitTime {
			msg := level + "内交易量 : " + fmt.Sprintf("%f", pool[i].Vol) +
				"\\n" + level + "内交易笔数 : " + strconv.Itoa(pool[i].Txs) +
				"\\n流动性 : " + fmt.Sprintf("%f", pool[i].Liquidity) + " ETH" +
				"\\n名称 ：" + pool[i].CoinSymbol +
				"\\n推特 ：" + strconv.Itoa(twNum) +
				"\\n代币地址 : " + token +
				"\\n年龄 : " + pool[i].Created

			for j := 0; j < len(pushBot.Mix); j++ {
				sr.PushMsg(msg, pushBot.Mix[j], token)
			}

			HotPool[token] = PushLog{
				Time:  nowTime,
				Times: HotPool[token].Times + 1,
				Vol:   pool[i].Vol,
				Tx:    pool[i].Txs,
				Tweet: twNum,
			}
		}
	}

	logger.Log(logger.InfoLevel, "检测到合格代币: ", passToken, "\n检测到不合格代币: ", unPassToken)
	//vol push
	if len(tmpVolElements) > 0 {
		volCard.Elements = tmpVolElements
		for i := 0; i < len(pushBot.Vol); i++ {
			sr.PushCard(volCard, pushBot.Vol[i], len(tmpVolElements))
		}
	}
	//tx push
	if len(tmpTxElements) > 0 {
		txCard.Elements = tmpTxElements
		for i := 0; i < len(pushBot.Tx); i++ {
			sr.PushCard(txCard, pushBot.Tx[i], len(tmpTxElements))
		}
	}
	//trader push
	if len(tmpTraderElements) > 0 {
		traderCard.Elements = tmpTxElements
		for i := 0; i < len(pushBot.Trader); i++ {
			sr.PushCard(traderCard, pushBot.Trader[i], len(tmpTraderElements))
		}
	}

}

func volHandle(pushLog *db.PushLog, pool data.HotPoolData, params conf.HotParam, now time.Time, twNum int, level string, s *security.CheckTokenResponse) (data.CardElement, bool, bool) {
	success := false
	var element data.CardElement
	var singleLevel bool

	if pool.Vol >= params.MinVol {
		if now.Unix()-pushLog.PushTime.Unix() > LimitTime {
			success = true
			//发送消息
			times := pushLog.Times + 1
			vol := pushLog.Vol
			var trend string
			if pool.Vol > vol {
				trend = "⬆"
			} else if pool.Vol < vol {
				trend = "⬇"
			}

			element, singleLevel = GetHotCardElement(1, pool, times, twNum, trend, level, s)
		}
	}

	return element, singleLevel, success
}
func txHandle(pushLog *db.PushLog, pool data.HotPoolData, params conf.HotParam, now time.Time, twNum int, level string, s *security.CheckTokenResponse) (data.CardElement, bool, bool) {
	success := false
	var element data.CardElement
	var singleLevel bool

	if (pool.Txs) >= params.MinTx {
		if now.Unix()-pushLog.PushTime.Unix() > LimitTime {
			success = true
			//发送消息
			times := pushLog.Times + 1
			tx := pushLog.Tx
			var trend string
			if pool.Txs > tx {
				trend = "⬆"
			} else {
				trend = "⬇"
			}

			element, singleLevel = GetHotCardElement(2, pool, times, twNum, trend, level, s)
		}
	}

	return element, singleLevel, success
}
func traderHandle(pushLog *db.PushLog, pool data.HotPoolData, params conf.HotParam, now time.Time, twNum int, level string, s *security.CheckTokenResponse) (data.CardElement, bool, bool) {
	success := false
	var element data.CardElement
	var singleLevel bool

	if (pool.Traders) >= params.MinTrader {
		if now.Unix()-pushLog.PushTime.Unix() > LimitTime {
			//发送消息
			times := pushLog.Times + 1
			trader := pushLog.Trader
			var trend string
			if pool.Traders > trader {
				trend = "⬆"
			} else {
				trend = "⬇"
			}

			element, singleLevel = GetHotCardElement(3, pool, times, twNum, trend, level, s)
		}
	}

	return element, singleLevel, success
}

func GetHotCardElement(t int, pool data.HotPoolData, times int64, twNum int, trend string, lv string, s *security.CheckTokenResponse) (data.CardElement, bool) {
	var c data.CardElement
	var tmpColumns []data.Column
	c.Tag = "column_set"
	c.FlexMode = "none"
	c.BackgroundStyle = "grey"
	firstContent := pool.CoinSymbol + "/ETH"
	if s != nil {
		owner := "是"
		pause := "否"
		lpLock := "否"
		coolDown := "否"
		mint := "否"
		if s.Owner {
			owner = "否"
		}
		if s.Pause {
			pause = "是"
		}
		if s.Lock {
			lpLock = "是"
		}
		if s.CoolDown {
			coolDown = "是"
		}
		if s.Mint {
			mint = "是"
		}
		firstContent += "\n\n买: " + fmt.Sprintf("%f", s.BuyTax) +
			"\n卖: " + fmt.Sprintf("%f", s.SellTax) + "" +
			"\n丢弃权限: " + owner +
			"\n可暂停交易: " + pause +
			"\nlp锁: " + lpLock +
			"\n冷却期: " + coolDown +
			"\n增发: " + mint
	}

	tmpColumns = append(tmpColumns, data.Column{
		Tag:           "column",
		Width:         "weighted",
		Weight:        1,
		VerticalAlign: "top",
		Elements: append([]data.ColumnTextElement{}, data.ColumnTextElement{
			Tag: "div",
			Text: data.Text{
				Tag:     "plain_text",
				Content: firstContent,
			},
		},
		),
	})

	level := "🌟"
	if times == 1 {
		level = "🌟🌟🌟🌟🌟"
	} else if times == 3 {
		level = "🌟🌟🌟"
	} else if times > 5 {
		level = "🌟🌟🌟🌟🌟"
	}
	tmpColumns = append(tmpColumns, data.Column{
		Tag:           "column",
		Width:         "weighted",
		Weight:        1,
		VerticalAlign: "top",
		Elements: append([]data.ColumnTextElement{}, data.ColumnTextElement{
			Tag: "div",
			Text: data.Text{
				Tag:     "plain_text",
				Content: level,
			},
		})},
	)

	base := " " + trend +
		"\n流动性 : " + fmt.Sprintf("%f", pool.Liquidity) + " ETH" +
		"\n名称 ：" + pool.CoinSymbol + " " + strconv.Itoa(int(times)) + "次" +
		"\n推特 ：" + strconv.Itoa(twNum) +
		"\n代币地址 : " + pool.CoinAddr +
		"\n年龄 : " + pool.Created
	if t == 1 {
		tmpColumns = append(tmpColumns, data.Column{
			Tag:           "column",
			Width:         "weighted",
			Weight:        3,
			VerticalAlign: "top",
			Elements: append([]data.ColumnTextElement{}, data.ColumnTextElement{
				Tag: "div",
				Text: data.Text{
					Tag:     "plain_text",
					Content: lv + "内交易量 : " + fmt.Sprintf("%f", pool.VolSell+pool.VolBuy) + base,
				},
			})},
		)
	}

	if t == 2 {
		tmpColumns = append(tmpColumns, data.Column{
			Tag:           "column",
			Width:         "weighted",
			Weight:        3,
			VerticalAlign: "top",
			Elements: append([]data.ColumnTextElement{}, data.ColumnTextElement{
				Tag: "div",
				Text: data.Text{
					Tag:     "plain_text",
					Content: lv + "内交易笔数 : " + strconv.Itoa(pool.TxsBuy+pool.TxsSell) + base,
				},
			})},
		)
	}

	if t == 3 {
		tmpColumns = append(tmpColumns, data.Column{
			Tag:           "column",
			Width:         "weighted",
			Weight:        3,
			VerticalAlign: "top",
			Elements: append([]data.ColumnTextElement{}, data.ColumnTextElement{
				Tag: "div",
				Text: data.Text{
					Tag:     "plain_text",
					Content: lv + "内交易人数 : " + strconv.Itoa(pool.Traders) + base,
				},
			})},
		)
	}

	c.Columns = tmpColumns

	return c, level == "🌟🌟🌟🌟🌟"
}

func getHotPoolURI(level string) string {
	var uri string
	if level == "5m" {
		uri = conf.Config.FiveM.Hot.URI
	} else if level == "15m" {

	} else if level == "30m" {

	} else if level == "1h" {

	}

	return uri
}
func getLimitParams(level string) conf.HotParam {
	var param conf.HotParam
	if level == "5m" {
		param = conf.Config.FiveM.Hot
	} else if level == "15m" {

	} else if level == "30m" {

	} else if level == "1h" {

	}

	return param
}
func getBotInfo(level string) conf.Bot {
	var bot conf.Bot
	if level == "5m" {
		bot = conf.Config.FiveM.Bot
	} else if level == "15m" {

	} else if level == "30m" {

	} else if level == "1h" {

	}

	return bot
}

func (sr *Srv) PushCard(card data.Card, botInfo conf.BotInfo, num int) {
	c, err := json.MarshalIndent(card, "", "  ")
	if err != nil {
		logger.Log(logger.ErrorLevel, "Card json err: ", err)
	}
	send := `{"msg_type": "interactive","card":` + string(c) + `}`
	if err != nil {
		logger.Log(logger.ErrorLevel, "ChatId err: ", err)
	}
	if err != nil {
		fmt.Println(err.Error())
	}
	sr.Bs.Send(context.Background(), &bot.SendRequest{
		Msg:    send,
		Token:  botInfo.Token,
		ChatId: int64(botInfo.ChatID),
	})

	logger.Log(logger.InfoLevel, "推送消息至：", botInfo.Name, "  推送数量: ", num)
}
func (sr *Srv) PushMsg(msg string, botInfo conf.BotInfo, token string) {

	data := `{
				"msg_type": "text",
				"content": {"text": "` + msg + `"}
			}`

	sr.Bs.Send(context.Background(), &bot.SendRequest{
		Msg:    data,
		Token:  botInfo.Token,
		ChatId: int64(botInfo.ChatID),
	})

	logger.Log(logger.InfoLevel, "推送消息至: ", botInfo.Name, " 代币地址: ", token)
}

func (sr *Srv) checkToken(token string) (*security.CheckTokenResponse, error) {
	//代币安全检查
	s, err := sr.Ss.CheckToken(context.Background(), &security.CheckTokenRequest{
		Token: token,
	})

	if err != nil {
		return nil, err
	}

	return s, nil
}

func (sr *Srv) AutoTrade(pool *data.HotPoolData, s *security.CheckTokenResponse, twNum int) {
	us := sr.Us
	uniPool := uni.Pool{
		Liquidity:            float32(pool.Liquidity),
		Vols:                 float32(pool.Vol),
		Txs:                  int64(pool.Txs),
		Traders:              int64(pool.Traders),
		Age:                  "0",
		IsTransferPause:      s.TransferPause,
		IsSlippageModifiable: s.SlippageModifiable,
		IsBlackList:          s.BlackList,
		IsWhiteList:          s.WhiteList,
		IsCoolDown:           s.CoolDown,
		BuyTax:               s.BuyTax,
		SellTax:              s.SellTax,
		IsBuy:                s.IsBuy,
		IsSellAll:            s.IsSellAll,
		Token:                pool.CoinAddr,
		Price:                float32(pool.Price),
		Twitter:              int64(twNum),
		HiddenOwner:          s.HiddenOwner,
	}
	us.FlashBuyAuto(context.Background(), &uni.FlashBuyAutoRequest{
		Pool: &uniPool,
	})
}
