package binarysearch

// https://leetcode.com/problems/smallest-substring-with-identical-characters-i/

// Approach #2: Binary Search on Answer
// Time: O(n * logn)
// Space: O(1)
func minLengthI(s string, numOps int) int {
	n := len(s)

	// Special case: alternating pattern
	// Try both "010101..." and "101010..."
	pattern1, pattern2 := 0, 0

	var digit int
	for i := range n {
		digit = int(s[i] - '0')
		if digit != i%2 {
			pattern1++
		}

		if digit != 1-i%2 {
			pattern2++
		}
	}

	if min(pattern1, pattern2) <= numOps {
		return 1
	}

	canAchieve := func(maxLen int) bool {
		ops, i := 0, 0
		for i < n {
			j := i
			for j < n && s[j] == s[i] {
				j++
			}

			segmentLen := j - i
			if segmentLen > maxLen {
				ops += segmentLen / (maxLen + 1)
			}

			i = j
		}

		return ops <= numOps
	}

	left, right := 2, n

	for left < right {
		mid := left + (right-left)/2
		if canAchieve(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return right
}

// // Approach #1: Brute Force - Try all possible lengths from 1 to n
// // Time: O(n^2)
// // Space: O(1)
// func minLengthI(s string, numOps int) int {
// 	n := len(s)

// 	// Special case: alternating pattern
// 	pattern1, pattern2 := 0, 0

// 	for i := range n {
// 		if s[i] != byte('0'+i%2) {
// 			pattern1++
// 		}

// 		if s[i] != byte('0'+(1-i%2)) {
// 			pattern2++
// 		}
// 	}

// 	if min(pattern1, pattern2) <= numOps {
// 		return 1
// 	}

// 	canAchieve := func(maxLen int) bool {
// 		ops, i := 0, 0
// 		for i < n {
// 			j := i
// 			for j < n && s[j] == s[i] {
// 				j++
// 			}

// 			segmentLen := j - i

// 			if segmentLen > maxLen {
// 				ops += segmentLen / (maxLen + 1)
// 			}

// 			i = j
// 		}

// 		return ops <= numOps
// 	}

// 	for length := 2; length <= n; length++ {
// 		if canAchieve(length) {
// 			return length
// 		}
// 	}

// 	return n
// }
