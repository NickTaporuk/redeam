package db

import (
	"github.com/NickTaporuk/redeam/src/models"
	"github.com/jinzhu/gorm"
)

var (
	// books use for initiate books model
	books models.Books
	// MigrateModels is collection of migrate models
	MigrateModels = []models.RedeamModel{
		&books,
	}
)

// DatabaseMigrate gorm migration wrapper
func DatabaseMigrate(db *gorm.DB, mds models.RedeamModels) error {
	tx := db.Begin()

	if tx.Error != nil {
		return tx.Error
	}

	for _, md := range mds {
		err := db.AutoMigrate(md).Error

		if err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit().Error
}
