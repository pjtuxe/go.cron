package main

import (
	"flag"
	"fmt"
	"go/types"
	"io/ioutil"
	"net/http"
)

var (
	apiUrl = flag.String("u", "https://httpbin.org/ip", "API URL")
)

type job struct {
	name                string
	image               string
	entrypoint          string
	commandcron_pattern string
	variables           types.Array
	command             types.Array
}

func log(msg string) {
	fmt.Println(msg)
}

func errorHandler(msg string, err error) {
	if err != nil {
		fmt.Printf(msg+"%s\n", err)
		panic(err)
	}
}

func getJobs(apiUrl string) {
	response, err := http.Get(apiUrl)
	errorHandler("The HTTP request failed with error", err)
	data, _ := ioutil.ReadAll(response.Body)
	log(string(data))
}

func main() {
	log("go.cron started")
	getJobs(*apiUrl)
}
