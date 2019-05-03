package utils

import (
	"os"
)

// Exists is wrapper for check file exist
func FileExists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsExist(err) {
			return false
		}
	}
	return true
}
