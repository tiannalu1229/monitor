package op

import (
	"bot/client/db/model"
	"go-micro.dev/v4/logger"
	"gorm.io/gorm"
)

func InsetHot(db *gorm.DB, hot model.Hot) {
	if err := db.Model(model.Hot{}).Save(&hot).Error; err != nil {
		logger.Log(logger.ErrorLevel, "insertPair failed:%s", err.Error())
	}
}
