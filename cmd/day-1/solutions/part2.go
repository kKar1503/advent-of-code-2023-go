package solutions

import (
	"strconv"
	"strings"

	"2023-go/pkg/utils"
)

func Part2(input []string) []byte {
	valueMap := map[string]int{
		"one":   1,
		"1":     1,
		"two":   2,
		"2":     2,
		"three": 3,
		"3":     3,
		"four":  4,
		"4":     4,
		"five":  5,
		"5":     5,
		"six":   6,
		"6":     6,
		"seven": 7,
		"7":     7,
		"eight": 8,
		"8":     8,
		"nine":  9,
		"9":     9,
	}
	var total uint64 = 0
	for _, line := range input {
		iStr := part2Helper1(line)
		i := valueMap[iStr]
		jStr := part2Helper2(line)
		j := valueMap[jStr]
		total += uint64(i*10 + j)
	}

	return []byte(strconv.FormatUint(total, 10))
}

func part2Helper1(line string) string {
	arr := []string{
		"one", "1", "two", "2",
		"three", "3", "four", "4",
		"five", "5", "six", "6",
		"seven", "7", "eight", "8",
		"nine", "9",
	}

	smallest := -1
	digit := ""

	for _, str := range arr {
		index := strings.Index(line, str)
		if index == -1 {
			continue
		}
		if smallest == -1 || index < smallest {
			smallest = index
			digit = str
		}
	}

	return digit
}

func part2Helper2(line string) string {
	reverse := utils.ReverseString(line)
	arr := []string{
		"eno", "1", "owt", "2",
		"eerht", "3", "ruof", "4",
		"evif", "5", "xis", "6",
		"neves", "7", "thgie", "8",
		"enin", "9",
	}

	smallest := -1
	digit := ""

	for _, str := range arr {
		index := strings.Index(reverse, str)
		if index == -1 {
			continue
		}
		if smallest == -1 || index < smallest {
			smallest = index
			digit = str
		}
	}

	return utils.ReverseString(digit)
}
