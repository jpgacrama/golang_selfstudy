package poker

import (
	"bufio"
	"io"
	"strings"
)

type CLI struct {
	playerStore PlayerStore
	in          io.Reader
}

func (cli *CLI) PlayPoker() {
	reader := bufio.NewScanner(cli.in)
	reader.Scan()
	cli.playerStore.RecordWin(extractWinner(reader.Text()))
}

func (c *CLI) InitializeCLI(p PlayerStore, i io.Reader) {
	c.playerStore = p
	c.in = i
}

func extractWinner(userInput string) string {
	return strings.Replace(userInput, " wins", "", 1)
}
