package poker_test

import (
	"golang_selfstudy/webApp"
	"strings"
	"testing"
)

func TestCLI(t *testing.T) {
	t.Run("record chris win from user input", func(t *testing.T) {
		in := strings.NewReader("Chris wins\n")
		playerStore := &StubPlayerStore{}

		cli := &poker.CLI{}
		cli.InitializeCLI(playerStore, in)
		cli.PlayPoker()

		AssertPlayerWin(t, playerStore, "Chris")
	})
	t.Run("record cleo win from user input", func(t *testing.T) {
		in := strings.NewReader("Cleo wins\n")
		playerStore := &StubPlayerStore{}

		cli := &poker.CLI{}
		cli.InitializeCLI(playerStore, in)
		cli.PlayPoker()

		AssertPlayerWin(t, playerStore, "Cleo")
	})
}
