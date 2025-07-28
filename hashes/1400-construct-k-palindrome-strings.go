package hashes

// https://leetcode.com/problems/construct-k-palindrome-strings/

func canConstruct(s string, k int) bool {
	if len(s) < k {
		return false
	}

	if len(s) == k {
		return true
	}

	var odds uint = 0
	for _, c := range s {
		odds ^= 1 << int(c-'a')
	}

	for odds > 0 {
		odds &= odds - 1
		k--
	}

	return k >= 0
}
