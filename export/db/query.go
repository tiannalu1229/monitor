package db

import (
	"encoding/json"
	"export/conf"
	"export/data"
	"fmt"
	"github.com/tealeg/xlsx"
	"go-micro.dev/v4/logger"
	"gorm.io/gorm"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

const (
	RGB_Light_Green = "FFC6EFCE"
	RGB_Dark_Green  = "FF006100"
	RGB_Light_Red   = "FFFFC7CE"
	RGB_Dark_Red    = "FF9C0006"
	RGB_White       = "FFFFFFFF"
	RGB_Black       = "00000000"
)

type Query struct {
	DataBase *gorm.DB
}

type FirstDailyResult struct {
	Symbol    string
	Token     string
	Pair      string
	Vol       float64
	Tx        int
	Trader    int
	Liquidity float64
	Age       string
	Tweet     int
	Price     float64
	PushTime  time.Time
	Score     int
	Today     time.Time
}

func (q *Query) FirstDaily(score int64) string {

	logger.Log(logger.InfoLevel, "首次推送数据导出")
	sql := "select DISTINCT current_timestamp today, symbol, first.token, pair, vol, tx, trader, tweet, price, age, liquidity, push_time, score from (select symbol, ha.token, pair, vol, tx, trader, liquidity, age, tweet, price, ha.push_time from (select token, min(push_time) push_time from hots where times = 1 and push_time >= current_date - 2 and liquidity > 1 and tx > 10 and trader > 10 group by token) ha left join hots on ha.token = hots.token and ha.push_time = hots.push_time) first left join securities on first.token = securities.token where score >= ?"
	//sql := "select * from (select DISTINCT * from (select symbol, ha.token, pair, vol, tx, trader, liquidity, age, tweet, price, ha.push_time from (select token, min(push_time) push_time from hots where times = 1 group by token) ha left join hots on ha.token = hots.token and ha.push_time = hots.push_time) first left join securities on first.token = securities.token) list where score = 100"
	var results []FirstDailyResult
	q.DataBase.Raw(sql, score).Scan(&results)
	fmt.Println(results[0].Today)

	year, month, day := time.Now().Date()
	ys := strconv.Itoa(year)
	ms := month.String()
	ds := strconv.Itoa(day)
	file := xlsx.NewFile()
	sheet, _ := file.AddSheet("第一次推送统计")
	formHeadRow := sheet.AddRow()
	formHeadRow.AddCell().SetString("代币名称")
	formHeadRow.AddCell().SetString("代币地址")
	formHeadRow.AddCell().SetString("5m交易量")
	formHeadRow.AddCell().SetString("5m交易次数")
	formHeadRow.AddCell().SetString("5m交易人数")
	formHeadRow.AddCell().SetString("流动性")
	formHeadRow.AddCell().SetString("池子年龄")
	formHeadRow.AddCell().SetString("推特提及次数")
	formHeadRow.AddCell().SetString("推送时价格")
	formHeadRow.AddCell().SetString("推送时间")
	formHeadRow.AddCell().SetString("盈亏")
	formHeadRow.AddCell().SetString("安全评分")

	var win int
	var lose int
	var rug int
	var totalProfit float64
	for i := 0; i < len(results); i++ {
		result := results[i]

		row := sheet.AddRow()
		row.AddCell().SetString(result.Symbol)
		row.AddCell().SetString(result.Token)
		row.AddCell().SetFloat(result.Vol)
		row.AddCell().SetInt(result.Tx)
		row.AddCell().SetInt(result.Trader)
		row.AddCell().SetFloat(result.Liquidity)
		row.AddCell().SetString(result.Age)
		row.AddCell().SetInt(result.Tweet)
		row.AddCell().SetFloat(result.Price)
		row.AddCell().SetDateTime(result.PushTime)
		pairNow := GetPair(result.Pair)

		var profit float64
		cell := row.AddCell()
		style := xlsx.NewStyle()
		fmt.Println(pairNow.Pair)
		if pairNow.StableReserve < 0.1 {
			rug++
			style = SetStyle(0)
			profit = 0
		} else if pairNow.Price > result.Price {
			win++
			style = SetStyle(1)
			profit = (pairNow.Price - result.Price) / result.Price
		} else if pairNow.Price < result.Price {
			lose++
			style = SetStyle(2)
			profit = (pairNow.Price - result.Price) / result.Price
		}
		cell.SetStyle(style)
		cell.SetFloat(profit)
		totalProfit += profit

		row.AddCell().SetInt(result.Score)
	}
	rowLast := sheet.AddRow()
	rowLast.AddCell().SetString("win:")
	rowLast.AddCell().SetString(strconv.Itoa(win))
	rowLast.AddCell().SetString("lose:")
	rowLast.AddCell().SetString(strconv.Itoa(lose))
	rowLast.AddCell().SetString("rug:")
	rowLast.AddCell().SetString(strconv.Itoa(rug))
	rowLast.AddCell().SetString("total profit:")
	rowLast.AddCell().SetFloat(totalProfit)
	file.Save("./excel/" + ys + "-" + ms + "-" + ds + ".xlsx")
	logger.Log(logger.InfoLevel, ys+"-"+ms+"-"+ds+".xlsx"+"导出完毕!")

	return ys + "-" + ms + "-" + ds + ".xlsx"
}

func SetStyle(result int) *xlsx.Style {
	style := xlsx.NewStyle()
	if result == 0 {
		style.Fill = *xlsx.NewFill("solid", RGB_Light_Red, "00FF0000")
	}

	if result == 1 {
		style.Fill = *xlsx.NewFill("solid", RGB_Light_Green, "00FF0000")
	}

	if result == 2 {
		style.Fill = *xlsx.NewFill("solid", RGB_White, "00FF0000")
	}

	return style
}

func GetPair(token string) data.PairData {
	url := conf.Config.Uri.Pair + token
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return data.PairData{}
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return data.PairData{}
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return data.PairData{}
	}

	var pair data.Pair
	json.Unmarshal(body, &pair)

	return pair.Data
}
