package strings

// https://leetcode.com/problems/length-of-last-word/

func lengthOfLastWord(s string) int {
	right := len(s) - 1
	for s[right] == ' ' {
		right--
	}

	left := right
	for left >= 0 && s[left] != ' ' {
		left--
	}

	return right - left
}

// func lengthOfLastWord(s string) int {
// 	var count = 0
// 	for idx := len(s) - 1; idx >= 0; idx-- {
// 		switch s[idx] {
// 		case ' ':
// 			if count > 0 {
// 				return count
// 			}
// 		default:
// 			count++
// 		}
// 	}
// 	return count
// }

// func lengthOfLastWord(s string) int {
// 	words := strings.Fields(s)
// 	return len(words[len(words)-1])
// }
