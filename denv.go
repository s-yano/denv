package denv

import (
	"os"
	"strconv"
)

func GetStr(key string, default_val string) string {
	if env := os.Getenv(key); env != "" {
		return env
	}
	return default_val
}

func GetInt(key string, default_val int) int {
	if env := os.Getenv(key); env != "" {
		if num, err := strconv.Atoi(env); err == nil {
			return num
		}
	}
	return default_val
}
