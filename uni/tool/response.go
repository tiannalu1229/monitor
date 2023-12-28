package tool

import (
	"google.golang.org/protobuf/types/known/timestamppb"
	"uni/data"
	uni "uni/proto"
)

func UserResultToUser(user *data.UserResult) *uni.User {
	u := uni.User{
		Address:              user.Address,
		Turn:                 user.Turn,
		PrivateKey:           user.PrivateKey,
		MaxPerSwap:           user.MaxPerSwap,
		PoolPercent:          user.PoolPercent,
		Liquidity:            user.Liquidity,
		Vols:                 user.Vols,
		Txs:                  user.Txs,
		Traders:              user.Traders,
		GasPrice:             user.GasPrice,
		Rpc:                  user.Rpc,
		Receipt:              user.Receipt,
		Age:                  user.Age,
		Level:                user.Level,
		IsTransferPause:      user.IsTransferPause,
		IsSlippageModifiable: user.IsSlippageModifiable,
		IsBlackList:          user.IsBlackList,
		IsWhiteList:          user.IsWhiteList,
		IsCoolDown:           user.IsCoolDown,
		BuyTax:               user.BuyTax,
		SellTax:              user.SellTax,
		IsBuy:                user.IsBuy,
		IsSellAll:            user.IsSellAll,
		Slippage:             user.Slippage,
		HiddenOwner:          user.HiddenOwner,
		Twitter:              user.Twitter,
	}

	return &u
}

func UserResultToSwap(swaps []data.SwapResult) []*uni.Swap {

	var us []*uni.Swap
	for i := 0; i < len(swaps); i++ {
		swap := swaps[i]
		u := uni.Swap{
			Token:   swap.Token,
			Symbol:  swap.Symbol,
			Cost:    swap.Cost,
			Buy:     swap.Buy,
			Now:     swap.Now,
			Price:   swap.Price,
			BuyTime: timestamppb.New(swap.BuyTime),
			Hash:    swap.Hash,
			IsSell:  swap.IsSell,
		}
		us = append(us, &u)
	}

	return us
}
