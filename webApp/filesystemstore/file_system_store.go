package filesystemstore

import (
	"fmt"
	"golang_selfstudy/webApp/player"
	"io"
)

type FileSystemPlayerStore struct {
	database io.Reader
}

func (f *FileSystemPlayerStore) SetDatabase(d io.Reader) {
	f.database = d
}

func (f *FileSystemPlayerStore) GetLeague() []player.Player {
	league, err := NewLeague(f.database)
	if err != nil {
		fmt.Println(fmt.Errorf("Unable to parse response from server %q into slice of Player, '%v'", league, err))
	}
	return league
}
