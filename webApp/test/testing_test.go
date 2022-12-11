package poker_test

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"sort"
	"testing"
	"webApp/src"
)

type StubPlayerStore struct {
	scores  map[string]int
	winners []string
	league  []poker.Player
}

func AssertPlayerWin(t testing.TB, game *poker.TexasHoldem, winner string) {
	t.Helper()
	winnersList := game.GetStore().GetWinnerList()
	gotWinner := sort.SearchStrings(winnersList, winner)

	// This is only true if the string is NOT found
	if gotWinner == len(winnersList) {
		t.Errorf("did not store correct winner: got: %v, want: %v", winnersList, winner)
	}
}

func AssertPlayerWinUsingStore(t testing.TB, store poker.PlayerStore, winner string) {
	t.Helper()
	winnersList := store.GetWinnerList()
	gotWinner := sort.SearchStrings(winnersList, winner)

	// This is only true if the string is NOT found
	if gotWinner == len(winnersList) {
		t.Errorf("did not store correct winner: got: %v, want: %v", winnersList, winner)
	}
}

func AssertScoreEquals(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func AssertNoError(t testing.TB, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("didn't expect an error but got one, %v", err)
	}
}

func AssertStatus(t testing.TB, response *httptest.ResponseRecorder, want int) {
	t.Helper()
	got := response.Code
	if got != want {
		t.Errorf("did not get correct status, got %d, want %d", got, want)
	}
}

func AssertResponseBody(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("response body is wrong, got %q want %q", got, want)
	}
}

func AssertLeague(t testing.TB, got, want poker.GroupOfPlayers) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func AssertContentType(t testing.TB, response *httptest.ResponseRecorder, want string) {
	t.Helper()
	if response.Result().Header.Get("content-type") != want {
		t.Errorf("response did not have content-type of %s, got %v", want, response.Result().Header)
	}
}

// This is an Adapter pattern. Call the function returned from this one to close the file safely
func CreateTempFile(t testing.TB, initialData string) (*os.File, func()) {
	t.Helper()
	tmpfile, err := os.CreateTemp("", "db")
	if err != nil {
		t.Fatalf("could not create temp file %v", err)
	}

	tmpfile.Write([]byte(initialData))
	removeFile := func() {
		tmpfile.Close()
		os.Remove(tmpfile.Name())
	}

	return tmpfile, removeFile
}

func NewGetScoreRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", name), nil)
	return req
}

func NewPostWinRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/players/%s", name), nil)
	return req
}

func NewLeagueRequest() *http.Request {
	req, _ := http.NewRequest(http.MethodGet, "/league", nil)
	return req
}

func NewGameRequest() (*http.Request, error) {
	return http.NewRequest(http.MethodGet, "/game", nil)
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	score := s.scores[name]
	return score
}

func (s *StubPlayerStore) RecordWin(name string) {
	s.winners = append(s.winners, name)
	sort.Strings(s.winners)
}

func (s *StubPlayerStore) GetWinnerList() []string {
	return s.winners
}

func (s *StubPlayerStore) GetLeague() poker.GroupOfPlayers {
	return s.league
}

func GetLeagueFromResponse(t testing.TB, body io.Reader) (l poker.GroupOfPlayers) {
	t.Helper()
	newLeague, err := poker.NewLeague(body)
	if err != nil {
		t.Fatalf("Unable to parse response from server %q into slice of Player, '%v'", body, err)
	}
	return newLeague
}
