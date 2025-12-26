package prefixsums

// https://leetcode.com/problems/count-partitions-with-even-sum-difference/

// Approach 1: Brute Force - Calculate Sums for Each Partition
// For each partition index, calculate left sum and right sum
// Check if difference is even
// Time: O(n²) - for each partition, sum takes O(n)
// Space: O(1)

// // Approach 2: Prefix Sum (Optimal)
// // Use running sums to track left and right partition sums
// // For each partition, check if (leftSum - rightSum) is even
// // Key insight: leftSum - rightSum is even when both have same parity
// // or equivalently when (leftSum - rightSum) % 2 == 0
// // Time: O(n)
// // Space: O(1)
// func countPartitions(nums []int) int {
// 	n := len(nums)
// 	if n < 2 {
// 		return 0
// 	}

// 	// Calculate total sum for right partition
// 	rightSum := 0
// 	for _, num := range nums {
// 		rightSum += num
// 	}

// 	leftSum := 0
// 	count := 0

// 	// Try each partition point
// 	for i := range n - 1 {
// 		leftSum += nums[i]
// 		rightSum -= nums[i]

// 		// Check if difference is even
// 		diff := leftSum - rightSum
// 		if diff%2 == 0 {
// 			count++
// 		}
// 	}

// 	return count
// }

// Alternative: Check parity directly
func countPartitions(nums []int) int {
	n := len(nums)

	totalSum := 0
	for _, num := range nums {
		totalSum += num
	}

	leftSum := 0
	count := 0

	for i := range n - 1 {
		leftSum += nums[i]
		rightSum := totalSum - leftSum

		// leftSum - rightSum is even if both have same parity
		if (leftSum-rightSum)%2 == 0 {
			count++
		}
	}

	return count
}
