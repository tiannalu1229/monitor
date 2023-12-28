package core

import (
	bot "bot/proto"
	"context"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/shopspring/decimal"
	"go-micro.dev/v4/logger"
	"math"
	"math/big"
	"uni/conf"
	"uni/data"
	"uni/handler/contract/erc20"
	"uni/handler/contract/universal"
	"uni/handler/contract/util"
)

type Core struct {
	Client   *ethclient.Client
	Token    string
	Session  *universal.UniversalSession
	Commands []byte
	Bs       *bot.BotService
}

type PoolInfo struct {
	Liquidity            float64
	Vols                 float64
	Txs                  int64
	Traders              int64
	Age                  string
	IsOpenSource         bool
	IsProxy              bool
	IsMintAble           bool
	IsHoney              bool
	IsLock               bool
	IsOwnerShip          bool
	IsTransferPause      bool
	IsSlippageModifiable bool
	IsBlackList          bool
	IsWhiteList          bool
	IsCoolDown           bool
	BuyTax               float64
	SellTax              float64
	IsBuy                bool
	IsSellAll            bool
}

func GetValue(user *data.UserResult, liquidity float32) *big.Int {
	amount := liquidity * user.PoolPercent / 100
	if amount > user.MaxPerSwap {
		amount = user.MaxPerSwap
	}
	a := decimal.NewFromFloat32(amount * 1e18)
	return a.BigInt()
}

func (c *Core) MakeSession(user *data.UserResult, value *big.Int, gasPrice float64) error {
	client, err := ethclient.Dial(user.Rpc)
	if err != nil {
		logger.Log(logger.ErrorLevel, "eth client err: ", err)
		return err
	}
	uv, err := universal.NewUniversal(conf.Config.Param.Universal, client)
	if err != nil {
		logger.Log(logger.ErrorLevel, "create session err: ", err)
		return err
	}
	session := new(universal.UniversalSession)
	session.Contract = uv
	privateKey, err := crypto.HexToECDSA(user.PrivateKey)
	if err != nil {
		return err
	}
	transactOpts, err := bind.NewKeyedTransactorWithChainID(privateKey, new(big.Int).SetInt64(conf.Config.Param.ChainId))
	if err != nil {
		return err
	}
	transactOpts.From = common.HexToAddress(user.Address)
	transactOpts.NoSend = false
	transactOpts.Value = value

	if gasPrice != 0 {
		gasPrice = math.Ceil(gasPrice)
		useGas := int64(gasPrice) * (user.GasPrice + 100) / 100
		transactOpts.GasPrice = big.NewInt(useGas * 1e9)
	}
	session.TransactOpts = *transactOpts
	c.Session = session
	c.Client = client
	return nil
}

