package main

import (
	pb "consul/proto"
	"context"
	"go-micro.dev/v4"
	"go-micro.dev/v4/logger"
)

func main() {
	// create a new service
	service := micro.NewService()

	// parse command line flags
	service.Init()

	c := pb.NewConsulService("consul-register", service.Client())
	//c.Deregister(context.Background(), &pb.DeregisterRequest{
	//	Addr:      "127.0.0.1:8500",
	//	ServiceID: "consul-register-127.0.0.1-12000",
	//})
	c.RegisterServiceStream(context.Background(), &pb.RegisterServiceStreamRequest{
		Addr:        "127.0.0.1:8500",
		ServiceName: "consul-register",
		Tag:         "register",
		Ip:          "127.0.0.1",
		Port:        12000,
	})

	logger.Info("服务注册完毕, 端口号：", 12000)
}
