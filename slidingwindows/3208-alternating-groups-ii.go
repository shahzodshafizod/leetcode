package slidingwindows

// https://leetcode.com/problems/alternating-groups-ii/

// Time: O(N)
// Space: O(1)
func numberOfAlternatingGroups(colors []int, k int) int {
	var groups, tails = 0, 1
	var n = len(colors)
	var next = (n + k - 2) % n
	var curr int
	for idx := n + k - 3; idx >= 0; idx-- {
		curr = idx % n
		if colors[curr] == colors[next] {
			tails = 0
		}
		tails++
		if tails >= k {
			groups++
		}
		next = curr
	}
	return groups
}
