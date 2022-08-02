package poker_test

import (
	"bytes"
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
		stdout := &bytes.Buffer{}
		in := strings.NewReader("5\nChris wins\n")
		blindAlerter := &poker.SpyBlindAlerter{}
		dummyPlayerStore := &StubPlayerStore{}
		game := poker.NewGame(blindAlerter, dummyPlayerStore)

		cli := poker.NewCLI(in, stdout, game)
		cli.PlayPoker()
		AssertPlayerWin(t, game, "Chris")
	})
	t.Run("record cleo win from user input", func(t *testing.T) {
		in := strings.NewReader("5\nCleo wins\n")
		var dummyStdOut = &bytes.Buffer{}
		blindAlerter := &poker.SpyBlindAlerter{}
		dummyPlayerStore := &StubPlayerStore{}
		game := poker.NewGame(blindAlerter, dummyPlayerStore)

		cli := poker.NewCLI(in, dummyStdOut, game)
		cli.PlayPoker()

		AssertPlayerWin(t, game, "Cleo")
	})
	t.Run("it schedules printing of blind values", func(t *testing.T) {
		in := strings.NewReader("5\nChris wins\n")
		var dummyStdOut = &bytes.Buffer{}
		blindAlerter := &poker.SpyBlindAlerter{}
		dummyPlayerStore := &StubPlayerStore{}
		game := poker.NewGame(blindAlerter, dummyPlayerStore)
		cli := poker.NewCLI(in, dummyStdOut, game)
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

				alerts := game.GetBlindAlerter().GetAlerts()
				if len(alerts) <= i {
					t.Fatalf("alert %d was not scheduled %v", i, alerts)
				}

				got := alerts[i]
				assertScheduledAlert(t, got, want)
			})
		}
	})
	t.Run("it prompts the user to enter the number of players", func(t *testing.T) {
		stdout := &bytes.Buffer{}
		in := strings.NewReader("7\n")
		blindAlerter := &poker.SpyBlindAlerter{}
		dummyPlayerStore := &StubPlayerStore{}
		game := poker.NewGame(blindAlerter, dummyPlayerStore)

		cli := poker.NewCLI(in, stdout, game)
		cli.PlayPoker()

		got := stdout.String()
		want := poker.PlayerPrompt

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}

		cases := []wantScheduledAlert{
			{0 * time.Second, 100},
			{12 * time.Minute, 200},
			{24 * time.Minute, 300},
			{36 * time.Minute, 400},
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
	if got.GetAmount() != want.amount {
		t.Errorf("Amount NOT the same got: %v, want: %v", got, want)
	}

	if got.GetScheduledAlertAt() != want.at {
		t.Errorf("Scheduled At NOT the same got: %v, want: %v", got, want)
	}
}
