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
	var jobs jobs
	response, requestErr := http.Get(utils.GetConfig().ApiUrl)
	utils.ErrorHandler("The HTTP request failed with error", requestErr, utils.GetConfig().Debug)
	if response != nil {
		data, readErr := ioutil.ReadAll(response.Body)
		utils.ErrorHandler("The HTTP request failed with error", readErr, utils.GetConfig().Debug)
		parseErr := json.Unmarshal(data, &jobs)
		utils.ErrorHandler("Parse error", parseErr, utils.GetConfig().Debug)
	}
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
