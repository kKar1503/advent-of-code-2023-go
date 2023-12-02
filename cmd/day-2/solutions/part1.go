package solutions

import (
	"fmt"
	"os"
	"strconv"
)

const (
	RED_MAX   = 12
	GREEN_MAX = 13
	BLUE_MAX  = 14
)

func Part1(input []string) []byte {
	games, err := splitIntoGames(input)
	if err != nil {
		fmt.Printf("Failed to split data into games: %s", err.Error())
		os.Exit(1)
	}

	total := 0

	for _, game := range games {
		valid := true
		for _, round := range game.Rounds {
			if round.B > BLUE_MAX {
				valid = false
				break
			}
			if round.R > RED_MAX {
				valid = false
				break
			}
			if round.G > GREEN_MAX {
				valid = false
				break
			}
		}
		if valid {
			total += game.GameNum
		}
	}

	return []byte(strconv.Itoa(total))
}
