package solitaire

import (
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := NewDeck("")
	for i := range d.cards {
		if i > 0 {
			if d.cards[i] != d.cards[i-1]+1 {
				t.Error("Expected ordered deck, got", d)
			}
		}
	}
}

func TestEncrypt(t *testing.T) {
	pt := "hello, world"
	var deck Deck
	var ct string
	ct, err := Encrypt(pt, deck)
	if err != nil {
		t.Error("Unexpected error: ", err)
	} else if ct != "hello, world" {
		t.Error("Got ", ct)
	}
}

func TestDecrypt(t *testing.T) {

}
