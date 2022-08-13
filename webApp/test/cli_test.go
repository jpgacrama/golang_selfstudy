package poker_test

import (
	"bytes"
	"golang_selfstudy/webApp/src"
	"io"
	"reflect"
	"strings"
	"testing"
	"time"
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
		assertFinishWith(t, game, "Chris")
	})
	t.Run("start game with 8 players and record 'Cleo' as winner", func(t *testing.T) {
		game := &poker.GameSpy{}
		in := userSends("8", "Cleo wins")
		dummyStdOut := &bytes.Buffer{}
		cli := poker.NewCLI(in, dummyStdOut, game)
		cli.PlayPoker()
		assertGameStartedWith(t, game, 8)
		assertFinishWith(t, game, "Cleo")
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
			{At: 10 * time.Minute, Amount: 200, To: to},
			{At: 20 * time.Minute, Amount: 300, To: to},
			{At: 30 * time.Minute, Amount: 400, To: to},
			{At: 40 * time.Minute, Amount: 500, To: to},
			{At: 50 * time.Minute, Amount: 600, To: to},
			{At: 60 * time.Minute, Amount: 800, To: to},
			{At: 70 * time.Minute, Amount: 1000, To: to},
			{At: 80 * time.Minute, Amount: 2000, To: to},
			{At: 90 * time.Minute, Amount: 4000, To: to},
			{At: 100 * time.Minute, Amount: 8000, To: to},
		}

		checkSchedulingCases(t, cases, blindAlerter)
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
	store := &StubPlayerStore{}
	dummyBlindAlerter := &poker.SpyBlindAlerter{}
	game := poker.NewTexasHoldem(dummyBlindAlerter, store)
	winner := "Ruth"
	game.Finish(winner)
	AssertPlayerWin(t, game, winner)
}

func checkSchedulingCases(t *testing.T, cases []poker.ScheduledAlert, blindAlerter *poker.SpyBlindAlerter) {
	t.Helper()
	gotAlerts, err := blindAlerter.GetAlerts()
	if err != nil {
		t.Fatalf("There are no alerts obtained.")
	}
	isEqual := reflect.DeepEqual(cases, gotAlerts)
	if !isEqual {
		t.Errorf("%v is not the same as %v", cases, gotAlerts)
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

func assertFinishWith(t testing.TB, game *poker.GameSpy, winner string) {
	t.Helper()
	got_winner := game.FinishCalledWith
	if got_winner != winner {
		t.Fatalf("Winner is wrong. Expected: %s, Got:%s", winner, got_winner)
	}
}

func assertGameError(t testing.TB, game *poker.GameSpy) {
	t.Helper()
	gotFinishCalledWith := game.FinishCalledWith
	if gotFinishCalledWith != "" {
		t.Fatalf("Game should have thrown an error.")
	}
}
