package hashes

// https://leetcode.com/problems/ransom-note/

// Approach: Counting via Hash Table
// Time: O(n)
// Space: O(26) = O(1)
func canConstruct0383(ransomNote string, magazine string) bool {
	var count [26]int
	for _, c := range magazine {
		count[c-'a']++
	}
	var code int
	for _, c := range ransomNote {
		code = int(c - 'a')
		count[code]--
		if count[code] < 0 {
			return false
		}
	}
	return true
}
