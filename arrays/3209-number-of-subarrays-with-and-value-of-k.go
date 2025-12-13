package arrays

// https://leetcode.com/problems/number-of-subarrays-with-and-value-of-k/

// Approach #2: Optimized with AND Property
// Time: O(n * log(max_val)) - for each position, update ~30 distinct values
// Space: O(log(max_val)) - store distinct AND values
func countSubarraysWithAndK(nums []int, k int) int64 {
	var count int64
	// prevAnds stores (and_value -> count_of_subarrays_with_this_value)
	prevAnds := make(map[int]int64)

	for _, num := range nums {
		// Start new subarray at current position
		currAnds := make(map[int]int64)
		currAnds[num] = 1

		// Extend all previous subarrays
		for andVal, cnt := range prevAnds {
			newAnd := andVal & num
			currAnds[newAnd] += cnt
		}

		// Count subarrays with AND = k
		count += currAnds[k]

		prevAnds = currAnds
	}

	return count
}

// // Approach #1: Brute Force
// // Time: O(n^2) - check all n*(n+1)/2 subarrays, TLE for large n
// // Space: O(1) - constant space
// func countSubarraysWithAndK(nums []int, k int) int64 {
// 	n := len(nums)
// 	count := int64(0)

// 	for i := range n {
// 		andVal := nums[i]
// 		for j := i; j < n; j++ {
// 			if i == j {
// 				andVal = nums[i]
// 			} else {
// 				andVal &= nums[j]
// 			}

// 			if andVal == k {
// 				count++
// 			}
// 		}
// 	}

// 	return count
// }
