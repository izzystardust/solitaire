/*
Package solitaire implements Bruce Schneier's Solitaire cipher, as seen in Cryptonomicon.
*/
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

// NewDeck creates an initialized deck keyed by the string key..
// The empty string, "", creates an ordered deck.
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

// Encrypt does what it says on the tin.
// Returns the encrypted string.
func Encrypt(plaintext string, deck Deck) string {
	//TODO
	return ""
}

// Decrypt does what is says on the tin.
// Returns the decrypted string.
func Decrypt(ciphertext string, deck Deck) string {
	//TODO
	return ""
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
	var ji1, ji2 int
	var j1, j2 uint8
	for {
		if isJoker(d.cards[ji1]) {
			j1 = d.cards[ji1]
			break
		}
		ji1++
	}

	ji2 = ji1 + 1
	for {
		if isJoker(d.cards[ji2]) {
			j2 = d.cards[ji2]
			break
		}
		ji2++
	}
	first := append([]uint8{j2}, d.cards[:ji1]...)
	middle := d.cards[ji1 + 1:ji2]
	last := append(d.cards[ji2+1:], j1)
	d.cards = append(last, middle...)
	d.cards = append(d.cards, first...)


	return d
}
