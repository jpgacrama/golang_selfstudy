package poker_test

import (
	"golang_selfstudy/webApp/src"
	"testing"
)

func TestFileSystemStore(t *testing.T) {
	t.Run("get player score", func(t *testing.T) {
		database, cleanDatabase := CreateTempFile(t, `[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 33}]`)
		defer cleanDatabase()

		store, err := poker.NewFileSystemPlayerStore(database)
		AssertNoError(t, err)

		got := store.GetPlayerScore("Chris")
		want := 33
		AssertScoreEquals(t, got, want)
	})
	t.Run("store wins for existing players", func(t *testing.T) {
		database, cleanDatabase := CreateTempFile(t, `[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 33}]`)
		defer cleanDatabase()
		store, err := poker.NewFileSystemPlayerStore(database)
		AssertNoError(t, err)
		store.RecordWin("Chris")

		got := store.GetPlayerScore("Chris")
		want := 34
		AssertScoreEquals(t, got, want)
	})
	t.Run("store wins for new players", func(t *testing.T) {
		database, cleanDatabase := CreateTempFile(t, `[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 33}]`)
		defer cleanDatabase()

		store, err := poker.NewFileSystemPlayerStore(database)
		AssertNoError(t, err)
		store.RecordWin("Pepper")

		got := store.GetPlayerScore("Pepper")
		want := 1
		AssertScoreEquals(t, got, want)
	})
	t.Run("works with an empty file", func(t *testing.T) {
		database, cleanDatabase := CreateTempFile(t, "")
		defer cleanDatabase()
		_, err := poker.NewFileSystemPlayerStore(database)
		AssertNoError(t, err)
	})
	t.Run("league sorted", func(t *testing.T) {
		database, cleanDatabase := CreateTempFile(t, `[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 33}]`)
		defer cleanDatabase()
		store, err := poker.NewFileSystemPlayerStore(database)
		AssertNoError(t, err)

		got := store.GetLeague()
		want := []poker.Player{
			{Name: "Chris", Wins: 33},
			{Name: "Cleo", Wins: 10},
		}

		AssertLeague(t, got, want)

		// read again
		got = store.GetLeague()
		AssertLeague(t, got, want)
	})
}
