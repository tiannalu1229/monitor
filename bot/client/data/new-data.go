package data

type NewPair struct {
	Code int           `json:"code"`
	Data []NewPairData `json:"data"`
	Msg  string        `json:"msg"`
}
type NewPairData struct {
	Id              string  `json:"_id"`
	CoinAddr        string  `json:"coin_addr"`
	CoinDecimal     int     `json:"coin_decimal"`
	CoinSymbol      string  `json:"coin_symbol"`
	CoinTotalSupply int     `json:"coin_total_supply"`
	CreateBlock     int     `json:"create_block"`
	CreateTime      int     `json:"create_time"`
	CreateTx        string  `json:"create_tx"`
	Creator         string  `json:"creator"`
	Age             string  `json:"age"`
	Eid             string  `json:"eid"`
	Name            string  `json:"name"`
	Pair            string  `json:"pair"`
	Pindex          int     `json:"pindex"`
	StableAddr      string  `json:"stable_addr"`
	StableDecimal   int     `json:"stable_decimal"`
	StableIndex     int     `json:"stable_index"`
	StableSymbol    string  `json:"stable_symbol"`
	CoinReserve     float64 `json:"coin_reserve"`
	StableReserve   float64 `json:"stable_reserve"`
}
