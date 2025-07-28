package hashes

// https://leetcode.com/problems/number-of-equivalent-domino-pairs/

func numEquivDominoPairs(dominoes [][]int) int {
	count := make(map[int]int)

	var a, b, key int

	var total int

	for idx := range dominoes {
		a, b = dominoes[idx][0], dominoes[idx][1]
		if a > b {
			a, b = b, a
		}

		key = a*10 + b
		total += count[key]
		count[key]++
	}

	return total
}
