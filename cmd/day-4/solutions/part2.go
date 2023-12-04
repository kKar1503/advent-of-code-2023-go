package solutions

import (
	"slices"
	"strconv"
)

func Part2(input []string) []byte {
	cards := splitIntoTwoSets(input)

	scratchCards := map[int]int{}

	for _, card := range cards {
		point := 0
		for _, winningNumber := range card.WinningNumbers {
			if slices.Contains(card.NumbersBought, winningNumber) {
				point++
			}
		}
		countOfCards := 1
		if count, ok := scratchCards[card.CardNumber]; ok {
			countOfCards += count
		}
		for i := 1; i <= point; i++ {
			if card2, ok := scratchCards[card.CardNumber+i]; ok {
				scratchCards[card.CardNumber+i] = card2 + countOfCards
			} else {
				scratchCards[card.CardNumber+i] = countOfCards
			}
		}
	}

	total := len(cards)

	for _, c := range scratchCards {
		total += c
	}

	return []byte(strconv.Itoa(total))
}
