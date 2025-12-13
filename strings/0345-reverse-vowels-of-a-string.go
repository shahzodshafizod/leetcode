package strings

// https://leetcode.com/problems/reverse-vowels-of-a-string/

// Approach: Two Pointers
// Use two pointers from both ends, swap when both point to vowels.
// Time: O(n) - single pass
// Space: O(n) - convert to rune slice
func reverseVowels(s string) string {
	vowels := map[rune]bool{
		'a': true, 'e': true, 'i': true, 'o': true, 'u': true,
		'A': true, 'E': true, 'I': true, 'O': true, 'U': true,
	}

	chars := []rune(s)
	left, right := 0, len(chars)-1

	for left < right {
		if !vowels[chars[left]] {
			left++
		} else if !vowels[chars[right]] {
			right--
		} else {
			chars[left], chars[right] = chars[right], chars[left]
			left++
			right--
		}
	}

	return string(chars)
}
