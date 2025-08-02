package hashes

// https://leetcode.com/problems/verifying-an-alien-dictionary/

func isAlienSorted(words []string, order string) bool {
	indices := make(map[byte]int)
	for idx := range order {
		indices[order[idx]] = idx
	}

	var n1, n2, limit int

	var sorted bool
	for idx, n := 1, len(words); idx < n; idx++ {
		sorted = false
		n1, n2 = len(words[idx-1]), len(words[idx])

		limit = min(n1, n2)
		for c := 0; c < limit; c++ {
			if indices[words[idx-1][c]] > indices[words[idx][c]] {
				return false
			} else if indices[words[idx-1][c]] < indices[words[idx][c]] {
				sorted = true

				break
			}
		}

		if !sorted && n1 > n2 {
			return false
		}
	}

	return true
}
