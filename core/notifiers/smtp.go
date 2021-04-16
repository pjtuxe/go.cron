package notifiers

type SMTP struct {
}

func (smtp SMTP) getName() string {
	return "SMTP"
}

func (smtp SMTP) notify() bool {
	return true
}
