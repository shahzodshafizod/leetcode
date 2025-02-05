package strings

// https://leetcode.com/problems/check-if-one-string-swap-can-make-strings-equal/

func areAlmostEqual(s1 string, s2 string) bool {
	var count = 0
	var index1, index2 = 0, 0
	for idx := len(s1) - 1; idx >= 0; idx-- {
		if s1[idx] == s2[idx] {
			continue
		}
		count++
		switch count {
		case 1:
			index2 = idx
		case 2:
			index1 = idx
		default:
			return false
		}
	}
	return s1[index1] == s2[index2] && s1[index2] == s2[index1]
}
