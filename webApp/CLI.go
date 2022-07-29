package poker

type CLI struct {
	playerStore PlayerStore
}

func (cli *CLI) PlayPoker() {
	cli.playerStore.RecordWin("Cleo")
}

func (c *CLI) SetPlayerStore(p PlayerStore) {
	c.playerStore = p
}
