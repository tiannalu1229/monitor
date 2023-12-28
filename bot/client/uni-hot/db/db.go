package db

import (
	"bot/client/data"
	"bot/client/db/model"
	"bot/client/db/op"
	"gorm.io/gorm"
	"time"
)

type PushLog struct {
	PushTime time.Time `json:"push_time"`
	Times    int64     `json:"times"`
	Vol      float64   `json:"vol"`
	Tx       int       `json:"tx"`
	Trader   int       `json:"trader"`
	Tweet    int       `json:"tweet"`
}

func GetHotLatest(database *gorm.DB, token string, typ int64) *PushLog {
	sql := "select push_time, times, vol, tx, trader, tweet from hots where token = ? and times = (select max(times) from hots where token = ? and type = ?) and type = ?"
	var pushLog PushLog
	database.Raw(sql, token, token, typ, typ).Scan(&pushLog)
	return &pushLog
}

func SaveHot(database *gorm.DB, pool *data.HotPoolData, twNum int, now time.Time, times int64, typ int64, level string) {
	op.InsetHot(database, model.Hot{
		Symbol:    pool.CoinSymbol,
		Token:     pool.CoinAddr,
		Price:     pool.Price,
		Pair:      pool.Pair,
		Liquidity: pool.Liquidity,
		Age:       pool.Created,
		Trader:    pool.Traders,
		Tweet:     twNum,
		PushTime:  now,
		Vol:       pool.Vol,
		Tx:        pool.Txs,
		Times:     times,
		Type:      typ,
		Level:     level,
	})
}
