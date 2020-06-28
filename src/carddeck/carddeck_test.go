package carddeck

import "testing"

func TestNewdeck(t *testing.T) {
	d := generateNewDeck();
	if (len(d) != 52) {
		t.Error("Length must be 52. Cards are missing from the deck but got", len(d))
	}
}