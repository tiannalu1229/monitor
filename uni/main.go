package main

import (
	"go-micro.dev/v4"
	"go-micro.dev/v4/logger"
	"uni/conf"
	"uni/db/model"
	utils "uni/db/tool"
	"uni/handler"
)

const (
	ConsulAddr = "127.0.0.1:8500"
)

var (
	service = "flashswap"
	host    = "127.0.0.1"
	port    = ":12003"
	version = "latest"
)

func main() {
	conf.InitConfig("./config/config.toml")

	db, err := utils.ConnectPg(conf.Config.Db.Host, conf.Config.Db.User, conf.Config.Db.Password, conf.Config.Db.DbName, conf.Config.Db.Port)
	if err != nil {
		logger.Log(logger.ErrorLevel, "db conn err: ", err)
	}

	db.AutoMigrate(
		&model.FlashSwap{},
		&model.FlashUser{},
	)

	rdb := utils.ConnectRedis(conf.Config.Redis.Host, conf.Config.Redis.Port, conf.Config.Redis.Password)
	//Create service
	srv := micro.NewService(
		micro.Address(port),
		micro.Name(service),
		micro.Handle(&handler.Uni{Db: db, Rdb: rdb}),
	)
	srv.Init(
		micro.Name(service),
		micro.Version(version),
	)
	//if err := pb.RegisterUniHandler(srv.Server(), new(handler.Uni)); err != nil {
	//	logger.Fatal(err)
	//}
	// Run service
	if err = srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
