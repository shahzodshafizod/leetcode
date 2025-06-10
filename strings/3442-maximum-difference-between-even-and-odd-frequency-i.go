package strings

// https://leetcode.com/problems/maximum-difference-between-even-and-odd-frequency-i/

func maxDifference(s string) int {
	var count [26]int
	for _, c := range s {
		count[int(c-'a')]++
	}
	var a1, a2 = 0, len(s)
	for _, cnt := range count {
		if cnt&1 == 1 {
			a1 = max(a1, cnt)
		} else if cnt > 0 {
			a2 = min(a2, cnt)
		}
	}
	return a1 - a2
}
