package greedy

// https://leetcode.com/problems/minimum-operations-to-make-array-equal-to-target/

// Approach #2: Greedy
// Time: O(n) - single pass through the array
// Space: O(1) - only using constant extra space
func minimumOperations(nums []int, target []int) int64 {
	abs := func(x int64) int64 {
		if x < 0 {
			return -x
		}

		return x
	}

	n := len(nums)
	currDiff := int64(target[0] - nums[0])
	operations := abs(currDiff)

	var prevDiff int64
	for i := 1; i < n; i++ {
		prevDiff = currDiff
		currDiff = int64(target[i] - nums[i])

		if prevDiff*currDiff >= 0 { // if signs are same
			// Same direction: add only the increase in magnitude
			operations += max(0, abs(currDiff)-abs(prevDiff))
		} else {
			// Different directions: need full magnitude
			operations += abs(currDiff)
		}
	}

	return operations
}

// // Approach #1: Brute-force (Simulation-based)
// // Time: O(n * max_diff) where max_diff is the maximum difference
// // Space: O(n) for creating a copy of nums array
// func minimumOperations(nums []int, target []int) int64 {
// 	n := len(nums)

// 	var operations int64

// 	var start, end, i, diff, increment int
// 	for !slices.Equal(nums, target) {
// 		i = 0
// 		for i < n && nums[i] == target[i] {
// 			i++
// 		}

// 		diff = target[i] - nums[i]

// 		if diff > 0 {
// 			increment = 1
// 		} else {
// 			increment = -1
// 		}

// 		start, end = i, i
// 		for end < n && (target[end]-nums[end])*increment > 0 {
// 			end++
// 		}

// 		for k := start; k < end; k++ {
// 			nums[k] += increment
// 		}

// 		operations++
// 	}

// 	return operations
// }
