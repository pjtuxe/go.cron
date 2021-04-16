package services

import (
	"context"
	"github.com/docker/docker/client"
	"go.cron/utils"
)

type DockerContext struct {
	Context context.Context
	Cli     client.Client
}

func InitDockerContext() DockerContext {
	utils.LogDebug("initializing docker context")
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	utils.ErrorHandler("docker initialization error", err, true)
	utils.LogDebug("docker context initialized")
	return DockerContext{Cli: *cli, Context: ctx}
}
