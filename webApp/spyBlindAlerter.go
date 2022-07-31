package poker

import (
	"time"
)

type SpyBlindAlerter struct {
	alerts []struct {
		scheduledAt time.Duration
		amount      int
	}
}

func (s *SpyBlindAlerter) ScheduleAlertAt(duration time.Duration, amount int) {
	s.alerts = append(s.alerts, struct {
		scheduledAt time.Duration
		amount      int
	}{duration, amount})
}

func (s *SpyBlindAlerter) GetAlerts() []struct {
	scheduledAt time.Duration
	amount      int
} {
	return s.alerts
}
