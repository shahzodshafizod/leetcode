package prefixsums

// https://leetcode.com/problems/number-of-submatrices-that-sum-to-target/

// Approach: Prefix Sum + Hash Table
// Time: O(nnm)
// Space: O(m)
func numSubmatrixSumTarget(matrix [][]int, target int) int {
	var m, n = len(matrix), len(matrix[0])
	// 560. Subarray Sum Equals K
	// https://leetcode.com/problems/subarray-sum-equals-k/
	var subarraySum = func(nums []int, k int) int {
		var counter = map[int]int{0: 1}
		var count, presum = 0, 0
		for idx := 0; idx < m; idx++ {
			presum += nums[idx]
			count += counter[presum-k]
			counter[presum]++
		}
		return count
	}
	var count = 0
	for start := 0; start < n; start++ {
		var presum = make([]int, m)
		for end := start; end < n; end++ {
			for row := 0; row < m; row++ {
				presum[row] += matrix[row][end]
			}
			count += subarraySum(presum, target)
		}
	}
	return count
}

// // Approach: Prefix Sum
// // Time: O(nnmm)
// // Space: O(m)
// func numSubmatrixSumTarget(matrix [][]int, target int) int {
// 	var m, n = len(matrix), len(matrix[0])
//  // 560. Subarray Sum Equals K
// 	// https://leetcode.com/problems/subarray-sum-equals-k/
// 	var subarraySum = func(nums []int, k int) int {
// 		var count, presum = 0, 0
// 		for idx := 0; idx < m; idx++ {
// 			presum += nums[idx]
// 			subsum := presum
// 			for start := 0; start <= idx; start++ {
// 				if subsum == k {
// 					count++
// 				}
// 				subsum -= nums[start]
// 			}
// 		}
// 		return count
// 	}
// 	var count = 0
// 	for start := 0; start < n; start++ {
// 		var presum = make([]int, m)
// 		for end := start; end < n; end++ {
// 			for row := 0; row < m; row++ {
// 				presum[row] += matrix[row][end]
// 			}
// 			count += subarraySum(presum, target)
// 		}
// 	}
// 	return count
// }
