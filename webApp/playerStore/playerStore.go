package playerstore

import (
	"golang_selfstudy/webApp/league"
	"golang_selfstudy/webApp/player"
)

type PlayerStore interface {
	GetPlayerScore(name string) int
	RecordWin(name string)
	GetLeague() league.GroupOfPlayers
}

type InMemoryPlayerStore struct {
	store map[string]int
}

func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{map[string]int{}}
}

func (i *InMemoryPlayerStore) RecordWin(name string) {
	i.store[name]++
}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return i.store[name]
}

func (i *InMemoryPlayerStore) GetLeague() league.GroupOfPlayers {
	var league league.GroupOfPlayers
	for name, wins := range i.store {
		league = append(league, player.Player{Name: name, Wins: wins})
	}
	return league
}
