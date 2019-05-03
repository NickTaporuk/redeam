package db

import (
	"github.com/NickTaporuk/redeam/src/models"
	"github.com/jinzhu/gorm"
)

// DatabaseMigrate gorm migration wrapper
func DatabaseMigrateSeeds(db *gorm.DB, mds models.RedeamModels) error {
	var err error

	tx := db.Begin()

	if tx.Error != nil {
		tx.Rollback()
		return tx.Error
	}

	for _, md := range mds {
		if err = db.FirstOrInit(md).Error; gorm.IsRecordNotFoundError(err) {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit().Error
}
