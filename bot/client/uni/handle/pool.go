package handle

import (
	"bot/client/uni/data"
)

var TokenNum map[string]int

func InitMap() {
	TokenNum = make(map[string]int)
}

func HandleNewPair(pairs []data.NewPairData) {
	for i := 0; i < len(pairs); i++ {
		//token := pairs[i].CoinAddr
		//pair := pairs[i].Pair
		//numToken := GetTweetNum(token)
		//numPair := GetTweetNum(pair)
		//num := numToken + numPair
		//fmt.Println("代币地址 : ", token, "推特提及次数 : ", num)
		//times, result := CheckNum(num, token)
		//if result {
		//	//发送消息
		//	fmt.Print("即将推送交易对 : ", pairs[i].Pair)
		//	send.Send(pairs[i], num, times)
		//}
	}
}

func CheckNum(num int, token string) (int, bool) {
	times := 0
	result := false
	if num >= 10 {
		result = true
		if TokenNum[token] < 3 {
			times = 1
		} else if TokenNum[token] < 6 {
			times = 2
		} else if TokenNum[token] < 10 {
			times = 3
		}
	} else if num >= 6 {
		if TokenNum[token] < 6 {
			result = true
			if TokenNum[token] < 3 {
				times = 1
			} else if TokenNum[token] < 6 {
				times = 2
			}
		}
	} else if num >= 3 {
		if TokenNum[token] < 3 {
			result = true
			times = 1
		}
	}
	TokenNum[token] = num

	return times, result
}
