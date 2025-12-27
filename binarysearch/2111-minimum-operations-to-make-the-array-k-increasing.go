package binarysearch

// https://leetcode.com/problems/minimum-operations-to-make-the-array-k-increasing/

// Approach #2: Binary Search + Longest Non-Decreasing Subsequence (LIS variant)
// Time: O(n log(n/k)) - for each of k subsequences of length n/k, we do binary search
// Space: O(n/k) - for storing the LIS array
func kIncreasing(arr []int, k int) int {
	longestNonDecreasingSubsequence := func(subseq []int) int {
		// Find length of longest non-decreasing subsequence using binary search
		if len(subseq) == 0 {
			return 0
		}

		// lis[i] stores the smallest tail element for non-decreasing subseq of length i+1
		lis := make([]int, 0)

		for _, num := range subseq {
			// Binary search for the position to insert/replace
			// We want rightmost position where lis[pos] <= num
			left, right := 0, len(lis)
			for left < right {
				mid := left + (right-left)/2
				if lis[mid] > num {
					right = mid
				} else {
					left = mid + 1
				}
			}

			// If we're at the end, append
			if right == len(lis) {
				lis = append(lis, num)
			} else {
				// Replace element at position right
				lis[right] = num
			}
		}

		return len(lis)
	}

	totalOps := 0
	// Split array into k subsequences based on i % k
	for remainder := range k {
		subseq := []int{}
		for idx := remainder; idx < len(arr); idx += k {
			subseq = append(subseq, arr[idx])
		}

		// Minimum operations = total - longest non-decreasing subsequence
		totalOps += len(subseq) - longestNonDecreasingSubsequence(subseq)
	}

	return totalOps
}

// // Approach #1: Brute-Force (Dynamic Programming for each subsequence)
// // Time: O(k * (n/k)^2) = O(n^2 / k)
// // Space: O(n/k)
// func kIncreasing(arr []int, k int) int {
// 	minOpsForNonDecreasing := func(subseq []int) int {
// 		// Find minimum operations to make subsequence non-decreasing using DP
// 		n := len(subseq)
// 		if n <= 1 {
// 			return 0
// 		}

// 		// dp[i] = length of longest non-decreasing subsequence ending at index i
// 		dp := make([]int, n)
// 		for i := range dp {
// 			dp[i] = 1
// 		}

// 		for i := 1; i < n; i++ {
// 			for j := 0; j < i; j++ {
// 				if subseq[j] <= subseq[i] {
// 					if dp[j]+1 > dp[i] {
// 						dp[i] = dp[j] + 1
// 					}
// 				}
// 			}
// 		}

// 		// Maximum length of non-decreasing subsequence
// 		maxLen := 0
// 		for _, length := range dp {
// 			if length > maxLen {
// 				maxLen = length
// 			}
// 		}

// 		// Minimum operations = total elements - elements we can keep
// 		return n - maxLen
// 	}

// 	totalOps := 0
// 	// Split array into k subsequences based on i % k
// 	for remainder := 0; remainder < k; remainder++ {
// 		subseq := []int{}
// 		for idx := remainder; idx < len(arr); idx += k {
// 			subseq = append(subseq, arr[idx])
// 		}
// 		totalOps += minOpsForNonDecreasing(subseq)
// 	}

// 	return totalOps
// }
