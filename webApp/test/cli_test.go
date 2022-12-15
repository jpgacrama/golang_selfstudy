package poker_test

import (
	"bytes"
	"io"
	"reflect"
	"strings"
	"testing"
	"time"
	"webApp/src"
)

func TestCLI(t *testing.T) {
	t.Run("start game with 3 players and finish game with 'Chris' as winner", func(t *testing.T) {
		game := &poker.GameSpy{}
		stdout := &bytes.Buffer{}
		in := userSends("3", "Chris wins")
		cli := poker.NewCLI(in, stdout, game)
		cli.PlayPoker()
		assertMessagesSentToUser(t, stdout, poker.PlayerPrompt)
		assertGameStartedWith(t, game, 3)
		assertFinishCalledWith(t, game, "Chris")
	})
	t.Run("start game with 8 players and record 'Cleo' as winner", func(t *testing.T) {
		game := &poker.GameSpy{}
		in := userSends("8", "Cleo wins")
		dummyStdOut := &bytes.Buffer{}
		cli := poker.NewCLI(in, dummyStdOut, game)
		cli.PlayPoker()
		assertGameStartedWith(t, game, 8)
		assertFinishCalledWith(t, game, "Cleo")
	})
}

func TestGame_Start(t *testing.T) {
	t.Run("schedules alerts on game start for 5 players", func(t *testing.T) {
		blindAlerter := &poker.SpyBlindAlerter{}
		dummyPlayerStore := &StubPlayerStore{}
		game := poker.NewTexasHoldem(blindAlerter, dummyPlayerStore)
		to := io.Discard
		game.Start(5, to)

		cases := []poker.ScheduledAlert{
			{At: 0 * time.Second, Amount: 100, To: to},
			{At: 10 * time.Second, Amount: 200, To: to},
			{At: 20 * time.Second, Amount: 300, To: to},
			{At: 30 * time.Second, Amount: 400, To: to},
			{At: 40 * time.Second, Amount: 500, To: to},
			{At: 50 * time.Second, Amount: 600, To: to},
			{At: 60 * time.Second, Amount: 800, To: to},
			{At: 70 * time.Second, Amount: 1000, To: to},
			{At: 80 * time.Second, Amount: 2000, To: to},
			{At: 90 * time.Second, Amount: 4000, To: to},
			{At: 100 * time.Second, Amount: 8000, To: to},
		}

		checkSchedulingCases(t, cases, blindAlerter)
	})

	t.Run("schedules alerts on game starts for 7 players", func(t *testing.T) {
		blindAlerter := &poker.SpyBlindAlerter{}
		dummyPlayerStore := &StubPlayerStore{}
		game := poker.NewTexasHoldem(blindAlerter, dummyPlayerStore)

		to := io.Discard
		game.Start(7, to)

		cases := []poker.ScheduledAlert{
			{At: 0 * time.Second, Amount: 100, To: to},
			{At: 12 * time.Second, Amount: 200, To: to},
			{At: 24 * time.Second, Amount: 300, To: to},
			{At: 36 * time.Second, Amount: 400, To: to},
			{At: 48 * time.Second, Amount: 500, To: to},
			{At: 60 * time.Second, Amount: 600, To: to},
			{At: 72 * time.Second, Amount: 800, To: to},
			{At: 84 * time.Second, Amount: 1000, To: to},
			{At: 96 * time.Second, Amount: 2000, To: to},
			{At: 108 * time.Second, Amount: 4000, To: to},
			{At: 120 * time.Second, Amount: 8000, To: to},
		}
		checkSchedulingCases(t, cases, blindAlerter)
	})
}

func TestGame_EnterNumberOfPlayers_AndStart(t *testing.T) {
	t.Run("it prompts the user to enter the number of players and starts the game", func(t *testing.T) {
		stdout := &bytes.Buffer{}
		in := userSends("7", "Jonas wins")
		game := &poker.GameSpy{}
		cli := poker.NewCLI(in, stdout, game)
		cli.PlayPoker()
		gotPrompt := stdout.String()
		wantPrompt := poker.PlayerPrompt
		if gotPrompt != wantPrompt {
			t.Errorf("got %q, want %q", gotPrompt, wantPrompt)
		}
		if game.StartCalledWith != 7 {
			t.Errorf("wanted Start called with 7 but got %d", game.StartCalledWith)
		}
	})
}

