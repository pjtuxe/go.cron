package services

import (
	"context"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"go.cron/utils"
)

type DockerContext struct {
	Context context.Context
	Cli     client.Client
}

func InitDockerContext() DockerContext {
	utils.LogInfo("initializing docker context")

	ctx := context.Background()
	// TODO: use docker api url
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())

	if err != nil {
		utils.LogError(err.Error())
		panic(err)
	}

	utils.LogInfo("pulling image: " + utils.GetConfig().RunnerBaseImage)

	// TODO: image pull policy from env
	_, err = cli.ImagePull(ctx, utils.GetConfig().RunnerBaseImage, types.ImagePullOptions{})

	if err != nil {
		utils.LogError(err.Error())
		panic(err)
	}

	utils.LogInfo("docker context initialized")
	return DockerContext{Cli: *cli, Context: ctx}
}
