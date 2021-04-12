package main

import (
	"github.com/golobby/container/v2"
	"go.cron/utils"
)

func bootstrap() {
	di := utils.Di{}
	di.InitFor(container.New())
}

func main() {
	// Bootstrapping the application
	bootstrap()
}
