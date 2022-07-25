package playerstore

type PlayerStore interface {
	GetPlayerScore(name string) int
	RecordWin(name string)
}

type InMemoryPlayerStore struct{}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return 123
}

func (i *InMemoryPlayerStore) RecordWin(name string) {}
