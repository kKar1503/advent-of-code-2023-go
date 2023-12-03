package solutions

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func Part1(input []string) []byte {
	field := make([][]string, len(input)-1)
	for i, line := range input {
		if line == "" {
			continue
		}
		splitLine := strings.Split(line, "")
		field[i] = splitLine
	}

	// Traverse the 2d array
	sum := traversePart1(field)

	return []byte(strconv.Itoa(sum))
}

func traversePart1(field [][]string) int {
	number := ""             // This stores the concurrent number strings
	coordinates := [][]int{} // this marks the starting xy and ending xy
	total := 0               // This stores the sum
	for y := range field {
		for x := range field[0] {
			n := field[y][x]
			if !isNumeric(n) {
				continue // if this is non numeric we skip
			}

			number += n
			coordinates = append(coordinates, []int{x, y})

			if x == (len(field[0])-1) || !isNumeric(field[y][x+1]) {
				// if this is the end or the next one is no longer number
				// Check around it
				if symbolIsAround(
					coordinates[0][0],
					coordinates[len(coordinates)-1][0],
					coordinates[0][1],
					field,
				) {
					num, err := strconv.Atoi(number)
					if err != nil {
						fmt.Printf(
							"Something went wrong during conversion, (num:%s) (err: %s)",
							number,
							err.Error(),
						)
						os.Exit(1)
					}
					fmt.Println(number)
					total += num
				}
				// resets no matter num has surround or not
				number = ""
				coordinates = [][]int{}
			}
		}
	}
	return total
}

func isNumeric(word string) bool {
	return regexp.MustCompile(`\d`).MatchString(word)
}

func symbolIsAround(xStart, xEnd, y int, field [][]string) bool {
	// Check above
	if y > 0 { // check if not first row
		x2, x2End, y2 := xStart, xEnd, y-1
		if xStart > 0 {
			// check if xStart is not most left
			x2 = xStart - 1
		}
		if xEnd < (len(field[0]) - 1) {
			// check if xEnd is not most right
			x2End = xEnd + 1
		}
		for x2 <= x2End {
			if isNotNumberOrPeriod(field[y2][x2]) {
				return true
			}
			x2++
		}
	}
	// Check below
	if y < (len(field) - 1) { // check if not last row
		x2, x2End, y2 := xStart, xEnd, y+1
		if xStart > 0 {
			// check if xStart is not most left
			x2 = xStart - 1
		}
		if xEnd < (len(field[0]) - 1) {
			// check if xEnd is not most right
			x2End = xEnd + 1
		}
		for x2 <= x2End {
			if isNotNumberOrPeriod(field[y2][x2]) {
				return true
			}
			x2++
		}
	}
	// Check left
	if xStart > 0 {
		xLeft := xStart - 1
		if isNotNumberOrPeriod(field[y][xLeft]) {
			return true
		}
	}
	// Check right
	if xEnd < (len(field[0]) - 1) {
		xRight := xEnd + 1
		if isNotNumberOrPeriod(field[y][xRight]) {
			return true
		}
	}
	return false
}

func isNotNumberOrPeriod(word string) bool {
	return regexp.MustCompile(`[^0-9.]`).MatchString(word)
}
