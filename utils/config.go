package utils

import (
	"os"
	"strconv"
)

type Config struct {
	RunnerBaseImage string
	Debug           bool
	ApiUrl          string
	RunnerEnv       string
}

func GetConfig() Config {
	debug, _ := strconv.ParseBool(os.Getenv("DEBUG"))

	return Config{
		RunnerBaseImage: os.Getenv("RUNNER_BASE_IMAGE"),
		Debug:           debug,
		ApiUrl:          os.Getenv("API_URL"),
		RunnerEnv:       os.Getenv("RUNNER_ENV"),
	}
}
