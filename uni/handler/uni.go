package handler

import (
	bot "bot/proto"
	"context"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-redis/redis/v8"
	"github.com/shopspring/decimal"
	"go-micro.dev/v4"
	"go-micro.dev/v4/logger"
	"go-micro.dev/v4/metadata"
	"gorm.io/gorm"
	"io"
	"log"
	"math/big"
	"net"
	"strconv"
	"strings"
	"sync"
	"time"
	"uni/conf"
	"uni/core"
	"uni/db"
	"uni/db/model"
	"uni/handler/contract/universal"
	"uni/handler/contract/util"
	pb "uni/proto"
	"uni/tool"
)

const (
	E18              = 1e18
	SwapToETHInV2    = "080c"
	SwapToETHOutV2   = "090c"
	SwapToTokenInV2  = "0b08"
	SwapToTokenOutV2 = "0b09"
)

type Uni struct {
	Db  *gorm.DB
	Rdb *redis.Client
}

func (e *Uni) CheckSign(ctx context.Context, request *pb.CheckSignRequest, response *pb.CheckSignResponse) error {
	md, ok := metadata.FromContext(ctx)
	if ok {
		if addr, ok := md["Remote"]; ok {
			host, _, _ := net.SplitHostPort(addr)
			logger.Log(logger.InfoLevel, "接收到: ", host, "请求接口CheckSign")
		}
	}
	sign, err := e.Rdb.Get(ctx, strings.ToLower(request.Address)+"-sign").Result()
	if errors.Is(err, redis.Nil) {
		// 缓存未命中，使用默认值
		sign = ""
		response.Result = false
	} else if err != nil {
		// 发生其他错误
		sign = ""
		response.Result = false
	}

	if sign != "" {
		response.Result = true
	} else {
		response.Result = false
	}

	return nil
}

func (e *Uni) Sign(ctx context.Context, request *pb.SignRequest, response *pb.SignResponse) error {

	md, ok := metadata.FromContext(ctx)
	if ok {
		if addr, ok := md["Remote"]; ok {
			host, _, _ := net.SplitHostPort(addr)
			logger.Log(logger.InfoLevel, "接收到: ", host, "请求接口Sign")
		}
	}
	message := request.Message

	prefix := []byte("\x19Ethereum Signed Message:\n")

	data := append(prefix, append([]byte(fmt.Sprintf("%d", len(message))), []byte(message)...)...)
	hash := crypto.Keccak256Hash(data)

	signatureHex := request.Sign

	signature := common.Hex2Bytes(signatureHex[2:])

	sigWithoutRecoveryID := signature[:len(signature)-1]

	recoveryID := signature[len(signature)-1] - 27

	pubKey, err := crypto.SigToPub(hash.Bytes(), append(sigWithoutRecoveryID, recoveryID))
	if err != nil {
		log.Fatalf("Could not recover public key: %v", err)
		return err
	}

	recoveredAddr := crypto.PubkeyToAddress(*pubKey)

	response.Result = recoveredAddr.Hex() == request.Address
	e.Rdb.Set(ctx, strings.ToLower(request.Address)+"-sign", request.Sign, time.Hour)
	return nil
}

func (e *Uni) Approve(ctx context.Context, request *pb.ApproveRequest, response *pb.ApproveResponse) error {
	md, ok := metadata.FromContext(ctx)
	if ok {
		if addr, ok := md["Remote"]; ok {
			host, _, _ := net.SplitHostPort(addr)
			logger.Log(logger.InfoLevel, "接收到: ", host, "请求接口Approve")
		}
	}

	srq := pb.SignRequest{
		Sign:    request.Sign,
		Message: request.Message,
		Address: request.Address,
	}
	srp := pb.SignResponse{}
	err := e.Sign(ctx, &srq, &srp)

	if err != nil {
		logger.Log(logger.ErrorLevel, "sign err: ", err)
		response.Msg = "701"
		return err
	}

	cdb := core.Database{
		Db: e.Db,
	}

	token := request.Token
	address := strings.ToLower(request.Address)
	userResult := cdb.GetFlashUser(address)
	gasPrice, err := tool.GetGas()
	if err != nil {
		logger.Log(logger.ErrorLevel, "gas获取失败: ", err)
	}
	core.Approve(token, userResult, gasPrice)
	return nil
}

