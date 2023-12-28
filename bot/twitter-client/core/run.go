package core

import (
	bot "bot/proto"
	"bot/twitter-client/conf"
	"bot/twitter-client/core/handle"
	"bot/twitter-client/db/model"
	utils "bot/twitter-client/db/tool"
	"fmt"
	"go-micro.dev/v4"
	"go-micro.dev/v4/logger"
	"gorm.io/gorm"
	"time"
)

type task struct {
	conf *conf.TomlConfig `json:"conf"`
	db   *gorm.DB         `json:"db"`
	bs   bot.BotService   `json:"bs"`
}

func Start() {
	t, err := newTask()
	if err != nil {
		logger.Log(logger.ErrorLevel, "start err: ", err)
	}

	t.initTable()
	for {
		t.twitterNew()
		time.Sleep(time.Second * 600)
	}
}

func newTask() (*task, error) {
	db, err := utils.ConnectPg(conf.Config.Db.Host, conf.Config.Db.User, conf.Config.Db.Password, conf.Config.Db.DbName, uint32(conf.Config.Db.Port))
	if err != nil {
		logger.Log(logger.ErrorLevel, "db conn err: ", err)
		return nil, err
	}

	srv := micro.NewService()
	srv.Init()
	bs := bot.NewBotService("msg-bot", srv.Client())

	return &task{
		conf: &conf.Config,
		db:   db,
		bs:   bs,
	}, nil
}

func (t *task) initTable() {
	err := t.db.AutoMigrate(
		model.TweetAnalysePushLog{},
	)

	if err != nil {
		panic(fmt.Sprintf("can not init table ,error:%s\n", err.Error()))
		return
	}
	logger.Log(logger.InfoLevel, "init table finish")
}

func (t *task) twitterNew() {
	th := handle.TwitterHandle{
		Bs:    &t.bs,
		Param: &t.conf.Param,
		Db:    t.db,
	}
	th.TwitterNewHandle()
}
