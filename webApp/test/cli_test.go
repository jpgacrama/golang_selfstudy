package poker_test

import (
	"fmt"
	"golang_selfstudy/webApp"
	"strings"
	"testing"
	"time"
)

type wantScheduledAlert struct {
	at     time.Duration
	amount int
}

func TestCLI(t *testing.T) {
	t.Run("record chris win from user input", func(t *testing.T) {
		in := strings.NewReader("Chris wins\n")
		playerStore := &StubPlayerStore{}

		var dummySpyAlerter = &poker.SpyBlindAlerter{}
		cli := poker.NewCLI(playerStore, in, dummySpyAlerter)
		cli.PlayPoker()

		AssertPlayerWin(t, playerStore, "Chris")
	})
	t.Run("record cleo win from user input", func(t *testing.T) {
		in := strings.NewReader("Cleo wins\n")
		playerStore := &StubPlayerStore{}

		var dummySpyAlerter = &poker.SpyBlindAlerter{}
		cli := poker.NewCLI(playerStore, in, dummySpyAlerter)
		cli.PlayPoker()

		AssertPlayerWin(t, playerStore, "Cleo")
	})
	t.Run("it schedules printing of blind values", func(t *testing.T) {
		in := strings.NewReader("Chris wins\n")
		playerStore := &StubPlayerStore{}
		blindAlerter := &poker.SpyBlindAlerter{}

		cli := poker.NewCLI(playerStore, in, blindAlerter)
		cli.PlayPoker()

		cases := []wantScheduledAlert{
			{0 * time.Second, 100},
			{10 * time.Minute, 200},
			{20 * time.Minute, 300},
			{30 * time.Minute, 400},
			{40 * time.Minute, 500},
			{50 * time.Minute, 600},
			{60 * time.Minute, 800},
			{70 * time.Minute, 1000},
			{80 * time.Minute, 2000},
			{90 * time.Minute, 4000},
			{100 * time.Minute, 8000},
		}

		for i, want := range cases {
			t.Run(fmt.Sprint(want), func(t *testing.T) {

				if len(blindAlerter.GetAlerts()) <= i {
					t.Fatalf("alert %d was not scheduled %v", i, blindAlerter.GetAlerts())
				}

				got := blindAlerter.GetAlerts()[i]
				assertScheduledAlert(t, got, want)
			})
		}
	})
}

func assertScheduledAlert(t *testing.T, got poker.ScheduledAlert, want wantScheduledAlert) {
	t.Helper()
	if (got.GetAmount() != want.amount) && (got.GetScheduledAlertAt() != want.at) {
		t.Errorf("got: %v, want: %v", got, want)
	}
}
