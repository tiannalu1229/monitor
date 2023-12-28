package handler

import (
	"context"
	"export/db"
	"gorm.io/gorm"
	"io"
	"os"
	"time"

	"go-micro.dev/v4/logger"

	pb "export/proto"
)

type Export struct {
	DataBase *gorm.DB
}

func (e *Export) Export(ctx context.Context, request *pb.ExportRequest, response *pb.ExportResponse) error {

	q := &db.Query{DataBase: e.DataBase}
	fileName := q.FirstDaily(request.Score)
	response.ChunkData, _ = os.ReadFile("./excel/" + fileName)
	os.WriteFile(fileName, response.ChunkData, 0644)
	return nil
}

func (e *Export) Call(ctx context.Context, req *pb.CallRequest, rsp *pb.CallResponse) error {
	logger.Infof("Received Export.Call request: %v", req)
	rsp.Msg = "Hello " + req.Name
	return nil
}

func (e *Export) ClientStream(ctx context.Context, stream pb.Export_ClientStreamStream) error {
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

func (e *Export) ServerStream(ctx context.Context, req *pb.ServerStreamRequest, stream pb.Export_ServerStreamStream) error {
	logger.Infof("Received Export.ServerStream request: %v", req)
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

func (e *Export) BidiStream(ctx context.Context, stream pb.Export_BidiStreamStream) error {
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
