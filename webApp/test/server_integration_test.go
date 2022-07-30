package poker_test

import (
	"golang_selfstudy/webApp"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	database, cleanDatabase := CreateTempFile(t, "")
	defer cleanDatabase()
	store, err := poker.NewFileSystemPlayerStore(database)
	AssertNoError(t, err)

	server := poker.NewPlayerServer(store)
	singlePlayer := "Pepper"

	server.ServeHTTP(httptest.NewRecorder(), NewPostWinRequest(singlePlayer))
	server.ServeHTTP(httptest.NewRecorder(), NewPostWinRequest(singlePlayer))
	server.ServeHTTP(httptest.NewRecorder(), NewPostWinRequest(singlePlayer))

	t.Run("get score", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, NewGetScoreRequest(singlePlayer))
		AssertStatus(t, response.Code, http.StatusOK)

		AssertResponseBody(t, response.Body.String(), "3")
	})

	t.Run("get league", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, NewLeagueRequest())
		AssertStatus(t, response.Code, http.StatusOK)

		got := GetLeagueFromResponse(t, response.Body)
		want := poker.GroupOfPlayers{
			{Name: "Pepper", Wins: 3},
		}
		AssertLeague(t, got, want)
	})
}
