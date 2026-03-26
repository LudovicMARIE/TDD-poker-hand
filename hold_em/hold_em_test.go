package holdem_test

import (
	"testing"

	holdem "github.com/LudovicMARIE/TDD-poker-hand/hold_em"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseCard(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		wantRank    holdem.Rank
		wantSuit    holdem.Suit
		expectedErr error
	}{
		{"Invalid suit", "Ax", 0, 0, holdem.ErrInvalidSuit},
		{"Invalid rank", "Zs", 0, 0, holdem.ErrInvalidRank},
		{"Invalid Size", "hbdfqsuhbdhqsuh", 0, 0, holdem.ErrInvalidLength},
		{"Ace of Hearts", "Ah", 14, 'h', nil},
		{"Ten of Spades", "Ts", 10, 's', nil},
		{"Two of Clubs", "2c", 2, 'c', nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			card, err := holdem.ParseCard(tt.input)

			if tt.expectedErr != nil {
				require.ErrorIs(t, err, tt.expectedErr)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tt.wantRank, card.Rank, "Rank mismatch")
				assert.Equal(t, tt.wantSuit, card.Suit, "Suit mismatch")
			}
		})
	}
}

func TestParseHandCategory(t *testing.T) {
	tests := []struct {
		name         string
		handCards    []holdem.Card
		expectedHand holdem.HandCategory
		expectedErr  error
	}{
		{"Invalid hand size", []holdem.Card{{Rank: 3, Suit: holdem.Diamonds}}, holdem.HighCard, holdem.ErrInvalidLength},
		{"Duplicate Cards", []holdem.Card{{Rank: 10, Suit: holdem.Hearts}, {Rank: 10, Suit: holdem.Hearts}, {Rank: 9, Suit: holdem.Clubs}, {Rank: 8, Suit: holdem.Spades}, {Rank: 2, Suit: holdem.Diamonds}}, holdem.HighCard, holdem.ErrDuplicateCard},
		{"High Card", []holdem.Card{{Rank: 14, Suit: holdem.Hearts}, {Rank: 13, Suit: holdem.Diamonds}, {Rank: 12, Suit: holdem.Clubs}, {Rank: 11, Suit: holdem.Spades}, {Rank: 9, Suit: holdem.Hearts}}, holdem.HighCard, nil},
		{"One Pair", []holdem.Card{{Rank: 14, Suit: holdem.Hearts}, {Rank: 12, Suit: holdem.Clubs}, {Rank: 11, Suit: holdem.Spades}, {Rank: 14, Suit: holdem.Diamonds}, {Rank: 9, Suit: holdem.Hearts}}, holdem.OnePair, nil},
		{"Two Pair", []holdem.Card{{Rank: 14, Suit: holdem.Diamonds}, {Rank: 12, Suit: holdem.Clubs}, {Rank: 14, Suit: holdem.Hearts}, {Rank: 12, Suit: holdem.Spades}, {Rank: 9, Suit: holdem.Hearts}}, holdem.TwoPair, nil},
		{"Three of a Kind", []holdem.Card{{Rank: 8, Suit: holdem.Hearts}, {Rank: 8, Suit: holdem.Diamonds}, {Rank: 13, Suit: holdem.Spades}, {Rank: 8, Suit: holdem.Clubs}, {Rank: 2, Suit: holdem.Hearts}}, holdem.ThreeOfAKind, nil},
		{"Full House", []holdem.Card{{Rank: 10, Suit: holdem.Hearts}, {Rank: 10, Suit: holdem.Diamonds}, {Rank: 10, Suit: holdem.Clubs}, {Rank: 4, Suit: holdem.Spades}, {Rank: 4, Suit: holdem.Hearts}}, holdem.FullHouse, nil},
		{"Four of a Kind", []holdem.Card{{Rank: 9, Suit: holdem.Hearts}, {Rank: 9, Suit: holdem.Diamonds}, {Rank: 9, Suit: holdem.Clubs}, {Rank: 2, Suit: holdem.Hearts}, {Rank: 9, Suit: holdem.Spades}}, holdem.FourOfAKind, nil},
		{"Flush", []holdem.Card{{Rank: 14, Suit: holdem.Hearts}, {Rank: 11, Suit: holdem.Hearts}, {Rank: 9, Suit: holdem.Hearts}, {Rank: 6, Suit: holdem.Hearts}, {Rank: 4, Suit: holdem.Hearts}}, holdem.Flush, nil},
		{"Straight (Standard)", []holdem.Card{{Rank: 9, Suit: holdem.Hearts}, {Rank: 8, Suit: holdem.Diamonds}, {Rank: 7, Suit: holdem.Clubs}, {Rank: 6, Suit: holdem.Spades}, {Rank: 5, Suit: holdem.Hearts}}, holdem.Straight, nil},
		{"Straight (Ace-Low Wheel)", []holdem.Card{{Rank: 14, Suit: holdem.Hearts}, {Rank: 4, Suit: holdem.Clubs}, {Rank: 3, Suit: holdem.Spades}, {Rank: 5, Suit: holdem.Diamonds}, {Rank: 2, Suit: holdem.Hearts}}, holdem.Straight, nil},
		{"Straight Flush", []holdem.Card{{Rank: 11, Suit: holdem.Hearts}, {Rank: 10, Suit: holdem.Hearts}, {Rank: 8, Suit: holdem.Hearts}, {Rank: 9, Suit: holdem.Hearts}, {Rank: 7, Suit: holdem.Hearts}}, holdem.StraightFlush, nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handCategory, err := holdem.ParseHandCategory(tt.handCards)

			if tt.expectedErr != nil {
				require.ErrorIs(t, err, tt.expectedErr)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tt.expectedHand, handCategory, "Hand category mismatch")
			}
		})
	}
}
