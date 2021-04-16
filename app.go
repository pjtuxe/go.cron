package main

import (
	"encoding/json"
	"github.com/docker/docker/api/types"
	"go.cron/models"
	"go.cron/models/job"
	"go.cron/services"
	"go.cron/utils"
	"io/ioutil"
	"net/http"
	"time"
)

type jobs []models.JobModel

func getJobs() jobs {
	response, requestErr := http.Get(utils.GetConfig().ApiUrl)
	utils.ErrorHandler("The HTTP request failed with error", requestErr)
	data, readErr := ioutil.ReadAll(response.Body)
	utils.ErrorHandler("The HTTP request failed with error", readErr)
	var Response jobs
	parseErr := json.Unmarshal(data, &Response)
	utils.ErrorHandler("Parse error", parseErr)
	return Response
}

func getJobsMock() []models.JobModel {
	var resp []models.JobModel
	resp = append(
		resp,
		models.JobModel{
			Image:           "alpine:latest",
			ImagePullPolicy: types.ImagePullOptions{},
			ID:              "6077324217c1a973b708f95e",
			CronPattern:     "* * * * *",
			Name:            "test",
			Command:         []string{"echo", "hello world"},
			Variables: []job.VariableModel{
				{Key: "TestKeyFromModel", Value: "TestValueFromModel"},
			},
		})
	return resp
}

func main() {
	utils.GetConfig()
	ctx := services.InitDockerContext()
	utils.LogInfo("go.cron started")
	// Using thread-safe Tick facility
	tick := time.Tick(time.Second)

	for range tick {
		services.Runner{Ctx: ctx}.Run(getJobsMock())
	}
}
