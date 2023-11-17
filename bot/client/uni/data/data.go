package data

type NewPair struct {
	Code int           `json:"code"`
	Data []NewPairData `json:"data"`
	Msg  string        `json:"msg"`
}
type NewPairData struct {
	Id            string  `json:"_id"`
	CoinAddr      string  `json:"coin_addr"`
	CoinDecimal   int     `json:"coin_decimal"`
	CoinSymbol    string  `json:"coin_symbol"`
	CreateBlock   int     `json:"create_block"`
	CreateTime    int     `json:"create_time"`
	CreateTx      string  `json:"create_tx"`
	Creator       string  `json:"creator"`
	Eid           string  `json:"eid"`
	Name          string  `json:"name"`
	Pair          string  `json:"pair"`
	Pindex        int     `json:"pindex"`
	StableAddr    string  `json:"stable_addr"`
	StableDecimal int     `json:"stable_decimal"`
	StableIndex   int     `json:"stable_index"`
	StableSymbol  string  `json:"stable_symbol"`
	CoinReserve   float64 `json:"coin_reserve"`
	StableReserve float64 `json:"stable_reserve"`
	Price         float64 `json:"price"`
	UpdateTime    int     `json:"update_time"`
}

type TweetResult struct {
	Code int       `json:"code"`
	Data TweetData `json:"data"`
	Msg  string    `json:"msg"`
}
type TweetData struct {
	Count int `json:"count"`
	Elemt []struct {
		Id            int64  `json:"_id"`
		CreatedAt     string `json:"created_at"`
		CreatedTime   int    `json:"created_time"`
		FavoriteCount int    `json:"favorite_count"`
		IdStr         string `json:"id_str"`
		Lang          string `json:"lang"`
		Name          string `json:"name"`
		RetweetCount  int    `json:"retweet_count"`
		ScreenName    string `json:"screen_name"`
		Text          string `json:"text"`
	} `json:"elemt"`
}

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
	Liquidaty  float64 `json:"liquidaty"`
	Created    string  `json:"created"`
	CreateTime int     `json:"create_time"`
	Vol        int     `json:"vol"`
	VolBuy     int     `json:"vol_buy"`
	VolSell    int     `json:"vol_sell"`
	Txs        int     `json:"txs"`
	TxsBuy     int     `json:"txs_buy"`
	TxsSell    int     `json:"txs_sell"`
}
