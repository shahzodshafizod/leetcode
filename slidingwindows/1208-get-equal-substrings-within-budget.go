package slidingwindows

// https://leetcode.com/problems/get-equal-substrings-within-budget/

// time: O(n)
// space: O(1)
func equalSubstring(s string, t string, maxCost int) int {
	getcost := func(idx int) int {
		cost := int(s[idx]) - int(t[idx])
		if cost < 0 {
			return -cost
		}

		return cost
	}
	start := 0
	cost := 0
	maxlen := 0

	n := len(s)
	for end := 0; end < n; end++ {
		cost += getcost(end)
		for cost > maxCost {
			cost -= getcost(start)
			start++
		}

		maxlen = max(maxlen, end-start+1)
	}

	return maxlen
}

// // time: O(n)
// // space: O(n)
// func equalSubstring(s string, t string, maxCost int) int {
// 	var n = len(s)
// 	var diff = make([]int, n)
// 	for idx := 0; idx < n; idx++ {
// 		diff[idx] = int(s[idx]) - int(t[idx])
// 		if diff[idx] < 0 {
// 			diff[idx] = -diff[idx]
// 		}
// 	}
// 	var start = 0
// 	var cost = 0
// 	var maxlen = 0
// 	for end := range diff {
// 		cost += diff[end]
// 		for cost > maxCost {
// 			cost -= diff[start]
// 			start++
// 		}
// 		maxlen = max(maxlen, end-start+1)
// 	}
// 	return maxlen
// }
