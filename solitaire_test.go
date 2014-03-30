package solitaire

import (
	"testing"
)

func createCardList(size int) []uint8 {
	var d []uint8
	for i := uint8(0); i < uint8(size)-2; i++ {
		d = append(d, i)
	}
	d = append(d, jokerA)
	d = append(d, jokerB)
	return d
}

func slicesEqual(a []uint8, b []uint8) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestNewDeck(t *testing.T) {
	d := NewDeck("")
	for i := range d {
		if i > 0 {
			if d[i] != d[i-1]+1 {
				t.Error("Expected ordered deck, got", d)
			}
		}
	}
}

func TestMoveJoker(t *testing.T) {
	var d Deck
	d = []uint8{1, 2, 3, 4, jokerA, jokerB}
	d = moveJoker(d, jokerA)
	if !slicesEqual(d, []uint8{1, 2, 3, 4, jokerB, jokerA}) {
		t.Error("Expected {[1, 2, 3, 4, 53, 52]}, got ", d)
	}
	d = moveJoker(d, jokerA)
	if !slicesEqual(d, []uint8{1, jokerA, 2, 3, 4, jokerB}) {
		t.Error("Expected {[1, 52, 2, 3, 4, 53]}, got ", d)
	}
}

func TestTripleCut(t *testing.T) {
	var d Deck
	d = []uint8{2, 4, 6, jokerB, 5, 8, 7, 1, jokerA, 3, 9}
	expect := []uint8{3, 9, jokerB, 5, 8, 7, 1, jokerA, 3, 9}
	d = tripleCut(d)
	if !slicesEqual(d, expect) {
		t.Error("Got ", d)
	}
}

func TestCountedCut(t *testing.T) {
	var d Deck
	d = createCardList(54)
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
