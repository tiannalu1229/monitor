package db

import (
	"gorm.io/gorm"
	"uni/db/model"
)

func InsertSwap(db *gorm.DB, swap model.FlashSwap) error {
	if err := db.Model(model.FlashSwap{}).Save(&swap).Error; err != nil {
		return err
	}
	return nil
}
func InsertUser(db *gorm.DB, user model.FlashUser) error {
	if err := db.Model(model.FlashUser{}).Save(&user).Error; err != nil {
		return err
	}
	return nil
}
func ModifyUser(db *gorm.DB, user model.FlashUser) error {
	var u model.FlashUser
	db.First(&u, user.ID)
	if err := db.Session(&gorm.Session{FullSaveAssociations: true}).Save(&user).Error; err != nil {
		return err
	}
	return nil
}
