package filesystemstore

import (
	"fmt"
	"golang_selfstudy/webApp/player"
	"io"
)

type FileSystemPlayerStore struct {
	database io.ReadWriteSeeker
}

func (f *FileSystemPlayerStore) SetDatabase(d io.ReadWriteSeeker) {
	f.database = d
}

func (f *FileSystemPlayerStore) GetLeague() []player.Player {
	f.database.Seek(0, 0)
	league, err := NewLeague(f.database)
	if err != nil {
		fmt.Println(fmt.Errorf("Unable to parse response from server %q into slice of Player, '%v'", league, err))
	}
	return league
}

func (f *FileSystemPlayerStore) GetPlayerScore(name string) int {
	var wins int
	for _, player := range f.GetLeague() {
		if player.Name == name {
			wins = player.Wins
			break
		}
	}
	return wins
}
