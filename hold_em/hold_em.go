package holdem

import (
	"errors"
	"sort"
)

type Suit rune
type Rank int

var (
	ErrInvalidLength = errors.New("invalid card length")
	ErrInvalidRank   = errors.New("invalid rank")
	ErrInvalidSuit   = errors.New("invalid suit")
	ErrDuplicateCard = errors.New("duplicate card in hand")
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
	if len(cards) != 5 {
		return HighCard, ErrInvalidLength
	}
	for i := 0; i < len(cards); i++ {
		for j := i + 1; j < len(cards); j++ {
			if cards[i] == cards[j] {
				return HighCard, ErrDuplicateCard
			}
		}
	}

	sort.Slice(cards, func(i, j int) bool {
		return cards[i].Rank > cards[j].Rank // Descending order
	})

	isFlush := true
	isStraight := true

	// check for flush
	for i := 1; i < len(cards); i++ {
		if cards[i].Suit != cards[0].Suit {
			isFlush = false
			break
		}
	}

	// check for consecutive ranks (straigth)
	if cards[0].Rank == 14 && cards[1].Rank == 5 && cards[2].Rank == 4 && cards[3].Rank == 3 && cards[4].Rank == 2 {
		// Ace-low straight (A-5-4-3-2)
		isStraight = true
	} else {
		for j := 1; j < len(cards); j++ {
			if cards[j].Rank != cards[j-1].Rank-1 {
				isStraight = false
				break
			}
		}
	}

	// map to count occurrences of each rank
	counts := make(map[Rank]int)
	for _, card := range cards {
		counts[card.Rank]++
	}
	pairsCount := 0
	threeOfAKindCount := 0
	fourOfAKindCount := 0

	// count pairs, three of a kind, and four of a kind
	for _, count := range counts {
		switch count {
		case 2:
			pairsCount++
		case 3:
			threeOfAKindCount++
		case 4:
			fourOfAKindCount++
		}
	}

	if isStraight && isFlush {
		return StraightFlush, nil
	}
	if fourOfAKindCount == 1 {
		return FourOfAKind, nil
	}
	if threeOfAKindCount == 1 && pairsCount == 1 {
		return FullHouse, nil
	}
	if isFlush {
		return Flush, nil
	}
	if isStraight {
		return Straight, nil
	}
	if threeOfAKindCount == 1 {
		return ThreeOfAKind, nil
	}
	if pairsCount == 2 {
		return TwoPair, nil
	}
	if pairsCount == 1 {
		return OnePair, nil
	}

	return HighCard, nil
}
