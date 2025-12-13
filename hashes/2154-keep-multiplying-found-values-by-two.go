package hashes

// https://leetcode.com/problems/keep-multiplying-found-values-by-two/

// Approach: Hash Set Lookup
// Time: O(n + log(max_value)) - build map + lookups
// Space: O(n) - hash map
func findFinalValue(nums []int, original int) int {
	numSet := make(map[int]bool)
	for _, num := range nums {
		numSet[num] = true
	}

	for numSet[original] {
		original *= 2
	}

	return original
}
