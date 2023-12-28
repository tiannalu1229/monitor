package handler

import (
	"context"
	"db/model"
	utils "db/tool"
	"go-micro.dev/v4/metadata"
	"gorm.io/gorm"
	"io"
	"net"
	"time"

	"go-micro.dev/v4/logger"

	pb "db/proto"
)

type Db struct {
	DataBase *gorm.DB
}

func (e *Db) GetPush(ctx context.Context, request *pb.GetPushRequest, response *pb.GetPushResponse) error {
	md, ok := metadata.FromContext(ctx)
	if ok {
		if addr, ok := md["Remote"]; ok {
			host, _, _ := net.SplitHostPort(addr)
			logger.Log(logger.InfoLevel, "接收到: ", host, "请求接口GetPush")
		}
	}
	params := request.Param
	where := utils.ParamHandle(params)
	sql := "select symbol, hva.token, vol, trader, tx, tweet, liquidity, age, hva.times from (select token, min(push_time) push_time, count(*) times from (select * from hots" + where + ") hvc group by token) hva" +
		" left join (select * from hots) hvb on hva.token = hvb.token and hva.push_time = hvb.push_time"

	var pushResults []model.PushResult
	e.DataBase.Raw(sql).Scan(&pushResults)

	var pushData *pb.PushData
	for i := 0; i < len(pushResults); i++ {
		pushData = utils.PushResultToPushData(pushResults[i])
		response.PushData = append(response.PushData, pushData)
	}

	return nil
}

func (e *Db) GetPushDetail(ctx context.Context, request *pb.GetPushDetailRequest, response *pb.GetPushDetailResponse) error {
	md, ok := metadata.FromContext(ctx)
	if ok {
		if addr, ok := md["Remote"]; ok {
			host, _, _ := net.SplitHostPort(addr)
			logger.Log(logger.InfoLevel, "接收到: ", host, "请求接口GetPushDetail")
		}
	}
	params := request.Param
	where := utils.ParamHandle(params)

	sql := "select push_time time, level, vol, trader, tx, tweet, liquidity, price, age, times, type from hots" + where

	var pushDetailResults []model.PushDetailResult
	e.DataBase.Raw(sql).Scan(&pushDetailResults)

	var pushDetailData *pb.PushDetailData
	for i := 0; i < len(pushDetailResults); i++ {
		pushDetailData = utils.PushDetailResultToPushDetailData(pushDetailResults[i])
		response.PushDetailData = append(response.PushDetailData, pushDetailData)
	}
	return nil
}

func (e *Db) Call(ctx context.Context, req *pb.CallRequest, rsp *pb.CallResponse) error {
	logger.Infof("Received Db.Call request: %v", req)
	rsp.Msg = "Hello " + req.Name
	return nil
}

func (e *Db) ClientStream(ctx context.Context, stream pb.Db_ClientStreamStream) error {
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

func (e *Db) ServerStream(ctx context.Context, req *pb.ServerStreamRequest, stream pb.Db_ServerStreamStream) error {
	logger.Infof("Received Db.ServerStream request: %v", req)
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

func (e *Db) BidiStream(ctx context.Context, stream pb.Db_BidiStreamStream) error {
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
