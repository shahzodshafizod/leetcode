package greedy

// https://leetcode.com/problems/minimum-deletions-to-make-string-k-special/

func minimumDeletions(word string, k int) int {
	var freq [26]int
	for _, c := range word {
		freq[int(c-'a')]++
	}

	res := len(word)

	var high, dels int
	for _, low := range freq {
		high, dels = low+k, 0

		for _, cnt := range freq {
			if cnt < low {
				dels += cnt
			} else if cnt > high {
				dels += cnt - high
			}
		}

		res = min(res, dels)
	}

	return res
}
