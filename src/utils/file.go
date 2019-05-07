package utils

import (
	"os"
)

// Exists is wrapper for check file exist
func FileExists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if err != nil {
			return false
		}
	}
	return true
}
