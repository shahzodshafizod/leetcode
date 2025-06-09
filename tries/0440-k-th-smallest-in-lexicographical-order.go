package tries

// https://leetcode.com/problems/k-th-smallest-in-lexicographical-order/

// Approach: Prefix Tree
// Time: O((log(10)n)^2)
// Space: O(1)
func findKthNumber(n int, k int) int {
	var count = func(curr int) int { // O(log(10)n)
		var nei = curr + 1 // neighboring branch of tree
		var cnt = 0
		for curr <= n {
			cnt += min(n+1, nei) - curr
			curr *= 10
			nei *= 10
		}
		return cnt
	}
	var curr = 1
	k-- // account for first element: curr=1
	var cnt int
	for k > 0 { // O(log(10)n)
		cnt = count(curr)
		if cnt <= k {
			curr++ // go to the next branch of tree
			k -= cnt
		} else {
			curr *= 10 // go down deeper in this branch
			k--        // accouting this node (curr)
		}
	}
	return curr
}
