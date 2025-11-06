package binarysearch

import "sort"

// https://leetcode.com/problems/maximum-frequency-of-an-element-after-performing-operations-ii/

// // Approach 1: Brute Force
// // For each possible target value, count how many elements can be transformed to it
// // Time: O(n * range), where range is max(nums) - min(nums) + 2*k
// // Space: O(1)
// func maxFrequency(nums []int, k int, numOperations int) int {
// 	abs := func(x int) int {
// 		if x < 0 {
// 			return -x
// 		}

// 		return x
// 	}

// 	// Try all possible target values
// 	minVal := slices.Min(nums) - k
// 	maxVal := slices.Max(nums) + k

// 	var freq int

// 	for target := minVal; target <= maxVal; target++ {
// 		// Count elements already at target (no operation needed)
// 		atTarget := 0
// 		// Count elements that can be transformed to target (need operation)
// 		transformable := 0

// 		for _, num := range nums {
// 			if num == target {
// 				atTarget++
// 			} else if abs(num-target) <= k {
// 				transformable++
// 			}
// 		}

// 		// We can use up to numOperations to transform elements
// 		transformable = min(transformable, numOperations)

// 		// Total frequency = already at target + transformed to target
// 		freq = max(freq, atTarget+transformable)
// 	}

// 	return freq
// }

// Approach 2: Optimized with Sorting + Binary Search
// Time: O(n log n), Space: O(n)
func maxFrequency(nums []int, k int, numOperations int) int {
	// bisectLeft finds leftmost position where nums[i] >= target
	bisectLeft := func(nums []int, target int) int {
		left, right := 0, len(nums)
		for left < right {
			mid := (left + right) / 2
			if nums[mid] < target {
				left = mid + 1
			} else {
				right = mid
			}
		}

		return left
	}

	// bisectRight finds leftmost position where nums[i] > target
	bisectRight := func(nums []int, target int) int {
		left, right := 0, len(nums)
		for left < right {
			mid := (left + right) / 2
			if nums[mid] <= target {
				left = mid + 1
			} else {
				right = mid
			}
		}

		return left
	}

	// Sort the array
	sort.Ints(nums)

	var candidates []int

	counter := make(map[int]int) // Count original frequencies
	set := make(map[int]bool)

	for _, num := range nums {
		counter[num]++
		for _, val := range []int{num - k, num, num + k} {
			if !set[val] {
				candidates = append(candidates, val)
				set[val] = true
			}
		}
	}

	var freq int

	for _, target := range candidates {
		lid := bisectLeft(nums, target-k)
		rid := bisectRight(nums, target+k)

		atTarget := counter[target]
		inRange := rid - lid
		transformable := inRange - atTarget
		transformable = min(transformable, numOperations)

		freq = max(freq, atTarget+transformable)
	}

	return freq
}
