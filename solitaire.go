package solitaire

import (
	"strings"
)

type Deck [54]int

func NewDeck(key string) Deck {
	var d Deck
	key = strings.ToUpper(key)
	for i := 0; i < 54; i++ {
		d[i] = i
	}
	for i := 0; i < len(key); i++ {
		d := countedCut(d, key[i]-'A')
	}
	return d
}

func countedCut(deck Deck, pos int) Deck {
	return nil
}

func Encrypt(plaintext string, deck []uint8) (string, error) {
	return "", nil
}

func Decrypt(ciphertext string, deck []uint8) (string, error) {
	return "", nil
}
