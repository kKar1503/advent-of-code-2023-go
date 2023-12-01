package solutions

import "strconv"

func Part1(input []string) []byte {
	var total uint64 = 0
	for _, line := range input {
		arr := []uint8(line)
		i := part1Helper1(arr)
		j := part1Helper2(arr)
		total += uint64(i*10 + j)
	}

	return []byte(strconv.FormatUint(total, 10))
}

func part1Helper1(line []uint8) uint8 {
	i := 0
	for i < len(line) {
		if line[i] >= 48 && line[i] <= 57 {
			return line[i] - 48
		}
		i++
	}
	return 0
}

func part1Helper2(line []uint8) uint8 {
	i := len(line) - 1
	for i >= 0 {
		if line[i] >= 48 && line[i] <= 57 {
			return line[i] - 48
		}
		i--
	}
	return 0
}
