package solutions

import (
	"fmt"
	"slices"
	"strconv"
)

func Part2(input []string) []byte {
	seeds, maps := inputToMapsV2(input, 2)

	currentValues := seeds
	currentCategory := "seed"

	for {
		currentMap := maps[currentCategory]

		fmt.Printf("currentCategory: %v\n", currentCategory)
		fmt.Printf("currentValues: %v\n", currentValues)
		fmt.Printf("currentMap: %v\n", currentMap)

		for i, currentValue := range currentValues {
			for _, rangeMapping := range currentMap.Mapping {
				if currentValue >= rangeMapping.LowerBound &&
					currentValue <= rangeMapping.UpperBound {
					currentValues[i] = currentValue + rangeMapping.ValueChange
				}
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
