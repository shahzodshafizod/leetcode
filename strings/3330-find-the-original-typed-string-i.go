package strings

// https://leetcode.com/problems/find-the-original-typed-string-i/

func possibleStringCount(word string) int {
	count := 1
	for idx := len(word) - 1; idx > 0; idx-- {
		if word[idx-1] == word[idx] {
			count++
		}
	}
	return count
}
