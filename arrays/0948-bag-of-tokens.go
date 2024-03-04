package arrays

import "sort"

// https://leetcode.com/problems/bag-of-tokens/

func bagOfTokensScore(tokens []int, power int) int {
	var score = 0
	sort.Ints(tokens)
	for left, right := 0, len(tokens)-1; left <= right; {
		if power >= tokens[left] {
			power -= tokens[left]
			score++ // face-up
			left++
		} else if score > 0 && left < right {
			power += tokens[right]
			score-- // face-down
			right--
		} else {
			break
		}
	}
	return score
}
