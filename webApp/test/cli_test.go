package poker_test

import (
	"bytes"
	"fmt"
	"golang_selfstudy/webApp"
	"strings"
	"testing"
	"time"
)

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

		cases := []poker.ScheduledAlert{
			{At: 0 * time.Second, Amount: 100},
			{At: 10 * time.Minute, Amount: 200},
			{At: 20 * time.Minute, Amount: 300},
			{At: 30 * time.Minute, Amount: 400},
			{At: 40 * time.Minute, Amount: 500},
			{At: 50 * time.Minute, Amount: 600},
			{At: 60 * time.Minute, Amount: 800},
			{At: 70 * time.Minute, Amount: 1000},
			{At: 80 * time.Minute, Amount: 2000},
			{At: 90 * time.Minute, Amount: 4000},
			{At: 100 * time.Minute, Amount: 8000},
		}

		for i, want := range cases {
			t.Run(fmt.Sprint(want), func(t *testing.T) {

				alerts := game.GetBlindAlerter().GetAlerts()
				assertAlertWasScheduled(t, alerts, i)

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

		got := stdout
		want := poker.PlayerPrompt

		assertMessagesSentToUser(t, got, want)
		cases := []poker.ScheduledAlert{
			{At: 0 * time.Second, Amount: 100},
			{At: 12 * time.Minute, Amount: 200},
			{At: 24 * time.Minute, Amount: 300},
			{At: 36 * time.Minute, Amount: 400},
		}

		alerts := blindAlerter.GetAlerts()
		for i, want := range cases {
			t.Run(fmt.Sprint(want), func(t *testing.T) {
				assertAlertWasScheduled(t, alerts, i)
				got := alerts[i]
				assertScheduledAlert(t, got, want)
			})
		}
	})
	t.Run("it prompts the user to enter the number of players and starts the game", func(t *testing.T) {
		stdout := &bytes.Buffer{}
		in := strings.NewReader("7\n")
		game := &poker.GameSpy{}

		cli := poker.NewCLI(in, stdout, game)
		cli.PlayPoker()

		gotPrompt := stdout
		wantPrompt := poker.PlayerPrompt
		assertMessagesSentToUser(t, gotPrompt, wantPrompt)

		if game.StartedWith != 7 {
			t.Errorf("wanted Start called with 7 but got %d", game.StartedWith)
		}
	})
	t.Run("it prints an error when a non numeric value is entered and does not start the game", func(t *testing.T) {
		stdout := &bytes.Buffer{}
		in := strings.NewReader("Pies\n")
		game := &poker.GameSpy{}
		cli := poker.NewCLI(in, stdout, game)
		cli.PlayPoker()

		if game.StartCalled {
			t.Errorf("game should not have started")
		}

		gotPrompt := stdout
		wantPrompt := poker.PlayerPrompt + poker.BadPlayerInputErrMsg
		assertMessagesSentToUser(t, gotPrompt, wantPrompt)
	})

}

func TestGame_Start(t *testing.T) {
	t.Run("schedules alerts on game start for 5 players", func(t *testing.T) {
		blindAlerter := &poker.SpyBlindAlerter{}
		dummyPlayerStore := &StubPlayerStore{}
		game := poker.NewGame(blindAlerter, dummyPlayerStore)
		game.Start(5)

		cases := []poker.ScheduledAlert{
			{At: 0 * time.Second, Amount: 100},
			{At: 10 * time.Minute, Amount: 200},
			{At: 20 * time.Minute, Amount: 300},
			{At: 30 * time.Minute, Amount: 400},
			{At: 40 * time.Minute, Amount: 500},
			{At: 50 * time.Minute, Amount: 600},
			{At: 60 * time.Minute, Amount: 800},
			{At: 70 * time.Minute, Amount: 1000},
			{At: 80 * time.Minute, Amount: 2000},
			{At: 90 * time.Minute, Amount: 4000},
			{At: 100 * time.Minute, Amount: 8000},
		}

		checkSchedulingCases(cases, t, blindAlerter)
	})
	t.Run("schedules alerts on game start for 7 players", func(t *testing.T) {
		blindAlerter := &poker.SpyBlindAlerter{}
		dummyPlayerStore := &StubPlayerStore{}
		game := poker.NewGame(blindAlerter, dummyPlayerStore)

		game.Start(7)

		cases := []poker.ScheduledAlert{
			{At: 0 * time.Second, Amount: 100},
			{At: 12 * time.Minute, Amount: 200},
			{At: 24 * time.Minute, Amount: 300},
			{At: 36 * time.Minute, Amount: 400},
		}

		checkSchedulingCases(cases, t, blindAlerter)
	})
}

func TestGame_Finish(t *testing.T) {
	store := &StubPlayerStore{}
	dummyBlindAlerter := &poker.SpyBlindAlerter{}
	game := poker.NewGame(dummyBlindAlerter, store)
	winner := "Ruth"

	game.Finish(winner)
	AssertPlayerWin(t, game, winner)
}

func assertScheduledAlert(t *testing.T, got poker.ScheduledAlert, want poker.ScheduledAlert) {
	t.Helper()
	if got.GetAmount() != want.Amount {
		t.Errorf("Amount NOT the same got: %v, want: %v", got, want)
	}

	if got.GetScheduledAlertAt() != want.At {
		t.Errorf("Scheduled At NOT the same got: %v, want: %v", got, want)
	}
}

func checkSchedulingCases(cases []poker.ScheduledAlert, t *testing.T, blindAlerter *poker.SpyBlindAlerter) {
	fmt.Println("This function is not yet implemented")
}

func assertMessagesSentToUser(t testing.TB, obtained *bytes.Buffer, messages ...string) {
	t.Helper()
	want := strings.Join(messages, "")
	got := obtained.String()
	if got != want {
		t.Errorf("got %q sent to stdout but expected %+v", got, messages)
	}
}

func assertAlertWasScheduled(t testing.TB, alerts []poker.ScheduledAlert, i int) {
	t.Helper()
	if len(alerts) <= i {
		t.Fatalf("alert %d was not scheduled %v", i, alerts)
	}
}
