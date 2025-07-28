package greedy

// https://leetcode.com/problems/minimum-domino-rotations-for-equal-row/

func minDominoRotations(tops []int, bottoms []int) int {
	// If one side can be made equal,
	// another side can be made equal with at least the same effort,
	// not more optimal!!!
	// So the first found one is the optimal one.
	n := len(tops)

	var idx int

	for _, target := range []int{tops[0], bottoms[0]} {
		tswaps, bswaps := 0, 0

		for idx = 0; idx < n; idx++ {
			if tops[idx] != target && bottoms[idx] != target {
				break
			} else if tops[idx] != target {
				tswaps++
			} else if bottoms[idx] != target {
				bswaps++
			}
		}

		if idx == n {
			return min(tswaps, bswaps)
		}
	}

	return -1
}
