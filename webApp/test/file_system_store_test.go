package webApp_test

import (
	"golang_selfstudy/webApp/player"
	"golang_selfstudy/webApp/playerstore"
	"strings"
	"testing"
)

func TestFileSystemStore(t *testing.T) {
	t.Run("league from a reader", func(t *testing.T) {
		database := strings.NewReader(`[
            {"Name": "Cleo", "Wins": 10},
            {"Name": "Chris", "Wins": 33}]`)

		store := playerstore.FileSystemPlayerStore{Database: database}
		got := store.GetLeague()
		want := []player.Player{
			{Name: "Cleo", Wins: 10},
			{Name: "Chris", Wins: 33},
		}
		assertLeague(t, got, want)
	})
}
