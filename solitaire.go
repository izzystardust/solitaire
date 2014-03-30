/*
Package solitaire implements Bruce Schneier's Solitaire cipher, as seen in Cryptonomicon.
*/
package solitaire

import (
	"fmt"
	"strings"
)

// A Deck is a deck of cards
type Deck []uint8

const (
	jokerA = 52
	jokerB = 53
)

// NewDeck creates an initialized deck keyed by the string key..
// The empty string, "", creates an ordered deck.
func NewDeck(key string) Deck {
	var d Deck
	key = strings.ToUpper(key)

	for i := uint8(0); i < 54; i++ {
		d = append(d, i)
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
	if key == 53 || key == 'A'+53 {
		return d
	}
	if key >= 'A' {
		d = append(d[key-'A'+1:], d[:key-'A'+1]...)
	} else {
		d = append(d[key:], d[:key]...)
	}
	return d
}

func doDeckMove(d Deck) Deck {
	moveJoker(d, jokerA)
	moveJoker(d, jokerB)
	moveJoker(d, jokerB)
	tripleCut(d)
	return d
}

func isJoker(c uint8) bool {
	return (c == jokerA || c == jokerB)
}

// moves the specified joker one card down.
// treats the deck as a loop (the last card is the first card).
func moveJoker(d Deck, j uint8) Deck {
	// get index of joker
	i := 0
	for {
		if d[i] == j {
			break
		}
		i++
	}

	if i+1 < len(d) {
		d[i], d[i+1] = d[i+1], d[i]
	} else {
		d = append(d[i:], d[:i]...)
		d = moveJoker(d, j)
	}
	return d
}

func tripleCut(d Deck) Deck {
	var ji1, ji2 int
	var j1, j2 uint8

	// find jokers
	for {
		if isJoker(d[ji1]) {
			j1 = d[ji1]
			break
		}
		ji1++
	}

	ji2 = ji1 + 1
	for {
		if isJoker(d[ji2]) {
			j2 = d[ji2]
			break
		}
		ji2++
	}
	first := append([]uint8{j2}, d[:ji1]...)
	fmt.Println(first)
	middle := d[ji1+1 : ji2]
	fmt.Println(middle)
	last := append(d[ji2+1:], j1)
	fmt.Println(last)
	d = append(last, middle...)
	d = append(d, first...)

	return d
}
