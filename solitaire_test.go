package solitaire

import (
	"testing"
	"fmt"
)

func createCardList() []uint8 {
	var d []uint8
	for i:= uint8(0); i < 54; i++ {
		d = append(d, i)
	}
	return d
}

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

func TestCountedCut(t *testing.T) {
	var d Deck
	d.cards = createCardList()
	d = countedCut(d, 'A')
	fmt.Println(d)
	fmt.Println(countedCut(d, 'Z'))
}

func TestEncrypt(t *testing.T) {
//	pt := "hello, world"
//	var deck Deck
//	var ct string
//	ct, err := Encrypt(pt, deck)
//	if err != nil {
//		t.Error("Unexpected error: ", err)
//	} else if ct != "hello, world" {
//		t.Error("Got ", ct)
//	}
}

func TestDecrypt(t *testing.T) {

}
