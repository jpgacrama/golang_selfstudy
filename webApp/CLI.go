package poker

import (
	"io"
)

type CLI struct {
	playerStore PlayerStore
	in          io.Reader
}

func (cli *CLI) PlayPoker() {
	cli.playerStore.RecordWin("Chris")
}

func (c *CLI) InitializeCLI(p PlayerStore, i io.Reader) {
	c.playerStore = p
	c.in = i
}
