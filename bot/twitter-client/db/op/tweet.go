package op

import (
	"bot/twitter-client/db/model"
	"go-micro.dev/v4/logger"
	"gorm.io/gorm"
)

type Database struct {
	Db *gorm.DB
}

func (db *Database) GetTweetAnalyseLog(key string) *model.TweetAnalysePushLog {
	sql := "select ta.keyword, ta.times, tweet, push_time from tweet_analyse_push_logs t left join (select keyword, max(times) times from tweet_analyse_push_logs group by keyword) ta on t.keyword = ta.keyword and t.times = ta.times where ta.keyword = ? and push_time < current_timestamp - interval '1day'"
	var pushResults model.TweetAnalysePushLog
	db.Db.Raw(sql, key).Scan(&pushResults)
	return &pushResults
}

func (db *Database) InsetTweetAnalyseLog(tweet *model.TweetAnalysePushLog) {
	if err := db.Db.Model(model.TweetAnalysePushLog{}).Save(&tweet).Error; err != nil {
		logger.Log(logger.ErrorLevel, "insertPair failed:%s", err.Error())
	}
}
