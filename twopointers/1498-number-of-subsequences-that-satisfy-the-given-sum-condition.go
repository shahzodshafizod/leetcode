package twopointers

import "sort"

// https://leetcode.com/problems/number-of-subsequences-that-satisfy-the-given-sum-condition/

// Approach: Two Pointer
// Time: O(nlogn)
// Space: O(1)
func numSubseq(nums []int, target int) int {
	n := len(nums)
	power2 := make([]int, n)
	power2[0] = 1

	const MOD int = 1e9 + 7

	for idx := 1; idx < n; idx++ {
		power2[idx] = (power2[idx-1] * 2) % MOD
	}

	sort.Ints(nums)

	count := 0

	left, right := 0, n-1
	for left <= right {
		if nums[left]+nums[right] > target {
			right--
		} else {
			// for each element in the subarray [left+1:right]
			// we can pick or not pick, so there are
			// 2 ^ (right - left) subsequences in total
			count = (count + power2[right-left]) % MOD
			left++
		}
	}

	return count
}
