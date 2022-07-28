package filesystemstore

import (
	"encoding/json"
	"golang_selfstudy/webApp/league"
	"golang_selfstudy/webApp/player"
	"os"
)

type FileSystemPlayerStore struct {
	database *json.Encoder
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

func NewFileSystemPlayerStore(file *os.File) *FileSystemPlayerStore {
	file.Seek(0, 0)
	league, _ := league.NewLeague(file)

	return &FileSystemPlayerStore{
		database: json.NewEncoder(&Tape{file}),
		league:   league,
	}
}

func (f *FileSystemPlayerStore) SetDatabase(d *json.Encoder) {
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
	f.database.Encode(f.league)
}
