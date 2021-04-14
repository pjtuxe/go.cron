package models

import "go.cron/models/job"
import "strings"

type JobModel struct {
	Name        string              `json:"name" validate:"empty=false > empty=false [empty=false] > ne=0"`
	ID          string              `json:"id" validate:"empty=false > empty=false [empty=false] > ne=0"`
	Image       string              `json:"image" validate:"empty=false > empty=false [empty=false] > ne=0"`
	Entrypoint  string              `json:"entrypoint"`
	Command     []string            `json:"command"`
	Variables   []job.VariableModel `json:"variables"`
	CronPattern string              `json:"cronPattern" validate:"empty=false > empty=false [empty=false] > ne=0"`
}

func (job JobModel) GetCronPattern() string {
	cron := strings.TrimLeft(job.CronPattern, " ")
	cron = strings.TrimRight(cron, " ")

	if 4 == strings.Count(cron, " ") {
		cron = "0 " + cron
	}

	return cron
}
