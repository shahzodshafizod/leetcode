package hashes

// hhttps://leetcode.com/problems/word-subsets/

func wordSubsets(words1 []string, words2 []string) []string {
	var maxFreq, freq [26]int
	for _, word := range words2 {
		for _, c := range word {
			freq[int(c-'a')]++
		}
		for idx := 0; idx < 26; idx++ {
			maxFreq[idx] = max(maxFreq[idx], freq[idx])
			freq[idx] = 0
		}
	}
	var subsets = make([]string, 0)
	var universal bool
	for _, word := range words1 {
		for _, c := range word {
			freq[int(c-'a')]++
		}
		universal = true
		for idx := 0; idx < 26; idx++ {
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
