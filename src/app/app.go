package app

import (
	"fmt"
	"os"
	"strconv"

	"github.com/NickTaporuk/redeam/src/db"
)

type (
	// Runner use for implement method Run
	Runner interface {
		Run() error
		Close()
	}
	// Main structure use for run app
	Main struct {
		Version string
	}
)

// Run method base runner of application
func (m *Main) Run() error {
	// initate db
	var err error
	fmt.Println(os.Environ())
	dbType, err := db.CheckEnvVar(db.EnvNameDatabaseType)
	if err != nil {
		return err
	}

	host, err := db.CheckEnvVar(db.EnvNameDatabaseHost)
	if err != nil {
		return err
	}

	user, err := db.CheckEnvVar(db.EnvNameDatabaseUser)
	if err != nil {
		return err
	}

	dbName, err := db.CheckEnvVar(db.EnvNameDatabaseName)
	if err != nil {
		return err
	}

	passw, err := db.CheckEnvVar(db.EnvNameDatabasePassword)
	if err != nil {
		return err
	}

	port, err := db.CheckEnvVar(db.EnvNameDatabasePort)
	if err != nil {
		return err
	}

	sslMode, err := db.CheckEnvVar(db.EnvNameDatabaseSSLMode)
	if err != nil {
		return err
	}

	migrate, err := db.CheckEnvVar(db.EnvNameDatabaseMigrate)
	if err != nil {
		return err
	}

	seeds, err := db.CheckEnvVar(db.EnvNameSeeds)
	if err != nil {
		return err
	}

	convertedMigrate, err := strconv.ParseBool(migrate)
	if err != nil {
		return err
	}

	convertedSeeds, err := strconv.ParseBool(seeds)

	if err != nil {
		return err
	}

	dbConf := db.Init(dbType, host, user, dbName, passw, port, sslMode, convertedMigrate, convertedSeeds)
	conn, err := db.NewConnection(dbConf)

	if err != nil {
		return err
	}

	fmt.Println("Connection : ", conn)

	return nil
}

func (m *Main) Close() {}
