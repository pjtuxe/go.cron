package src

import (
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"go.cron/core/models"
	"go.cron/core/services"
	"go.cron/core/utils"
	"strconv"
	"strings"
	"time"
)

type JobRunner struct {
	Ctx services.DockerContext
}

func ImagePull(name string, ctx services.DockerContext) {
	// TODO: image pull policy from job object
	_, pullErr := ctx.Cli.ImagePull(ctx.Context, name, types.ImagePullOptions{})

	utils.ErrorHandler("Image Pull error", pullErr, utils.GetConfig().Debug)
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

	return env
}

func (runner JobRunner) Run(job *models.JobModel) {
	utils.LogInfo("running job \"" + job.ID + "\"")
	ImagePull(job.Image, runner.Ctx)
	utils.ObjDebugger(job, "Container Created from: ")
	resp, err := runner.Ctx.Cli.ContainerCreate(
		runner.Ctx.Context,
		&container.Config{
			Image: job.Image,
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

	utils.ObjDebugger(job, "Container Started from: ")
	err = runner.Ctx.Cli.ContainerStart(
		runner.Ctx.Context,
		resp.ID,
		types.ContainerStartOptions{})

	if err != nil {
		utils.LogError(err.Error())
		return
	}
}
