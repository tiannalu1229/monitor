package handle

import (
	bot "bot/proto"
	"bot/twitter-client/conf"
	"bot/twitter-client/core/result"
	"bot/twitter-client/db/model"
	"bot/twitter-client/db/op"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"go-micro.dev/v4/logger"
	"gorm.io/gorm"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"sync"
	"time"
)

type TwitterHandle struct {
	Db    *gorm.DB
	Bs    *bot.BotService
	Param *conf.Param
}

// TwitterNewHandle 观察列表高频提及/**
func (t *TwitterHandle) TwitterNewHandle() {
	uri := t.Param.New.URI
	bot := t.Param.New.Bot
	db := op.Database{Db: t.Db}
	for i := 0; i < len(t.Param.New.Check); i++ {
		var wg sync.WaitGroup
		errChan := make(chan error, 1)
		wg.Add(1)
		go func(checkId int) {
			check := t.Param.New.Check[checkId]
			twUri := strings.Replace(uri, "%min%", strconv.Itoa(check.N), -1)
			twUri = strings.Replace(twUri, "%list%", check.List, -1)
			data := GetTweetAnalyse(twUri)
			for key, analyseData := range data {
				for j := 0; j < len(analyseData); j++ {
					keyword := prefixExchange(key) + analyseData[j].Key
					log := db.GetTweetAnalyseLog(keyword)
					if log.Tweet < int64(analyseData[j].Count) {
						send := "list name: " + check.Name +
							"\\n提及: " + keyword + " +" + strconv.Itoa(analyseData[j].Count)
						t.PushMsg(send, bot[0])
						newLog := model.TweetAnalysePushLog{
							Keyword:  keyword,
							ListId:   check.List,
							ListName: check.Name,
							Times:    log.Times + 1,
							Tweet:    int64(analyseData[j].Count),
							PushTime: time.Now(),
						}
						db.InsetTweetAnalyseLog(&newLog)
					}
				}
			}
		}(i)
		close(errChan) // 关闭channel
		// 处理可能发生的错误
		for err := range errChan {
			if err != nil {
				logger.Log(logger.ErrorLevel, "交易失败: ", err)
			}
		}
	}
}

// TwitterListHandle 观察列表实时消息推送/**
func (t *TwitterHandle) TwitterListHandle(list string) {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	u := url.URL{Scheme: "ws", Host: "8.210.97.145:5007", Path: "/tweet/new/" + list}
	logger.Log(logger.InfoLevel, "connecting to %s", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		logger.Log(logger.ErrorLevel, "dial:", err)
	}
	defer c.Close()

	done := make(chan struct{})

	go func() {
		defer close(done)
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				logger.Log(logger.ErrorLevel, "read:", err)
				return
			}
			logger.Log(logger.InfoLevel, "recv: %s", message)

			var newTweet result.TweetListNew
			json.Unmarshal(message, &newTweet)

			if newTweet.Event == "new" {
				send := "@" + newTweet.Data.ScreenName + " 发推提醒" +
					"\\n" + "tweet内容:" +
					"\\n" + newTweet.Data.Text
				bot := conf.Bot{
					Token:  "",
					ChatID: 0,
					Name:   "",
				}
				t.PushMsg(send, bot)
			}
		}
	}()
}

func GetTweetAnalyse(url string) map[string][]result.TweetAnalyseData {
	method := "GET"

	payload := strings.NewReader(``)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return nil
	}
	req.Header.Add("Content-Type", "application/json")

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

	var tweetAnalyse result.TweetAnalyse
	json.Unmarshal(body, &tweetAnalyse)
	if tweetAnalyse.Code == 1 {
		return tweetAnalyse.Data
	}

	return nil
}

func (t *TwitterHandle) PushMsg(msg string, botInfo conf.Bot) {

	data := `{
				"msg_type": "text",
				"content": {"text": "` + msg + `"}
			}`

	(*t.Bs).Send(context.Background(), &bot.SendRequest{
		Msg:    data,
		Token:  botInfo.Token,
		ChatId: int64(botInfo.ChatID),
	})

	logger.Log(logger.InfoLevel, "推送消息至: ", botInfo.Name)
}

func prefixExchange(prefix string) string {
	if prefix == "at" {
		return "@"
	}
	if prefix == "dollar" {
		return "$"
	}
	if prefix == "address" {
		return ""
	}
	return ""
}
