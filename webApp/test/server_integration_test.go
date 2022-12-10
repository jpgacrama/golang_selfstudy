package poker_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"webApp/src"
)

func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	database, cleanDatabase := CreateTempFile(t, "")
	defer cleanDatabase()
	store, err := poker.NewFileSystemPlayerStore(database)
	AssertNoError(t, err)

	dummyGame := &poker.GameSpy{}
	server, err := poker.NewPlayerServer(store, dummyGame)
	if err != nil {
		t.Fatalf("cannot create NewPlayerServer")
	}
	singlePlayer := "Pepper"

	server.ServeHTTP(httptest.NewRecorder(), NewPostWinRequest(singlePlayer))
	server.ServeHTTP(httptest.NewRecorder(), NewPostWinRequest(singlePlayer))
	server.ServeHTTP(httptest.NewRecorder(), NewPostWinRequest(singlePlayer))

	t.Run("get score", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, NewGetScoreRequest(singlePlayer))
		AssertStatus(t, response, http.StatusOK)

		AssertResponseBody(t, response.Body.String(), "3")
	})

	t.Run("get league", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, NewLeagueRequest())
		AssertStatus(t, response, http.StatusOK)

		got := GetLeagueFromResponse(t, response.Body)
		want := poker.GroupOfPlayers{
			{Name: "Pepper", Wins: 3},
		}
		AssertLeague(t, got, want)
	})
}
