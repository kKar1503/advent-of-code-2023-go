package solutions

import (
	"fmt"
	"strconv"
	"strings"
)

type Range struct {
	Start  int
	Length int
}

func inputToMapsPart2V3(
	input []string,
) (seeds []Range, mappings map[string]ConversionMapV2) {
	// First line is the seeds:
	seeds = getSeedPart2V3(input[0])

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

func getSeedPart2V3(s string) []Range {
	seedsStr := strings.Split(strings.TrimPrefix(s, "seeds: "), " ")
	seedRanges := make([]Range, len(seedsStr)/2)
	for i := 0; i < len(seedsStr); i += 2 {
		seedStart, _ := strconv.Atoi(seedsStr[i])
		seedRange, _ := strconv.Atoi(seedsStr[i+1])
		r := Range{
			Start:  seedStart,
			Length: seedRange,
		}
		seedRanges[i/2] = r
	}

	return seedRanges
}
