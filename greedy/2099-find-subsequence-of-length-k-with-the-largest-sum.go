package greedy

import "sort"

// https://leetcode.com/problems/find-subsequence-of-length-k-with-the-largest-sum/

func maxSubsequence(nums []int, k int) []int {
	n := len(nums)
	indices := make([]int, n)

	for idx := 0; idx < n; idx++ {
		indices[idx] = idx
	}

	sort.Slice(indices, func(i int, j int) bool {
		return nums[indices[i]] > nums[indices[j]]
	})

	pick := make([]bool, n)
	for idx := 0; idx < k; idx++ {
		pick[indices[idx]] = true
	}

	subseq := make([]int, 0, k)

	for idx, num := range nums {
		if pick[idx] {
			subseq = append(subseq, num)
		}
	}

	return subseq
}
