package greedy

// https://leetcode.com/problems/minimum-length-of-string-after-operations/

// Approach: Greedy
// Time: O(n)
// Space: O(1)
func minimumLength(s string) int {
	var count [26]int
	for _, c := range s {
		count[int(c-'a')]++
	}
	length := 0
	for _, cnt := range count {
		if cnt&1 == 1 {
			length += 1
		} else if cnt != 0 {
			length += 2
		}
	}
	return length
}
