package notifiers

type Notifier interface {
	getName() string
	notify() bool
}
