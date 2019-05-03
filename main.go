package main

import (
	"github.com/NickTaporuk/redeam/src/app"
	"github.com/NickTaporuk/redeam/src/utils"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

const (
	// DotEnvFilePath is environment file path
	DotEnvFilePath = "./docker/.env"
	// AppVersion is version application
	AppVersion = "0.0.1"
)

//nolint
func init() {
	// check has env file
	// if i work by docker container, i have to check only os environment
	// but if i work by developing i wont to work by .env file
	if utils.FileExists(DotEnvFilePath) {
		err := godotenv.Load(DotEnvFilePath)
		if err != nil {
			log.Fatalf("Error loading %s file", DotEnvFilePath)
		}
	}
}

func main() {

	m := app.Main{
		Version: AppVersion,
	}

	defer m.Close()

	err := m.Run()

	if err != nil {
		log.Fatal(err)
	}
}
