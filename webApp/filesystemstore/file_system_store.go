package filesystemstore

import (
	"encoding/json"
	"golang_selfstudy/webApp/league"
	"golang_selfstudy/webApp/player"
	"io"
	"os"
)

type FileSystemPlayerStore struct {
	database io.Writer
	league   league.GroupOfPlayers
}

type Tape struct {
	file *os.File
}

func (t *Tape) Write(p []byte) (n int, err error) {
	t.file.Truncate(0)
	t.file.Seek(0, 0)
	return t.file.Write(p)
}

func (t *Tape) SetFile(f *os.File) {
	t.file = f
}

func NewFileSystemPlayerStore(database *os.File) *FileSystemPlayerStore {
	database.Seek(0, 0)
	league, _ := league.NewLeague(database)
	return &FileSystemPlayerStore{
		database: &Tape{file: database},
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

	json.NewEncoder(f.database).Encode(f.league)
}
