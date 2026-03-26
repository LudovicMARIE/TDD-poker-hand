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

func TestCompareHands(t *testing.T) {
	tests := []struct {
		name     string
		handA    holdem.EvaluatedHand
		handB    holdem.EvaluatedHand
		expected int // 1 for A wins, -1 for B wins, 0 for tie
	}{
		{
			name:     "Different Categories: Flush beats One Pair",
			handA:    holdem.EvaluatedHand{Category: holdem.Flush},
			handB:    holdem.EvaluatedHand{Category: holdem.OnePair},
			expected: 1,
		},
		{
			name:     "Different Categories: High Card loses to Straight",
			handA:    holdem.EvaluatedHand{Category: holdem.HighCard},
			handB:    holdem.EvaluatedHand{Category: holdem.Straight},
			expected: -1,
		},
		{
			name: "High Card Tie-Breaker: A-high beats K-high",
			handA: holdem.EvaluatedHand{
				Category: holdem.HighCard,
				Cards:    []holdem.Card{{Rank: 14}, {Rank: 10}, {Rank: 9}, {Rank: 5}, {Rank: 2}},
			},
			handB: holdem.EvaluatedHand{
				Category: holdem.HighCard,
				Cards:    []holdem.Card{{Rank: 13}, {Rank: 12}, {Rank: 11}, {Rank: 9}, {Rank: 8}},
			},
			expected: 1,
		},
		{
			name: "High Card Tie-Breaker: Deep tie-break (3rd card)",
			handA: holdem.EvaluatedHand{
				Category: holdem.HighCard,
				Cards:    []holdem.Card{{Rank: 14}, {Rank: 13}, {Rank: 9}, {Rank: 5}, {Rank: 2}},
			},
			handB: holdem.EvaluatedHand{
				Category: holdem.HighCard,
				Cards:    []holdem.Card{{Rank: 14}, {Rank: 13}, {Rank: 8}, {Rank: 7}, {Rank: 6}},
			},
			expected: 1, // 9 beats 8
		},
		{
			name: "Absolute Tie",
			handA: holdem.EvaluatedHand{
				Category: holdem.HighCard,
				Cards:    []holdem.Card{{Rank: 14}, {Rank: 13}, {Rank: 9}, {Rank: 5}, {Rank: 2}},
			},
			handB: holdem.EvaluatedHand{
				Category: holdem.HighCard,
				Cards:    []holdem.Card{{Rank: 14}, {Rank: 13}, {Rank: 9}, {Rank: 5}, {Rank: 2}},
			},
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := holdem.CompareHands(tt.handA, tt.handB)
			assert.Equal(t, tt.expected, result, "Comparison result mismatch")
		})
	}
}

func TestTexasHoldem(t *testing.T) {
	t.Run("Example E - Quads on board kicker decides", func(t *testing.T) {
		board := parseCards(t, []string{"7s", "7h", "7d", "7c", "4s"})
		p1Hole := parseCards(t, []string{"As", "Ks"})
		p2Hole := parseCards(t, []string{"Qs", "Js"})

		b1, _ := holdem.GetBest5From7(append(board, p1Hole...))
		b2, _ := holdem.GetBest5From7(append(board, p2Hole...))

		results := []holdem.PlayerResult{
			{PlayerID: "Player1", BestHand: b1},
			{PlayerID: "Player2", BestHand: b2},
		}

		winners := holdem.DetermineWinners(results)
		assert.Equal(t, []string{"Player1"}, winners)
		assert.Equal(t, holdem.FourOfAKind, b1.Category)
	})

	t.Run("Example D - Board plays (Tie)", func(t *testing.T) {
		board := parseCards(t, []string{"5s", "6h", "7d", "8c", "9s"})
		p1Hole := parseCards(t, []string{"2s", "2h"})
		p2Hole := parseCards(t, []string{"3s", "3h"})

		b1, _ := holdem.GetBest5From7(append(board, p1Hole...))
		b2, _ := holdem.GetBest5From7(append(board, p2Hole...))

		winners := holdem.DetermineWinners([]holdem.PlayerResult{
			{PlayerID: "P1", BestHand: b1},
			{PlayerID: "P2", BestHand: b2},
		})

		assert.Len(t, winners, 2)
		assert.Contains(t, winners, "P1")
		assert.Contains(t, winners, "P2")
	})
}

func parseCards(t *testing.T, inputs []string) []holdem.Card {
	res := make([]holdem.Card, len(inputs))
	for i, s := range inputs {
		c, err := holdem.ParseCard(s)
		require.NoError(t, err)
		res[i] = c
	}
	return res
}
