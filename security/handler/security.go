package handler

import (
	"context"
	"gorm.io/gorm"
	"io"
	"security/core"
	"strings"
	"time"

	"go-micro.dev/v4/logger"

	pb "security/proto"
)

type Security struct {
	Db *gorm.DB
}

func (e *Security) CheckToken(ctx context.Context, request *pb.CheckTokenRequest, response *pb.CheckTokenResponse) error {
	token := strings.ToLower(request.Token)
	db := core.Database{Db: e.Db}
	result := db.GetSecurity(token)
	if result.ID <= 0 || result.CheckTime.Unix() < time.Now().Unix()-300 {
		logger.Log(logger.InfoLevel, "开始审计代币: ", token)
		gp := core.GetTokenSecurityGoPlus(token)
		if gp.Code != 1 {
			logger.Log(logger.ErrorLevel, "审计代币问题: ", gp.Message)
			return nil
		}
		g := gp.Result
		s := db.SaveSecurity(g, result.ID)
		response.Owner = s.IsOwnerShip
		response.Honey = s.IsHoney
		response.BuyTax = float32(s.BuyTax)
		response.SellTax = float32(s.SellTax)
		response.Pause = s.IsTransferPause
		response.Score = int64(s.Score)
		response.OpenSource = s.IsOpenSource
		response.HiddenOwner = s.HiddenOwner
		response.Proxy = s.IsProxy
		response.CoolDown = s.IsCoolDown
		response.Lock = s.IsLock
		response.Mint = s.IsMintAble
		response.TransferPause = s.IsTransferPause
		response.SlippageModifiable = s.IsSlippageModifiable
		response.BlackList = s.IsBlackList
		response.WhiteList = s.IsWhiteList
		response.IsBuy = s.IsBuy
		response.IsSellAll = s.IsSellAll
		response.HiddenOwner = s.HiddenOwner
	} else {
		response.Owner = result.IsOwnerShip
		response.Honey = result.IsHoney
		response.BuyTax = float32(result.BuyTax)
		response.SellTax = float32(result.SellTax)
		response.Pause = result.IsTransferPause
		response.Score = int64(result.Score)
		response.OpenSource = result.IsOpenSource
		response.HiddenOwner = result.HiddenOwner
		response.Proxy = result.IsProxy
		response.CoolDown = result.IsCoolDown
		response.Lock = result.IsLock
		response.Mint = result.IsMintAble
		response.TransferPause = result.IsTransferPause
		response.SlippageModifiable = result.IsSlippageModifiable
		response.BlackList = result.IsBlackList
		response.WhiteList = result.IsWhiteList
		response.IsBuy = result.IsBuy
		response.IsSellAll = result.IsSellAll
		response.HiddenOwner = result.HiddenOwner
	}

	return nil
}

func (e *Security) Call(ctx context.Context, req *pb.CallRequest, rsp *pb.CallResponse) error {
	logger.Infof("Received Security.Call request: %v", req)
	rsp.Msg = "Hello " + req.Name
	return nil
}

func (e *Security) ClientStream(ctx context.Context, stream pb.Security_ClientStreamStream) error {
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

func (e *Security) ServerStream(ctx context.Context, req *pb.ServerStreamRequest, stream pb.Security_ServerStreamStream) error {
	logger.Infof("Received Security.ServerStream request: %v", req)
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

func (e *Security) BidiStream(ctx context.Context, stream pb.Security_BidiStreamStream) error {
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
