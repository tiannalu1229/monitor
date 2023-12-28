package util

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"uni/conf"
	"uni/handler/contract/uniswapv2"
)

func GetAmountIn(ethClient *ethclient.Client, amountOut *big.Int, tokenA common.Address, tokenB common.Address) *big.Int {
	u, _ := uniswapv2.NewUniswapv2(conf.Config.Param.UniRouter, ethClient)
	session := new(uniswapv2.Uniswapv2Session)
	session.Contract = u
	privateKey, _ := crypto.HexToECDSA(conf.Config.Param.PrivateKey)
	transactOpts, _ := bind.NewKeyedTransactorWithChainID(privateKey, new(big.Int).SetInt64(conf.Config.Param.ChainId))
	session.TransactOpts = *transactOpts

	var path []common.Address
	path = append(append(path, tokenA), tokenB)
	amountIns, err := session.GetAmountsIn(amountOut, path)
	if err != nil {
		fmt.Println(err)
	}

	return amountIns[0]
}

func GetAmountOut(ethClient *ethclient.Client, amountIn *big.Int, tokenA common.Address, tokenB common.Address) *big.Int {
	u, _ := uniswapv2.NewUniswapv2(conf.Config.Param.UniRouter, ethClient)
	session := new(uniswapv2.Uniswapv2Session)
	session.Contract = u
	privateKey, _ := crypto.HexToECDSA(conf.Config.Param.PrivateKey)
	transactOpts, _ := bind.NewKeyedTransactorWithChainID(privateKey, new(big.Int).SetInt64(conf.Config.Param.ChainId))
	session.TransactOpts = *transactOpts

	var path []common.Address
	path = append(append(path, tokenA), tokenB)
	amountOuts, err := session.GetAmountsOut(amountIn, path)
	if err != nil {
		fmt.Println(err)
	}

	return amountOuts[len(amountOuts)-1]
}
