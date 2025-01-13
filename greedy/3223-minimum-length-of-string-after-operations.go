package greedy

// https://leetcode.com/problems/minimum-length-of-string-after-operations/

// Approach: Greedy
// Time: O(n)
// Space: O(1)
func minimumLength(s string) int {
	var count [26]int
	var length = len(s)
	var key int
	for _, c := range s {
		key = int(c - 'a')
		count[key]++
		if count[key] > 2 {
			count[key] -= 2
			length -= 2
		}
	}
	return length
}

