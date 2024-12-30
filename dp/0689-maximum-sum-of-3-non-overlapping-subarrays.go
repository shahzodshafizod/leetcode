package dp

// https://leetcode.com/problems/maximum-sum-of-3-non-overlapping-subarrays/

// Approach #3: Prefix Sum + Sliding Window
// Time: O(n)
// Space: O(n)
func maxSumOfThreeSubarrays(nums []int, k int) []int {
	var n = len(nums)
	var sums = make([]int, n)
	for idx := 0; idx < k; idx++ {
		sums[0] += nums[idx]
	}
	var right = make([][2]int, n) // {sum, start_idx}
	right[0][0] = sums[0]
	for idx := 1; idx <= n-k; idx++ {
		sums[idx] = sums[idx-1] - nums[idx-1] + nums[idx+k-1]
		right[idx][0] = sums[idx]
		right[idx][1] = idx
	}
	for idx := n - k - 1; idx >= 0; idx-- {
		if sums[idx] >= right[idx+1][0] {
			right[idx] = [2]int{sums[idx], idx}
		} else {
			right[idx] = right[idx+1]
		}
	}
	var total, indices = 0, []int{}
	var left, lid = 0, -1
	for idx := k; idx <= n-2*k; idx++ {
		if sums[idx-k] > left {
			left = sums[idx-k]
			lid = idx - k
		}
		if left+sums[idx]+right[idx+k][0] > total {
			total = left + sums[idx] + right[idx+k][0]
			indices = []int{lid, idx, right[idx+k][1]}
		}
	}
	return indices
}

// // Approach #2: Top-Down Dynamic Programming
// // Time: O(3n) = O(n)
// // Space: O(3n) = O(n)
// func maxSumOfThreeSubarrays(nums []int, k int) []int {
// 	var n = len(nums)
// 	var sums = make([]int, n)
// 	for idx := 0; idx < k; idx++ {
// 		sums[0] += nums[idx]
// 	}
// 	for idx := 1; idx <= n-k; idx++ {
// 		sums[idx] = sums[idx-1] - nums[idx-1] + nums[idx+k-1]
// 	}
// 	type SubArr struct {
// 		total   int
// 		indices []int
// 	}
// 	var memo = make([][4]*SubArr, n)
// 	var dp func(idx int, quota int) SubArr
// 	dp = func(idx int, quota int) SubArr {
// 		if idx+k > n || quota == 0 {
// 			return SubArr{}
// 		}
// 		if memo[idx][quota] != nil {
// 			return *memo[idx][quota]
// 		}
// 		var include = dp(idx+k, quota-1)
// 		include.total += sums[idx]
// 		include.indices = append([]int{idx}, include.indices...)
// 		var skip = dp(idx+1, quota)
// 		if include.total >= skip.total {
// 			memo[idx][quota] = &include
// 		} else {
// 			memo[idx][quota] = &skip
// 		}
// 		return *memo[idx][quota]
// 	}
// 	return dp(0, 3).indices
// }

// // Approach #1: Brute-Force
// // Time: O(n^3)
// // Space: O(1)
// func maxSumOfThreeSubarrays(nums []int, k int) []int {
// 	var n = len(nums)
// 	var sums = make([]int, n)
// 	for idx := 0; idx < k; idx++ {
// 		sums[0] += nums[idx]
// 	}
// 	for idx := 1; idx <= n-k; idx++ {
// 		sums[idx] = sums[idx-1] - nums[idx-1] + nums[idx+k-1]
// 	}
// 	var total, indices = 0, []int{}
// 	for a := 0; a < n-2*k; a++ {
// 		for b := a + k; b <= n-k; b++ {
// 			for c := b + k; c < n; c++ {
// 				if sums[a]+sums[b]+sums[c] > total {
// 					total = sums[a] + sums[b] + sums[c]
// 					indices = []int{a, b, c}
// 				}
// 			}
// 		}
// 	}
// 	return indices
// }
