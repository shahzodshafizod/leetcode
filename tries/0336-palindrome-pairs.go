package tries

// https://leetcode.com/problems/palindrome-pairs/

// Approach: HashMap with Reversed Words and Palindrome Checking
// Key insight: For words[i] + words[j] to be a palindrome, we need:
// - Case 1: reverse(words[i]) == words[j] (equal length palindromes)
// - Case 2: words[i] = prefix + suffix, where suffix is palindrome and reverse(prefix) exists
// - Case 3: words[j] = prefix + suffix, where prefix is palindrome and reverse(suffix) exists
//
// Algorithm:
// 1. Store all words in a hashmap with their reversed form as key
// 2. For each word, split it into all possible (prefix, suffix) combinations
// 3. Check:
//   - If suffix is palindrome and reverse(prefix) exists in map → valid pair
//   - If prefix is palindrome and reverse(suffix) exists in map → valid pair
//
// 4. Handle edge cases with empty strings
//
// Example: words = ["abcd", "dcba", "lls", "s", "sssll"]
// - "abcd" + "dcba" = "abcddcba" (palindrome)
// - "dcba" + "abcd" = "dcbaabcd" (palindrome)
// - "lls" split into prefixes and check if reverse exists
//
// Time: O(n * k^2) where n is number of words, k is average word length
//
//	For each word (n), we check all splits (k) and verify palindrome (k)
//
// Space: O(n * k) for the hashmap storing all words
func palindromePairs(words []string) [][]int {
	isPalindrome := func(s string) bool {
		for left, right := 0, len(s)-1; left < right; left, right = left+1, right-1 {
			if s[left] != s[right] {
				return false
			}
		}

		return true
	}

	reverseString := func(s string) string {
		b := []byte(s)
		for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
			b[i], b[j] = b[j], b[i]
		}

		return string(b)
	}

	// Build hashmap: word -> index
	wordMap := make(map[string]int, len(words))
	for i, word := range words {
		wordMap[word] = i
	}

	result := make([][]int, 0)

	for i, word := range words {
		n := len(word)
		// Check all possible splits of current word
		for j := 0; j <= n; j++ {
			// Case 1: suffix is palindrome, check if reverse(prefix) exists
			if isPalindrome(word[j:]) {
				reversedPrefix := reverseString(word[:j])
				if idx, exists := wordMap[reversedPrefix]; exists && idx != i {
					result = append(result, []int{i, idx})
				}
			}

			// Case 2: prefix is palindrome, check if reverse(suffix) exists
			// Avoid duplicate when j == len(word) (empty suffix creates same pair as Case 1 with j==0)
			if j > 0 && isPalindrome(word[:j]) {
				reversedSuffix := reverseString(word[j:])
				if idx, exists := wordMap[reversedSuffix]; exists && idx != i {
					result = append(result, []int{idx, i})
				}
			}
		}
	}

	return result
}
