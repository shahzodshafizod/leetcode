package strings

import "strings"

// https://leetcode.com/problems/reverse-prefix-of-word/

func reversePrefix(word string, ch byte) string {
	end := strings.IndexByte(word, ch)
	if end >= 0 {
		reversed := make([]byte, 0, end+1)
		for idx := end; idx >= 0; idx-- {
			reversed = append(reversed, word[idx])
		}
		word = string(reversed) + word[end+1:]
	}
	return word
}
