package solutions

import (
	"strconv"
	"strings"
)

type ConversionMap struct {
	Src     string
	Dst     string
	Mapping map[int]int
}

func inputToMaps(input []string) (seeds []int, mappings map[string]ConversionMap) {
	// First line is the seeds:
	seedsStr := strings.Split(strings.TrimPrefix(input[0], "seeds: "), " ")
	seeds = []int{}
	for _, seedStr := range seedsStr {
		seed, _ := strconv.Atoi(seedStr)
		seeds = append(seeds, seed)
	}

	mappings = map[string]ConversionMap{}

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
		conversionMap := ConversionMap{
			Src: split[0],
			Dst: split[1],
		}

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

		conversionMap.Mapping = createMap(mappingArr)
		mappings[split[0]] = conversionMap
	}

	return seeds, mappings
}

func createMap(m []string) map[int]int {
	output := map[int]int{}
	for _, line := range m {
		split := strings.Split(line, " ")
		dst, _ := strconv.Atoi(split[0])
		src, _ := strconv.Atoi(split[1])
		ran, _ := strconv.Atoi(split[2])

		for i := 0; i < ran; i++ {
			output[src+i] = dst + i
		}
	}

	return output
}
