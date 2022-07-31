package poker

import (
	"fmt"
	"time"
)

type scheduledAlert struct {
	at     time.Duration
	amount int
}

type SpyBlindAlerter struct {
	alerts []scheduledAlert
}

func (s *SpyBlindAlerter) ScheduleAlertAt(at time.Duration, amount int) {
	s.alerts = append(s.alerts, scheduledAlert{at, amount})
}

func (s *SpyBlindAlerter) GetAlerts() []scheduledAlert {
	return s.alerts
}

func (a *scheduledAlert) GetScheduledAlertAt() time.Duration {
	return a.at
}

func (a *scheduledAlert) GetAmount() int {
	return a.amount
}

func (s scheduledAlert) String() string {
	return fmt.Sprintf("%d chips at %v", s.amount, s.at)
}
