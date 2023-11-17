package main

import (
	"bot/client/uni/handle"
	"fmt"
	"time"
)

func main() {
	for {
		fmt.Println("当前时间 : ", time.Now())
		//newPair := pool.GetNewPair()
		//fmt.Println("查询到交易对个数 : ", len(newPair))
		//pool.HandleNewPair(newPair)
		hotPool := handle.GetHotPool()
		fmt.Println("查询到热门交易对个数 : ", len(hotPool))
		handle.HotPoolHandle(hotPool)
		time.Sleep(time.Second * 30)
	}
}
