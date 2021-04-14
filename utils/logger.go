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

func LogDebug(msg string) {
	if GetConfig().Debug {
		payload, _ := json.Marshal(NewError(msg, "debug"))
		log(string(payload))
	}
}
