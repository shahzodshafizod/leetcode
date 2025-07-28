package hashes

// https://leetcode.com/problems/minimum-number-of-operations-to-make-elements-in-array-distinct/

// Approach: Reverse traversal
// Time: O(n)
// Space: O(n)
func minimumOperations(nums []int) int {
	visited := make(map[int]struct{})

	var exists bool

	for idx := len(nums) - 1; idx >= 0; idx-- {
		if _, exists = visited[nums[idx]]; exists {
			return (idx + 3) / 3
		}

		visited[nums[idx]] = struct{}{}
	}

	return 0
}
