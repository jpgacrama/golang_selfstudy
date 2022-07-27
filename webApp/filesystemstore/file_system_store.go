package filesystemstore

import (
	"encoding/json"
	"golang_selfstudy/webApp/league"
	"golang_selfstudy/webApp/player"
	"io"
)

type FileSystemPlayerStore struct {
	database io.ReadWriteSeeker
	league   league.GroupOfPlayers
}

func NewFileSystemPlayerStore(database io.ReadWriteSeeker) *FileSystemPlayerStore {
	database.Seek(0, 0)
	league, _ := league.NewLeague(database)
	return &FileSystemPlayerStore{
		database: database,
		league:   league,
	}
}
func (f *FileSystemPlayerStore) SetDatabase(d io.ReadWriteSeeker) {
	f.database = d
}

func (f *FileSystemPlayerStore) GetLeague() league.GroupOfPlayers {
	return f.league
}

func (f *FileSystemPlayerStore) GetPlayerScore(name string) int {
	player := f.league.Find(name)
	if player != nil {
		return player.Wins
	}
	return 0
}

func (f *FileSystemPlayerStore) RecordWin(name string) {
	person := f.league.Find(name)
	if person != nil {
		person.Wins++
	} else {
		f.league = append(f.league, player.Player{Name: name, Wins: 1})
	}

	f.database.Seek(0, 0)
	json.NewEncoder(f.database).Encode(f.league)
}
