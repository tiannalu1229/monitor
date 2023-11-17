package main

import (
	"consul/handler"
	pb "consul/proto"
	"go-micro.dev/v4"
	"go-micro.dev/v4/logger"
)

var (
	service = "consul-register"
	host    = "127.0.0.1"
	port    = ":12000"
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
	if err := pb.RegisterConsulHandler(srv.Server(), new(handler.Consul)); err != nil {
		logger.Fatal(err)
	}
	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
