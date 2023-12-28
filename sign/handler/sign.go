package handler

import (
	"context"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/go-redis/redis/v8"
	"go-micro.dev/v4/metadata"
	"io"
	"net"
	"strings"
	"time"

	"go-micro.dev/v4/logger"

	pb "sign/proto"
)

type Sign struct {
	Rdb *redis.Client
}

func (e *Sign) CheckSign(ctx context.Context, request *pb.CheckSignRequest, response *pb.CheckSignResponse) error {
	md, ok := metadata.FromContext(ctx)
	if ok {
		if addr, ok := md["Remote"]; ok {
			host, _, _ := net.SplitHostPort(addr)
			logger.Log(logger.InfoLevel, "接收到: ", host, "请求接口CheckSign")
		}
	}
	sign, err := e.Rdb.Get(ctx, strings.ToLower(request.Address)+"-sign").Result()
	if errors.Is(err, redis.Nil) {
		// 缓存未命中，使用默认值
		sign = ""
		response.Result = false
	} else if err != nil {
		// 发生其他错误
		sign = ""
		response.Result = false
	}

	if sign != "" {
		response.Result = true
	} else {
		response.Result = false
	}

	return nil
}

func (e *Sign) Sign(ctx context.Context, request *pb.SignRequest, response *pb.SignResponse) error {

	md, ok := metadata.FromContext(ctx)
	if ok {
		if addr, ok := md["Remote"]; ok {
			host, _, _ := net.SplitHostPort(addr)
			logger.Log(logger.InfoLevel, "接收到: ", host, "请求接口Sign")
		}
	}
	message := request.Message

	prefix := []byte("\x19Ethereum Signed Message:\n")

	data := append(prefix, append([]byte(fmt.Sprintf("%d", len(message))), []byte(message)...)...)
	hash := crypto.Keccak256Hash(data)

	signatureHex := request.Sign

	signature := common.Hex2Bytes(signatureHex[2:])

	sigWithoutRecoveryID := signature[:len(signature)-1]

	recoveryID := signature[len(signature)-1] - 27

	pubKey, err := crypto.SigToPub(hash.Bytes(), append(sigWithoutRecoveryID, recoveryID))
	if err != nil {
		logger.Log(logger.ErrorLevel, "Could not recover public key: %v", err)
		return err
	}

	recoveredAddr := crypto.PubkeyToAddress(*pubKey)

	response.Result = recoveredAddr.Hex() == request.Address
	e.Rdb.Set(ctx, strings.ToLower(request.Address)+"-sign", request.Sign, time.Hour)
	return nil
}

func (e *Sign) Call(ctx context.Context, req *pb.CallRequest, rsp *pb.CallResponse) error {
	logger.Infof("Received Sign.Call request: %v", req)
	rsp.Msg = "Hello " + req.Name
	return nil
}

func (e *Sign) ClientStream(ctx context.Context, stream pb.Sign_ClientStreamStream) error {
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

func (e *Sign) ServerStream(ctx context.Context, req *pb.ServerStreamRequest, stream pb.Sign_ServerStreamStream) error {
	logger.Infof("Received Sign.ServerStream request: %v", req)
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

func (e *Sign) BidiStream(ctx context.Context, stream pb.Sign_BidiStreamStream) error {
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
