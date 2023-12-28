package main

import (
	"export/conf"
	"export/handler"
	utils "export/tool"
	"go-micro.dev/v4"

	"go-micro.dev/v4/logger"
)

var (
	service = "export"
	host    = "127.0.0.1"
	port    = ":12099"
	version = "latest"
)

func main() {
	conf.InitConfig("./config/config.toml")

	db, err := utils.ConnectPg(conf.Config.Db.Host, conf.Config.Db.User, conf.Config.Db.Password, conf.Config.Db.DbName, conf.Config.Db.Port)
	if err != nil {
		logger.Log(logger.ErrorLevel, "db conn err: ", err)
	}
	//Create service
	srv := micro.NewService(
		micro.Address(port),
		micro.Name(service),
		micro.Handle(&handler.Export{DataBase: db}),
	)
	// Create service
	srv.Init(
		micro.Name(service),
		micro.Version(version),
	)

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
