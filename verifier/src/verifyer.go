package src

import (
	"go.cron/core/services"
)

type JobVerifier struct {
	Ctx services.DockerContext
}

func (runner JobVerifier) Run() {
}
