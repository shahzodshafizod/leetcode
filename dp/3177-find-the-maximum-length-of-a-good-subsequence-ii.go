package dp

// https://leetcode.com/problems/find-the-maximum-length-of-a-good-subsequence-ii/

// Approach #3: Bottom-Up Dynamic Programming (Space Optimized)
// Time: O(nk)
// Space: O(unique_nums * k)
func maximumLength3177(nums []int, k int) int {
	dp := make([]map[int]int, k+1)
	for i := range dp {
		dp[i] = make(map[int]int)
	}

	res := make([]int, k+1)
	for _, num := range nums {
		for count := k; count >= 0; count-- {
			dp[count][num]++
			if count > 0 {
				dp[count][num] = max(dp[count][num], res[count-1]+1)
			}

			res[count] = max(res[count], dp[count][num])
		}
	}

	return res[k]
}

// // Approach #2: Top-Down Dynamic Programming
// // Time: O(nk)
// // Space: O(nk)
// func maximumLength3177(nums []int, k int) int {
// 	indices := make(map[int][]int)
// 	for idx, num := range nums {
// 		indices[num] = append(indices[num], idx)
// 	}
// 	n := len(nums)
// 	memo := make([][][2]*int, n)
// 	for i := range memo {
// 		memo[i] = make([][2]*int, k+1)
// 	}
// 	var dp func(i int, k int, same int) int
// 	dp = func(i int, k int, same int) int {
// 		if i == n {
// 			return 0
// 		}
// 		if memo[i][k][same] != nil {
// 			return *memo[i][k][same]
// 		}
// 		var take, skip int
// 		v := indices[nums[i]]
// 		left, right := 0, len(v)-1
// 		var mid int
// 		for left <= right {
// 			mid = left + (right-left)/2
// 			if v[mid] > i {
// 				right = mid - 1
// 			} else {
// 				left = mid + 1
// 			}
// 		}
// 		// check if i was the last index of nums[i]
// 		if left == len(v) {
// 			take = 1
// 		} else {
// 			take = 1 + dp(v[left], k, 1)
// 		}
// 		if k > 0 {
// 			take = max(take, 1+dp(i+1, k-1, 0))
// 		}
// 		if same == 0 {
// 			skip = dp(i+1, k, 0)
// 		}
// 		res := max(skip, take)
// 		memo[i][k][same] = &res
// 		return res
// 	}
// 	return dp(0, k, 0)
// }

// // Approach #1: Top-Down Dynamic Programming (TLE)
// // Time: O(nnk)
// // Space: O(nk)
// func maximumLength3177(nums []int, k int) int {
// 	n := len(nums)
// 	memo := make([][]*int, n)
// 	for i := range memo {
// 		memo[i] = make([]*int, k+1)
// 	}
// 	var dp func(i int, k int) int
// 	dp = func(i, k int) int {
// 		if i == n {
// 			return 0
// 		}
// 		if memo[i][k] != nil {
// 			return *memo[i][k]
// 		}
// 		length := 1
// 		for j := i + 1; j < n; j++ {
// 			if nums[j] == nums[i] {
// 				length = max(length, 1+dp(j, k))
// 			} else if k >= 1 {
// 				length = max(length, 1+dp(j, k-1))
// 			}
// 		}
// 		memo[i][k] = &length
// 		return length
// 	}
// 	return dp(0, k)
// }
