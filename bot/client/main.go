package main

import (
	"bot/client/conf"
	"bot/client/core"
)

func main() {
	conf.InitConfig("./config/config.toml")
	core.Start()
}
