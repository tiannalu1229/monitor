package tool

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

type Gas struct {
	Code string    `json:"code"`
	Msg  string    `json:"msg"`
	Data []GasData `json:"data"`
}

type GasData struct {
	ChainFullName         string `json:"chainFullName"`
	ChainShortName        string `json:"chainShortName"`
	Symbol                string `json:"symbol"`
	BestTransactionFee    string `json:"bestTransactionFee"`
	BestTransactionFeeSat string `json:"bestTransactionFeeSat"`
	RecommendedGasPrice   string `json:"recommendedGasPrice"`
	RapidGasPrice         string `json:"rapidGasPrice"`
	StandardGasPrice      string `json:"standardGasPrice"`
	SlowGasPrice          string `json:"slowGasPrice"`
}

func GetGas() (float64, error) {
	url := "https://www.oklink.com/api/v5/explorer/blockchain/fee?chainShortName=ETH"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	req.Header.Add("Ok-Access-Key", "5c62e866-16b0-4e43-95b4-a6cb27ae97f8")
	req.Header.Add("Cookie", "aliyungf_tc=cc484f41569dc279b3082bbb08e84112848bc952be90d487ebdce410c957c3fc")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	var gasResult Gas
	json.Unmarshal(body, &gasResult)
	if len(gasResult.Data) > 0 {
		gas, _ := strconv.ParseFloat(gasResult.Data[0].RecommendedGasPrice, 64)
		return gas, nil
	} else {
		return 0, nil
	}
}
