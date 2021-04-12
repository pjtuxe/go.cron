package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

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
	getJobs("https://httpbin.org/ip")
}
