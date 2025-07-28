package strings

// https://leetcode.com/problems/valid-word/

func isValid(word string) bool {
	if len(word) < 3 {
		return false
	}

	isVowel := map[rune]bool{
		'a': true, 'e': true, 'i': true, 'o': true, 'u': true,
		'A': true, 'E': true, 'I': true, 'O': true, 'U': true,
	}
	isWrong := map[rune]bool{'@': true, '#': true, '$': true}
	hasVowel, hasConsonent := false, false

	for _, c := range word {
		if isWrong[c] {
			return false
		}

		if isVowel[c] {
			hasVowel = true
		} else {
			hasConsonent = true
		}
	}

	return hasVowel && hasConsonent
}
