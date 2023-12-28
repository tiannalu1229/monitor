package conf

import (
	"fmt"
	"github.com/BurntSushi/toml"
)

type TomlConfig struct {
	Db  Db  `toml:"db"`
	Uri Uri `toml:"uri"`
}

type Db struct {
	Host     string `toml:"host"`
	Port     uint32 `toml:"port"`
	User     string `toml:"user"`
	Password string `toml:"password"`
	DbName   string `toml:"db_name"`
}

type Uri struct {
	Pair string `toml:"pair"`
}

var Config TomlConfig

func InitConfig(filePath string) {
	if _, err := toml.DecodeFile(filePath, &Config); err != nil {
		fmt.Errorf("failed to init config file: %v", err)
	}
}
