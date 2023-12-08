package solutions

import "sort"

func giveRank(hands []Hand, start int) ([]Hand, int) {
	res := make([]Hand, len(hands))

	// assign total values
	for i, hand := range hands {
		cards := hand.Cards
		hand.Values = cards[0] * 1_00_00_00_00
		hand.Values += cards[1] * 1_00_00_00
		hand.Values += cards[2] * 1_00_00
		hand.Values += cards[3] * 1_00
		hand.Values += cards[4]
		res[i] = hand
	}

	// sort by values
	sort.Slice(res, func(i, j int) bool {
		return res[i].Values < res[j].Values
	})

	res2 := make([]Hand, len(hands))

	// assign rank
	for i, hand := range res {
		hand.Rank = i + start
		res2[i] = hand
	}

	return res2, start + len(hands)
}

func convertHandToCardValues(hand string, cardValues map[string]int) []int {
	values := make([]int, 5)
	for i, c := range hand {
		cStr := string(c)
		values[i] = cardValues[cStr]
	}

	return values
}
