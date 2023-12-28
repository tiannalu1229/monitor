package conf

import (
	"fmt"
	"github.com/BurntSushi/toml"
)

type TomlConfig struct {
	Out struct {
		Host  string `toml:"host"`
		Port  string `toml:"port"`
		Front string `toml:"front"`
	} `toml:"out"`
	Db struct {
		Host     string `toml:"host"`
		Port     int    `toml:"port"`
		User     string `toml:"user"`
		Password string `toml:"password"`
		DbName   string `toml:"db_name"`
	} `toml:"db"`
	Ctl struct {
		Level []string `toml:"level"`
	} `toml:"ctl"`
	FiveM struct {
		Hot HotParam `toml:"hot"`
		Bot Bot      `toml:"bot"`
	} `toml:"5m"`
	New struct {
		NewURI    string `toml:"new_uri"`
		NewToken  string `toml:"new_token"`
		NewChatID int    `toml:"new_chat_id"`
	} `toml:"new"`
	Tweet struct {
		KeywordURI string `toml:"keyword_uri"`
	} `toml:"tweet"`
}

type HotParam struct {
	URI           string  `toml:"uri"`
	MinVol        float64 `toml:"min_vol"`
	MinTx         int     `toml:"min_tx"`
	MinTrader     int     `toml:"min_trader"`
	HotPoolMinEth float64 `toml:"hot_pool_min_eth"`
	TweetIncrease int     `toml:"tweet_increase"`
}

type Bot struct {
	Vol    []BotInfo `toml:"vol"`
	Tx     []BotInfo `toml:"tx"`
	Trader []BotInfo `toml:"trader"`
	Tweet  []BotInfo `toml:"tweet"`
	Mix    []BotInfo `toml:"mix"`
	Safe   []BotInfo `toml:"safe"`
}

type BotInfo struct {
	Token  string `toml:"token"`
	ChatID int    `toml:"chat_id"`
	Name   string `toml:"name"`
}

var Config TomlConfig

func InitConfig(filePath string) {
	if _, err := toml.DecodeFile(filePath, &Config); err != nil {
		fmt.Errorf("failed to init config file: %v", err)
	}
}
