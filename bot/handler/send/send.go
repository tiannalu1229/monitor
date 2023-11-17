package send

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"net/http"
	"strings"
)

func TgSend(token string, chatID int64, message string) error {

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
		return err
	}

	bot.Debug = true

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	msg := tgbotapi.NewMessage(chatID, message)
	_, err = bot.Send(msg)
	if err != nil {
		log.Panic(err)
		return err
	}

	return nil
}

func BookSend(token string, message string) error {
	url := "https://open.feishu.cn/open-apis/bot/v2/hook/" + token //要访问的Url地址
	context := "application/json"

	sendData := `{
		"msg_type": "text",
		"content": {"text": "` + message + `"}
	}`

	result, err := http.Post(url, context, strings.NewReader(sendData))
	if err != nil {
		fmt.Printf("post failed, err:%v\n", err)
		return err
	}
	defer result.Body.Close()
	return nil
}
