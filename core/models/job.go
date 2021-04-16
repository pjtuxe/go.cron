package models

import (
	"go.cron/core/models/job"
	"strings"
)

type JobModel struct {
	Name        string              `json:"name" validate:"required"`
	ID          string              `json:"id" validate:"required"`
	Image       string              `json:"image" validate:"required"`
	Entrypoint  string              `json:"entrypoint" validate:"-"`
	Command     []string            `json:"command" validate:"-"`
	Variables   []job.VariableModel `json:"variables" validate:"-"`
	CronPattern string              `json:"cronPattern" validate:"required"`
}

func (job JobModel) GetCronPattern() string {
	cron := strings.TrimLeft(job.CronPattern, " ")
	cron = strings.TrimRight(cron, " ")

	if 4 == strings.Count(cron, " ") {
		cron = "0 " + cron
	}

	return cron
}
