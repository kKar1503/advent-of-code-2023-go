package solutions

import (
	"fmt"
	"os"
	"strconv"
)

func Part2(input []string) []byte {
	games, err := splitIntoGames(input)
	if err != nil {
		fmt.Printf("Failed to split data into games: %s", err.Error())
		os.Exit(1)
	}

	total := 0

	for _, game := range games {
		redMin := 0
		blueMin := 0
		greenMin := 0
		for _, round := range game.Rounds {
			if round.B > blueMin {
				blueMin = round.B
			}
			if round.R > redMin {
				redMin = round.R
			}
			if round.G > greenMin {
				greenMin = round.G
			}
		}
		total += greenMin * redMin * blueMin
	}

	return []byte(strconv.Itoa(total))
}
