package holdem_test

import (
	"testing"

	holdem "github.com/LudovicMARIE/TDD-poker-hand/hold_em"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseCard(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		wantRank  holdem.Rank
		wantSuit  holdem.Suit
		expectErr bool
	}{
		{"Invalid Card", "Xx", 0, 0, true},
		{"Invalid Size", "hbdfqsuhbdhqsuh", 0, 0, true},
		{"Ace of Hearts", "Ah", 14, 'h', false},
		{"Ten of Spades", "Ts", 10, 's', false},
		{"Two of Clubs", "2c", 2, 'c', false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			card, err := holdem.ParseCard(tt.input)

			if tt.expectErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tt.wantRank, card.Rank, "Rank mismatch")
				assert.Equal(t, tt.wantSuit, card.Suit, "Suit mismatch")
			}
		})
	}
}
