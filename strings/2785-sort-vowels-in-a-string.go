package strings

// https://leetcode.com/problems/sort-vowels-in-a-string/

func sortVowels(s string) string {
	indices := map[rune]int{
		'A': 1, 'E': 2, 'I': 3, 'O': 4, 'U': 5,
		'a': 6, 'e': 7, 'i': 8, 'o': 9, 'u': 10,
	}
	vowels := []rune{' ', 'A', 'E', 'I', 'O', 'U', 'a', 'e', 'i', 'o', 'u'}
	rs := []rune(s)

	var freq [11]int
	for _, c := range rs {
		freq[indices[c]]++
	}

	idx := 1

	for i, c := range rs {
		if indices[c] > 0 {
			for freq[idx] == 0 {
				idx++
			}

			rs[i] = vowels[idx]
			freq[idx]--
		}
	}

	return string(rs)
}
