package core

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"security/data"
)

func GetTokenSecurityGoPlus(token string) *data.GoPlus {
	url := "https://api.gopluslabs.io/api/v1/token_security/1?contract_addresses=" + token
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return nil
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	var security data.GoPlus
	json.Unmarshal(body, &security)
	return &security
}

func GetTokenSecurityTokenSniffer(token string) *data.TokenSniffer {
	url := "https://tokensniffer.com/api/v2/tokens/1/" + token + "?apikey=9f798a02488992f4ca09f9f2580ae445d3a0fe7d&include_metrics=true&include_tests=true&block_until_ready=false"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("accept", "application/json")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	var result data.TokenSniffer
	json.Unmarshal(body, &result)
	return &result
}
