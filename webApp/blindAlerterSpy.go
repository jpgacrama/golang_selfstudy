package poker

import (
	"fmt"
	"time"
)

type ScheduledAlert struct {
	At     time.Duration
	Amount int
}
type SpyBlindAlerter struct {
	alerts []ScheduledAlert
}

func (s *SpyBlindAlerter) ScheduleAlertAt(at time.Duration, amount int) {
	s.alerts = append(s.alerts, ScheduledAlert{at, amount})
}

func (s *SpyBlindAlerter) GetAlerts() []ScheduledAlert {
	return s.alerts
}

func (a *ScheduledAlert) GetScheduledAlertAt() time.Duration {
	return a.At
}

func (a *ScheduledAlert) GetAmount() int {
	return a.Amount
}

func (s ScheduledAlert) String() string {
	return fmt.Sprintf("%d chips at %v", s.Amount, s.At)
}
