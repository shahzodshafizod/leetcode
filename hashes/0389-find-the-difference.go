package hashes

// https://leetcode.com/problems/find-the-difference/

// Approach: Bit Manipulation
// Time: O(n)
// Space: O(1)
func findTheDifference(s string, t string) byte {
	var extra byte = 0
	for idx := range s {
		extra ^= s[idx]
		extra ^= t[idx]
	}

	extra ^= t[len(t)-1]

	return extra
}

// // Approach: Hash Table
// // Time: O(n)
// // Space: O(26) = O(1)
// func findTheDifference(s string, t string) byte {
// 	var count = make(map[byte]int, 26)
// 	for idx := range s {
// 		count[s[idx]]++
// 		count[t[idx]]++
// 	}
// 	count[t[len(t)-1]]++
// 	var key byte
// 	for key = range count {
// 		if count[key] == 1 {
// 			break
// 		}
// 	}
// 	return key
// }
