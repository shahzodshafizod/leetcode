package hashes

// https://leetcode.com/problems/set-mismatch/

// Approach 3: In-place Marking (Most Optimal for Space)
// Intuition: Use array indices to mark seen numbers by negating values.
// When we encounter a number whose corresponding index is already negative,
// that number is the duplicate. After marking, the positive value's index
// indicates the missing number.
// Time Complexity: O(n) - two passes through array
// Space Complexity: O(1) - modifying input array in-place
func findErrorNums(nums []int) []int {
	abs := func(x int) int {
		if x < 0 {
			return -x
		}

		return x
	}

	duplicate := -1

	// Mark seen numbers by negating the value at their index
	for _, num := range nums {
		index := abs(num) - 1 // Convert to 0-indexed
		if nums[index] < 0 {
			// Already marked, this is the duplicate
			duplicate = abs(num)
		} else {
			// Mark as seen
			nums[index] = -nums[index]
		}
	}

	// Find the missing number (index that's still positive)
	missing := -1

	for i := range nums {
		if nums[i] > 0 {
			missing = i + 1

			break
		}
	}

	return []int{duplicate, missing}
}

// // Approach 2: Mathematical (Optimized)
// // Intuition: Use mathematical properties to find duplicate and missing.
// // Let duplicate = d, missing = m, array length = n
// // - Expected sum = n*(n+1)/2
// // - Expected sum of squares = n*(n+1)*(2n+1)/6
// // - actual_sum - expected_sum = d - m ... (equation 1)
// // - actual_sum_sq - expected_sum_sq = d^2 - m^2 = (d-m)(d+m) ... (equation 2)
// // From eq2/eq1: d + m = diff_sum_sq / diff_sum
// // Solving: d = ((d-m) + (d+m))/2, m = ((d+m) - (d-m))/2
// // Time Complexity: O(n) - single pass to calculate sums
// // Space Complexity: O(1) - only using a few variables
// func findErrorNums(nums []int) []int {
// 	n := len(nums)
// 	expectedSum := n * (n + 1) / 2
// 	expectedSumSq := n * (n + 1) * (2*n + 1) / 6

// 	actualSum := 0
// 	actualSumSq := 0

// 	for _, num := range nums {
// 		actualSum += num
// 		actualSumSq += num * num
// 	}

// 	diffSum := actualSum - expectedSum       // d - m
// 	diffSumSq := actualSumSq - expectedSumSq // d^2 - m^2

// 	sumDM := diffSumSq / diffSum // d + m

// 	duplicate := (diffSum + sumDM) / 2
// 	missing := sumDM - duplicate

// 	return []int{duplicate, missing}
// }

// // Approach 1: Hash Table (Brute Force)
// // Intuition: Count frequency of each number using a map.
// // The duplicate appears twice, and the missing number doesn't appear.
// // Time Complexity: O(n) - single pass to build frequency map, single pass to find results
// // Space Complexity: O(n) - hash map stores up to n entries
// func findErrorNums(nums []int) []int {
// 	freq := make(map[int]int)
// 	duplicate, missing := 0, 0

// 	// Count frequencies
// 	for _, num := range nums {
// 		freq[num]++
// 	}

// 	// Find duplicate and missing
// 	for i := 1; i <= len(nums); i++ {
// 		switch freq[i] {
// 		case 0:
// 			missing = i
// 		case 2:
// 			duplicate = i
// 		}
// 	}

// 	return []int{duplicate, missing}
// }
