package main

import (
	"bot/handler"
	pb "bot/proto"
	consul "consul/proto"
	"context"
	"fmt"
	"go-micro.dev/v4"
	"go-micro.dev/v4/logger"
	"os"
	"os/signal"
	"syscall"
)

const (
	ConsulAddr = "127.0.0.1:8500"
)

var (
	service = "bot"
	host    = "127.0.0.1"
	port    = ":12001"
	version = "latest"
)

func main() {
	// Create service
	srv := micro.NewService(
		micro.Address(port),
	)
	srv.Init(
		micro.Name(service),
		micro.Version(version),
	)

	// Register handler
	if err := pb.RegisterBotHandler(srv.Server(), new(handler.Bot)); err != nil {
		logger.Fatal(err)
	}

	// Register to consul
	c := consul.NewConsulService("consul-register", srv.Client())
	c.RegisterServiceStream(context.Background(), &consul.RegisterServiceStreamRequest{
		Addr:        ConsulAddr,
		ServiceName: service,
		Ip:          host,
		Port:        12001,
		Tag:         "push-bot",
	})

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	// 退出时注销服务
	serviceId := fmt.Sprintf("%s-%s-%d", service, host, 12001)
	c.Deregister(context.Background(), &consul.DeregisterRequest{
		Addr:      ConsulAddr,
		ServiceID: serviceId,
	})
}
