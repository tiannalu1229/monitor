package model

import (
	"time"
)

type Model struct {
	ID uint `gorm:"column:id; PRIMARY_KEY"`
}

type FlashSwap struct {
	Model
	Token     string    `gorm:"column:token"`
	Symbol    string    `gorm:"column:symbol"`
	User      string    `gorm:"column:user"`
	Cost      string    `gorm:"column:cost"`
	BuyAmount string    `gorm:"column:buy_amount"`
	NowAmount string    `gorm:"column:now_amount"`
	Price     string    `gorm:"column:price"`
	IsSell    int       `gorm:"column:is_sell"`
	BuyTime   time.Time `gorm:"column:buy_time"`
	Hash      string    `gorm:"column:hash"`
	Gas       string    `gorm:"column:gas"`
}

type FlashUser struct {
	Model
	Address              string  `gorm:"column:address"`
	Turn                 int64   `gorm:"column:turn"`
	PrivateKey           string  `gorm:"column:private_key"`
	MaxPerSwap           float64 `gorm:"column:max_per_swap"`
	PoolPercent          float64 `gorm:"column:pool_percent"`
	Liquidity            float64 `gorm:"column:liquidity"`
	Vols                 float64 `gorm:"column:vols"`
	Txs                  int64   `gorm:"column:txs"`
	Traders              int64   `gorm:"column:traders"`
	GasPrice             int64   `gorm:"column:gas_price"`
	Rpc                  string  `gorm:"column:rpc"`
	Receipt              string  `gorm:"column:receipt"`
	Age                  int64   `gorm:"column:age"`
	Level                int64   `gorm:"column:level"`
	BotToken             string  `gorm:"column:bot_token"`
	ChatId               int64   `gorm:"column:chat_id"`
	IsTransferPause      bool    `gorm:"column:is_transfer_pause"`
	IsSlippageModifiable bool    `gorm:"column:is_slippage_modifiable"`
	IsBlackList          bool    `gorm:"column:is_black_list"`
	IsWhiteList          bool    `gorm:"column:is_white_list"`
	IsCoolDown           bool    `gorm:"column:is_cool_down"`
	BuyTax               float32 `gorm:"column:buy_tax"`
	SellTax              float32 `gorm:"column:sell_tax"`
	IsBuy                bool    `gorm:"column:is_buy"`
	IsSellAll            bool    `gorm:"column:is_sell_all"`
	Slippage             int64   `gorm:"column:slippage"`
	HiddenOwner          bool    `gorm:"column:hidden_owner"`
	Twitter              int64   `gorm:"column:twitter"`
}
