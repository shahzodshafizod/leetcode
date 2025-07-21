package hashes

// https://leetcode.com/problems/minimum-operations-to-make-array-values-equal-to-k/

// Approach: Hash map
// Time: O(n)
// Space: O(n)
func minOperations(nums []int, k int) int {
	visited := make(map[int]struct{})
	for _, num := range nums {
		if num < k {
			return -1
		} else if num > k {
			visited[num] = struct{}{}
		}
	}
	return len(visited)
}
