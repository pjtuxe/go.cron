package utils

import (
	"github.com/golobby/container/v2/pkg/container"
	"go.cron/notifiers"
)

type Di struct {
}

func (di Di) InitFor(container container.Container) {
	_ = container.Singleton(func() notifiers.Notifier {
		// TODO: choose notifier based on environment variables
		return notifiers.SlackNotifier{}
	})
}
