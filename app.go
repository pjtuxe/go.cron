package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"go.cron/models"
	"go.cron/models/job"
	"log"
	"net/http"
)

func getJobsMock() []models.JobModel {
	var resp []models.JobModel
	resp = append(
		resp,
		models.JobModel{
			Image:       "alpine:latest",
			ID:          "6077324217c1a973b708f95e",
			CronPattern: "* * * * *",
			Name:        "test",
			Command:     []string{"echo", "hello world"},
			Variables: []job.VariableModel{
				{Key: "TestKeyFromModel", Value: "TestValueFromModel"},
			},
		})
	return resp
}

func returnAllJobs(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllJobs")
	json.NewEncoder(w).Encode(getJobsMock())
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/all", returnAllJobs)
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
	fmt.Println("Rest API v2.0 - Mux Routers")
	handleRequests()
}
