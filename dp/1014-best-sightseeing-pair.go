package dp

// https://leetcode.com/problems/best-sightseeing-pair/

func maxScoreSightseeingPair(values []int) int {
	var n = len(values)
	var maxval = values[n-1] - n + 1
	var score = 0
	for idx := n - 2; idx >= 0; idx-- {
		score = max(score, values[idx]+idx+maxval)
		maxval = max(maxval, values[idx]-idx)
	}
	return score
}
