package core

import (
	"fmt"
	"github.com/ethereum/go-ethereum/core/types"
	"gorm.io/gorm"
	"math/big"
	"strconv"
	"strings"
	"time"
	"uni/data"
	"uni/db"
	"uni/db/model"
	pb "uni/proto"
)

type Database struct {
	Db *gorm.DB
}

func (d *Database) GetOnUser() []data.UserResult {
	sql := "select * from flash_users where turn = 1"
	var result []data.UserResult
	d.Db.Raw(sql).Scan(&result)
	return result
}

func (d *Database) GetFlashUser(address string) *data.UserResult {
	sql := "select * from flash_users where address = ?"
	var result data.UserResult
	d.Db.Raw(sql, address).Scan(&result)
	return &result
}

func (d *Database) GetFlashSwap(address string) []data.SwapResult {
	sql := "select * from flash_swaps where user = ?"
	var result []data.SwapResult
	d.Db.Raw(sql, address).Scan(&result)
	return result
}

// GetBuyUser 单个池子可自动购买账户/**
func (d *Database) GetBuyUser(pool *pb.Pool) []data.UserResult {
	sql := getAutoBuyUserSql(pool)
	var result []data.UserResult
	d.Db.Raw(sql).Scan(&result)
	return result
}

func getAutoBuyUserSql(pool *pb.Pool) string {

	age := "0"
	if strings.Contains(pool.Age, "d") {
		ageArray := strings.Split(pool.Age, "d")
		age = ageArray[0]
	}

	sql := "select * from flash_users where turn = 1"
	sql += " and liquidity <= " + fmt.Sprintf("%f", pool.Liquidity) +
		" and vols <= " + fmt.Sprintf("%f", pool.Vols) +
		" and txs <= " + strconv.Itoa(int(pool.Txs)) +
		" and traders <= " + strconv.Itoa(int(pool.Traders)) +
		" and age <= " + age +
		" and twitter <= " + strconv.Itoa(int(pool.Twitter))

	fmt.Println(sql)
	return sql
}

func (d *Database) SaveSwap(pool *pb.Pool, user string, value *big.Int, r *types.Receipt) {

	//r.Logs
	var amount int64
	for i := 0; i < len(r.Logs); i++ {
		log := r.Logs[i]
		for j := 0; j < len(log.Topics); j++ {
			topic := log.Topics[j]
			if topic.Hex() == "0xd78ad95fa46c994b6551d0da85fc275fe613ce37657fb8d5e3d130840159d822" {
				data := log.Data
				amountOutString := string(data)[66+64 : 66+64+64]
				amount, _ = strconv.ParseInt(amountOutString, 16, 64)
			}
		}
	}
	swap := model.FlashSwap{
		Token:     pool.Token,
		User:      user,
		Cost:      value.String(),
		BuyAmount: strconv.Itoa(int(amount)),
		NowAmount: strconv.Itoa(int(amount)),
		Price:     fmt.Sprintf("%f", pool.Price),
		IsSell:    0,
		BuyTime:   time.Now(),
		Hash:      r.TxHash.Hex(),
		Gas:       strconv.Itoa(int(r.GasUsed)),
	}

	db.InsertSwap(d.Db, swap)
}
