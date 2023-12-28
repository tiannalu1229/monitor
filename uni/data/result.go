package data

import (
	"time"
	"uni/db/model"
)

type UserResult struct {
	model.Model
	Address              string
	Turn                 int64
	PrivateKey           string
	MaxPerSwap           float32
	PoolPercent          float32
	Liquidity            float32
	Vols                 float32
	Txs                  int64
	Traders              int64
	GasPrice             int64
	Rpc                  string
	Receipt              string
	Age                  int64
	Level                int64
	IsTransferPause      bool
	IsSlippageModifiable bool
	IsBlackList          bool
	IsWhiteList          bool
	IsCoolDown           bool
	BuyTax               float32
	SellTax              float32
	IsBuy                bool
	IsSellAll            bool
	Slippage             int64
	HiddenOwner          bool
	Twitter              int64
	BotToken             string
	ChatId               int64
}

type SwapResult struct {
	model.Model
	Token   string
	Symbol  string
	Cost    string
	Buy     string
	Now     string
	Price   string
	BuyTime time.Time
	Hash    string
	IsSell  int64
}
