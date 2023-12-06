package solutions

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Part1(input []string) []byte {
	races := createRaces(input)
	total := 1

	for _, race := range races {
		totalPossibilities := 0
		for i := 1; i <= race.Time; i++ {
			if calculateDistance(i, race.Time) > race.Distance {
				totalPossibilities++
			}
		}
		total = total * totalPossibilities
	}

	return []byte(strconv.Itoa(total))
}

type Race struct {
	Time     int
	Distance int
}

func createRaces(input []string) []Race {
	time := input[0]
	distance := input[1]
	time = strings.TrimPrefix(time, "Time:")
	times := strings.Fields(time)
	distance = strings.TrimPrefix(distance, "Distance:")
	distances := strings.Fields(distance)

	if len(times) != len(distances) {
		fmt.Println("time and distance count mismatch!")
		os.Exit(1)
	}

	races := make([]Race, len(times))

	for i := 0; i < len(races); i++ {
		t, _ := strconv.Atoi(times[i])
		d, _ := strconv.Atoi(distances[i])
		races[i] = Race{
			Time:     t,
			Distance: d,
		}
	}

	return races
}

func calculateDistance(holdingTime, totalDuration int) int {
	return (totalDuration - holdingTime) * holdingTime
}
