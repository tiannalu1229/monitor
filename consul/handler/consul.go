package handler

import (
	"context"
	"fmt"
	"github.com/hashicorp/consul/api"
	"go-micro.dev/v4/logger"
	"io"
	"time"

	pb "consul/proto"
)

type Consul struct {
}

func (e *Consul) Deregister(ctx context.Context, request *pb.DeregisterRequest, response *pb.DeregisterResponse) error {
	cfg := api.DefaultConfig()
	cfg.Address = request.Addr
	c, err := api.NewClient(cfg)
	if err != nil {
		return err
	}
	return c.Agent().ServiceDeregister(request.ServiceID)
}

func (e *Consul) RegisterServiceStream(ctx context.Context, request *pb.RegisterServiceStreamRequest, response *pb.RegisterServiceStreamResponse) error {
	cfg := api.DefaultConfig()
	cfg.Address = request.Addr
	c, err := api.NewClient(cfg)
	if err != nil {
		return err
	}

	fmt.Println("开始注册服务")

	//check := &api.AgentServiceCheck{
	//	GRPC:     fmt.Sprintf("%s:%d", request.Ip, request.Port), // 这里一定是外部可以访问的地址
	//	Timeout:  "10s",                                          // 超时时间
	//	Interval: "10s",                                          // 运行检查的频率
	//	// 指定时间后自动注销不健康的服务节点
	//	// 最小超时时间为1分钟，收获不健康服务的进程每30秒运行一次，因此触发注销的时间可能略长于配置的超时时间。
	//	DeregisterCriticalServiceAfter: "1m",
	//}

	srv := &api.AgentServiceRegistration{
		ID:      fmt.Sprintf("%s-%s-%d", request.ServiceName, request.Ip, request.Port), // 服务唯一ID
		Name:    request.ServiceName,                                                    // 服务名称
		Tags:    []string{request.Tag},                                                  // 为服务打标签
		Address: request.Ip,
		Port:    int(request.Port),
		//Check:   check,
	}
	return c.Agent().ServiceRegister(srv)
}

func (e *Consul) Call(ctx context.Context, req *pb.CallRequest, rsp *pb.CallResponse) error {
	logger.Infof("Received Consul.Call request: %v", req)
	rsp.Msg = "Hello " + req.Name
	return nil
}

func (e *Consul) ClientStream(ctx context.Context, stream pb.Consul_ClientStreamStream) error {
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

func (e *Consul) ServerStream(ctx context.Context, req *pb.ServerStreamRequest, stream pb.Consul_ServerStreamStream) error {
	logger.Infof("Received Consul.ServerStream request: %v", req)
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

func (e *Consul) BidiStream(ctx context.Context, stream pb.Consul_BidiStreamStream) error {
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
