package configuration

import (
	"fmt"
	"strconv"

	"github.com/NickTaporuk/redeam/src/utils"
)

type (
	// DatabaseConfiger
	DatabaseConfiger interface {
		DatabaseMigrate(cnf *DatabaseConfiger) error
	}
	// DatabaseConfig
	DatabaseConfig struct {
		Migrate          bool
		Seeds            bool
		SslMode          string
		DatabaseHost     string
		DatabaseType     string
		DatabaseUser     string
		DatabaseName     string
		DatabasePassword string
		DatabasePort     string
	}
)

// CheckEnvVar is environment checker
func CheckEnvVar(key string, data map[string]string) error {
	var v string
	var err error

	v, err = utils.EnvVarExists(key)
	if err != nil {
		return err
	}

	data[key] = v

	return nil
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
