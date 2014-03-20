package solitaire

import (
	"fmt"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := NewDeck()
	fmt.Println(d)
}

func TestEncrypt(t *testing.T) {
	pt := "hello, world"
	var deck []int
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
