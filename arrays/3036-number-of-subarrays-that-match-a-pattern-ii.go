package arrays

// https://leetcode.com/problems/number-of-subarrays-that-match-a-pattern-ii/

// Approach #2: Optimized - Convert to pattern and use KMP
// Convert nums to comparison pattern, then use KMP for pattern matching
// Time: O(n + m) using KMP algorithm
// Space: O(n + m) for converted array and LPS table
func countMatchingSubarrays(nums []int, pattern []int) int {
	// Convert nums to pattern representation
	n := len(nums)
	numsPattern := make([]int, n-1)

	for i := range n - 1 {
		if nums[i+1] > nums[i] {
			numsPattern[i] = 1
		} else if nums[i+1] == nums[i] {
			numsPattern[i] = 0
		} else {
			numsPattern[i] = -1
		}
	}

	// KMP search
	kmpSearch := func(text []int, pattern []int) int {
		if len(pattern) == 0 {
			return 0
		}

		// Build LPS array
		lps := make([]int, len(pattern))
		length := 0
		i := 1

		for i < len(pattern) {
			if pattern[i] == pattern[length] {
				length++
				lps[i] = length
				i++
			} else {
				if length != 0 {
					length = lps[length-1]
				} else {
					lps[i] = 0
					i++
				}
			}
		}

		// Search for pattern
		count := 0
		i = 0  // text index
		j := 0 // pattern index

		for i < len(text) {
			if text[i] == pattern[j] {
				i++
				j++
			}

			if j == len(pattern) {
				count++
				j = lps[j-1]
			} else if i < len(text) && text[i] != pattern[j] {
				if j != 0 {
					j = lps[j-1]
				} else {
					i++
				}
			}
		}

		return count
	}

	return kmpSearch(numsPattern, pattern)
}

// // Approach #1: Brute Force - Check each subarray
// // For each position, check if pattern matches the comparison pattern
// // Time: O(n * m) where n = len(nums), m = len(pattern)
// // Space: O(1)
// func countMatchingSubarrays(nums []int, pattern []int) int {
// 	n := len(nums)
// 	m := len(pattern)
// 	count := 0

// 	for i := 0; i <= n-m-1; i++ {
// 		match := true

// 		for j := range m {
// 			if pattern[j] == 1 {
// 				if nums[i+j+1] <= nums[i+j] {
// 					match = false

// 					break
// 				}
// 			} else if pattern[j] == 0 {
// 				if nums[i+j+1] != nums[i+j] {
// 					match = false

// 					break
// 				}
// 			} else { // pattern[j] == -1
// 				if nums[i+j+1] >= nums[i+j] {
// 					match = false

// 					break
// 				}
// 			}
// 		}

// 		if match {
// 			count++
// 		}
// 	}

// 	return count
// }
