package core

import (
	"bot/client/conf"
	"bot/client/db/model"
	utils "bot/client/db/tool"
	hot "bot/client/uni-hot/handle"
	new "bot/client/uni-new/handle"
	bot "bot/proto"
	"fmt"
	"go-micro.dev/v4"
	"go-micro.dev/v4/logger"
	"gorm.io/gorm"
	security "security/proto"
	"time"
	uni "uni/proto"
)

type task struct {
	conf *conf.TomlConfig         `json:"conf"`
	db   *gorm.DB                 `json:"db"`
	bs   bot.BotService           `json:"bs"`
	ss   security.SecurityService `json:"ss"`
	us   uni.UniService           `json:"us"`
}

func Start() {
	t, err := newTask()
	if err != nil {
		logger.Log(logger.ErrorLevel, "start err: ", err)
	}

	t.initTable()
	for {
		t.checkNew()
		for i := 0; i < len(conf.Config.Ctl.Level); i++ {
			t.checkHot(conf.Config.Ctl.Level[i])
		}
		time.Sleep(time.Second * 30)
	}
}

func newTask() (*task, error) {
	initMap()
	db, err := utils.ConnectPg(conf.Config.Db.Host, conf.Config.Db.User, conf.Config.Db.Password, conf.Config.Db.DbName, uint32(conf.Config.Db.Port))
	if err != nil {
		logger.Log(logger.ErrorLevel, "db conn err: ", err)
		return nil, err
	}

	srv := micro.NewService()
	srv.Init()
	bs := bot.NewBotService("msg-bot", srv.Client())
	ss := security.NewSecurityService("security", srv.Client())
	us := uni.NewUniService("flashswap", srv.Client())

	return &task{
		conf: &conf.Config,
		db:   db,
		bs:   bs,
		ss:   ss,
		us:   us,
	}, nil
}

func initMap() {
	hot.InitHotPoolMap()
	new.InitTokenNumMap()
}

func (t *task) initTable() {
	err := t.db.AutoMigrate(
		&model.Hot{},
		&model.New{},
	)

	if err != nil {
		panic(fmt.Sprintf("can not init table ,error:%s\n", err.Error()))
		return
	}
	logger.Log(logger.InfoLevel, "init table finish")
}

func (t *task) checkHot(level string) {
	hotPool := hot.GetHotPool(level)
	logger.Log(logger.InfoLevel, "检测到代币个数: ", len(hotPool))
	srv := hot.Srv{
		Bs: t.bs,
		Ss: t.ss,
		Us: t.us,
	}
	srv.HotPoolHandle(hotPool, t.db, level)
}

func (t *task) checkNew() {
	newPair := new.GetNewPair(t.conf)
	new.NewPairHandle(newPair, t.db, t.conf)
}
