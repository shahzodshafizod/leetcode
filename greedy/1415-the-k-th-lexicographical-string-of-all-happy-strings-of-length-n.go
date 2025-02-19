package greedy

// https://leetcode.com/problems/the-k-th-lexicographical-string-of-all-happy-strings-of-length-n/

// Approach: Combinatorics
// Time: O(n)
// Space: O(1)
func getHappyString(n int, k int) string {
	var happy = make([]rune, 0, n)
	var partitionLen = 1 << n // 2^n
	var template = "abc"
	var i int
	var prev rune = -1
	for idx := 0; idx < n; idx++ {
		partitionLen /= 2
		i = 0
		for _, c := range template {
			if prev == c {
				continue
			}
			if k <= (i+1)*partitionLen {
				happy = append(happy, c)
				prev = c
				k -= i * partitionLen
				break
			}
			i++
		}
	}
	if len(happy) == 0 {
		return ""
	}
	return string(happy)
}
