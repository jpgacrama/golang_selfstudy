package poker_test

import (
	"encoding/json"
	"golang_selfstudy/webApp"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	database, cleanDatabase := createTempFile(t, "")
	defer cleanDatabase()
	store := &poker.FileSystemPlayerStore{}
	tape := poker.Tape{}
	tape.SetFile(database)
	store.SetDatabase(json.NewEncoder(&tape))

	server := poker.NewPlayerServer(store)
	singlePlayer := "Pepper"

	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(singlePlayer))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(singlePlayer))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(singlePlayer))

	t.Run("get score", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, newGetScoreRequest(singlePlayer))
		assertStatus(t, response.Code, http.StatusOK)

		assertResponseBody(t, response.Body.String(), "3")
	})

	t.Run("get league", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, newLeagueRequest())
		assertStatus(t, response.Code, http.StatusOK)

		got := getLeagueFromResponse(t, response.Body)
		want := poker.GroupOfPlayers{
			{Name: "Pepper", Wins: 3},
		}
		assertLeague(t, got, want)
	})
}
