package poker

import (
	"encoding/json"
	"fmt"
	"io"
)

type GroupOfPlayers []Player

func (l GroupOfPlayers) Find(name string) *Player {
	for i, p := range l {
		if p.Name == name {
			return &l[i]
		}
	}
	return nil
}

func NewLeague(rdr io.Reader) (GroupOfPlayers, error) {
	var league GroupOfPlayers
	err := json.NewDecoder(rdr).Decode(&league)
	if err != nil {
		err = fmt.Errorf("problem parsing league, %v", err)
	}
	return league, err
}
