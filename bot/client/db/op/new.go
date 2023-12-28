package op

import (
	"bot/client/db/model"
	"go-micro.dev/v4/logger"
	"gorm.io/gorm"
)

func InsetNew(db *gorm.DB, new model.New) {
	if err := db.Model(model.New{}).Save(&new).Error; err != nil {
		logger.Log(logger.ErrorLevel, "insertPair failed:%s", err.Error())
	}
}
