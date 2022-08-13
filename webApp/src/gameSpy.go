package poker

import (
	"io"
)

type GameSpy struct {
	StartCalled      bool
	StartCalledWith  int
	BlindAlert       []byte
	FinishedCalled   bool
	FinishCalledWith string
}

func (g *GameSpy) Start(numberOfPlayers int, out io.Writer) {
	g.StartCalled = true
	g.StartCalledWith = numberOfPlayers
	out.Write(g.BlindAlert)
}

func (g *GameSpy) Finish(winner string) {
	g.FinishCalledWith = winner
}
