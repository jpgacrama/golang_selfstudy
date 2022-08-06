package poker

import (
	"io"
)

type GameSpy struct {
	StartedWith       int
	FinishedWith      string
	StartCalled       bool
	AlertsDestination io.Writer
}

func (g *GameSpy) Start(numberOfPlayers int, alertsDestination io.Writer) {
	g.StartedWith = numberOfPlayers
	g.StartCalled = true
	g.AlertsDestination = alertsDestination
}

func (g *GameSpy) Finish(winner string) {
	g.FinishedWith = winner
}
