package filesystemstore

import (
	"encoding/json"
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
	var league []player.Player
	json.NewDecoder(f.database).Decode(&league)
	return league
}
