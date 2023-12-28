package data

type TweetOnlyCount struct {
	Code int    `json:"code"`
	Data int    `json:"data"`
	Msg  string `json:"msg"`
}

type TweetResult struct {
	Code int       `json:"code"`
	Data TweetData `json:"data"`
	Msg  string    `json:"msg"`
}
type TweetData struct {
	Count int `json:"count"`
	Elemt []struct {
		Id            int64  `json:"_id"`
		CreatedAt     string `json:"created_at"`
		CreatedTime   int    `json:"created_time"`
		FavoriteCount int    `json:"favorite_count"`
		IdStr         string `json:"id_str"`
		Lang          string `json:"lang"`
		Name          string `json:"name"`
		RetweetCount  int    `json:"retweet_count"`
		ScreenName    string `json:"screen_name"`
		Text          string `json:"text"`
	} `json:"elemt"`
}
