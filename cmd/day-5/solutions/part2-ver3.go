package solutions

import (
	"fmt"
	"strconv"
)

func Part2V3(input []string) []byte {
	seeds, maps := inputToMapsPart2V3(input)

	currentRanges := seeds
	currentCategory := "seed"

	for {
		currentMap := maps[currentCategory]

		fmt.Printf("currentCategory: %v\n", currentCategory)
		fmt.Printf("currentRanges: %v\n", currentRanges)
		fmt.Printf("currentMap: %v\n", currentMap)

		newRanges := []Range{}

		for _, currentRange := range currentRanges {
			totalLength := currentRange.Length
			totalLengthCalculated := 0
			brokenDownRanges := []Range{} // This is broken down each range
			for _, rangeMapping := range currentMap.Mapping {
				// Exit condition for when the totalLengthCalculated == totalLength
				if totalLength == totalLengthCalculated {
					break
				}

				leftSide := currentRange.Start
				rightSide := currentRange.Start + currentRange.Length - 1

				// Check if the range is within
				if leftSide >= rangeMapping.LowerBound &&
					rightSide <= rangeMapping.UpperBound {
					// This signifies that it is within the range, and therefore this will
					// consume 100% of the totalLength
					currentRange.Start = leftSide + rangeMapping.ValueChange
					brokenDownRanges = append(brokenDownRanges, currentRange)
					currentRange.Length = 0
					// We break cause it only handles 1 condition each mapping
					break
				}

				// Check if the left side of range is overlapping
				if leftSide >= rangeMapping.LowerBound &&
					leftSide <= rangeMapping.UpperBound {
					// This signifies that we will trim off the right side
					// This consumes the length of the rangeMapping
					newRange := Range{
						Start:  leftSide + rangeMapping.ValueChange,
						Length: rangeMapping.UpperBound - currentRange.Start + 1,
					}
					currentRange.Start = rangeMapping.UpperBound + 1
					currentRange.Length = currentRange.Length - newRange.Length
					brokenDownRanges = append(brokenDownRanges, newRange)
					totalLengthCalculated += newRange.Length
				}

				// Check if the right side of range is overlapping
				if rightSide <= rangeMapping.UpperBound &&
					rightSide >= rangeMapping.LowerBound {
					// This signifies that we will trim off the left side
					// This consumes the length of the rangeMapping
					newRange := Range{
						Start:  rangeMapping.LowerBound + rangeMapping.ValueChange,
						Length: rightSide - rangeMapping.LowerBound + 1,
					}
					currentRange.Length = rangeMapping.LowerBound - currentRange.Start
					brokenDownRanges = append(brokenDownRanges, newRange)
					totalLengthCalculated += newRange.Length
				}
			}
			if currentRange.Length != 0 {
				newRanges = append(newRanges, currentRange)
			}
			newRanges = append(newRanges, brokenDownRanges...)
		}

		currentRanges = newRanges
		currentCategory = currentMap.Dst

		if currentCategory == "location" {
			break
		}
	}

	min := currentRanges[0].Start
	for _, locRange := range currentRanges {
		if locRange.Start < min {
			min = locRange.Start
		}
	}

	return []byte(strconv.Itoa(min))
}
