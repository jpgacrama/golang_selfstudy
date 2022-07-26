package webApp_test

import (
	"golang_selfstudy/webApp/filesystemstore"
	"golang_selfstudy/webApp/player"
	"strings"
	"testing"
)

func TestFileSystemStore(t *testing.T) {
	t.Run("league from a reader", func(t *testing.T) {
		database := strings.NewReader(`[
            {"Name": "Cleo", "Wins": 10},
            {"Name": "Chris", "Wins": 33}]`)

		store := filesystemstore.FileSystemPlayerStore{}
		store.SetDatabase(database)
		got := store.GetLeague()
		want := []player.Player{
			{Name: "Cleo", Wins: 10},
			{Name: "Chris", Wins: 33},
		}
		assertLeague(t, got, want)
	})
}
