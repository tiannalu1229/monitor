package main

import (
	"github.com/gorilla/websocket"
	"go-micro.dev/v4/logger"
	"net/url"
	"os"
	"os/signal"
	"time"
)

func main() {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	u := url.URL{Scheme: "ws", Host: "8.210.97.145:5007", Path: "/tweet/new/1729460081446863028"} // 更换为你的WebSocket服务地址
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
		}
	}()

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-done:
			return
		case t := <-ticker.C:
			err := c.WriteMessage(websocket.TextMessage, []byte(t.String()))
			if err != nil {
				logger.Log(logger.ErrorLevel, "write:", err)
				return
			}
		case <-interrupt:
			logger.Log(logger.InfoLevel, "interrupt")

			// Cleanly close the connection by sending a close message and then
			// waiting (with timeout) for the server to close the connection.
			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				logger.Log(logger.ErrorLevel, "write close:", err)
				return
			}
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			c.Close()
			return
		}
	}
}
