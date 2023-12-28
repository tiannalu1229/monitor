package conf

import (
	"fmt"
	"github.com/BurntSushi/toml"
)

type TomlConfig struct {
	Db    Db    `toml:"db"`
	Param Param `toml:"param"`
}

type Db struct {
	Host     string `toml:"host"`
	Port     int    `toml:"port"`
	User     string `toml:"user"`
	Password string `toml:"password"`
	DbName   string `toml:"db_name"`
}

type Param struct {
	New New `toml:"new"`
}

type New struct {
	Bot   []Bot   `toml:"bot"`
	Check []Check `toml:"check"`
	URI   string  `toml:"uri"`
}

type Bot struct {
	Token  string `toml:"token"`
	ChatID int    `toml:"chat_id"`
	Name   string `toml:"name"`
}

type Check struct {
	Name string `toml:"name"`
	List string `toml:"list"`
	N    int    `toml:"n"`
}

var Config TomlConfig

func InitConfig(filePath string) {
	if _, err := toml.DecodeFile(filePath, &Config); err != nil {
		fmt.Errorf("failed to init config file: %v", err)
	}
}
