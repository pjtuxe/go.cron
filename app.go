package main

import (
	"encoding/json"
	"go.cron/models"
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
	var jobs jobs
	parseErr := json.Unmarshal(data, &jobs)
	utils.ErrorHandler("Parse error", parseErr)
	return jobs
}

func main() {
	utils.GetConfig()
	ctx := services.InitDockerContext()
	utils.LogInfo("go.cron started")
	// Using thread-safe Tick facility
	tick := time.Tick(time.Second)

	for range tick {
		services.Runner{Ctx: ctx}.Run(getJobs())
	}
}
