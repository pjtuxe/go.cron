package main

import (
	"encoding/json"
	"fmt"
	"go.cron/models"
	"go.cron/models/job"
	"go.cron/services"
	"go.cron/utils"
	"io/ioutil"
	"net/http"
	"time"
)

type jobs []models.JobModel

func errorHandler(msg string, err error) {
	if err != nil {
		fmt.Printf(msg+"%s\n", err)
		panic(err)
	}
}

func getJobs(apiUrl string) jobs {
	response, err := http.Get(utils.GetConfig().ApiUrl)
	data, _ := ioutil.ReadAll(response.Body)
	errorHandler("The HTTP request failed with error", err)
	var Response jobs
	parseErr := json.Unmarshal(data, &Response)
	errorHandler("Parse error", parseErr)
	return Response
}

func getJobsMock() []models.JobModel {
	var resp []models.JobModel
	resp = append(
		resp,
		models.JobModel{
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

func main() {
	utils.LogInfo("go.cron started")
	ctx := services.InitDockerContext()

	// Using thread-safe Tick facility
	tick := time.Tick(time.Second)

	for range tick {
		services.Runner{Ctx: ctx}.Run(getJobsMock())
	}
}
