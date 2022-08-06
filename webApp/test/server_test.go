package poker_test

import (
	"github.com/gorilla/websocket"
	"golang_selfstudy/webApp/src"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestGETPlayers(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{
			"Pepper": 20,
			"Floyd":  10,
		},
		nil,
		nil,
	}
	dummyGame := &poker.GameSpy{}
	server := mustMakePlayerServer(t, &store, dummyGame)
	t.Run("returns Pepper's score", func(t *testing.T) {
		request := NewGetScoreRequest("Pepper")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		AssertStatus(t, response, http.StatusOK)
		AssertResponseBody(t, response.Body.String(), "20")
	})

	t.Run("returns Floyd's score", func(t *testing.T) {
		request := NewGetScoreRequest("Floyd")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		AssertStatus(t, response, http.StatusOK)
		AssertResponseBody(t, response.Body.String(), "10")
	})

	t.Run("returns 404 on missing players", func(t *testing.T) {
		request := NewGetScoreRequest("Apollo")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)
		AssertStatus(t, response, http.StatusNotFound)
	})
}

func TestStoreWins(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{},
		nil,
		nil,
	}
	dummyGame := &poker.GameSpy{}
	server := mustMakePlayerServer(t, &store, dummyGame)
	t.Run("it records wins on POST", func(t *testing.T) {
		player := "Pepper"
		request := NewPostWinRequest(player)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)
		AssertStatus(t, response, http.StatusAccepted)

		if len(store.winners) != 1 {
			t.Fatalf("got %d calls to RecordWin want %d", len(store.winners), 1)
		}

		if store.winners[0] != player {
			t.Errorf("did not store correct winner got %q want %q", store.winners[0], player)
		}
	})
}

func TestLeague(t *testing.T) {
	t.Run("it returns the league table as JSON", func(t *testing.T) {
		wantedLeague := poker.GroupOfPlayers{
			{Name: "Cleo", Wins: 32},
			{Name: "Chris", Wins: 20},
			{Name: "Tiest", Wins: 14},
		}

		store := StubPlayerStore{nil, nil, wantedLeague}
		dummyGame := &poker.GameSpy{}
		server := mustMakePlayerServer(t, &store, dummyGame)
		request := NewLeagueRequest()
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		got := GetLeagueFromResponse(t, response.Body)
		AssertStatus(t, response, http.StatusOK)
		AssertLeague(t, got, wantedLeague)
		AssertContentType(t, response, poker.JsonContentType)
	})
}

func TestGame(t *testing.T) {
	t.Run("GET game returns 200", func(t *testing.T) {
		dummyGame := &poker.GameSpy{}
		server := mustMakePlayerServer(t, &StubPlayerStore{}, dummyGame)
		request, err := NewGameRequest()
		if err != nil {
			t.Fatalf("cannot create a new game request")
		}
		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)
		AssertStatus(t, response, http.StatusOK)
	})
	t.Run("start a game with 3 players and declare Ruth the winner", func(t *testing.T) {
		dummyPlayerStore := &StubPlayerStore{}
		game := &poker.GameSpy{}
		winner := "Ruth"
		server := httptest.NewServer(mustMakePlayerServer(t, dummyPlayerStore, game))
		ws := mustDialWS(t, "ws"+strings.TrimPrefix(server.URL, "http")+"/ws")

		defer server.Close()
		defer ws.Close()

		writeWSMessage(t, ws, "3")
		writeWSMessage(t, ws, winner)

		// This is BAD. I should just have an async definition here
		time.Sleep(10 * time.Millisecond)
		assertGameStartedWith(t, game, 3)
		assertFinishCalledWith(t, game, winner)
	})
}

func mustMakePlayerServer(t *testing.T, store poker.PlayerStore, game poker.Game) *poker.PlayerServer {
	t.Helper()
	server, err := poker.NewPlayerServer(store, game)
	if err != nil {
		t.Fatal("problem creating player server", err)
	}
	return server
}

// Close must be invoked by the calling function
func mustDialWS(t *testing.T, url string) *websocket.Conn {
	ws, _, err := websocket.DefaultDialer.Dial(url, nil)

	if err != nil {
		t.Fatalf("could not open a ws connection on %s %v", url, err)
	}

	return ws
}

func writeWSMessage(t testing.TB, conn *websocket.Conn, message string) {
	t.Helper()
	if err := conn.WriteMessage(websocket.TextMessage, []byte(message)); err != nil {
		t.Fatalf("could not send message over ws connection %v", err)
	}
}
