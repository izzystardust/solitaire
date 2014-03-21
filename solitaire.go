package solitaire

import (
	"strings"
)

type Deck struct {
	cards []uint8
}

const (
	JokerA = 52
	JokerB = 53
)

func NewDeck(key string) Deck {
	var d Deck
	key = strings.ToUpper(key)

	for i := uint8(0); i < 54; i++ {
		d.cards = append(d.cards, i)
	}

	for i := 0; i < len(key); i++ {
		d = countedCut(d, key[i])
	}
	return d
}

func Encrypt(plaintext string, deck Deck) (string, error) {
	return "", nil
}

func Decrypt(ciphertext string, deck Deck) (string, error) {
	return "", nil
}

func countedCut(d Deck, key uint8) Deck {
	if key >= 'A' {
		d.cards = append(d.cards[key-'A'+1:], d.cards[:key-'A'+1]...)
	} else {
		d.cards = append(d.cards[key:], d.cards[:key]...)
	}
	return d
}

func doDeckMove(d Deck) Deck {
	moveJoker(d, JokerA)
	moveJoker(d, JokerB)
	moveJoker(d, JokerB)
	tripleCut(d)
	return d
}

func isJoker(c uint8) bool {
	return (c == JokerA || c == JokerB)
}

// moves the specified joker one card down.
// treats the deck as a loop (the last card is the first card).
func moveJoker(d Deck, j uint8) Deck {
	// get index of joker
	i := 0
	for {
		if d.cards[i] == j {
			break
		}
		i++
	}

	if i+1 < len(d.cards) {
		d.cards[i], d.cards[i+1] = d.cards[i+1], d.cards[i]
	} else {
		d.cards = append(d.cards[i:], d.cards[:i]...)
		d = moveJoker(d, j)
	}
	return d
}

func tripleCut(d Deck) Deck {
	//TODO
	return d
}
