package data

import (
	"time"
	"uni/db/model"
)

type SecurityResult struct {
	model.Model
	Token                string
	IsOpenSource         bool
	HiddenOwner          bool
	IsProxy              bool
	IsMintAble           bool
	IsHoney              bool
	IsLock               bool
	IsOwnerShip          bool
	IsTransferPause      bool
	IsSlippageModifiable bool
	IsBlackList          bool
	IsWhiteList          bool
	IsCoolDown           bool
	BuyTax               float64
	SellTax              float64
	IsBuy                bool
	IsSellAll            bool
	TotalSupply          string
	Score                int
	CheckTime            time.Time
}
