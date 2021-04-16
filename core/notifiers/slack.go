package notifiers

type SlackNotifier struct {
}

func (n SlackNotifier) getName() string {
	return "Slack"
}

func (n SlackNotifier) notify() bool {
	return true
}
