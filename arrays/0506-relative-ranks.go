package arrays

import (
	"strconv"
)

// https://leetcode.com/problems/relative-ranks/

// time: O(3*n) = O(n)
// space: O(n)
func findRelativeRanks(score []int) []string {
	max := -1
	for _, score := range score {
		if max == -1 || max < score {
			max = score
		}
	}
	// var max = slices.Max(score)
	indexes := make([]int, max+1)
	for idx, score := range score {
		indexes[score] = idx + 1
	}

	ranks := make([]string, len(score))
	place := 0

	var rank string

	for score := max; score >= 0; score-- {
		if indexes[score] != 0 {
			place++
			switch place {
			case 1:
				rank = "Gold Medal"
			case 2:
				rank = "Silver Medal"
			case 3:
				rank = "Bronze Medal"
			default:
				rank = strconv.Itoa(place)
			}

			ranks[indexes[score]-1] = rank
		}
	}

	return ranks
}

// // time: O(n log n)
// // space: O(n)
// func findRelativeRanks(score []int) []string {
// 	var pq = design.NewPQ[[2]int](make([][2]int, 0), func(x, y [2]int) bool { return x[0] < y[0] })
// 	for index, score := range score {
// 		pq.Push([2]int{score, index})
// 	}
// 	var ranks = make([]string, len(score))
// 	var rank string
// 	var place = 0
// 	for pq.Len() > 0 {
// 		place++
// 		switch place {
// 		case 1:
// 			rank = "Gold Medal"
// 		case 2:
// 			rank = "Silver Medal"
// 		case 3:
// 			rank = "Bronze Medal"
// 		default:
// 			rank = strconv.Itoa(place)
// 		}
// 		ranks[pq.Pop()[1]] = rank
// 	}
// 	return ranks
// }
