package league

import (
	"encoding/json"
	"fmt"
	"golang_selfstudy/webApp/player"
	"io"
)

type GroupOfPlayers []player.Player

func (l GroupOfPlayers) Find(name string) *player.Player {
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
