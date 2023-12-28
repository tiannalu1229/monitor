package handler

import (
	"bufio"
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"io"
	"os"
	"tg-bot/core"
	"time"

	"go-micro.dev/v4/logger"

	pb "tg-bot/proto"
)

type TgBot struct {
	Bot *tgbotapi.BotAPI
}

func (e *TgBot) ReceiveMsg(ctx context.Context, request *pb.ReceiveMsgRequest, response *pb.ReceiveMsgResponse) error {

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	ctx = context.Background()
	ctx, cancel := context.WithCancel(ctx)

	updates := e.Bot.GetUpdatesChan(u)
	go core.ReceiveUpdates(ctx, updates)

	// Tell the user the bot is online
	logger.Log(logger.InfoLevel, "Start listening for updates. Press enter to stop")

	// Wait for a newline symbol, then cancel handling updates
	bufio.NewReader(os.Stdin).ReadBytes('\n')
	cancel()

	return nil
}

func (e *TgBot) Call(ctx context.Context, req *pb.CallRequest, rsp *pb.CallResponse) error {
	logger.Infof("Received TgBot.Call request: %v", req)
	rsp.Msg = "Hello " + req.Name
	return nil
}

func (e *TgBot) ClientStream(ctx context.Context, stream pb.TgBot_ClientStreamStream) error {
	var count int64
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			logger.Infof("Got %v pings total", count)
			return stream.SendMsg(&pb.ClientStreamResponse{Count: count})
		}
		if err != nil {
			return err
		}
		logger.Infof("Got ping %v", req.Stroke)
		count++
	}
}

func (e *TgBot) ServerStream(ctx context.Context, req *pb.ServerStreamRequest, stream pb.TgBot_ServerStreamStream) error {
	logger.Infof("Received TgBot.ServerStream request: %v", req)
	for i := 0; i < int(req.Count); i++ {
		logger.Infof("Sending %d", i)
		if err := stream.Send(&pb.ServerStreamResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
		time.Sleep(time.Millisecond * 250)
	}
	return nil
}

func (e *TgBot) BidiStream(ctx context.Context, stream pb.TgBot_BidiStreamStream) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		logger.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&pb.BidiStreamResponse{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
