package hashes

// https://leetcode.com/problems/longest-palindrome-by-concatenating-two-letter-words/

// Approach #2: Array
func longestPalindrome(words []string) int {
	var count [26 * 26]int
	length, unpaired := 0, 0
	var a, b, word, revw int
	for idx := range words {
		a = int(words[idx][0] - 'a')
		b = int(words[idx][1] - 'a')
		word, revw = a*26+b, b*26+a

		if count[revw] > 0 {
			length += 4
			count[revw]--
			if word == revw {
				unpaired--
			}
		} else {
			count[word]++
			if word == revw {
				unpaired++
			}
		}
	}
	if unpaired > 0 {
		length += 2 // center
	}
	return length
}

// // Approach #1: Hash Table
// func longestPalindrome(words []string) int {
// 	var length = 0
// 	var count = make(map[[2]byte]int)
// 	var word, revw [2]byte
// 	for idx := range words {
// 		word[0], word[1] = words[idx][0], words[idx][1]
// 		revw[0], revw[1] = words[idx][1], words[idx][0]
// 		if count[revw] > 0 {
// 			count[revw]--
// 			length += 4
// 		} else {
// 			count[word]++
// 		}
// 	}
// 	for word, cnt := range count {
// 		if cnt > 0 && word[0] == word[1] {
// 			length += 2
// 			break
// 		}
// 	}
// 	return length
// }
