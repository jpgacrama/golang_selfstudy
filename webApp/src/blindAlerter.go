package poker

import (
	"fmt"
	"io"
	"time"
)

type BlindAlerter interface {
	ScheduleAlertAt(duration time.Duration, amount int, to io.Writer)
	GetAlerts() ([]ScheduledAlert, error)
}

type BlindAlerterFunc func(duration time.Duration, amount int, to io.Writer) BlindAlerter

func (a BlindAlerterFunc) ScheduleAlertAt(duration time.Duration, amount int, to io.Writer) {
	a(duration, amount, to)
}

func Alerter(duration time.Duration, amount int, to io.Writer) {
	time.AfterFunc(duration, func() {
		fmt.Fprintf(to, "Blind is now %d\n", amount)
	})
}

func (a BlindAlerterFunc) GetAlerts() ([]ScheduledAlert, error) {
	return nil, fmt.Errorf("BlindAlerter interface does not have alerts[]")
}
