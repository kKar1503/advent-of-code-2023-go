package solutions

import (
	"fmt"
	"slices"
	"strconv"
)

func Part1(input []string) []byte {
	seeds, maps := inputToMaps(input)

	currentValues := seeds
	currentCategory := "seed"

	for {
		currentMap := maps[currentCategory]

		fmt.Printf("currentCategory: %v\n", currentCategory)
		fmt.Printf("currentValues: %v\n", currentValues)
		fmt.Printf("currentMap: %v\n", currentMap)

		for i, currentValue := range currentValues {
			if dst, ok := currentMap.Mapping[currentValue]; ok {
				currentValues[i] = dst
			}
		}

		currentCategory = currentMap.Dst
		if currentCategory == "location" {
			break
		}
	}

	minLocation := slices.Min(currentValues)

	return []byte(strconv.Itoa(minLocation))
}
