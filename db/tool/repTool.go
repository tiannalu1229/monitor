package utils

import (
	"db/model"
	pb "db/proto"
)

func PushResultToPushData(result model.PushResult) *pb.PushData {
	var pushData pb.PushData
	pushData.Vol = result.Vol
	pushData.Tweet = result.Tweet
	pushData.Tx = result.Tx
	pushData.Token = result.Token
	pushData.Age = result.Age
	pushData.Times = result.Times
	pushData.Symbol = result.Symbol
	pushData.Liquidity = result.Liquidity
	pushData.Trader = result.Trader

	return &pushData
}

func PushDetailResultToPushDetailData(result model.PushDetailResult) *pb.PushDetailData {
	var pushDetailData pb.PushDetailData
	pushDetailData.Time = result.Time
	pushDetailData.Level = result.Level
	pushDetailData.Vol = result.Vol
	pushDetailData.Trader = result.Trader
	pushDetailData.Tx = result.Tx
	pushDetailData.Tweet = result.Tweet
	pushDetailData.Liquidity = result.Liquidity
	pushDetailData.Price = result.Price
	pushDetailData.Age = result.Age
	pushDetailData.Times = result.Times
	pushDetailData.Type = result.Type

	return &pushDetailData
}
