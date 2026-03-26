package holdem

import (
	"errors"
)

type Suit rune
type Rank int

var (
	ErrInvalidLength = errors.New("invalid card length")
	ErrInvalidRank   = errors.New("invalid rank")
	ErrInvalidSuit   = errors.New("invalid suit")
)

type Card struct {
	Rank Rank
	Suit Suit
}

const (
	Hearts   Suit = 'h'
	Diamonds Suit = 'd'
	Clubs    Suit = 'c'
	Spades   Suit = 's'
)

type HandCategory int

const (
	HighCard HandCategory = iota
	OnePair
	TwoPair
	ThreeOfAKind
	Straight
	Flush
	FullHouse
	FourOfAKind
	StraightFlush
)

// ParseCard takes a string like "Ah" or "Ts" and returns a Card.
func ParseCard(s string) (Card, error) {
	if len(s) != 2 {
		return Card{}, ErrInvalidLength
	}

	rankByte := s[0] // rank
	suitByte := s[1] // suit

	var rank Rank

	switch rankByte {
	case 'A':
		rank = 14
	case 'K':
		rank = 13
	case 'Q':
		rank = 12
	case 'J':
		rank = 11
	case 'T':
		rank = 10
	case '9':
		rank = 9
	case '8':
		rank = 8
	case '7':
		rank = 7
	case '6':
		rank = 6
	case '5':
		rank = 5
	case '4':
		rank = 4
	case '3':
		rank = 3
	case '2':
		rank = 2
	default:
		return Card{}, ErrInvalidRank
	}

	switch suitByte {
	case 'h', 'H': // hearts
		return Card{Rank: rank, Suit: Hearts}, nil
	case 'd', 'D': // diamonds
		return Card{Rank: rank, Suit: Diamonds}, nil
	case 'c', 'C': // clubs
		return Card{Rank: rank, Suit: Clubs}, nil
	case 's', 'S': // spades
		return Card{Rank: rank, Suit: Spades}, nil
	default:
		return Card{}, ErrInvalidSuit
	}

}

func ParseHandCategory(cards []Card) (HandCategory, error) {
	panic("not implemented yet")
}
