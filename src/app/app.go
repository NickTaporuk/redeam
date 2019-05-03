package app

import (
	"fmt"

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

	conn, err := db.Init()

	if err != nil {
		return err
	}

	fmt.Println("Connection : ", conn)

	return nil
}

func (m *Main) Close() {}
