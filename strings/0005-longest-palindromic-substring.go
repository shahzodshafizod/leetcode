package strings

// https://leetcode.com/problems/longest-palindromic-substring/

// time: O(n^2)
// space: O(n) (recursion)
func longestPalindrome(s string) string {
	n := len(s)
	var expand func(left int, right int) string
	expand = func(left int, right int) string {
		if left < 0 || right >= n || s[left] != s[right] {
			return ""
		}
		res := s[left : right+1]
		nextRes := expand(left-1, right+1)
		if len(nextRes) > len(res) {
			res = nextRes
		}
		return res
	}
	var longRes string
	radius := 0
	for i := 0; i < n-radius; i++ {
		for _, res := range []string{
			expand(i, i),   // odd length
			expand(i, i+1), // even length
		} {
			if len(res) > len(longRes) {
				longRes = res
				radius = len(longRes) / 2
			}
		}
	}
	return longRes
}
