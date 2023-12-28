package handle

import (
	"bot/client/conf"
	"bot/client/data"
	"bot/client/tweet"
	"bot/client/uni-new/db"
	bot "bot/proto"
	"context"
	"encoding/json"
	"fmt"
	"go-micro.dev/v4"
	"go-micro.dev/v4/logger"
	"gorm.io/gorm"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

var TokenNum map[string]int64

func InitTokenNumMap() {
	TokenNum = make(map[string]int64)
}

func GetNewPair(c *conf.TomlConfig) []data.NewPairData {
	url := c.New.NewURI
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

	var newPair data.NewPair
	json.Unmarshal(body, &newPair)

	return newPair.Data
}

func NewPairHandle(pairs []data.NewPairData, database *gorm.DB, c *conf.TomlConfig) {

	t := tweet.Tweet{
		Url: conf.Config.Tweet.KeywordURI,
	}
	for i := 0; i < len(pairs); i++ {
		token := pairs[i].CoinAddr
		pair := pairs[i]
		numToken := t.GetTweetNum(token)
		numPair := t.GetTweetNum(pair.Pair)
		num := numToken + numPair
		pushLog := db.GetNewLatest(database, token)
		result := CheckNum(num, pushLog.Tweet)
		if result {
			//发送消息
			logger.Log(logger.InfoLevel, "推送消息至：", "【V2】新池+推特提及")
			msg := "名称 ：" + pair.CoinSymbol + "/ETH" + "(" + strconv.Itoa(num) + "次)" +
				"\\n流动性 : " + fmt.Sprintf("%f", pair.StableReserve) + " ETH" +
				"\\n代币地址 : " + pair.CoinAddr +
				"\\n年龄 : " + pair.Age +
				"\\n推送次数 : " + strconv.Itoa(int(pushLog.Times)+1)
			sendMsg := `{
							"msg_type": "text",
							"content": {"text": "` + msg + `"}
						}`
			service := micro.NewService()
			service.Init()
			bs := bot.NewBotService("msg-bot", service.Client())
			bs.Send(context.Background(), &bot.SendRequest{
				Msg:    sendMsg,
				Token:  c.New.NewToken,
				ChatId: int64(c.New.NewChatID),
			})

			db.SaveNew(database, &pairs[i], num, time.Now(), pushLog.Times+1)
		}
	}
}

func CheckNum(num int, oldNum int) bool {
	result := false
	if num >= 10 {
		if num-oldNum > 5 {
			result = true
		}
	} else if num >= 6 {
		if oldNum < 6 {
			result = true
		}
	} else if num >= 3 {
		if oldNum < 3 {
			result = true
		}
	}

	return result
}
