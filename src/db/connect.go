package db

import (
	"github.com/NickTaporuk/redeam/src/configuration"
	"github.com/NickTaporuk/redeam/src/seeds"
	"github.com/jinzhu/gorm"
)

// Init initialize db
func Init(conf *configuration.Config) (*gorm.DB, error) {
	var err error
	var conn *gorm.DB

	conn, err = NewConnection(conf)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

// NewConnection initiate db
func NewConnection(conf *configuration.Config) (*gorm.DB, error) {

	var err error
	var db *gorm.DB

	uri := conf.String()

	db, err = gorm.Open(conf.DatabaseType, uri)
	if err != nil {
		return nil, err
	}

	if conf.Migrate {
		err = DatabaseMigrate(db, MigrateModels)
		if err != nil {
			return nil, err
		}
	}

	if conf.Seeds {
		err = DatabaseMigrateSeeds(db, seeds.Seeds())
		if err != nil {
			return nil, err
		}
	}

	return db, nil
}
