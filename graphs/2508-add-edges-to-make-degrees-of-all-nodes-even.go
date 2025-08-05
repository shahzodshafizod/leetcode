package graphs

// https://leetcode.com/problems/add-edges-to-make-degrees-of-all-nodes-even/

func isPossible(n int, edges [][]int) bool {
	adjSet := make([]map[int]bool, n+1)
	for idx := range adjSet {
		adjSet[idx] = make(map[int]bool)
	}

	var a, b int
	for _, edge := range edges {
		a, b = edge[0], edge[1]
		adjSet[a][b] = true
		adjSet[b][a] = true
	}

	var odds []int

	for node := 1; node <= n; node++ {
		if len(adjSet[node])&1 == 1 {
			odds = append(odds, node)
		}
	}

	if len(odds) == 2 {
		for node := 1; node <= n; node++ {
			if !adjSet[odds[0]][node] && !adjSet[odds[1]][node] {
				return true
			}
		}
	} else if len(odds) == 4 {
		a, b, c, d := odds[0], odds[1], odds[2], odds[3]
		if !adjSet[a][b] && !adjSet[c][d] ||
			!adjSet[a][c] && !adjSet[b][d] ||
			!adjSet[a][d] && !adjSet[b][c] {
			return true
		}
	}

	return len(odds) == 0
}
