package conf

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

type TomlConfig struct {
	Db    Db    `toml:"db"`
	Redis Redis `toml:"redis"`
	Param Param `toml:"param"`
}

type Redis struct {
	Host     string `toml:"host"`
	Port     int    `toml:"port"`
	Password string `toml:"password"`
	Db       int    `toml:"db"`
}

type Db struct {
	Host     string `toml:"host"`
	Port     uint32 `toml:"port"`
	User     string `toml:"user"`
	Password string `toml:"password"`
	DbName   string `toml:"db_name"`
}

type Param struct {
	Universal  common.Address `toml:"universal"`
	UniRouter  common.Address `toml:"uni_router"`
	PrivateKey string         `toml:"private_key"`
	HttpClient string         `toml:"http_client"`
	User       common.Address `toml:"user"`
	GasPrice   *big.Int       `toml:"gas_price"`
	GasLimit   uint64         `toml:"gas_limit"`
	ChainId    int64          `toml:"chain_id"`
	EthAddress common.Address `toml:"eth_address"`
	EthByte    string         `toml:"eth_byte"`
	TxBot      string         `toml:"tx_bot"`
}

var Config TomlConfig

func InitConfig(filePath string) {
	if _, err := toml.DecodeFile(filePath, &Config); err != nil {
		fmt.Errorf("failed to init config file: %v", err)
	}
}
