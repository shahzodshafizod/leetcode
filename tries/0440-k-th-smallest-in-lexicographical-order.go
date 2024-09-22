package tries

// https://leetcode.com/problems/k-th-smallest-in-lexicographical-order/

// Approach: Prefix Tree
// Time: O((Log N)^2)
// Space: O(1)
func findKthNumber(n int, k int) int {
	var count = func(curr int) int { // O(Log N)
		neighbor := curr + 1 // neighboring branch of tree
		items := 0
		for curr <= n {
			// n+1: means include the last element
			// if it's in the left side, it works only once
			items += min(n+1, neighbor) - curr
			curr *= 10
			neighbor *= 10
		}
		return items
	}
	var curr = 1
	k--         // account for first element: curr=1
	for k > 0 { // O(Log N)
		items := count(curr)
		if k >= items {
			curr++ // go to the next branch of tree
			k -= items
		} else {
			curr *= 10
			k--
		}
	}
	return curr
}
