package utils

import (
	"github.com/go-playground/validator"
)

func Validate(obj interface{}, msg string, fatal bool) bool {
	validate := validator.New()
	err := validate.Struct(obj)
	if err != nil {
		if fatal {
			LogError(msg)
			panic(err)
		} else {
			LogWarn(msg)
			return false
		}
	}
	return true
}
