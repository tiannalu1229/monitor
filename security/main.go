package main

import (
	"go-micro.dev/v4"
	"go-micro.dev/v4/logger"
	"security/conf"
	"security/db/model"
	utils "security/db/tool"
	"security/handler"
)

var (
	service = "security"
	host    = "127.0.0.1"
	port    = ":12088"
	version = "latest"
)

func main() {

	conf.InitConfig("./config/config.toml")

	db, err := utils.ConnectPg(conf.Config.Db.Host, conf.Config.Db.User, conf.Config.Db.Password, conf.Config.Db.DbName, conf.Config.Db.Port)
	if err != nil {
		logger.Log(logger.ErrorLevel, "db conn err: ", err)
	}

	db.AutoMigrate(
		&model.Security{},
	)

	// Create service
	srv := micro.NewService(
		micro.Address(port),
		micro.Name(service),
		micro.Handle(&handler.Security{Db: db}),
	)
	srv.Init(
		micro.Name(service),
		micro.Version(version),
	)

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
