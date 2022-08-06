package poker

import (
	"io"
	"os"
	"time"
)

type Game interface {
	Start(numberOfPlayers int, alertsDestination io.Writer)
	Finish(winner string)
}

type TexasHoldem struct {
	alerter BlindAlerter
	store   PlayerStore
	dest    io.Writer
}

func NewGame(alerter BlindAlerter, store PlayerStore) *TexasHoldem {
	return &TexasHoldem{
		alerter: alerter,
		store:   store,
	}
}

func (p *TexasHoldem) Start(numberOfPlayers int, alertsDestination io.Writer) {
	blindIncrement := time.Duration(5+numberOfPlayers) * time.Minute

	blinds := []int{100, 200, 300, 400, 500, 600, 800, 1000, 2000, 4000, 8000}
	blindTime := 0 * time.Second
	for _, blind := range blinds {
		p.alerter.ScheduleAlertAt(blindTime, blind, os.Stdout)
		blindTime = blindTime + blindIncrement
	}
	p.dest = alertsDestination
}

func (p *TexasHoldem) Finish(winner string) {
	p.store.RecordWin(winner)
}

func (g *TexasHoldem) GetStore() PlayerStore {
	return g.store
}

func (g *TexasHoldem) GetBlindAlerter() BlindAlerter {
	return g.alerter
}
