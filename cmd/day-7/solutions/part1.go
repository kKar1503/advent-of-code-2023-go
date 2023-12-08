package solutions

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

type HandType = uint8

const (
	HighCard HandType = iota
	OnePair
	TwoPair
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind
)

func cardValuesPart1() map[string]int {
	return map[string]int{
		"A": 13,
		"K": 12,
		"Q": 11,
		"J": 10,
		"T": 9,
		"9": 8,
		"8": 7,
		"7": 6,
		"6": 5,
		"5": 4,
		"4": 3,
		"3": 2,
		"2": 1,
	}
}

type Hand struct {
	CardFaces string
	Cards     []int
	Values    int
	Bid       int
	Rank      int
}

type HandTypesMapping = map[HandType][]Hand

func Part1(input []string) []byte {
	handTypesMapping := splitIntoHandTypesMappingPart1(input)

	start, total := 1, 0

	for i := 0; i < 7; i++ {
		hands := handTypesMapping[uint8(i)]
		hands, end := giveRank(hands, start)
		for _, hand := range hands {
			fmt.Printf("hand: %#v\n", hand)
			total += hand.Bid * hand.Rank
		}
		start = end
	}

	return []byte(strconv.Itoa(total))
}

func splitIntoHandTypesMappingPart1(inputs []string) HandTypesMapping {
	handTypesMapping := HandTypesMapping{}
	cardValues := cardValuesPart1()

	for _, input := range inputs {
		if input == "" {
			continue
		}
		split := strings.Split(input, " ")
		bid, _ := strconv.Atoi(split[1])
		hand := Hand{
			CardFaces: split[0],
			Cards:     convertHandToCardValues(split[0], cardValues),
			Bid:       bid,
		}
		handType := determineTypePart1(split[0])
		if handsByHandType, ok := handTypesMapping[handType]; ok {
			handTypesMapping[handType] = append(handsByHandType, hand)
		} else {
			handTypesMapping[handType] = []Hand{hand}
		}
	}

	return handTypesMapping
}

func determineTypePart1(hand string) HandType {
	if isFiveOfAKindPart1(hand) {
		return FiveOfAKind
	}

	if isFourOfAKindPart1(hand) {
		return FourOfAKind
	}

	if isFullHousePart1(hand) {
		return FullHouse
	}

	if isThreeOfAKindPart1(hand) {
		return ThreeOfAKind
	}

	if isTwoPairPart1(hand) {
		return TwoPair
	}

	if isOnePairPart1(hand) {
		return OnePair
	}

	return HighCard
}

func isFiveOfAKindPart1(s string) bool {
	for i := 1; i < len(s); i++ {
		if s[i] != s[i-1] {
			return false
		}
	}

	return true
}

func isFourOfAKindPart1(s string) bool {
	str1 := ""
	counter1 := 0
	str2 := ""
	counter2 := 0

	for i := 0; i < len(s); i++ {
		c := string(s[i])
		if str1 == c {
			counter1++
			continue
		}

		if str2 == c {
			counter2++
			continue
		}

		if str1 == "" {
			// sets the first char
			str1 = c
			counter1++
			continue
		}

		if str2 == "" {
			// sets the second char
			str2 = c
			counter2++
			continue
		}
	}

	return (counter1 == 4) || (counter2 == 4)
}

func isFullHousePart1(s string) bool {
	str1 := ""
	counter1 := 0
	str2 := ""
	counter2 := 0

	for i := 0; i < len(s); i++ {
		c := string(s[i])
		if str1 == c {
			counter1++
			continue
		}

		if str2 == c {
			counter2++
			continue
		}

		if str1 == "" {
			// sets the first char
			str1 = c
			counter1++
			continue
		}

		if str2 == "" {
			// sets the second char
			str2 = c
			counter2++
			continue
		}
	}

	return (counter1 == 3 && counter2 == 2) || (counter1 == 2 && counter2 == 3)
}

func isThreeOfAKindPart1(s string) bool {
	str1 := ""
	counter1 := 0
	str2 := ""
	counter2 := 0
	str3 := ""
	counter3 := 0

	for i := 0; i < len(s); i++ {
		c := string(s[i])
		if str1 == c {
			counter1++
			continue
		}

		if str2 == c {
			counter2++
			continue
		}

		if str3 == c {
			counter3++
			continue
		}

		if str1 == "" {
			// sets the first char
			str1 = c
			counter1++
			continue
		}

		if str2 == "" {
			// sets the second char
			str2 = c
			counter2++
			continue
		}

		if str3 == "" {
			// sets the third char
			str3 = c
			counter3++
			continue
		}
	}

	return (counter1 == 3 && counter2 == 1 && counter3 == 1) ||
		(counter1 == 1 && counter2 == 3 && counter3 == 1) ||
		(counter1 == 1 && counter2 == 1 && counter3 == 3)
}

func isTwoPairPart1(s string) bool {
	str1 := ""
	counter1 := 0
	str2 := ""
	counter2 := 0
	str3 := ""
	counter3 := 0

	for i := 0; i < len(s); i++ {
		c := string(s[i])
		if str1 == c {
			counter1++
			continue
		}

		if str2 == c {
			counter2++
			continue
		}

		if str3 == c {
			counter3++
			continue
		}

		if str1 == "" {
			// sets the first char
			str1 = c
			counter1++
			continue
		}

		if str2 == "" {
			// sets the second char
			str2 = c
			counter2++
			continue
		}

		if str3 == "" {
			// sets the third char
			str3 = c
			counter3++
			continue
		}
	}

	return (counter1 == 2 && counter2 == 2 && counter3 == 1) ||
		(counter1 == 2 && counter2 == 1 && counter3 == 2) ||
		(counter1 == 1 && counter2 == 2 && counter3 == 2)
}

func isOnePairPart1(s string) bool {
	appeared := []string{}
	paired := false

	for i := 0; i < len(s); i++ {
		c := string(s[i])
		if !slices.Contains(appeared, c) {
			appeared = append(appeared, c)
			continue
		}
		if paired {
			return false
		}
		paired = true
	}

	return paired
}
