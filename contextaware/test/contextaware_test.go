package contextaware_test

import (
	"strings"
	"testing"
)

func TestContextAwareReader(t *testing.T) {
	t.Run("lets just see how a normal reader works", func(t *testing.T) {
		rdr := strings.NewReader("123456")
		got := make([]byte, 3)
		_, err := rdr.Read(got)
		if err != nil {
			t.Fatal(err)
		}
		assertBufferHas(t, got, "123")
		_, err = rdr.Read(got)

		if err != nil {
			t.Fatal(err)
		}
		assertBufferHas(t, got, "456")
	})
}

func assertBufferHas(t testing.TB, buf []byte, want string) {
	t.Helper()
	got := string(buf)
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
