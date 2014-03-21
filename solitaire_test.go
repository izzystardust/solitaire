package solitaire

import (
	"testing"
)

func createCardList(size int) []uint8 {
	var d []uint8
	for i:= uint8(0); i < uint8(size)-2; i++ {
		d = append(d, i)
	}
	d = append(d, JokerA)
	d = append(d, JokerB)
	return d
}

func slicesEqual(a []uint8, b []uint8) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range(a) {
		if a[i] != b[i] {
			return false
		}
	}
	return true
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

func TestMoveJoker(t *testing.T) {
	var d Deck
	d.cards = []uint8{1,2,3,4,JokerA, JokerB}
	d = moveJoker(d, JokerA)
	if !slicesEqual(d.cards ,[]uint8{1,2,3,4,JokerB, JokerA}) {
		t.Error("Expected {[1, 2, 3, 4, 53, 52]}, got ", d)
	}
	d = moveJoker(d, JokerA)
	if !slicesEqual(d.cards, []uint8{1, JokerA, 2, 3, 4, JokerB}) {
		t.Error("Expected {[1, 52, 2, 3, 4, 53]}, got ", d)
	}
}

func TestTripleCut(t *testing.T) {
	var d Deck
	d.cards = []uint8{2, 4, 6, JokerB, 5, 8, 7, 1, JokerA, 3, 9}
	expect := []uint8{3, 9, JokerB, 5, 8, 7, 1, JokerA, 3, 9}
	d = tripleCut(d)
	if !slicesEqual(d.cards, expect) {
		t.Error("Got ", d)
	}
}

func TestCountedCut(t *testing.T) {
	var d Deck
	d.cards = createCardList(54)
	d = countedCut(d, 'A')
	//fmt.Println(d)
	//fmt.Println(countedCut(d, 'Z'))
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
