package binarysearch

// https://leetcode.com/problems/smallest-substring-with-identical-characters-ii/

// Approach: Binary Search + Greedy Verification
// Time: O(n * log n)
// Space: O(n)
func minLength(s string, numOps int) int {
	// check k=1: 010101... or 101010...
	n := len(s)

	var cnt int

	for i := range n {
		if int(s[i]-'0') == i%2 {
			cnt++
		}
	}

	if min(cnt, n-cnt) <= numOps {
		return 1
	}

	// groupby
	var (
		lens []int
		size = 0
	)

	prev := '?'
	for _, c := range s {
		if c == prev {
			lens[size-1]++
		} else {
			lens = append(lens, 1)
			size++
			prev = c
		}
	}

	check := func(k int) bool {
		var ops int
		for _, l := range lens {
			ops += l / (k + 1)
		}

		return ops <= numOps
	}

	low, high := 2, n

	var mid int
	for low < high {
		mid = low + (high-low)/2
		if check(mid) {
			high = mid
		} else {
			low = mid + 1
		}
	}

	return low
}
