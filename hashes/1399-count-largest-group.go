package hashes

// https://leetcode.com/problems/count-largest-group/

func countLargestGroup(n int) int {
	groups := make(map[int]int)

	var sum, tmp int

	size := 0

	for num := 1; num <= n; num++ {
		tmp, sum = num, 0
		for ; tmp > 0; tmp /= 10 {
			sum += tmp % 10
		}

		groups[sum]++
		size = max(size, groups[sum])
	}

	count := 0

	for _, cnt := range groups {
		if cnt == size {
			count++
		}
	}

	return count
}
