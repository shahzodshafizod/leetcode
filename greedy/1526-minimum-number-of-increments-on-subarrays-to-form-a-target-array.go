package greedy

// https://leetcode.com/problems/minimum-number-of-increments-on-subarrays-to-form-a-target-array/

func minNumberOperations(target []int) int {
	ops := target[0]
	for i, n := 1, len(target); i < n; i++ {
		ops += max(target[i]-target[i-1], 0)
	}

	return ops
}
