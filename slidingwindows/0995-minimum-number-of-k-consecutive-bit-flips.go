package slidingwindows

// https://leetcode.com/problems/minimum-number-of-k-consecutive-bit-flips/

// Approach #3: Sliding Window + Dynamic Programming: Tabulation (Space Optimized)
// time: O(n)
// space: O(1)
func minKBitFlips(nums []int, k int) int {
	totalFlips := 0
	flips := 0
	n := len(nums)
	for idx := range nums {
		if nums[idx]&1 == flips&1 {
			if idx+k > n {
				return -1
			}
			nums[idx+k-1] |= 2
			flips++
			totalFlips++
		}
		flips -= nums[idx] >> 1
		// nums[idx] &= 1 // restore the original bit
		// nums[idx] = 1  // set the bit
	}
	return totalFlips
}

// // Approach #2: Sliding Window + Dynamic Programming: Tabulation
// // time: O(n)
// // space: O(n)
// func minKBitFlips(nums []int, k int) int {
// 	var totalFlips = 0
// 	var flips = 0
// 	var n = len(nums)
// 	var dp = make([]int, n)
// 	for idx := range nums {
// 		if nums[idx] == flips&1 {
// 			if idx+k > n {
// 				return -1
// 			}
// 			dp[idx+k-1] = 1
// 			flips++
// 			totalFlips++
// 		}
// 		flips -= dp[idx]
// 	}
// 	return totalFlips
// }

// // Approach #1: Brute Force
// // time: O(nk)
// // space: O(1)
// func minKBitFlips(nums []int, k int) int {
// 	var nexti int
// 	var flips = 0
// 	for i, n := 0, len(nums); i < n; i++ {
// 		if nums[i] == 0 {
// 			if i+k > n {
// 				return -1
// 			}
// 			nexti = i
// 			for j := i + k - 1; j >= i; j-- {
// 				if nums[j] == 0 {
// 					nums[j] = 1
// 				} else {
// 					nums[j] = 0
// 					nexti = j
// 				}
// 			}
// 			flips++
// 			i = nexti - 1
// 		}
// 	}
// 	return flips
// }
