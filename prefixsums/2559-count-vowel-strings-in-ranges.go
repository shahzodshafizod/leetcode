package prefixsums

// https://leetcode.com/problems/count-vowel-strings-in-ranges/

// Approach: Prefix Sum
// Time: O(W+Q), W=len(words), Q=len(queries)
// Space: O(W)
func vowelStrings(words []string, queries [][]int) []int {
	vowels := map[byte]bool{'a': true, 'e': true, 'i': true, 'o': true, 'u': true}

	precount := make([]int, len(words)+1)
	for idx, word := range words {
		precount[idx+1] = precount[idx]
		if vowels[word[0]] && vowels[word[len(word)-1]] {
			precount[idx+1]++
		}
	}

	ans := make([]int, len(queries))

	var left, right int
	for idx := range queries {
		left, right = queries[idx][0], queries[idx][1]
		ans[idx] = precount[right+1] - precount[left]
	}

	return ans
}
