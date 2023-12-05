package main

import (
	"fmt"
	"os"

	"2023-go/cmd/day-5/solutions"
	"2023-go/pkg/utils"
)

const day = 5

var parts = map[string]solutionFn{
	"part1":   solutions.Part1,
	"part1v2": solutions.Part1V2,
	"part2":   solutions.Part2,
}

type solutionFn = func([]string) []byte

func main() {
	input, err := utils.ImportAsArray(fmt.Sprintf("cmd/day-%d/input", day))
	if err != nil {
		fmt.Printf("Failed reading the file: %s", err.Error())
		os.Exit(1)
	}

	part := os.Args[1]

	output := parts[part](input)

	utils.Output(output, fmt.Sprintf("cmd/day-%d/output-%s", day, part))
}
