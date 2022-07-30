package poker_test

import (
	"golang_selfstudy/webApp"
	"strings"
	"testing"
)

func TestCLI(t *testing.T) {
	in := strings.NewReader("Chris wins\n")
	playerStore := &StubPlayerStore{}

	cli := &poker.CLI{}
	cli.InitializeCLI(playerStore, in)
	cli.PlayPoker()

	if len(playerStore.winCalls) != 1 {
		t.Fatal("expected a win call but didn't get any")
	}

	got := playerStore.winCalls[0]
	want := "Chris"
	if got != want {
		t.Errorf("didn't record correct winner, got %q, want %q", got, want)
	}
}
