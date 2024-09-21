package strings

// https://leetcode.com/problems/find-the-index-of-the-first-occurrence-in-a-string/

// Approach: Rabin Karp Algorithm
// Time: O(M+N)
// Space: O(1)
func strStr(haystack string, needle string) int {
	const base, mod int = 29, 1e9 + 7
	var subs = 0
	var digit int
	var power = 1
	for idx, c := range needle {
		digit = int(c-'a') + 1
		subs = (subs*base + digit) % mod
		if idx > 0 {
			power = (power * base) % mod
		}
	}
	var m = len(needle)
	var s = 0
	for idx, c := range haystack {
		if idx >= m {
			digit = int(haystack[idx-m]-'a') + 1
			s = (s + mod - (digit*power)%mod) % mod
		}
		digit = int(c-'a') + 1
		s = (s*base + digit) % mod
		if s == subs {
			return idx + 1 - m
		}
	}
	return -1
}

// // Approach: KMP (Knuth-Morris-Pratt) Algorithm
// // Time: O(M+N)
// // Space: O(N)
// func strStr(haystack string, needle string) int {
// 	var m = len(needle)
// 	lps := make([]int, m) // longest prefix suffix
// 	lps[0] = 0            // there's no prefix for first char
// 	var prevLPS = 0
// 	for curr := 1; curr < m; { // O(2*m)
// 		if needle[curr] == needle[prevLPS] {
// 			lps[curr] = prevLPS + 1
// 			prevLPS++
// 			curr++
// 		} else if prevLPS == 0 {
// 			lps[curr] = 0
// 			curr++
// 		} else {
// 			prevLPS = lps[prevLPS-1]
// 		}
// 	}
// 	var n = len(haystack)
// 	var hi, ni = 0, 0
// 	for hi < n {
// 		if haystack[hi] == needle[ni] {
// 			hi++
// 			ni++
// 		} else if ni == 0 {
// 			hi++
// 		} else {
// 			ni = lps[ni-1]
// 		}
// 		if ni == m {
// 			return hi - m
// 		}
// 	}
// 	return -1
// }

// func strStr(haystack string, needle string) int {
// 	return strings.Index(haystack, needle)
// }
