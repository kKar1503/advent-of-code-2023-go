package solutions

import (
	"strconv"
	"strings"
)

type Round struct {
	R int
	B int
	G int
}

type Game struct {
	GameNum int
	Rounds  []Round
}

func splitIntoGames(input []string) (games []Game, err error) {
	games = []Game{}
	for _, gameInput := range input {
		if gameInput == "" {
			continue
		}

		gameTmp := Game{}
		firstSplit := strings.Split(gameInput, ": ")
		gameNumStr := strings.Split(firstSplit[0], " ")
		gameNum, err2 := strconv.Atoi(gameNumStr[1])
		err = err2
		if err != nil {
			return []Game{}, err
		}

		gameTmp.GameNum = gameNum
		secondSplit := strings.Split(firstSplit[1], "; ")
		gameTmp.Rounds = []Round{}

		for _, roundStr := range secondSplit {
			round := Round{
				R: 0,
				B: 0,
				G: 0,
			}
			colorDraws := strings.Split(roundStr, ", ")
			for _, colorDraw := range colorDraws {
				numAndColor := strings.Split(colorDraw, " ")
				num, err2 := strconv.Atoi(numAndColor[0])
				err = err2
				if err != nil {
					return []Game{}, err
				}
				color := numAndColor[1]
				switch color {
				case "red":
					round.R += num
				case "blue":
					round.B += num
				case "green":
					round.G += num
				}
			}
			gameTmp.Rounds = append(gameTmp.Rounds, round)
		}
		games = append(games, gameTmp)
	}
	return games, nil
}
