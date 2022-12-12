package payload_test

import (
	"osexec/src"
	"testing"
)

func TestGetDataIntegration(t *testing.T) {
	got := osexec.GetData(osexec.GetXMLFromCommand())
	want := "HAPPY NEW YEAR!"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
