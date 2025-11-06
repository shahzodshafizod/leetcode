package dp

import "math/bits"

// https://leetcode.com/problems/maximize-the-number-of-partitions-after-operations/

func maxPartitionsAfterOperations(s string, k int) int {
	memo := make(map[[3]int]*int)

	var dp func(i int, mask int, changed int) int

	dp = func(i int, mask int, changed int) int {
		if i == len(s) {
			return 1
		}

		key := [3]int{i, mask, changed}
		if memo[key] != nil {
			return *memo[key]
		}

		var best int

		curr := int(s[i] - 'a')

		nmask := mask | (1 << curr)
		if bits.OnesCount(uint(nmask)) <= k {
			best = max(best, dp(i+1, nmask, changed))
		} else {
			best = max(best, 1+dp(i+1, 1<<curr, changed))
		}

		if changed == 0 {
			for next := range 26 {
				if next == curr {
					continue
				}

				nmask = mask | (1 << next)
				if bits.OnesCount(uint(nmask)) <= k {
					best = max(best, dp(i+1, nmask, 1))
				} else {
					best = max(best, 1+dp(i+1, 1<<next, 1))
				}
			}
		}

		memo[key] = &best

		return best
	}

	return dp(0, 0, 0)
}