func (c *Core) CoreOperate(value *big.Int, slippage int64) (*types.Transaction, error) {
	inputs := c.makeTx(value, slippage)
	tx, err := c.Session.Execute(c.Commands, inputs)
	if err != nil {
		return nil, err
	} else {
		return tx, nil
	}
}
func (c *Core) makeTx(value *big.Int, slippage int64) (inputs [][]byte) {
	commands := c.Commands
	tokenBS := fmt.Sprintf("%064s", c.Token[len(c.Token)-40:])
	amountBS := fmt.Sprintf("%064s", hex.EncodeToString(value.Bytes()))
	if hex.EncodeToString(commands) == "0b08" {
		wrap := "0000000000000000000000000000000000000000000000000000000000000002" +
			amountBS
		wrapInput, _ := hex.DecodeString(wrap)
		inputs = append(inputs, wrapInput)

		amountOut := util.GetAmountOut(c.Client, value, conf.Config.Param.EthAddress, common.HexToAddress(c.Token))
		amountOut.Mod(amountOut, big.NewInt(100-slippage))
		amountOut.Div(amountOut, big.NewInt(100))
		amountOutBS := fmt.Sprintf("%064s", hex.EncodeToString(amountOut.Bytes()))

		swap := "0000000000000000000000000000000000000000000000000000000000000001" +
			amountBS +
			amountOutBS +
			"00000000000000000000000000000000000000000000000000000000000000a0" +
			"0000000000000000000000000000000000000000000000000000000000000000" +
			"0000000000000000000000000000000000000000000000000000000000000002" +
			conf.Config.Param.EthByte +
			tokenBS

		swapInput, _ := hex.DecodeString(swap)
		inputs = append(inputs, swapInput)
	}
	if hex.EncodeToString(commands) == "0b09" {
		wrap := "0000000000000000000000000000000000000000000000000000000000000002" +
			amountBS
		wrapInput, _ := hex.DecodeString(wrap)
		inputs = append(inputs, wrapInput)

		amountIn := util.GetAmountIn(c.Client, value, conf.Config.Param.EthAddress, common.HexToAddress(c.Token))
		amountIn.Mod(amountIn, big.NewInt(100))
		amountIn.Div(amountIn, big.NewInt(100-slippage))
		amountInBS := fmt.Sprintf("%064s", hex.EncodeToString(amountIn.Bytes()))

		swap := "0000000000000000000000000000000000000000000000000000000000000001" +
			amountInBS +
			amountBS +
			"00000000000000000000000000000000000000000000000000000000000000a0" +
			"0000000000000000000000000000000000000000000000000000000000000000" +
			"0000000000000000000000000000000000000000000000000000000000000002" +
			conf.Config.Param.EthByte +
			tokenBS

		swapInput, _ := hex.DecodeString(swap)
		inputs = append(inputs, swapInput)
	}
	if hex.EncodeToString(commands) == "080c" {
		amountOut := util.GetAmountOut(c.Client, value, common.HexToAddress(c.Token), conf.Config.Param.EthAddress)
		amountOut.Mod(amountOut, big.NewInt(100-slippage))
		amountOut.Div(amountOut, big.NewInt(100))
		amountOutBS := fmt.Sprintf("%064s", hex.EncodeToString(amountOut.Bytes()))

		swap := "0000000000000000000000000000000000000000000000000000000000000001" +
			amountBS +
			amountOutBS +
			"00000000000000000000000000000000000000000000000000000000000000a0" +
			"0000000000000000000000000000000000000000000000000000000000000000" +
			"0000000000000000000000000000000000000000000000000000000000000002" +
			tokenBS +
			conf.Config.Param.EthByte

		swapInput, _ := hex.DecodeString(swap)
		inputs = append(inputs, swapInput)

		unwrap := "0000000000000000000000000000000000000000000000000000000000000002" +
			amountBS
		unwrapInput, _ := hex.DecodeString(unwrap)
		inputs = append(inputs, unwrapInput)
	}
	if hex.EncodeToString(commands) == "090c" {
		amountIn := util.GetAmountIn(c.Client, value, common.HexToAddress(c.Token), conf.Config.Param.EthAddress)
		amountIn.Mod(amountIn, big.NewInt(100))
		amountIn.Div(amountIn, big.NewInt(100-slippage))
		amountInBS := fmt.Sprintf("%064s", hex.EncodeToString(amountIn.Bytes()))

		swap := "0000000000000000000000000000000000000000000000000000000000000001" +
			amountInBS +
			amountBS +
			"00000000000000000000000000000000000000000000000000000000000000a0" +
			"0000000000000000000000000000000000000000000000000000000000000000" +
			"0000000000000000000000000000000000000000000000000000000000000002" +
			conf.Config.Param.EthByte +
			tokenBS

		swapInput, _ := hex.DecodeString(swap)
		inputs = append(inputs, swapInput)

		unwrap := "0000000000000000000000000000000000000000000000000000000000000002" +
			amountBS
		unwrapInput, _ := hex.DecodeString(unwrap)
		inputs = append(inputs, unwrapInput)
	}

	return inputs
}

func (c *Core) GetTxResult(tx *types.Transaction) (*types.Receipt, error) {
	receipt, err := bind.WaitMined(context.Background(), c.Client, tx)
	return receipt, err
}
func (c *Core) SwapAlert() {
	(*c.Bs).Send(context.Background(), &bot.SendRequest{
		Msg:    "sendData",
		Token:  conf.Config.Param.TxBot,
		ChatId: 0,
	})
}

func Approve(token string, user *data.UserResult, gasPrice float64) error {
	client, err := ethclient.Dial(user.Rpc)
	if err != nil {
		logger.Log(logger.ErrorLevel, "eth client err: ", err)
		return err
	}
	erc, err := erc20.NewErc20(common.HexToAddress(token), client)
	if err != nil {
		logger.Log(logger.ErrorLevel, "create session err: ", err)
		return err
	}
	session := new(erc20.Erc20Session)
	session.Contract = erc
	privateKey, err := crypto.HexToECDSA(user.PrivateKey)
	if err != nil {
		return err
	}
	transactOpts, err := bind.NewKeyedTransactorWithChainID(privateKey, new(big.Int).SetInt64(conf.Config.Param.ChainId))
	if err != nil {
		return err
	}
	transactOpts.From = common.HexToAddress(user.Address)
	transactOpts.NoSend = false
	transactOpts.Value = big.NewInt(0)

	if gasPrice != 0 {
		useGas := int64(gasPrice) * (user.GasPrice + 100) / 100
		transactOpts.GasPrice = big.NewInt(useGas * 1e9)
	}
	session.TransactOpts = *transactOpts

	maxUint256 := new(big.Int).Sub(new(big.Int).Lsh(big.NewInt(1), 256), big.NewInt(1))
	session.Approve(conf.Config.Param.Universal, maxUint256)
	return nil
}
