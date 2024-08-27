package slidingwindows

// https://leetcode.com/problems/substring-with-concatenation-of-all-words/

// Time: O(N)
// Space: O(N)
func findSubstring(s string, words []string) []int {
	var getCounter = func() map[string]int {
		var count = make(map[string]int)
		for _, word := range words {
			count[word]++
		}
		return count
	}
	var wordsLen = len(words)
	var wordLen = len(words[0])
	var windowLen = wordsLen * wordLen
	var n = len(s)
	var wordsFound int
	var substrings = make([]int, 0)
	for idx := 0; idx < wordLen; idx++ {
		var count = getCounter()
		wordsFound = 0
		var left = idx
		for right := left + wordLen; right <= n; right += wordLen {
			var word = s[right-wordLen : right]
			if _, exists := count[word]; exists {
				if count[word] > 0 {
					wordsFound++
				}
				count[word]--
			}
			if right-left == windowLen {
				if wordsFound == wordsLen {
					substrings = append(substrings, left)
				}
				word = s[left : left+wordLen]
				if _, exists := count[word]; exists {
					count[word]++
					if count[word] > 0 {
						wordsFound--
					}
				}
				left += wordLen
			}
		}
	}
	return substrings
}
