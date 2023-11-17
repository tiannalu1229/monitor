package handler

import (
	"bot/handler/send"
	"context"
	"io"
	"time"

	"go-micro.dev/v4/logger"

	pb "bot/proto"
)

type Bot struct{}

func (e *Bot) Send(ctx context.Context, request *pb.SendRequest, response *pb.SendResponse) error {
	token := request.Token
	chatId := request.ChatId
	msg := request.Msg

	var err error
	if chatId < 0 {
		err = send.TgSend(token, chatId, msg)
	} else {
		err = send.BookSend(token, msg)
	}

	if err != nil {
		return err
	}

	return nil
}

func (e *Bot) Call(ctx context.Context, req *pb.CallRequest, rsp *pb.CallResponse) error {
	logger.Infof("Received Bot.Call request: %v", req)
	rsp.Msg = "Hello " + req.Name
	return nil
}

func (e *Bot) ClientStream(ctx context.Context, stream pb.Bot_ClientStreamStream) error {
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

func (e *Bot) ServerStream(ctx context.Context, req *pb.ServerStreamRequest, stream pb.Bot_ServerStreamStream) error {
	logger.Infof("Received Bot.ServerStream request: %v", req)
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

func (e *Bot) BidiStream(ctx context.Context, stream pb.Bot_BidiStreamStream) error {
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
