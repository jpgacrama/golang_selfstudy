package filesystemstore

import (
	"encoding/json"
	"fmt"
	"golang_selfstudy/webApp/league"
	"io"
)

type FileSystemPlayerStore struct {
	database io.ReadWriteSeeker
}

func (f *FileSystemPlayerStore) SetDatabase(d io.ReadWriteSeeker) {
	f.database = d
}

func (f *FileSystemPlayerStore) GetLeague() league.GroupOfPlayers {
	f.database.Seek(0, 0)
	league, err := league.NewLeague(f.database)
	if err != nil {
		fmt.Println(fmt.Errorf("Unable to parse response from server %q into slice of Player, '%v'", league, err))
	}
	return league
}

func (f *FileSystemPlayerStore) GetPlayerScore(name string) int {
	player := f.GetLeague().Find(name)
	if player != nil {
		return player.Wins
	}
	return 0
}

func (f *FileSystemPlayerStore) RecordWin(name string) {
	league := f.GetLeague()
	player := league.Find(name)

	if player != nil {
		player.Wins++
	}

	f.database.Seek(0, 0)
	json.NewEncoder(f.database).Encode(league)
}
