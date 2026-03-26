package main

import (
	"fmt"
	"math/rand"
	"time"

	holdem "github.com/LudovicMARIE/TDD-poker-hand/hold_em"
)

func main() {
	// 1. Initialize random seed
	rand.Seed(time.Now().UnixNano())

	// 2. Create and shuffle a deck of 52 cards
	deck := createDeck()
	rand.Shuffle(len(deck), func(i, j int) {
		deck[i], deck[j] = deck[j], deck[i]
	})

	// 3. Deal cards
	// Texas Hold'em: 2 cards per player, 5 for the board
	p1Hole := deck[0:2]
	p2Hole := deck[2:4]
	board := deck[4:9]

	fmt.Println("--- Texas Hold'em Reveal ---")
	fmt.Printf("Board:    %v\n", formatCards(board))
	fmt.Printf("Player 1: %v\n", formatCards(p1Hole))
	fmt.Printf("Player 2: %v\n", formatCards(p2Hole))
	fmt.Println("----------------------------")

	// 4. Determine best 5-card hand for each player from their 7 available cards
	// Requirement: 5 board cards + 2 hole cards [cite: 8, 9, 11]
	p1Best, _ := holdem.GetBest5From7(append(board, p1Hole...))
	p2Best, _ := holdem.GetBest5From7(append(board, p2Hole...))

	// 5. Wrap results for the comparison engine
	results := []holdem.PlayerResult{
		{PlayerID: "Player 1", BestHand: p1Best},
		{PlayerID: "Player 2", BestHand: p2Best},
	}

	// 6. Determine the winner(s) (supports ties/split pots) [cite: 12, 21]
	winners := holdem.DetermineWinners(results)

	// 7. Display detailed output
	for _, res := range results {
		fmt.Printf("%s Category: %s\n", res.PlayerID, getCategoryName(res.BestHand.Category))
		fmt.Printf("%s Best 5:   %v\n", res.PlayerID, formatCards(res.BestHand.Cards))
	}

	fmt.Println("----------------------------")
	if len(winners) > 1 {
		fmt.Printf("Result: Split Pot between %v\n", winners)
	} else {
		fmt.Printf("Result: %s Wins!\n", winners[0])
	}
}

// Helper to create a standard 52-card deck
func createDeck() []holdem.Card {
	suits := []holdem.Suit{holdem.Hearts, holdem.Diamonds, holdem.Clubs, holdem.Spades}
	deck := make([]holdem.Card, 0, 52)
	for _, s := range suits {
		for r := 2; r <= 14; r++ {
			deck = append(deck, holdem.Card{Rank: holdem.Rank(r), Suit: s})
		}
	}
	return deck
}

// Helper to make card printing look pretty
func formatCards(cards []holdem.Card) string {
	rankMap := map[holdem.Rank]string{
		14: "A", 13: "K", 12: "Q", 11: "J", 10: "T",
	}
	res := ""
	for _, c := range cards {
		rStr := fmt.Sprintf("%d", c.Rank)
		if val, ok := rankMap[c.Rank]; ok {
			rStr = val
		}
		res += fmt.Sprintf("[%s%c] ", rStr, c.Suit)
	}
	return res
}

// Helper to map category iotas to strings [cite: 24]
func getCategoryName(cat holdem.HandCategory) string {
	names := []string{
		"High Card", "One Pair", "Two Pair", "Three of a Kind",
		"Straight", "Flush", "Full House", "Four of a Kind", "Straight Flush",
	}
	if int(cat) < len(names) {
		return names[cat]
	}
	return "Unknown"
}
