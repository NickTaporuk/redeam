package db

import (
	"fmt"
	"strconv"

	"github.com/NickTaporuk/redeam/src/seeds"
	"github.com/NickTaporuk/redeam/src/utils"
	"github.com/jinzhu/gorm"
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
func (dbCnf *DatabaseConfig) String() string {
	databaseURI := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		dbCnf.DatabaseHost, dbCnf.DatabasePort,
		dbCnf.DatabaseUser,
		dbCnf.DatabaseName,
		dbCnf.DatabasePassword,
		dbCnf.SslMode,
	)

	return databaseURI
}

func CheckEnvVar(key string, data map[string]string) (err error) {
	var v string

	v, err = utils.EnvVarExists(key)
	if err != nil {
		return err
	}

	data[key] = v

	return nil
}

//nolint
// Init intialize db
func Init() (*gorm.DB, error) {
	var err error
	data := make(map[string]string)

	err = CheckEnvVar(EnvNameDatabaseType, data)
	if err != nil {
		return nil, err
	}

	err = CheckEnvVar(EnvNameDatabaseHost, data)
	if err != nil {
		return nil, err
	}
	err = CheckEnvVar(EnvNameDatabaseUser, data)
	if err != nil {
		return nil, err
	}

	err = CheckEnvVar(EnvNameDatabaseName, data)
	if err != nil {
		return nil, err
	}

	err = CheckEnvVar(EnvNameDatabasePssCode, data)
	if err != nil {
		return nil, err
	}
	err = CheckEnvVar(EnvNameDatabasePort, data)
	if err != nil {
		return nil, err
	}

	err = CheckEnvVar(EnvNameDatabaseSSLMode, data)
	if err != nil {
		return nil, err
	}

	err = CheckEnvVar(EnvNameDatabaseMigrate, data)
	if err != nil {
		return nil, err
	}

	err = CheckEnvVar(EnvNameSeeds, data)
	if err != nil {
		return nil, err
	}

	dbConf, err := NewDatabaseConfig(data)
	if err != nil {
		return nil, err
	}

	conn, err := NewConnection(dbConf)
	if err != nil {
		return nil, err
	}

	return conn, nil

}

// NewDatabaseConfig initiate database configuration data
func NewDatabaseConfig(cnfData map[string]string) (*DatabaseConfig, error) {
	var err error
	convertedMigrate, err := strconv.ParseBool(cnfData[EnvNameDatabaseMigrate])
	if err != nil {
		return nil, err
	}

	convertedSeeds, err := strconv.ParseBool(cnfData[EnvNameSeeds])

	if err != nil {
		return nil, err
	}

	return &DatabaseConfig{
		DatabaseType:     cnfData[EnvNameDatabaseType],
		DatabaseHost:     cnfData[EnvNameDatabaseHost],
		DatabaseUser:     cnfData[EnvNameDatabaseUser],
		DatabaseName:     cnfData[EnvNameDatabaseName],
		DatabasePassword: cnfData[EnvNameDatabasePssCode],
		DatabasePort:     cnfData[EnvNameDatabasePort],
		SslMode:          cnfData[EnvNameDatabaseSSLMode],
		Migrate:          convertedMigrate,
		Seeds:            convertedSeeds,
	}, nil
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
