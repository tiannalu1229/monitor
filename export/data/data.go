package data

type Pair struct {
	Code int      `json:"code"`
	Data PairData `json:"data"`
	Msg  string   `json:"msg"`
}

type PairData struct {
	Id              string  `json:"_id"`
	Pindex          int     `json:"pindex"`
	Pair            string  `json:"pair"`
	Name            string  `json:"name"`
	CoinAddr        string  `json:"coin_addr"`
	CoinSymbol      string  `json:"coin_symbol"`
	CoinDecimal     int     `json:"coin_decimal"`
	CoinTotalSupply int64   `json:"coin_total_supply"`
	StableAddr      string  `json:"stable_addr"`
	StableSymbol    string  `json:"stable_symbol"`
	StableDecimal   int     `json:"stable_decimal"`
	StableIndex     int     `json:"stable_index"`
	CreateBlock     int     `json:"create_block"`
	CreateTime      int     `json:"create_time"`
	CreateTx        string  `json:"create_tx"`
	Creator         string  `json:"creator"`
	Eid             string  `json:"eid"`
	CoinReserve     float64 `json:"coin_reserve"`
	StableReserve   float64 `json:"stable_reserve"`
	Price           float64 `json:"price"`
	UpdateTime      int     `json:"update_time"`
}
