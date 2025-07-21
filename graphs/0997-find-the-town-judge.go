package graphs

// https://leetcode.com/problems/find-the-town-judge/

func findJudge(n int, trust [][]int) int {
	trusters := make([]int, n)
	trustees := make([]int, n)
	for _, tr := range trust {
		trusters[tr[0]-1]++
		trustees[tr[1]-1]++
	}
	for label := 0; label < n; label++ {
		if trusters[label] == 0 && trustees[label] == n-1 {
			return label + 1
		}
	}
	return -1
}
