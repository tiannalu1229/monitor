package model

type PushResult struct {
	Symbol    string
	Token     string
	Vol       float32
	Trader    int64
	Tx        int64
	Tweet     int64
	Liquidity float32
	Age       string
	Times     int64
}

type PushDetailResult struct {
	Time      string
	Symbol    string
	Token     string
	Vol       float32
	Trader    int64
	Tx        int64
	Tweet     int64
	Liquidity float32
	Age       string
	Times     int64
	Level     string
	Price     float32
	Type      int64
}
