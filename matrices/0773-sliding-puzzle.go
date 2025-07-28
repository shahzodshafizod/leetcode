package matrices

import (
	"bytes"
)

// https://leetcode.com/problems/sliding-puzzle/

// Approach: Breadth-First Search
// Time: O((m⋅n)!×(m⋅n))
// Space: O((m⋅n)!)
func slidingPuzzle(board [][]int) int {
	neighbors := [][]int{{1, 3}, {0, 2, 4}, {1, 5}, {0, 4}, {1, 3, 5}, {2, 4}}
	state := make([]byte, 0, 6)

	for _, row := range board {
		for _, cell := range row {
			state = append(state, byte('0'+cell))
		}
	}

	queue := []string{string(state)}
	seen := map[string]bool{string(state): true}
	length := 0

	for size := len(queue); size > 0; size = len(queue) {
		for idx := 0; idx < size; idx++ {
			state = []byte(queue[idx])
			if string(state) == "123450" {
				return length
			}

			curr := bytes.IndexByte(state, '0')
			for _, next := range neighbors[curr] {
				state[curr], state[next] = state[next], state[curr]
				if nextState := string(state); !seen[nextState] {
					queue = append(queue, nextState)
					seen[nextState] = true
				}

				state[curr], state[next] = state[next], state[curr]
			}
		}

		queue = queue[size:]
		length++
	}

	return -1
}

// // Approach: Depth-First Search
// // Time: O((m⋅n)!×(m⋅n)^2)
// // Space: O((m⋅n)!)
// func slidingPuzzle(board [][]int) int {
// 	var state = make([]byte, 0, 6)
// 	for _, row := range board {
// 		for _, cell := range row {
// 			state = append(state, byte('0'+cell))
// 		}
// 	}
// 	var neighbors = [6][]int{{1, 3}, {0, 2, 4}, {1, 5}, {0, 4}, {1, 3, 5}, {2, 4}}
// 	var dist = make(map[string]int)
// 	var dfs func(state []byte, curr int, moves int)
// 	dfs = func(state []byte, curr int, moves int) {
// 		var key = string(state)
// 		if d, ok := dist[key]; ok && d <= moves {
// 			return
// 		}
// 		dist[key] = moves
// 		for _, next := range neighbors[curr] {
// 			state[curr], state[next] = state[next], state[curr]
// 			dfs(state, next, moves+1)
// 			state[curr], state[next] = state[next], state[curr]
// 		}
// 	}
// 	dfs(state, bytes.IndexByte(state, '0'), 0)
// 	if d, ok := dist["123450"]; ok {
// 		return d
// 	}
// 	return -1
// }
