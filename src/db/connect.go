package db

import (
	"fmt"

	"github.com/NickTaporuk/redeam/src/seeds"
	"github.com/NickTaporuk/redeam/src/utils"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

type (
	// DatabaseConfiger
	DatabaseConfiger interface {
		DatabaseMigrate(dbCnf *DatabaseConfig) error
	}
	// DatabaseConfig
	DatabaseConfig struct {
		SslMode          string
		DatabaseHost     string
		DatabaseType     string
		DatabaseUser     string
		DatabaseName     string
		DatabasePassword string
		DatabasePort     string
		Migrate          bool
		Seeds            bool
	}
)

// String() struct to string
func (dbCnf DatabaseConfig) String() string {
	databaseURI := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", dbCnf.DatabaseHost, dbCnf.DatabasePort, dbCnf.DatabaseUser, dbCnf.DatabaseName, dbCnf.DatabasePassword, dbCnf.SslMode)

	return databaseURI
}

func CheckEnvVar(key string) (val string, err error) {
	val, err = utils.EnvVarExists(key)
	if err != nil {
		return "", err
	}

	return val, nil
}

// Init initiate database configuration data
func Init(dbType string, host string, user string, dbName string, passw string, port string, sslMode string, migrate bool, seeds bool) *DatabaseConfig {

	return &DatabaseConfig{
		DatabaseType:     dbType,
		DatabaseHost:     host,
		DatabaseUser:     user,
		DatabaseName:     dbName,
		DatabasePassword: passw,
		DatabasePort:     port,
		SslMode:          sslMode,
		Migrate:          migrate,
		Seeds:            seeds,
	}
}

// NewConnection initiate db
func NewConnection(dbCnf *DatabaseConfig) (*gorm.DB, error) {

	var err error
	uri := dbCnf.String()

	db, err := gorm.Open(dbCnf.DatabaseType, uri)
	if err != nil {
		return nil, err
	}

	if dbCnf.Migrate {
		err = DatabaseMigrate(db, MigrateModels)
		if err != nil {
			return nil, err
		}
	}

	if dbCnf.Seeds {
		err = DatabaseMigrateSeeds(db, seeds.Seeds())
		if err != nil {
			return nil, err
		}
	}

	return db, nil
}
