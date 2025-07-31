package slidingwindows

// https://leetcode.com/problems/count-of-substrings-containing-every-vowel-and-k-consonants-ii/

// Time: O(N)
// Space: O(1)
func countOfSubstrings(word string, k int) int64 {
	isVowel := map[byte]bool{
		'a': true,
		'e': true,
		'i': true,
		'o': true,
		'u': true,
	}
	n := len(word)
	atLeast := func(k int) int64 {
		var result int64

		var count [26]int

		var code int

		vowels, consonents := 0, 0
		left := 0

		for right := 0; right < n; right++ {
			code = int(word[right] - 'a')
			count[code]++

			if !isVowel[word[right]] {
				consonents++
			} else if count[code] == 1 {
				vowels++
			}

			for ; vowels == 5 && consonents >= k; left++ {
				result += int64(n - right)
				code = int(word[left] - 'a')
				count[code]--

				if !isVowel[word[left]] {
					consonents--
				} else if count[code] == 0 {
					vowels--
				}
			}
		}

		return result
	}

	return atLeast(k) - atLeast(k+1)
}
