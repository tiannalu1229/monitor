package main

import (
	consul "consul/proto"
	"context"
	"db/conf"
	"db/handler"
	utils "db/tool"
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
	service = "database"
	host    = "127.0.0.1"
	port    = ":12004"
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
		micro.Handle(&handler.Db{DataBase: db}),
	)
	srv.Init(
		micro.Name(service),
		micro.Version(version),
	)

	// Register to consul
	c := consul.NewConsulService("consul-register", srv.Client())
	c.RegisterServiceStream(context.Background(), &consul.RegisterServiceStreamRequest{
		Addr:        ConsulAddr,
		ServiceName: service,
		Ip:          host,
		Port:        12004,
		Tag:         "database",
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
