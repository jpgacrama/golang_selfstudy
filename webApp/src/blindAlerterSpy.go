package poker

import (
	"fmt"
	"io"
	"time"
)

type ScheduledAlert struct {
	At     time.Duration
	Amount int
	To     io.Writer
}
type SpyBlindAlerter struct {
	alerts []ScheduledAlert
}

func (s *SpyBlindAlerter) ScheduleAlertAt(duration time.Duration, amount int, to io.Writer) {
	s.alerts = append(s.alerts, ScheduledAlert{duration, amount, to})
}

func (s *SpyBlindAlerter) GetAlerts() ([]ScheduledAlert, error) {
	if s.alerts != nil {
		return s.alerts, nil
	}
	return nil, fmt.Errorf("alerts[] do not exist")
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