func (e *Uni) GetFlashUser(ctx context.Context, request *pb.GetFlashUserRequest, response *pb.GetFlashUserResponse) error {
	md, ok := metadata.FromContext(ctx)
	if ok {
		if addr, ok := md["Remote"]; ok {
			host, _, _ := net.SplitHostPort(addr)
			logger.Log(logger.InfoLevel, "接收到: ", host, "请求接口GetFlashUser")
		}
	}
	srq := pb.SignRequest{
		Sign:    request.Sign,
		Message: request.Message,
		Address: request.Address,
	}
	srp := pb.SignResponse{}
	err := e.Sign(ctx, &srq, &srp)

	if err != nil {
		logger.Log(logger.ErrorLevel, "sign err: ", err)
		response.Msg = "701"
		return err
	}
	cdb := core.Database{
		Db: e.Db,
	}

	address := strings.ToLower(request.Address)
	userResult := cdb.GetFlashUser(address)
	response.User = tool.UserResultToUser(userResult)
	return nil
}

func (e *Uni) SaveFlashUser(ctx context.Context, request *pb.SaveFlashUserRequest, response *pb.SaveFlashUserResponse) error {
	md, ok := metadata.FromContext(ctx)
	if ok {
		if addr, ok := md["Remote"]; ok {
			host, _, _ := net.SplitHostPort(addr)
			logger.Log(logger.InfoLevel, "接收到: ", host, "请求接口SaveFlashUser")
		}
	}
	srq := pb.SignRequest{
		Sign:    request.Sign,
		Message: request.Message,
		Address: request.Address,
	}
	srp := pb.SignResponse{}
	err := e.Sign(ctx, &srq, &srp)
	if err != nil {
		logger.Log(logger.ErrorLevel, "sign err: ", err)
		return err
	}

	if !srp.Result {
		logger.Log(logger.InfoLevel, "verify failed")
		return nil
	}

	user := request.User
	cdb := core.Database{
		Db: e.Db,
	}

	address := strings.ToLower(user.Address)
	receipt := strings.ToLower(user.Receipt)

	modelUser := model.FlashUser{
		Address:              address,
		Turn:                 user.Turn,
		PrivateKey:           user.PrivateKey,
		MaxPerSwap:           float64(user.MaxPerSwap),
		PoolPercent:          float64(user.PoolPercent),
		Liquidity:            float64(user.Liquidity),
		Vols:                 float64(user.Vols),
		Txs:                  user.Txs,
		Traders:              user.Traders,
		GasPrice:             user.GasPrice,
		Rpc:                  user.Rpc,
		Receipt:              receipt,
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

	userResult := cdb.GetFlashUser(address)
	if userResult.Model.ID > 0 {
		modelUser.ID = userResult.Model.ID
		db.ModifyUser(e.Db, modelUser)
	} else {
		db.InsertUser(e.Db, modelUser)
	}

	response.Msg = "1"
	return nil
}

func (e *Uni) GetFlashSwap(ctx context.Context, request *pb.GetFlashSwapRequest, response *pb.GetFlashSwapResponse) error {
	md, ok := metadata.FromContext(ctx)
	if ok {
		if addr, ok := md["Remote"]; ok {
			host, _, _ := net.SplitHostPort(addr)
			logger.Log(logger.InfoLevel, "接收到: ", host, "请求接口GetFlashSwap")
		}
	}
	srq := pb.SignRequest{
		Sign:    request.Sign,
		Message: request.Message,
		Address: request.Address,
	}
	srp := pb.SignResponse{}
	err := e.Sign(ctx, &srq, &srp)

	if err != nil {
		logger.Log(logger.ErrorLevel, "sign err: ", err)
		response.Msg = "701"
		return err
	}
	cdb := core.Database{
		Db: e.Db,
	}

	address := strings.ToLower(request.Address)
	swapResults := cdb.GetFlashSwap(address)
	response.Swap = tool.UserResultToSwap(swapResults)
	return nil
}

func (e *Uni) Flash(ctx context.Context, request *pb.FlashRequest, response *pb.FlashResponse) error {
	md, ok := metadata.FromContext(ctx)
	if ok {
		if addr, ok := md["Remote"]; ok {
			host, _, _ := net.SplitHostPort(addr)
			logger.Log(logger.InfoLevel, "接收到: ", host, "请求接口Flash")
		}
	}
	srq := pb.SignRequest{
		Sign:    request.Sign,
		Message: request.Message,
		Address: request.Address,
	}
	srp := pb.SignResponse{}
	err := e.Sign(ctx, &srq, &srp)

	if err != nil {
		logger.Log(logger.ErrorLevel, "sign err: ", err)
		response.Msg = "701"
		return err
	}

	token := request.Token
	amount := request.Amount
	a := decimal.NewFromFloat32(amount * E18)
	amountInt := a.BigInt()

	command := request.Command
	commands, _ := hex.DecodeString(command)

	service := micro.NewService()
	service.Init()
	bs := bot.NewBotService("msg-bot", service.Client())

	c := core.Core{
		Token:    token,
		Commands: commands,
		Bs:       &bs,
	}

	cdb := core.Database{
		Db: e.Db,
	}
	user := cdb.GetFlashUser(request.Address)
	var value *big.Int
	if strings.Contains(request.Command, "0c") {
		value = big.NewInt(0)
	} else {
		value = amountInt
	}
	gasPrice, err := tool.GetGas()
	if err != nil {
		logger.Log(logger.ErrorLevel, "gas获取失败: ", err)
	}
	err = c.MakeSession(user, value, gasPrice)
	if err != nil {
		logger.Log(logger.ErrorLevel, "用户: ", user.Address, "交易封装失败: ", err)
		return err
	}
	tx, err := c.CoreOperate(amountInt, user.Slippage)
	if err != nil {
		logger.Log(logger.ErrorLevel, "用户: ", user.Address, "交易发送失败: ", err)
		return err
	} else {
		logger.Log(logger.InfoLevel, "用户: ", user.Address, "正在swap代币: ", c.Token, "\n购买hash: ", tx.Hash(), "\ngas: ", tx.GasPrice().String())
	}

	return nil
}

func (e *Uni) FlashBuy(ctx context.Context, request *pb.FlashBuyRequest, response *pb.FlashBuyResponse) error {

	token := request.Token
	tokenBS := fmt.Sprintf("%064s", token[len(token)-40:])

	amount := request.Amount
	a := decimal.NewFromFloat32(amount * E18)
	amountInt := a.BigInt()
	amountBS := fmt.Sprintf("%064s", hex.EncodeToString(amountInt.Bytes()))

	var commands []byte
	commands, _ = hex.DecodeString("0b08")

	ethClient, _ := ethclient.Dial(conf.Config.Param.HttpClient)
	uv, _ := universal.NewUniversal(conf.Config.Param.Universal, ethClient)
	session := new(universal.UniversalSession)
	session.Contract = uv
	privateKey, _ := crypto.HexToECDSA(conf.Config.Param.PrivateKey)
	transactOpts, err := bind.NewKeyedTransactorWithChainID(privateKey, new(big.Int).SetInt64(conf.Config.Param.ChainId))
	transactOpts.NoSend = false
	transactOpts.Value = amountInt

	gasPrice, err := tool.GetGas()
	if gasPrice != 0 {
		useGas := int64(gasPrice)
		transactOpts.GasPrice = big.NewInt(useGas * 1e9)
	}
	session.TransactOpts = *transactOpts

	service := micro.NewService()
	service.Init()
	bs := bot.NewBotService("msg-bot", service.Client())

	var inputs [][]byte
	var wrapInput []byte
	var swapInput []byte

	wrap := "0000000000000000000000000000000000000000000000000000000000000002" + amountBS
	wrapInput, _ = hex.DecodeString(wrap)
	inputs = append(inputs, wrapInput)

	amountOut := util.GetAmountOut(ethClient, amountInt, conf.Config.Param.EthAddress, common.HexToAddress(token))
	amountOutBS := fmt.Sprintf("%064s", hex.EncodeToString(amountOut.Bytes()))

	swap := "0000000000000000000000000000000000000000000000000000000000000001" +
		amountBS +
		amountOutBS +
		"00000000000000000000000000000000000000000000000000000000000000a0" +
		"0000000000000000000000000000000000000000000000000000000000000000" +
		"0000000000000000000000000000000000000000000000000000000000000002" +
		conf.Config.Param.EthByte +
		tokenBS

	swapInput, _ = hex.DecodeString(swap)
	inputs = append(inputs, swapInput)

	tx, err := session.Execute(commands, inputs)
	if err != nil {
		logger.Log(logger.ErrorLevel, "交易发送失败: ", err, "\n参数: \n", wrap, "\n", swap)
	}
	logger.Log(logger.InfoLevel, "正在购买代币: ", token, "\n购买hash: ", tx.Hash())

	msg := "购买代币: " + token +
		"\\n金额:" + fmt.Sprintf("%f", amount) +
		"\\n消耗gas:" + strconv.Itoa(int(tx.GasPrice().Int64())) + "*" + strconv.Itoa(int(tx.Gas())) +
		"\\n交易hash: " + tx.Hash().String()
	sendData := `{
				"msg_type": "text",
				"content": {"text": "` + msg + `"}
			}`

	bs.Send(context.Background(), &bot.SendRequest{
		Msg:    sendData,
		Token:  conf.Config.Param.TxBot,
		ChatId: 0,
	})
	response.Msg = "success"

	return nil
}

func (e *Uni) FlashSell(ctx context.Context, request *pb.FlashBuyRequest, response *pb.FlashBuyResponse) error {
	token := request.Token
	tokenBS := fmt.Sprintf("%064s", token[len(token)-40:])

	amount := request.Amount
	a := decimal.NewFromFloat32(amount * E18)
	amountInt := a.BigInt()
	amountBS := fmt.Sprintf("%064s", hex.EncodeToString(amountInt.Bytes()))

	var commands []byte
	commands, _ = hex.DecodeString("0b08")

	ethClient, _ := ethclient.Dial(conf.Config.Param.HttpClient)
	uv, _ := universal.NewUniversal(conf.Config.Param.Universal, ethClient)
	session := new(universal.UniversalSession)
	session.Contract = uv
	privateKey, _ := crypto.HexToECDSA(conf.Config.Param.PrivateKey)
	transactOpts, err := bind.NewKeyedTransactorWithChainID(privateKey, new(big.Int).SetInt64(conf.Config.Param.ChainId))
	transactOpts.From = conf.Config.Param.User
	transactOpts.NoSend = false
	transactOpts.Value = amountInt

	gasPrice, err := tool.GetGas()
	useGas := int64(gasPrice) + 5
	transactOpts.GasPrice = big.NewInt(useGas * 1e9)
	transactOpts.GasLimit = conf.Config.Param.GasLimit
	session.TransactOpts = *transactOpts

	var inputs [][]byte
	var wrapInput []byte
	var swapInput []byte

	wrap := "0000000000000000000000000000000000000000000000000000000000000002" + amountBS
	wrapInput, _ = hex.DecodeString(wrap)
	inputs = append(inputs, wrapInput)

	amountOut := util.GetAmountOut(ethClient, amountInt, conf.Config.Param.EthAddress, common.HexToAddress(token))

	fmt.Println(amountOut.String())
	amountOutBS := fmt.Sprintf("%064s", hex.EncodeToString(amountOut.Bytes()))

	swap := "0000000000000000000000000000000000000000000000000000000000000001" +
		amountBS +
		amountOutBS +
		"00000000000000000000000000000000000000000000000000000000000000a0" +
		"0000000000000000000000000000000000000000000000000000000000000000" +
		"0000000000000000000000000000000000000000000000000000000000000002" +
		tokenBS +
		conf.Config.Param.EthByte

	swapInput, _ = hex.DecodeString(swap)
	inputs = append(inputs, swapInput)

	tx, err := session.Execute(commands, inputs)
	if err != nil {
		logger.Log(logger.ErrorLevel, "交易发送失败: ", err, "\n参数: \n", wrap, "\n", swap)
	}
	logger.Log(logger.InfoLevel, "正在购买代币: ", token, "\n购买hash: ", tx.Hash())

	msg := "购买代币: " + token +
		"\\n金额:" + fmt.Sprintf("%f", amount) +
		"\\n交易hash: " + tx.Hash().String()
	sendData := `{
				"msg_type": "text",
				"content": {"text": "` + msg + `"}
			}`

	service := micro.NewService()
	service.Init()
	bs := bot.NewBotService("msg-bot", service.Client())
	bs.Send(context.Background(), &bot.SendRequest{
		Msg:    sendData,
		Token:  conf.Config.Param.TxBot,
		ChatId: 0,
	})

	return nil
}

func (e *Uni) FlashBuyAuto(ctx context.Context, request *pb.FlashBuyAutoRequest, response *pb.FlashBuyAutoResponse) error {
	//TODO 加密
	md, ok := metadata.FromContext(ctx)
	if ok {
		if addr, ok := md["Remote"]; ok {
			host, _, _ := net.SplitHostPort(addr)
			logger.Log(logger.InfoLevel, "接收到: ", host, "请求接口FlashBuyAuto")
		}
	}
	cdb := core.Database{
		Db: e.Db,
	}
	users := cdb.GetBuyUser(request.Pool)
	logger.Log(logger.InfoLevel, len(users), "用户需要购买")

	service := micro.NewService()
	service.Init()
	bs := bot.NewBotService("msg-bot", service.Client())
	var commands []byte
	//TODO
	commands, _ = hex.DecodeString("0b08")
	for i := 0; i < len(users); i++ {
		var wg sync.WaitGroup
		errChan := make(chan error, 1)
		wg.Add(1)
		go func(workerId int) error {
			user := users[workerId]
			c := core.Core{
				Token:    request.Pool.Token,
				Commands: commands,
				Bs:       &bs,
			}
			value := core.GetValue(&user, request.Pool.Liquidity)
			gasPrice, err := tool.GetGas()
			if err != nil {
				logger.Log(logger.ErrorLevel, "gas获取失败: ", err)
			}
			err = c.MakeSession(&user, value, gasPrice)
			if err != nil {
				logger.Log(logger.ErrorLevel, "用户: ", user.Address, "交易封装失败: ", err)
				return err
			}
			tx, err := c.CoreOperate(value, user.Slippage)
			if err != nil {
				logger.Log(logger.ErrorLevel, "用户: ", user.Address, "交易发送失败: ", err)
				return err
			} else {
				logger.Log(logger.InfoLevel, "用户: ", user.Address, "正在购买代币: ", c.Token, "\n购买hash: ", tx.Hash(), "\ngas: ", tx.GasPrice().String())
			}

			var message string
			r, err := c.GetTxResult(tx)
			if err != nil {
				message = "购买代币: " + request.Pool.Token + " 失败"
				logger.Log(logger.ErrorLevel, "用户: ", user.Address, "交易失败: ", err)
			} else {
				if r.Status == 0 {
					message = "购买代币: " + request.Pool.Token + " 失败" +
						"\\n交易hash: " + r.TxHash.Hex()
					logger.Log(logger.ErrorLevel, "用户: ", user.Address, "交易失败, hash: ", r.TxHash)
				} else if r.Status == 1 {
					message = "购买代币: " + request.Pool.Token + " 成功" +
						"\\n交易hash: " + r.TxHash.Hex() +
						"\\n购买花费: " + value.String() + "ETH" +
						"\\n花费gas: " + strconv.Itoa(int(r.GasUsed))
					cdb.SaveSwap(request.Pool, user.Address, value, r)
				}
			}
			sendData := `{
				"msg_type": "text",
				"content": {"text": "` + message + `"}
			}`

			if user.BotToken != "" {
				bs.Send(context.Background(), &bot.SendRequest{
					Msg:    sendData,
					Token:  user.BotToken,
					ChatId: user.ChatId,
				})
			}
			return nil
		}(i)
		close(errChan) // 关闭channel
		// 处理可能发生的错误
		for err := range errChan {
			if err != nil {
				logger.Log(logger.ErrorLevel, "交易失败: ", err)
			}
		}
	}

	return nil
}

func (e *Uni) Call(ctx context.Context, req *pb.CallRequest, rsp *pb.CallResponse) error {
	logger.Infof("Received Uni.Call request: %v", req)
	rsp.Msg = "Hello " + req.Name
	return nil
}

func (e *Uni) ClientStream(ctx context.Context, stream pb.Uni_ClientStreamStream) error {
	var count int64
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			logger.Infof("Got %v pings total", count)
			return stream.SendMsg(&pb.ClientStreamResponse{Count: count})
		}
		if err != nil {
			return err
		}
		logger.Infof("Got ping %v", req.Stroke)
		count++
	}
}

func (e *Uni) ServerStream(ctx context.Context, req *pb.ServerStreamRequest, stream pb.Uni_ServerStreamStream) error {
	logger.Infof("Received Uni.ServerStream request: %v", req)
	for i := 0; i < int(req.Count); i++ {
		logger.Infof("Sending %d", i)
		if err := stream.Send(&pb.ServerStreamResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
		time.Sleep(time.Millisecond * 250)
	}
	return nil
}

func (e *Uni) BidiStream(ctx context.Context, stream pb.Uni_BidiStreamStream) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		logger.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&pb.BidiStreamResponse{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
