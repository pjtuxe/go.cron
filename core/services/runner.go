package services

import (
	"github.com/robfig/cron"
	"go.cron/core/models"
	"go.cron/core/utils"
	"time"
)

type Runner struct {
	Ctx DockerContext
}

func (runner Runner) Run(jobs []models.JobModel) {
	c := cron.New()

	for _, job := range jobs {
		go func(job models.JobModel) {
			if utils.Validate(job, "Invalid job: "+job.ID, utils.GetConfig().Debug) {
				err := c.AddFunc(job.GetCronPattern(), func() {
					JobRunner{Ctx: runner.Ctx}.Run(&job)
				})

				if err != nil {
					utils.LogError(err.Error())
				}
			} else {
				utils.LogWarn(utils.ObjParser(job))
			}
		}(job)
	}

	// Running all the jobs
	c.Start()
	time.Sleep(45 * time.Second)
	// Cleanup for the next stage
	c.Stop()
}
