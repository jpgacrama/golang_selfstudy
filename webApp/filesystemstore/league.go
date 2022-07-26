package filesystemstore

import (
	"encoding/json"
	"fmt"
	"golang_selfstudy/webApp/player"
	"io"
)

func NewLeague(rdr io.Reader) ([]player.Player, error) {
	var league []player.Player
	err := json.NewDecoder(rdr).Decode(&league)
	if err != nil {
		err = fmt.Errorf("problem parsing league, %v", err)
	}
	return league, err
}
