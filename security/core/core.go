package core

import (
	"gorm.io/gorm"
	"security/data"
	"security/db/model"
	"security/db/op"
	"strconv"
	"strings"
	"time"
)

type Database struct {
	Db *gorm.DB
}

func (db *Database) GetSecurity(token string) *data.SecurityResult {
	// and check_time > current_timestamp - interval '5 minutes'
	sql := "select * from securities where token = ?"
	var result data.SecurityResult
	db.Db.Raw(sql, token).Scan(&result)
	return &result
}

func (db *Database) SaveSecurity(g map[string]data.GoPlusResult, id uint) *model.Security {
	var s model.Security
	for key, result := range g {
		s.ID = id
		s.Token = key
		s.IsOpenSource = stringToBool(result.IsOpenSource)
		s.IsProxy = stringToBool(result.IsProxy)
		s.IsMintAble = stringToBool(result.IsMintable)
		s.IsHoney = stringToBool(result.IsHoneypot)
		s.IsLock = judgeLock(result.LpHolders)
		//true: 没丢权限 false: 丢权限
		s.IsOwnerShip = judgeOwnerShip(result.OwnerAddress)
		s.HiddenOwner = stringToBool(result.HiddenOwner)
		s.IsTransferPause = stringToBool(result.TransferPausable)
		s.IsSlippageModifiable = stringToBool(result.SlippageModifiable)
		s.IsBlackList = stringToBool(result.IsBlacklisted)
		s.IsWhiteList = stringToBool(result.IsWhitelisted)
		s.IsCoolDown = stringToBool(result.TradingCooldown)
		buyTax, _ := strconv.ParseFloat(result.BuyTax, 64)
		s.BuyTax = buyTax * 100
		sellTax, _ := strconv.ParseFloat(result.SellTax, 64)
		s.SellTax = sellTax * 100
		s.IsBuy = !stringToBool(result.CannotBuy)
		s.IsSellAll = !stringToBool(result.CannotSellAll)
		s.TotalSupply = result.TotalSupply
		s.CheckTime = time.Now()

		score := 100
		if s.IsOwnerShip || !s.IsLock || s.IsHoney || s.IsMintAble || s.BuyTax > 50 || s.SellTax > 50 {
			score = score - 100
		} else if s.IsSlippageModifiable {
			score -= 10
		} else if s.IsBlackList || s.IsWhiteList {
			score -= 10
		} else if s.IsTransferPause {
			score -= 10
		} else if s.IsCoolDown {
			score -= 10
		} else {
			score = 100
		}

		if score > 50 {
			tsSecurity := GetTokenSecurityTokenSniffer(key)
			s.Score = tsSecurity.Score
		}
	}

	op.InsetSecurity(db.Db, s)

	return &s
}

func stringToBool(result string) bool {
	if result == "1" {
		return true
	} else {
		return false
	}
}

func judgeLock(holders []data.GoPlusLpHolder) bool {
	var lockLpPercent float64
	for i := 0; i < len(holders); i++ {
		holder := holders[i]
		if holder.IsLocked == 1 {
			percent, _ := strconv.ParseFloat(holder.Percent, 64)
			lockLpPercent += percent
		}
	}
	return lockLpPercent > 0.5
}

func judgeOwnerShip(owner string) bool {
	owner = strings.ToLower(owner)
	return owner != "0x0000000000000000000000000000000000000000" && owner != "0x000000000000000000000000000000000000dead"
}
