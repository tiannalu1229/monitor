package result

type TweetAnalyse struct {
	Code int                           `json:"code"`
	Data map[string][]TweetAnalyseData `json:"data"`
	Msg  string                        `json:"msg"`
}

type TweetAnalyseData struct {
	Key   string `json:"key"`
	Count int    `json:"count"`
}

type TweetListNew struct {
	Topic string `json:"topic"`
	Event string `json:"event"`
	Data  struct {
		Id         string        `json:"_id"`
		ListId     string        `json:"list_id"`
		TweetId    string        `json:"tweet_id"`
		Ts         int           `json:"ts"`
		Bookmarks  int           `json:"bookmarks"`
		CreatedAt  int           `json:"created_at"`
		Favorites  int           `json:"favorites"`
		Text       string        `json:"text"`
		Lang       string        `json:"lang"`
		Views      string        `json:"views"`
		ScreenName string        `json:"screen_name"`
		Quotes     int           `json:"quotes"`
		Replies    int           `json:"replies"`
		Retweets   int           `json:"retweets"`
		Media      []interface{} `json:"media"`
		Retweeted  interface{}   `json:"retweeted"`
	} `json:"data"`
}
