package data

type HotPool struct {
	Code int           `json:"code"`
	Data []HotPoolData `json:"data"`
	Msg  string        `json:"msg"`
}

type HotPoolData struct {
	Pair       string  `json:"pair"`
	CoinAddr   string  `json:"coin_addr"`
	CoinSymbol string  `json:"coin_symbol"`
	Price      float64 `json:"price"`
	Liquidity  float64 `json:"liquidity"`
	Created    string  `json:"created"`
	CreateTime int     `json:"create_time"`
	Vol        float64 `json:"vol"`
	VolBuy     float64 `json:"vol_buy"`
	VolSell    float64 `json:"vol_sell"`
	Txs        int     `json:"txs"`
	TxsBuy     int     `json:"txs_buy"`
	TxsSell    int     `json:"txs_sell"`
	Traders    int     `json:"traders"`
}
