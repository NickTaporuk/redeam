package main

import (
	"log"

	"github.com/NickTaporuk/redeam/src/app"
	"github.com/NickTaporuk/redeam/src/utils"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

const (
	// DotEnvFilePath is environment file path
	DotEnvFilePath = "./docker/.env"
	// AppVersion is version application
	AppVersion = "0.0.1"
)

//nolint
func init() {
	var err error
	// check has env file
	// if i work by docker container, i have to check only os environment
	// but if i work by developing i wont to work by .env file
	if utils.FileExists(DotEnvFilePath) {
		err = godotenv.Load(DotEnvFilePath)
		if err != nil {
			log.Fatalf("Error loading %s file", DotEnvFilePath)
		}
	}
}

func main() {
	var err error
	var m app.Main

	m = app.Main{}

	m.SetVersion(AppVersion)

	err = m.Init()

	if err != nil {
		log.Fatal(err)
	}

	err = m.Run()

	if err != nil {
		log.Fatal(err)
	}

	defer m.Close()
}
