package payload_test

import (
	"osexec/src"
	"strings"
	"testing"
)

func TestGetDataIntegration(t *testing.T) {
	got := osexec.GetData(osexec.GetXMLFromCommand())
	want := "HAPPY NEW YEAR!"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestGetData(t *testing.T) {
	input := strings.NewReader(`
	<payload>
    <message>Cats are the best animal</message>
	</payload>`)

	got := osexec.GetData(input)
	want := "CATS ARE THE BEST ANIMAL"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
