package solutions

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Card struct {
	CardNumber     int
	WinningNumbers []int
	NumbersBought  []int
}

func splitIntoTwoSets(cards []string) []Card {
	res := make([]Card, len(cards)-1)
	for i, card := range cards {
		if card == "" {
			continue
		}
		c := Card{
			WinningNumbers: []int{},
			NumbersBought:  []int{},
		}
		split1 := strings.Split(card, ": ")
		cardNumberStr := strings.TrimPrefix(split1[0], "Card")
		cardNumberStr = strings.TrimLeft(cardNumberStr, " ")

		cardNumber, err := strconv.Atoi(cardNumberStr)
		if err != nil {
			fmt.Printf("failed to convert card number : %s", cardNumberStr)
			os.Exit(1)
		}
		c.CardNumber = cardNumber

		numbers := split1[1]
		numbers = strings.ReplaceAll(numbers, "  ", " ")
		numbers = strings.Trim(numbers, " ")
		split2 := strings.Split(numbers, " | ")
		winningNumbers := strings.Split(split2[0], " ")
		numbersBought := strings.Split(split2[1], " ")

		for _, winningNumber := range winningNumbers {
			n, err := strconv.Atoi(winningNumber)
			if err != nil {
				fmt.Printf("Winning: Something went wrong parsing %s\n", winningNumber)
				os.Exit(1)
			}
			c.WinningNumbers = append(c.WinningNumbers, n)
		}

		for _, numberBought := range numbersBought {
			n, err := strconv.Atoi(numberBought)
			if err != nil {
				fmt.Printf("Bought: Something went wrong parsing %s\n", numberBought)
				os.Exit(1)
			}
			c.NumbersBought = append(c.NumbersBought, n)
		}

		res[i] = c
	}

	return res
}
