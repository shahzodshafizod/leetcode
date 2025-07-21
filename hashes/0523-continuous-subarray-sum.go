package hashes

// https://leetcode.com/problems/continuous-subarray-sum/

// time: O(n)
// space: O(n)
func checkSubarraySum(nums []int, k int) bool {
	modsum := 0
	prefixes := map[int]int{modsum: -1}
	var prev int
	var ok bool
	for idx := range nums {
		modsum = (modsum + nums[idx]) % k
		if prev, ok = prefixes[modsum]; !ok {
			prefixes[modsum] = idx
		} else if idx-prev > 1 {
			return true
		}
	}
	return false
}

// // time: O(2 ^ n)
// // space: O(2 ^ n)
// func checkSubarraySum(nums []int, k int) bool {
// 	var n = len(nums)
// 	var dfs func(start int, end int, sum int) bool
// 	dfs = func(start int, end int, sum int) bool {
// 		if end == n {
// 			return false
// 		}
// 		sum += nums[end]
// 		if sum%k == 0 {
// 			return true
// 		}
// 		// decision to include nums[start]
// 		return dfs(start, end+1, sum) ||
// 			// decision NOT to include nums[start]
// 			dfs(start+1, end+1, sum-nums[start])
// 	}
// 	return dfs(0, 1, nums[0])
// }

// // time: O(n ^ 2)
// // space: O(1)
// func checkSubarraySum(nums []int, k int) bool {
// 	var sum = 0
// 	for i, n := 0, len(nums); i < n; i++ {
// 		sum = nums[i]
// 		for j := i + 1; j < n; j++ {
// 			sum += nums[j]
// 			if sum%k == 0 {
// 				return true
// 			}
// 		}
// 	}
// 	return false
// }