func TestGame_ErrorCases(t *testing.T) {
	t.Run("number of players is wrong game should not start", func(t *testing.T) {
		game := &poker.GameSpy{}
		in := userSends("john", "This string is useless")
		dummyStdOut := &bytes.Buffer{}
		cli := poker.NewCLI(in, dummyStdOut, game)
		cli.PlayPoker()
		assertGameError(t, game)
	})
	t.Run("number of players is correct but winner statement is not", func(t *testing.T) {
		game := &poker.GameSpy{}
		in := userSends("8", "Lloyd is a killer")
		dummyStdOut := &bytes.Buffer{}
		cli := poker.NewCLI(in, dummyStdOut, game)
		cli.PlayPoker()
		assertGameStartedWith(t, game, 8)
		assertGameError(t, game)
	})
	t.Run("it prints an error when a non numeric value is entered and does not start the game", func(t *testing.T) {
		game := &poker.GameSpy{}
		stdout := &bytes.Buffer{}
		in := userSends("pies")
		cli := poker.NewCLI(in, stdout, game)
		cli.PlayPoker()
		assertGameNotStarted(t, game)
		assertMessagesSentToUser(t, stdout, poker.PlayerPrompt, poker.BadPlayerInputErrMsg)
	})
}

func TestGame_Finish(t *testing.T) {
	t.Run("start game with 3 players and finish game with 'Chris' as winner", func(t *testing.T) {
		game := &poker.GameSpy{}
		stdout := &bytes.Buffer{}
		in := userSends("3", "Chris wins")
		cli := poker.NewCLI(in, stdout, game)
		cli.PlayPoker()
		assertMessagesSentToUser(t, stdout, poker.PlayerPrompt)
		assertGameStartedWith(t, game, 3)
		assertFinishCalledWith(t, game, "Chris")
	})
	t.Run("start game with 8 players and record 'Cleo' as winner", func(t *testing.T) {
		game := &poker.GameSpy{}
		in := userSends("8", "Cleo wins")
		dummyStdOut := &bytes.Buffer{}
		cli := poker.NewCLI(in, dummyStdOut, game)
		cli.PlayPoker()
		assertGameStartedWith(t, game, 8)
		assertFinishCalledWith(t, game, "Cleo")
	})
	t.Run("it prints an error when a non numeric value is entered and does not start the game", func(t *testing.T) {
		game := &poker.GameSpy{}
		stdout := &bytes.Buffer{}
		in := userSends("pies")
		cli := poker.NewCLI(in, stdout, game)
		cli.PlayPoker()
		assertGameNotStarted(t, game)
		assertMessagesSentToUser(t, stdout, poker.PlayerPrompt, poker.BadPlayerInputErrMsg)
	})
}

func checkSchedulingCases(t *testing.T, cases []poker.ScheduledAlert, blindAlerter *poker.SpyBlindAlerter) {
	t.Helper()
	gotAlerts, err := blindAlerter.GetAlerts()
	if err != nil {
		t.Fatalf("There are no alerts obtained.")
	}
	isEqual := reflect.DeepEqual(cases, gotAlerts)
	if !isEqual {
		t.Errorf("\n\t%v \n is not the same as \n\t%v", cases, gotAlerts)
	}
}

func assertMessagesSentToUser(t testing.TB, obtained *bytes.Buffer, messages ...string) {
	t.Helper()
	want := strings.Join(messages, "")
	got := obtained.String()
	if got != want {
		t.Errorf("got %q sent to stdout but expected %+v", got, messages)
	}
}

func userSends(messages ...string) *strings.Reader {
	want := strings.Join(messages, "\n")
	return strings.NewReader(want)
}

func assertGameNotStarted(t testing.TB, game *poker.GameSpy) {
	t.Helper()
	if game.StartCalled {
		t.Fatalf("Game should not have started yet.")
	}
}

func assertGameStartedWith(t testing.TB, game *poker.GameSpy, numPlayers int) {
	t.Helper()
	got_players := game.StartCalledWith
	if got_players != numPlayers {
		t.Fatalf("Game started with the wrong number of players. Expected: %d, Got:%d", numPlayers, got_players)
	}
}

func assertFinishCalledWith(t testing.TB, game *poker.GameSpy, winner string) {
	t.Helper()
	passed := retryUntil(500*time.Millisecond, func() bool {
		return game.FinishCalledWith == winner
	})
	if !passed {
		t.Errorf("expected finish called with %q but got %q", winner, game.FinishCalledWith)
	}
}

func assertGameError(t testing.TB, game *poker.GameSpy) {
	t.Helper()
	gotFinishCalledWith := game.FinishCalledWith
	if gotFinishCalledWith != "" {
		t.Fatalf("Game should have thrown an error.")
	}
}

func retryUntil(d time.Duration, f func() bool) bool {
	deadline := time.Now().Add(d)
	for time.Now().Before(deadline) {
		if f() {
			return true
		}
	}
	return false
}
