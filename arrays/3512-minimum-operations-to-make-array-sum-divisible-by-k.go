package arrays

// https://leetcode.com/problems/minimum-operations-to-make-array-sum-divisible-by-k/

// Approach: Calculate remainder and operations needed
// Time: O(n) - sum calculation
// Space: O(1) - constant space
func minOperations(nums []int, k int) int {
	total := 0
	for _, num := range nums {
		total += num
	}

	return total % k
}
