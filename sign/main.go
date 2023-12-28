package main

import (
	"sign/conf"
	"sign/handler"
	pb "sign/proto"
	utils "uni/db/tool"

	"go-micro.dev/v4"
	"go-micro.dev/v4/logger"
)

var (
	service = "sign"
	host    = "127.0.0.1"
	port    = ":12010"
	version = "latest"
)

func main() {
	conf.InitConfig("./config/config.toml")
	rdb := utils.ConnectRedis(conf.Config.Redis.Host, conf.Config.Redis.Port, conf.Config.Redis.Password)

	// Create service
	srv := micro.NewService(
		micro.Address(port),
		micro.Name(service),
		micro.Handle(&handler.Sign{Rdb: rdb}),
	)
	srv.Init(
		micro.Name(service),
		micro.Version(version),
	)

	// Register handler
	if err := pb.RegisterSignHandler(srv.Server(), new(handler.Sign)); err != nil {
		logger.Fatal(err)
	}
	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
