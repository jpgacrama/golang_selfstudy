package poker

import (
	"fmt"
	"os"
	"time"
)

type BlindAlerter interface {
	ScheduleAlertAt(duration time.Duration, amount int)
	GetAlerts() []ScheduledAlert
}

type BlindAlerterFunc func(duration time.Duration, amount int) BlindAlerter

func (a BlindAlerterFunc) ScheduleAlertAt(duration time.Duration, amount int) {
	a(duration, amount)
}

func StdOutAlerter(duration time.Duration, amount int) BlindAlerter {
	time.AfterFunc(duration, func() {
		fmt.Fprintf(os.Stdout, "Blind is now %d\n", amount)
	})
	return nil
}

func (a BlindAlerterFunc) GetAlerts() []ScheduledAlert {
	return nil
}
