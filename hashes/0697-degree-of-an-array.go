package hashes

// https://leetcode.com/problems/degree-of-an-array/

// Approach #2: Optimized - Hash Table with first/last occurrence tracking
// Track frequency, first occurrence, and last occurrence for each element
// For elements with max frequency, find minimum (last - first + 1)
// N = length of array
// Time: O(N) - single pass through array
// Space: O(N) - hash tables for tracking
func findShortestSubArray(nums []int) int {
	// Track frequency, first occurrence, last occurrence
	freq := make(map[int]int)
	first := make(map[int]int)
	last := make(map[int]int)

	for i, num := range nums {
		freq[num]++
		if _, exists := first[num]; !exists {
			first[num] = i
		}

		last[num] = i
	}

	// Find the degree (max frequency)
	degree := 0
	for _, count := range freq {
		degree = max(degree, count)
	}

	// Find minimum length subarray for elements with max frequency
	minLength := len(nums)

	for num, count := range freq {
		if count == degree {
			length := last[num] - first[num] + 1
			minLength = min(minLength, length)
		}
	}

	return minLength
}

// // Approach #1: Brute Force - Try all subarrays
// // For each subarray, calculate its degree and check if it matches array degree
// // N = length of array
// // Time: O(N^3) - N^2 subarrays, O(N) to calculate degree for each
// // Space: O(N) - frequency map for each subarray
// func findShortestSubArray(nums []int) int {
// 	// First, find the degree of the full array
// 	fullFreq := make(map[int]int)
// 	for _, num := range nums {
// 		fullFreq[num]++
// 	}
// 	fullDegree := 0
// 	for _, count := range fullFreq {
// 		fullDegree = max(fullDegree, count)
// 	}

// 	n := len(nums)
// 	minLength := n // Worst case: entire array

// 	// Try all subarrays
// 	for start := range n {
// 		freq := make(map[int]int)
// 		for end := start; end < n; end++ {
// 			// Add nums[end] to current subarray
// 			freq[nums[end]]++

// 			// Check if this subarray has the same degree
// 			currentDegree := 0
// 			for _, count := range freq {
// 				currentDegree = max(currentDegree, count)
// 			}
// 			if currentDegree == fullDegree {
// 				minLength = min(minLength, end-start+1)
// 			}
// 		}
// 	}

// 	return minLength
// }
