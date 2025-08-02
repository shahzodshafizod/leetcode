package hashes

// https://leetcode.com/problems/find-common-characters/

// n = len(words)
// k = max(len(words[i]))
// time: O(n * k)
// space: O(26 * 2) = O(1)
func commonChars(words []string) []string {
	const n = 26

	var common [n]int
	for _, c := range words[0] {
		common[c-'a']++
	}

	var current [n]int

	for i, wn := 1, len(words); i < wn; i++ {
		for _, c := range words[i] {
			current[c-'a']++
		}

		for i := 0; i < n; i++ {
			common[i] = min(common[i], current[i])
			current[i] = 0
		}
	}

	var result []string

	for letter, count := range common {
		for count > 0 {
			count--

			result = append(result, string(byte('a'+letter)))
		}
	}

	return result
}
