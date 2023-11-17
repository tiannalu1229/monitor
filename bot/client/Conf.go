package conf

import (
	"fmt"
	"github.com/BurntSushi/toml"
)

type TomlConfig struct {
	Out Out `toml:"out"`
	Hot Hot `toml:"hot"`
}

type Out struct {
	Host string `toml:"host"`
	Port string `toml:"port"`
}

type Hot struct {
	Uri      string `toml:"ip"`
	BotToken string `toml:"bot_token"`
	ChatId   int64  `toml:"chat_id"`
}

var Config TomlConfig

func InitConfig(filePath string) {
	if _, err := toml.DecodeFile(filePath, &Config); err != nil {
		fmt.Errorf("failed to init config file: %v", err)
	}
}
