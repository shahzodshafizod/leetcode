package graphs

// https://leetcode.com/problems/jump-game-iv/

// Approach: Breadth-First Search (BFS, Shortest Path in Graph)
// Time: O(N)
// Space: O(N)
func minJumps(arr []int) int {
	n := len(arr)
	indices := make(map[int][]int)
	for idx := 0; idx < n; idx++ {
		indices[arr[idx]] = append(indices[arr[idx]], idx)
	}
	visited := make([]bool, n)
	queue := [][2]int{{0, 0}} // {node, jumps}
	idx := 0
	visited[0] = true
	for idx < len(queue) {
		curr := queue[idx]
		idx++
		if curr[0] == n-1 {
			return curr[1]
		}
		for _, next := range []int{curr[0] - 1, curr[0] + 1} {
			if next >= 0 && next < n && !visited[next] {
				queue = append(queue, [2]int{next, curr[1] + 1})
				visited[next] = true
			}
		}
		for _, next := range indices[arr[curr[0]]] {
			if !visited[next] {
				queue = append(queue, [2]int{next, curr[1] + 1})
				visited[next] = true
			}
		}
		indices[arr[curr[0]]] = nil
	}
	return n - 1
}
