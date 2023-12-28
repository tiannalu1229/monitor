package model

import "time"

type Model struct {
	ID uint `gorm:"column:id; PRIMARY_KEY"`
}

type Security struct {
	Model
	Token                string    `gorm:"column:token"`
	IsOpenSource         bool      `gorm:"column:is_open_source"`
	IsProxy              bool      `gorm:"column:is_proxy"`
	IsMintAble           bool      `gorm:"column:is_mint_able"`
	IsHoney              bool      `gorm:"column:is_honey"`
	IsLock               bool      `gorm:"column:is_lock"`
	IsOwnerShip          bool      `gorm:"column:is_owner_ship"`
	HiddenOwner          bool      `gorm:"column:hidden_owner"`
	IsTransferPause      bool      `gorm:"column:is_transfer_pause"`
	IsSlippageModifiable bool      `gorm:"column:slippage_modifiable"`
	IsBlackList          bool      `gorm:"column:is_black_list"`
	IsWhiteList          bool      `gorm:"column:is_white_list"`
	IsCoolDown           bool      `gorm:"column:is_cool_down"`
	BuyTax               float64   `gorm:"column:buy_tax"`
	SellTax              float64   `gorm:"column:sell_tax"`
	IsBuy                bool      `gorm:"column:is_buy"`
	IsSellAll            bool      `gorm:"column:is_sell_all"`
	TotalSupply          string    `gorm:"column:total_supply"`
	Score                int       `gorm:"column:score"`
	CheckTime            time.Time `gorm:"column:check_time"`
}
