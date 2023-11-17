package handle

import (
	conf "bot/client"
	"bot/client/uni/data"
	bot "bot/proto"
	"context"
	"encoding/json"
	"fmt"
	"go-micro.dev/v4"
	"io/ioutil"
	"net/http"
	"strconv"
)

var HotPool map[string]int

func InitHotPoolMap() {
	TokenNum = make(map[string]int)
}

func GetHotPool() []data.HotPoolData {
	url := conf.Config.Hot.Uri
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

func HotPoolHandle(pool []data.HotPoolData) {
	for i := 0; i < len(pool); i++ {
		token := pool[i].CoinAddr
		fmt.Println("代币地址 : ", token)
		if (pool[i].TxsBuy+pool[i].TxsSell) >= 50 || (pool[i].VolBuy+pool[i].VolSell) >= 5 {

			if HotPool[pool[i].CoinAddr] <= 0 {
				//发送消息
				fmt.Print("即将推送交易对 : ", pool[i].Pair)

				msg := "热门交易 ： " +
					"\\n名称 ：" + pool[i].CoinSymbol + "/ETH" +
					"\\n代币地址 : " + pool[i].CoinAddr +
					"\\n5m内交易笔数 : " + strconv.Itoa(pool[i].TxsBuy+pool[i].TxsSell) +
					"\\n5m内交易量 : " + strconv.Itoa(pool[i].VolSell+pool[i].VolBuy) +
					"\\n流动性 : " + fmt.Sprintf("%f", pool[i].Liquidaty) + " ETH" +
					"\\n年龄 : " + strconv.Itoa(pool[i].CreateTime)

				service := micro.NewService()
				service.Init()
				bs := bot.NewBotService("bot", service.Client())
				bs.Send(context.Background(), &bot.SendRequest{
					Msg:    msg,
					Token:  conf.Config.Hot.BotToken,
					ChatId: conf.Config.Hot.ChatId,
				})

				HotPool[pool[i].CoinAddr] += 1
			}

		}
	}
}
