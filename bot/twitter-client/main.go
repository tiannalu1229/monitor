package main

import (
	"bot/twitter-client/conf"
	"bot/twitter-client/core"
	"fmt"
)

func main() {
	conf.InitConfig("./config/config.toml")
	fmt.Println(conf.Config.Db.Port)
	core.Start()
}
