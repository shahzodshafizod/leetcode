package slidingwindows

// https://leetcode.com/problems/number-of-substrings-containing-all-three-characters/

func numberOfSubstrings(s string) int {
	count := 0
	a, b, c := -1, -1, -1

	for idx, n := 0, len(s); idx < n; idx++ {
		switch s[idx] {
		case 'a':
			a = idx
		case 'b':
			b = idx
		case 'c':
			c = idx
		default:
		}

		count += min(a, b, c) + 1
	}

	return count
}
