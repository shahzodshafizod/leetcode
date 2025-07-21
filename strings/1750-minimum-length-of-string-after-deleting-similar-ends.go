package strings

// https://leetcode.com/problems/minimum-length-of-string-after-deleting-similar-ends/

func minimumLength(s string) int {
	left, right := 0, len(s)-1
	var c byte
	for left < right && s[left] == s[right] {
		c = s[left]
		for left < right && s[left] == c {
			left++
		}
		for left <= right && s[right] == c {
			right--
		}
	}
	return right - left + 1
}
