package poker_test

import (
	"golang_selfstudy/webApp"
	"testing"
)

func TestCLI(t *testing.T) {
	playerStore := &StubPlayerStore{}
	cli := &poker.CLI{}
	cli.SetPlayerStore(playerStore)
	cli.PlayPoker()

	if len(playerStore.winCalls) != 1 {
		t.Fatal("expected a win call but didn't get any")
	}
}
