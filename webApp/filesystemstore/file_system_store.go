package filesystemstore

import (
	"encoding/json"
	"fmt"
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
	person := league.Find(name)

	if person != nil {
		person.Wins++
	} else {
		league = append(league, player.Player{Name: name, Wins: 1})
	}

	f.database.Seek(0, 0)
	json.NewEncoder(f.database).Encode(league)
}
