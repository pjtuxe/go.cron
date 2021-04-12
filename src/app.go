package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
)

var (
	apiUrl = flag.String("u", "https://6074d24f066e7e0017e7a5d5.mockapi.io/api/jobs/test", "API URL")
)

type job struct {
	Name       string `json:"name"`
	ID         string `json:"id"`
	Image      string `json:"image"`
	Entrypoint string `json:"entrypoint"`
	Command    string `json:"command"`
	Variables  struct {
		Key string `json:"key"`
	} `json:"variables"`
	Cronpattern string `json:"cronPattern"`
}

type jobs []job

func log(msg string) {
	fmt.Println(msg)
}

func errorHandler(msg string, err error) {
	if err != nil {
		fmt.Printf(msg+"%s\n", err)
		panic(err)
	}
}

func getJobs(apiUrl string) jobs {
	response, err := http.Get(apiUrl)
	data, _ := ioutil.ReadAll(response.Body)
	errorHandler("The HTTP request failed with error", err)
	var Response jobs
	parseErr := json.Unmarshal(data, &Response)
	errorHandler("Parse error", parseErr)
	return Response
}

func main() {
	log("go.cron started")
	fmt.Printf("%+v\n", getJobs(*apiUrl))
}
