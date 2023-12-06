package solutions

import (
	"strconv"
	"strings"
)

func Part2(input []string) []byte {
	time, distance := createRace(input)

	total := 0

	for i := 1; i <= time; i++ {
		if calculateDistance(i, time) > distance {
			total++
		}
	}

	return []byte(strconv.Itoa(total))
}

func createRace(input []string) (time int, distance int) {
	timeStr := strings.TrimPrefix(input[0], "Time:")
	timeStr = strings.ReplaceAll(timeStr, " ", "")
	time, _ = strconv.Atoi(timeStr)
	distanceStr := strings.TrimPrefix(input[1], "Distance:")
	distanceStr = strings.ReplaceAll(distanceStr, " ", "")
	distance, _ = strconv.Atoi(distanceStr)
	return
}
