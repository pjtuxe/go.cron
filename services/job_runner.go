package services

import (
	"encoding/json"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"go.cron/models"
	"go.cron/utils"
	"strconv"
	"strings"
	"time"
)

type JobRunner struct {
	Ctx DockerContext
}

func generateContainerNameFor(job *models.JobModel) string {
	return job.ID + "-" + job.Name + "-" + strconv.FormatInt(time.Now().Unix(), 10)
}

func prepareJobEnvironment(job *models.JobModel) []string {
	var env []string

	// Append job model variables
	for _, variable := range job.Variables {
		env = append(env, variable.Key+"="+variable.Value)
	}

	// Append default variables
	if len(utils.GetConfig().RunnerEnv) > 0 {
		for _, variable := range strings.Split(utils.GetConfig().RunnerEnv, "|") {
			env = append(env, variable)
		}
	}

	content, _ := json.Marshal(env)
	utils.LogInfo(string(content))

	return env
}

func (runner JobRunner) Run(job *models.JobModel) {
	utils.LogInfo("running job \"" + job.ID + "\"")

	// TODO: make API request to API to generate log id

	utils.LogDebug("create container")
	resp, err := runner.Ctx.Cli.ContainerCreate(
		runner.Ctx.Context,
		&container.Config{
			Image: utils.GetConfig().RunnerBaseImage,
			Cmd:   job.Command,
			Tty:   false,
			Env:   prepareJobEnvironment(job),
			Labels: map[string]string{
				"managed-by":     "go.cron",
				"go-cron-job-id": job.ID,
				"go-cron-log-id": "",
			},
		},
		nil,
		nil,
		nil,
		generateContainerNameFor(job))

	if err != nil {
		utils.LogError(err.Error())
		return
	}

	utils.LogDebug("start container")
	err = runner.Ctx.Cli.ContainerStart(
		runner.Ctx.Context,
		resp.ID,
		types.ContainerStartOptions{})

	if err != nil {
		utils.LogError(err.Error())
		return
	}
}
