package db

import (
	"bot/client/data"
	"bot/client/db/model"
	"bot/client/db/op"
	"gorm.io/gorm"
	"time"
)

type PushLog struct {
	Times int64 `json:"times"`
	Tweet int   `json:"tweet"`
}

func GetNewLatest(database *gorm.DB, token string) *PushLog {
	sql := "select times, tweet from news where token = ? and times = (select max(times) from news where token = ?)"
	var pushLog PushLog
	database.Raw(sql, token, token).Scan(&pushLog)
	return &pushLog
}

func SaveNew(database *gorm.DB, pool *data.NewPairData, twNum int, now time.Time, times int64) {
	op.InsetNew(database, model.New{
		Symbol:    pool.CoinSymbol,
		Token:     pool.CoinAddr,
		Pair:      pool.Pair,
		Liquidity: pool.StableReserve,
		Age:       pool.Age,
		Tweet:     twNum,
		PushTime:  now,
		Times:     times,
	})
}
