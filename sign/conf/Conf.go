package conf

import (
	"fmt"
	"github.com/BurntSushi/toml"
)

type TomlConfig struct {
	Redis Redis `toml:"redis"`
}

type Redis struct {
	Host     string `toml:"host"`
	Port     int    `toml:"port"`
	Password string `toml:"password"`
	Db       int    `toml:"db"`
}

var Config TomlConfig

func InitConfig(filePath string) {
	if _, err := toml.DecodeFile(filePath, &Config); err != nil {
		fmt.Errorf("failed to init config file: %v", err)
	}
}
