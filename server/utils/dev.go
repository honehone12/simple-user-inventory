package utils

import (
	"os"
	"strings"
)

func IsDev() bool {
	devEnv := os.Getenv("DEV")
	if len(devEnv) == 0 || strings.Compare(devEnv, "true") != 0 {
		return false
	}
	return true
}
