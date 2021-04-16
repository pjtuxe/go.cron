package utils

import (
	"encoding/json"
	"fmt"
	"time"
)

type Error struct {
	Severity string `json:"severity"`
	Message  string `json:"message"`
	Date     string `json:"date"`
}

func NewError(message string, severity string) Error {
	err := Error{}
	err.Message = message
	err.Severity = severity
	err.Date = time.Now().String()
	return err
}

func ErrorHandler(msg string, err error, fatal bool) bool {
	if err != nil {
		LogError(msg)
		if fatal {
			panic(err)
		}
		return false
	}

	return true
}

func log(msg string) {
	fmt.Println(msg)
}

func LogInfo(msg string) {
	payload, _ := json.Marshal(NewError(msg, "info"))
	log(string(payload))
}

func LogError(msg string) {
	payload, _ := json.Marshal(NewError(msg, "error"))
	log(string(payload))
}

func LogWarn(msg string) {
	payload, _ := json.Marshal(NewError(msg, "warn"))
	log(string(payload))
}

func LogDebug(msg string) {
	if GetConfig().Debug {
		payload, _ := json.Marshal(NewError(msg, "debug"))
		log(string(payload))
	}
}

func ObjDebugger(obj interface{}, msg string) {
	if GetConfig().Debug {
		LogDebug(msg + ObjParser(obj))
	}
}

func ObjParser(obj interface{}) string {
	content, _ := json.Marshal(obj)
	return string(content)
}
