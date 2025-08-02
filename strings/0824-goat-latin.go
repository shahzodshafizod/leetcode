package strings

import "strings"

// https://leetcode.com/problems/goat-latin/

func toGoatLatin(sentence string) string {
	isVowel := map[byte]bool{
		'a': true, 'e': true, 'i': true, 'o': true, 'u': true,
		'A': true, 'E': true, 'I': true, 'O': true, 'U': true,
	}

	var (
		words = make([]string, 0)
		err   error
	)

	var a strings.Builder
	for _, word := range strings.Split(sentence, " ") {
		err = a.WriteByte('a')
		_ = err

		if !isVowel[word[0]] {
			word = word[1:] + string(word[0])
		}

		words = append(words, word+"ma"+a.String())
	}

	return strings.Join(words, " ")
}
