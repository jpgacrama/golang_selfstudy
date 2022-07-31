package poker

import (
	"time"
)

type Alerts struct {
	scheduledAt time.Duration
	amount      int
}

type SpyBlindAlerter struct {
	alerts []Alerts
}

func (s *SpyBlindAlerter) ScheduleAlertAt(duration time.Duration, amount int) {
	s.alerts = append(s.alerts, Alerts{scheduledAt: duration, amount: amount})
}

func (s *SpyBlindAlerter) GetAlerts() []Alerts {
	return s.alerts
}

func (a *Alerts) GetScheduledAt() time.Duration {
	return a.scheduledAt
}

func (a *Alerts) GetAmount() int {
	return a.amount
}
