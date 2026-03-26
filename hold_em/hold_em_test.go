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
		name        string
		handCards   []holdem.Card
		expectHand  holdem.HandCategory
		expectedErr bool
	}{
		{"Invalid hand size", []holdem.Card{{Rank: 3, Suit: holdem.Diamonds}}, holdem.HighCard, true},
		{"High Card", []holdem.Card{{Rank: 14, Suit: holdem.Hearts}, {Rank: 13, Suit: holdem.Diamonds}, {Rank: 12, Suit: holdem.Clubs}, {Rank: 11, Suit: holdem.Spades}, {Rank: 9, Suit: holdem.Hearts}}, holdem.HighCard, false},
		{"One Pair", []holdem.Card{{Rank: 14, Suit: holdem.Hearts}, {Rank: 12, Suit: holdem.Clubs}, {Rank: 11, Suit: holdem.Spades}, {Rank: 14, Suit: holdem.Diamonds}, {Rank: 9, Suit: holdem.Hearts}}, holdem.OnePair, false},
		{"Two Pair", []holdem.Card{{Rank: 14, Suit: holdem.Diamonds}, {Rank: 12, Suit: holdem.Clubs}, {Rank: 14, Suit: holdem.Hearts}, {Rank: 12, Suit: holdem.Spades}, {Rank: 9, Suit: holdem.Hearts}}, holdem.TwoPair, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handCategory, err := holdem.ParseHandCategory(tt.handCards)

			if tt.expectedErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tt.expectHand, handCategory, "Hand category mismatch")
			}
		})
	}
}
