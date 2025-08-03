package greedy

import (
	"sort"
)

// https://leetcode.com/problems/rearranging-fruits/

func minCost(basket1 []int, basket2 []int) int64 {
	freq, minimal := make(map[int]int), int(1e9)
	for i, n := 0, len(basket1); i < n; i++ {
		freq[basket1[i]]++
		freq[basket2[i]]--
		minimal = min(minimal, basket1[i], basket2[i])
	}

	var merge []int

	for val, cnt := range freq {
		if cnt < 0 {
			cnt = -cnt
		}

		if cnt&1 == 1 {
			return -1
		}

		for cnt = cnt / 2; cnt > 0; cnt-- {
			merge = append(merge, val)
		}
	}

	if len(merge) == 0 {
		return 0
	}

	sort.Ints(merge)

	var cost int64
	for i := len(merge)/2 - 1; i >= 0; i-- {
		cost += int64(min(minimal*2, merge[i]))
	}

	return cost
}
