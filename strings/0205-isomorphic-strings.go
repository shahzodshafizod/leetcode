package strings

// https://leetcode.com/problems/isomorphic-strings/

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	// ASCII is a 7-bit character set containing 128 characters
	var stable, ttable [128]uint16
	for idx := range s {
		if stable[s[idx]] != ttable[t[idx]] {
			return false
		}
		// +1 needs to not save the default value: 0
		stable[s[idx]] = uint16(idx + 1)
		ttable[t[idx]] = uint16(idx + 1)
	}
	return true
}
