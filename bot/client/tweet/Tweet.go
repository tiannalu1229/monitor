package tweet

import (
	"bot/client/data"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

/*
*
Tweet Num
*/

type Tweet struct {
	Url string `json:"url"`
}

func (t *Tweet) GetTweetNum(keyword string) int {
	url := t.Url
	url = strings.Replace(url, "keyword", keyword, -1)
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return 0
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return 0
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return 0
	}

	var tweetResult data.TweetOnlyCount
	json.Unmarshal(body, &tweetResult)
	return tweetResult.Data
}
