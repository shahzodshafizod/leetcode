package backtracking

import (
	"math"
	"slices"
)

// https://leetcode.com/problems/24-game/

func judgePoint24(cards []int) bool {
	var backtracking func(cards []float64) bool

	backtracking = func(cards []float64) bool {
		n := len(cards)
		if n == 1 {
			return math.Abs(cards[0]-24) < 0.001
		}

		var card1, card2 float64

		for i := range n {
			for j := i + 1; j < n; j++ {
				card1 = cards[i]
				card2 = cards[j]

				values := []float64{card1 + card2, card1 - card2, card2 - card1, card1 * card2}
				if card1 != 0 {
					values = append(values, card2/card1)
				}

				if card2 != 0 {
					values = append(values, card1/card2)
				}

				cards = slices.Delete(cards, j, j+1)
				cards = slices.Delete(cards, i, i+1)

				for _, value := range values {
					cards = append(cards, value)
					if backtracking(cards) {
						return true
					}

					cards = cards[:len(cards)-1]
				}

				cards = slices.Insert(cards, i, card1)
				cards = slices.Insert(cards, j, card2)
			}
		}

		return false
	}

	fcards := make([]float64, 0, len(cards))
	for _, card := range cards {
		fcards = append(fcards, float64(card))
	}

	return backtracking(fcards)
}
