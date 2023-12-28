package op

import (
	"gorm.io/gorm"
	"security/db/model"
)

func InsetSecurity(db *gorm.DB, security model.Security) error {
	if security.ID != 0 {
		result := db.Save(&security)
		if result.Error != nil {
			return result.Error
		}
	} else {
		// 如果 user.ID 是零值，你可能需要处理这种情况，例如报错或者使用 Create
		db.Create(&security)
	}
	return nil
}
