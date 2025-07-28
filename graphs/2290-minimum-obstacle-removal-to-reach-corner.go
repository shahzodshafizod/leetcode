package graphs

// https://leetcode.com/problems/minimum-obstacle-removal-to-reach-corner/

// Approach #4: Greedy (Breadth-First Search + Two Arrays)
// Time: O(mn), m=# of rows, n=# of cols
// Space: O(mn)
func minimumObstacles(grid [][]int) int {
	type Cell struct {
		Obstacles int
		Row       int
		Col       int
	}

	front := []*Cell{{0, 0, 0}} // {obstacles, row, col}
	directions := [5]int{-1, 0, 1, 0, -1}
	m, n := len(grid), len(grid[0])
	seen := make([][]bool, m)

	for idx := range seen {
		seen[idx] = make([]bool, n)
	}

	obstacles := m + n - 1

	for len(front) > 0 {
		back := make([]*Cell, 0)

		for idx := 0; idx < len(front); idx++ {
			curr := front[idx]
			if curr.Row == m-1 && curr.Col == n-1 {
				return curr.Obstacles
			}

			var r, c int
			for dir := 1; dir < 5; dir++ {
				r = curr.Row + directions[dir-1]
				c = curr.Col + directions[dir]

				if min(r, c) >= 0 && r < m && c < n && !seen[r][c] {
					if grid[r][c] == 0 {
						front = append(front, &Cell{curr.Obstacles, r, c})
					} else {
						back = append(back, &Cell{curr.Obstacles + 1, r, c})
					}

					seen[r][c] = true
				}
			}
		}

		front = back
	}

	return obstacles
}

// // Approach #3: Greedy (Breadth-First Search + Doubly Linked List)
// // Time: O(mn), m=# of rows, n=# of cols
// // Space: O(mn)
// func minimumObstacles(grid [][]int) int {
// 	type Cell struct {
// 		Obstacles int
// 		Row       int
// 		Col       int
// 	}
// 	var queue = list.New()
// 	queue.PushBack(&Cell{0, 0, 0}) // {obstacles, row, col}
// 	var directions = [5]int{-1, 0, 1, 0, -1}
// 	var m, n = len(grid), len(grid[0])
// 	var seen = make([][]bool, m)
// 	for idx := range seen {
// 		seen[idx] = make([]bool, n)
// 	}
// 	var obstacles = m + n - 1
// 	for queue.Len() > 0 {
// 		var curr = queue.Remove(queue.Front()).(*Cell)
// 		if curr.Row == m-1 && curr.Col == n-1 {
// 			obstacles = curr.Obstacles
// 			break
// 		}
// 		var r, c int
// 		for dir := 1; dir < 5; dir++ {
// 			r = curr.Row + directions[dir-1]
// 			c = curr.Col + directions[dir]
// 			if min(r, c) >= 0 && r < m && c < n && !seen[r][c] {
// 				if grid[r][c] == 0 {
// 					queue.PushFront(&Cell{curr.Obstacles, r, c})
// 				} else {
// 					queue.PushBack(&Cell{curr.Obstacles + 1, r, c})
// 				}
// 				seen[r][c] = true
// 			}
// 		}
// 	}
// 	return obstacles
// }

// // Approach #2: Dijkstra's
// // Time: O(mnlog(mn)), m=# of rows, n=# of cols
// // Space: O(mn)
// func minimumObstacles(grid [][]int) int {
// 	type Cell struct {
// 		Obstacles int
// 		Row       int
// 		Col       int
// 	}
// 	var minHeap = pkg.NewHeap(
// 		[]*Cell{{0, 0, 0}}, // {obstacles, row, col}
// 		func(x, y *Cell) bool {
// 			return x.Obstacles < y.Obstacles
// 		},
// 	)
// 	var directions = [5]int{-1, 0, 1, 0, -1}
// 	var m, n = len(grid), len(grid[0])
// 	var seen = make([][]bool, m)
// 	for idx := range seen {
// 		seen[idx] = make([]bool, n)
// 	}
// 	var obstacles = m + n - 1
// 	for minHeap.Len() > 0 {
// 		var curr = heap.Pop(minHeap).(*Cell)
// 		if curr.Row == m-1 && curr.Col == n-1 {
// 			obstacles = curr.Obstacles
// 			break
// 		}
// 		var r, c int
// 		for dir := 1; dir < 5; dir++ {
// 			r = curr.Row + directions[dir-1]
// 			c = curr.Col + directions[dir]
// 			if min(r, c) >= 0 && r < m && c < n && !seen[r][c] {
// 				heap.Push(minHeap, &Cell{curr.Obstacles + grid[r][c], r, c})
// 				seen[r][c] = true
// 			}
// 		}
// 	}
// 	return obstacles
// }

// // Approach #1: Depth-First Search (Time Limit Exceeded)
// // Time: O(4^(m*n)), m=# of rows, n=# of cols
// // Space: O(m*n)
// func minimumObstacles(grid [][]int) int {
// 	var directions = [5]int{-1, 0, 1, 0, -1}
// 	var m, n = len(grid), len(grid[0])
// 	var seen = make([][]bool, m)
// 	for idx := range seen {
// 		seen[idx] = make([]bool, n)
// 	}
// 	var dfs func(row int, col int) int
// 	dfs = func(row int, col int) int {
// 		if row == m-1 && col == n-1 {
// 			return 0
// 		}
// 		var result int = 1e5
// 		var r, c int
// 		for dir := 1; dir < 5; dir++ {
// 			r, c = row+directions[dir-1], col+directions[dir]
// 			if min(r, c) >= 0 && r < m && c < n && !seen[r][c] {
// 				seen[r][c] = true
// 				result = min(result, grid[r][c]+dfs(r, c))
// 				seen[r][c] = false
// 			}
// 		}
// 		return result
// 	}
// 	return dfs(0, 0)
// }
