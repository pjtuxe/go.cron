package utils

import (
	"os"
	"strconv"
)

type Config struct {
	Debug     bool   `validate:"-"`
	ApiUrl    string `validate:"required,min=4"`
	RunnerEnv string `validate:"-"`
}

func GetConfig() Config {
	debug, _ := strconv.ParseBool(os.Getenv("DEBUG"))

	conf := Config{
		Debug:     debug,
		ApiUrl:    os.Getenv("API_URL"),
		RunnerEnv: os.Getenv("RUNNER_ENV"),
	}

	if Validate(conf, "Invalid Configuration", true) {
		return conf
	}

	return Config{}
}
