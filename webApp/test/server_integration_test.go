package webApp_test

import (
	"golang_selfstudy/webApp/player"
	"golang_selfstudy/webApp/playerstore"
	"golang_selfstudy/webApp/server"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	store := playerstore.NewInMemoryPlayerStore()
	server := server.NewPlayerServer(store)
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
		want := []player.Player{
			{Name: "Pepper", Wins: 3},
		}
		assertLeague(t, got, want)
	})
}
