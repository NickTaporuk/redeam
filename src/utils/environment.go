package utils

import (
	"errors"
	"fmt"
	"os"
)

const (
	// EnvVarNotFound
	EnvVarNotFound = "Requirenment environment variable key name %s not found"
)

func EnvVarExists(key string) (value string, err error) {
	val := os.Getenv(key)
	if val == "" {
		errText := fmt.Sprintf(EnvVarNotFound, key)
		return "", errors.New(errText)
	}

	return val, nil
}
