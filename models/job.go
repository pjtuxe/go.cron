package models

import (
	"github.com/docker/docker/api/types"
	"go.cron/models/job"
	"strings"
)

type JobModel struct {
	Name            string                 `json:"name" validate:"required"`
	ID              string                 `json:"id" validate:"required"`
	Image           string                 `json:"image" validate:"required"`
	ImagePullPolicy types.ImagePullOptions `json:"image" validate:"required"`
	Entrypoint      string                 `json:"entrypoint" validate:"-"`
	Command         []string               `json:"command" validate:"-"`
	Variables       []job.VariableModel    `json:"variables" validate:"-"`
	CronPattern     string                 `json:"cronPattern" validate:"required"`
}

func (job JobModel) GetCronPattern() string {
	cron := strings.TrimLeft(job.CronPattern, " ")
	cron = strings.TrimRight(cron, " ")

	if 4 == strings.Count(cron, " ") {
		cron = "0 " + cron
	}

	return cron
}
