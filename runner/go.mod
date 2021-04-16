module runner

go 1.16

replace go.cron/core => ../core

require (
	github.com/docker/docker v20.10.6+incompatible
	github.com/robfig/cron v1.2.0
	go.cron/core v0.0.0-00010101000000-000000000000
)
