package solutions

import (
	"slices"
	"strconv"
)

func Part1(input []string) []byte {
	cards := splitIntoTwoSets(input)

	total := 0

	for _, card := range cards {
		point := 0
		for _, winningNumber := range card.WinningNumbers {
			if slices.Contains(card.NumbersBought, winningNumber) {
				if point == 0 {
					point = 1
				} else {
					point = point * 2
				}
			}
		}
		total += point
	}

	return []byte(strconv.Itoa(total))
}
