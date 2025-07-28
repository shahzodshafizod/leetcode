package hashes

// https://leetcode.com/problems/valid-anagram/

// Approach: Counter
// Time: O(S+T), S=len(s), T=len(t)
// Space: O(26) = O(1)
func isAnagram(s string, t string) bool {
	counter := make(map[rune]int)
	for _, c := range s {
		counter[c]++
	}

	for _, c := range t {
		counter[c]--
	}

	for _, cnt := range counter {
		if cnt != 0 {
			return false
		}
	}

	return true
}

// // Approach: Sorting
// // Time: O(SlogS+TlogT), S=len(s), T=len(t)
// // Space: O(S+T)
// func isAnagram(s string, t string) bool {
// 	sb := []byte(s)
// 	tb := []byte(t)
// 	sort.Slice(sb, func(i, j int) bool { return sb[i] < sb[j] })
// 	sort.Slice(tb, func(i, j int) bool { return tb[i] < tb[j] })
// 	return string(sb) == string(tb)
// }
