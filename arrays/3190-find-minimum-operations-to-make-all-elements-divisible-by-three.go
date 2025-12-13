package arrays

// https://leetcode.com/problems/find-minimum-operations-to-make-all-elements-divisible-by-three/

// Approach: Count Remainders
// Time: O(n)
// Space: O(1)
func minimumOperations(nums []int) int {
	operations := 0

	for _, num := range nums {
		if num%3 != 0 {
			operations++
		}
	}

	return operations
}
