package greedy

// https://leetcode.com/problems/longest-palindrome/

// time: O(n)
// space: O(52) = O(1)
func longestPalindrome(s string) int {
	seen := make(map[rune]bool)
	length := 0

	for _, c := range s {
		if seen[c] {
			length += 2
		}

		seen[c] = !seen[c]
	}

	if len(s) > length {
		length++
	}

	return length
}

// // time: O(n)
// // space: O(52) = O(1)
// func longestPalindrome(s string) int {
// 	var freq = make(map[rune]int)
// 	var oddFreq = 0
// 	for _, c := range s {
// 		freq[c]++
// 		if freq[c]&1 == 1 {
// 			oddFreq++
// 		} else {
// 			oddFreq--
// 		}
// 	}
// 	if oddFreq == 0 {
// 		return len(s)
// 	}
// 	return len(s) - oddFreq + 1
// }

// // time: O(2n) = O(n)
// // space: O(52) = O(1)
// func longestPalindrome(s string) int {
// 	var freq = make(map[rune]int)
// 	for _, c := range s {
// 		freq[c]++
// 	}
// 	var maxlen = 0
// 	for _, freq := range freq {
// 		if freq&1 == 1 {
// 			if maxlen&1 == 0 {
// 				maxlen++
// 			}
// 			freq--
// 		}
// 		maxlen += freq
// 	}
// 	return maxlen
// }
