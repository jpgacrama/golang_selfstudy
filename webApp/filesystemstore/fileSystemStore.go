package filesystemstore

import (
	"encoding/json"
	"fmt"
	"golang_selfstudy/webApp/league"
	"golang_selfstudy/webApp/player"
	"os"
	"sort"
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

func initialisePlayerDBFile(file *os.File) error {
	file.Seek(0, 0)
	info, err := file.Stat()
	if err != nil {
		return fmt.Errorf("problem getting file info from file %s, %v", file.Name(), err)
	}

	if info.Size() == 0 {
		file.Write([]byte("[]"))
		file.Seek(0, 0)
	}
	return nil
}

func NewFileSystemPlayerStore(file *os.File) (*FileSystemPlayerStore, error) {
	err := initialisePlayerDBFile(file)
	if err != nil {
		return nil, fmt.Errorf("problem initialising player db file, %v", err)
	}

	league, err := league.NewLeague(file)
	if err != nil {
		return nil, fmt.Errorf("problem loading player store from file %s, %v", file.Name(), err)
	}

	return &FileSystemPlayerStore{
		database: json.NewEncoder(&Tape{file}),
		league:   league,
	}, nil
}

func (f *FileSystemPlayerStore) SetDatabase(d *json.Encoder) {
	f.database = d
}

func (f *FileSystemPlayerStore) GetLeague() league.GroupOfPlayers {
	sort.Slice(f.league, func(i, j int) bool {
		return f.league[i].Wins > f.league[j].Wins
	})
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
