package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"uni/tool"
)

func main() {
	//jsonString := "Token   string                 `protobuf:\"bytes,1,opt,name=token,proto3\" json:\"token,omitempty\"`\n\tSymbol  string                 `protobuf:\"bytes,2,opt,name=symbol,proto3\" json:\"symbol,omitempty\"`\n\tCost    string                 `protobuf:\"bytes,3,opt,name=cost,proto3\" json:\"cost,omitempty\"`\n\tBuy     string                 `protobuf:\"bytes,4,opt,name=buy,proto3\" json:\"buy,omitempty\"`\n\tNow     string                 `protobuf:\"bytes,5,opt,name=now,proto3\" json:\"now,omitempty\"`\n\tPrice   string                 `protobuf:\"bytes,6,opt,name=price,proto3\" json:\"price,omitempty\"`\n\tBuyTime *timestamppb.Timestamp `protobuf:\"bytes,7,opt,name=buyTime,proto3\" json:\"buyTime,omitempty\"`\n\tHash    string                 `protobuf:\"bytes,8,opt,name=hash,proto3\" json:\"hash,omitempty\"`"
	//jsonArray := strings.Split(jsonString, "\n")
	//for i := 0; i < len(jsonArray); i++ {
	//	p := strings.Split(jsonArray[i], "=")
	//	key := strings.Split(p[1], ",")
	//	fmt.Println("\"" + key[0] + "\":\"\",")
	//}
	//data := []byte("0x00000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000de0b6b3a764000000000000000000000000000000000000000000000000000005fc689f29612bfe0000000000000000000000000000000000000000000000000000000000000000")
	//result, _ := strconv.ParseInt(string(data)[66+64:66+64+64], 16, 64)
	//println(result)

	url := "https://www.oklink.com/api/v5/explorer/blockchain/fee?chainShortName=ETH"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("Ok-Access-Key", "5c62e866-16b0-4e43-95b4-a6cb27ae97f8")
	req.Header.Add("Cookie", "aliyungf_tc=cc484f41569dc279b3082bbb08e84112848bc952be90d487ebdce410c957c3fc")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}

	var gasResult tool.Gas
	json.Unmarshal(body, &gasResult)
	//gas, _ := strconv.ParseFloat(gasResult.Data[0].RecommendedGasPrice, 64)

	gas := 25.3
	fmt.Println(gasResult.Data[0].RecommendedGasPrice, int64(gas)*(5+100)/100)
}
