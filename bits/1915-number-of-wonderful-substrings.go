package bits

// https://leetcode.com/problems/number-of-wonderful-substrings/

func wonderfulSubstrings(word string) int64 {
	mask := 0 // current state
	freq := map[int]int64{mask: 1}
	var key int
	var result int64 = 0
	for _, letter := range word {
		// update state
		mask ^= 1 << (letter - 'a')
		// add count of same previous states
		result += freq[mask]
		for idx := 0; idx < 10; idx++ {
			key = mask ^ (1 << idx)
			if value, exists := freq[key]; exists {
				result += value
			}
		}
		// add 1 to count of times we've seen current state
		freq[mask]++
	}
	return result
}
