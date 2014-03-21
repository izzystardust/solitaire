package solitaire

import (
	"strings"
)

type Deck struct {
	cards []uint8
}

func NewDeck(key string) Deck {
	var d Deck
	key = strings.ToUpper(key)

	for i := uint8(0); i < 54; i++ {
		d.cards = append(d.cards, i)
	}

	for i := 0; i < len(key); i++ {
		// counted cut
		d.cards = append(d.cards[key[i]-'A':], d.cards[:key[i]-'A']...)
	}
	return d
}

func Encrypt(plaintext string, deck Deck) (string, error) {
	return "", nil
}

func Decrypt(ciphertext string, deck Deck) (string, error) {
	return "", nil
}
