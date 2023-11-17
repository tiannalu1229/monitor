package main

import (
	consul "consul/proto"
	"context"
	"go-micro.dev/v4"
	"uni/handler"
	pb "uni/proto"

	"go-micro.dev/v4/logger"
)

const (
	ConsulAddr = "127.0.0.1:8500"
)

var (
	service = "uni"
	host    = "127.0.0.1"
	port    = ":12011"
	version = "latest"
)

func main() {
	// Create service
	srv := micro.NewService()
	srv.Init(
		micro.Name(service),
		micro.Version(version),
	)

	// Register handler
	if err := pb.RegisterUniHandler(srv.Server(), new(handler.Uni)); err != nil {
		logger.Fatal(err)
	}

	// Register to consul
	c := consul.NewConsulService("consul-register", srv.Client())
	c.RegisterServiceStream(context.Background(), &consul.RegisterServiceStreamRequest{
		Addr:        ConsulAddr,
		ServiceName: service,
		Ip:          host,
		Port:        12011,
		Tag:         "uni",
	})
	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
