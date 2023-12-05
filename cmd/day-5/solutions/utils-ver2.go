package solutions

import (
	"fmt"
	"strconv"
	"strings"
)

type ConversionMapV2 struct {
	Src     string
	Dst     string
	Mapping []RangeMapping
}

type RangeMapping struct {
	LowerBound  int
	UpperBound  int
	ValueChange int
}

func inputToMapsV2(
	input []string,
	partNumber int,
) (seeds []int, mappings map[string]ConversionMapV2) {
	// First line is the seeds:
		seeds = getSeedPart1(input[0])

	mappings = map[string]ConversionMapV2{}

	// traverse the file lines to the bottom
	for i := 2; i < len(input); {
		line := input[i]
		if !strings.Contains(line, "map:") {
			i++
			continue
		}

		// This is the start of the mapping
		mappingName := strings.TrimSuffix(line, " map:")
		split := strings.Split(mappingName, "-to-")
		conversionMap := ConversionMapV2{
			Src: split[0],
			Dst: split[1],
		}

		fmt.Printf("conversionMap: %v\n", conversionMap)
		mappingArr := []string{}

		for {
			i++
			line = input[i]
			if line == "" {
				// we will end at the position of the empty line
				break
			}
			mappingArr = append(mappingArr, line)
		}

		conversionMap.Mapping = createMapV2(mappingArr)
		mappings[split[0]] = conversionMap
	}

	return seeds, mappings
}

func getSeedPart1(s string) []int {
	seedsStr := strings.Split(strings.TrimPrefix(s, "seeds: "), " ")
	seeds := []int{}
	for _, seedStr := range seedsStr {
		seed, _ := strconv.Atoi(seedStr)
		seeds = append(seeds, seed)
	}

	return seeds
}

	return seeds
}

func createMapV2(m []string) []RangeMapping {
	output := make([]RangeMapping, len(m))
	for i, line := range m {
		rm := RangeMapping{}
		split := strings.Split(line, " ")
		dst, _ := strconv.Atoi(split[0])
		src, _ := strconv.Atoi(split[1])
		ran, _ := strconv.Atoi(split[2])
		rm.LowerBound = src
		rm.UpperBound = src + ran - 1
		rm.ValueChange = dst - src

		output[i] = rm
	}

	return output
}
