package solutions

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func cardValuesPart2() map[string]int {
	return map[string]int{
		"A": 13,
		"K": 12,
		"Q": 11,
		"T": 10,
		"9": 9,
		"8": 8,
		"7": 7,
		"6": 6,
		"5": 5,
		"4": 4,
		"3": 3,
		"2": 2,
		"J": 1,
	}
}

const joker = "J"

func Part2(input []string) []byte {
	handTypesMapping := splitIntoHandTypesMappingPart2(input)

	for _, hand := range handTypesMapping[ThreeOfAKind] {
		if strings.Contains(hand.CardFaces, joker) {
			fmt.Printf("ThreeOfAKind: %v\n", hand)
		}
	}

	start, total := 1, 0

	for i := 0; i < 7; i++ {
		hands := handTypesMapping[uint8(i)]
		hands, end := giveRank(hands, start)
		for _, hand := range hands {
			// fmt.Printf("hand: %#v\n", hand)
			total += hand.Bid * hand.Rank
		}
		start = end
	}

	return []byte(strconv.Itoa(total))
}

func splitIntoHandTypesMappingPart2(inputs []string) HandTypesMapping {
	handTypesMapping := HandTypesMapping{}
	cardValues := cardValuesPart2()

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
		handType := determineTypePart2(split[0])
		if handsByHandType, ok := handTypesMapping[handType]; ok {
			handTypesMapping[handType] = append(handsByHandType, hand)
		} else {
			handTypesMapping[handType] = []Hand{hand}
		}
	}

	return handTypesMapping
}

func determineTypePart2(hand string) HandType {
	if isFiveOfAKindPart2(hand) {
		return FiveOfAKind
	}

	if isFourOfAKindPart2(hand) {
		return FourOfAKind
	}

	if isFullHousePart2(hand) {
		return FullHouse
	}

	if isThreeOfAKindPart2(hand) {
		return ThreeOfAKind
	}

	if isTwoPairPart2(hand) {
		return TwoPair
	}

	if isOnePairPart2(hand) {
		return OnePair
	}

	return HighCard
}

func isFiveOfAKindPart2(s string) bool {
	c := ""
	for i := 0; i < len(s); i++ {
		c2 := string(s[i])
		if c2 == joker {
			continue
		}
		if c == "" {
			c = c2
			continue
		}

		if c != c2 {
			return false
		}
	}

	return true
}

func isFourOfAKindPart2(s string) bool {
	jCount := 0
	str1 := ""
	counter1 := 0
	str2 := ""
	counter2 := 0

	for i := 0; i < len(s); i++ {
		c := string(s[i])
		if c == joker {
			jCount++
			continue
		}

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

	return jCount == 3 || counter1+jCount == 4 || counter2+jCount == 4
}

func isFullHousePart2(s string) bool {
	jCount := 0
	str1 := ""
	counter1 := 0
	str2 := ""
	counter2 := 0

	for i := 0; i < len(s); i++ {
		c := string(s[i])
		if c == joker {
			jCount++
			continue
		}

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

	if jCount == 0 {
		return (counter1 == 3 && counter2 == 2) ||
			(counter1 == 2 && counter2 == 3)
	} else if jCount == 1 {
		return (counter1 == 2 && counter2 == 2) ||
			(counter1 == 3 && counter2 == 1)
	} else {
		return (counter1 == 1 && counter2 == 2) ||
			(counter1 == 3) ||
			(counter1 == 2 && counter2 == 1)
	}
}

func isThreeOfAKindPart2(s string) bool {
	jCount := 0
	str1 := ""
	counter1 := 0
	str2 := ""
	counter2 := 0
	str3 := ""
	counter3 := 0

	for i := 0; i < len(s); i++ {
		c := string(s[i])
		if c == joker {
			jCount++
			continue
		}

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

	if jCount == 2 {
		return true
	} else if jCount == 0 {
		return counter1 == 3 ||
			counter2 == 3 ||
			counter3 == 3
	} else {
		return counter1 == 2 ||
			counter2 == 2 ||
			counter3 == 2
	}
}

func isTwoPairPart2(s string) bool {
	jCount := 0
	str1 := ""
	counter1 := 0
	str2 := ""
	counter2 := 0
	str3 := ""
	counter3 := 0

	for i := 0; i < len(s); i++ {
		c := string(s[i])
		if c == joker {
			jCount++
			continue
		}

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

	if jCount == 0 {
		return (counter1 == 2 && counter2 == 2) ||
			(counter1 == 2 && counter3 == 2) ||
			(counter2 == 2 && counter3 == 2)
	} else if jCount == 1 {
		return (counter1 == 2 && counter2 == 1) ||
			(counter1 == 2 && counter3 == 2) ||
			(counter2 == 2 && counter1 == 1) ||
			(counter2 == 2 && counter3 == 1) ||
			(counter3 == 2 && counter1 == 1) ||
			(counter3 == 2 && counter2 == 1)
	} else if jCount == 2 {
		return (counter1 == 1 && counter2 == 1) ||
			(counter1 == 1 && counter3 == 2) ||
			(counter2 == 1 && counter1 == 1) ||
			(counter2 == 1 && counter3 == 1) ||
			(counter3 == 1 && counter1 == 1) ||
			(counter3 == 1 && counter2 == 1) ||
			counter1 == 2 ||
			counter2 == 2 ||
			counter3 == 2
	}

	return false
}

func isOnePairPart2(s string) bool {
	appeared := []string{}
	paired := false

	for i := 0; i < len(s); i++ {
		c := string(s[i])
		if c == joker {
			// having a joker always can form a pair
			return true
		}
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
