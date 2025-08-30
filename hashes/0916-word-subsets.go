package hashes

// hhttps://leetcode.com/problems/word-subsets/

func wordSubsets(words1 []string, words2 []string) []string {
	var maxFreq, freq [26]int

	for _, word := range words2 {
		for _, c := range word {
			freq[int(c-'a')]++
		}

		for idx := range 26 {
			maxFreq[idx] = max(maxFreq[idx], freq[idx])
			freq[idx] = 0
		}
	}

	subsets := make([]string, 0)

	var universal bool

	for _, word := range words1 {
		for _, c := range word {
			freq[int(c-'a')]++
		}

		universal = true

		for idx := range 26 {
			if freq[idx] < maxFreq[idx] {
				universal = false
			}

			freq[idx] = 0
		}

		if universal {
			subsets = append(subsets, word)
		}
	}

	return subsets
}
