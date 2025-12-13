package graphs

// https://leetcode.com/problems/shortest-path-to-get-all-keys/

// Approach: BFS with State Compression (Bitmask)
// Time: O(m * n * 2^k) where m,n = grid dimensions, k = number of keys
// Space: O(m * n * 2^k) for visited set
func shortestPathAllKeys(grid []string) int {
	m, n := len(grid), len(grid[0])

	// Find start position and count total keys
	startR, startC := 0, 0
	keyCount := 0

	for r := range m {
		for c := range n {
			if grid[r][c] == '@' {
				startR, startC = r, c
			} else if grid[r][c] >= 'a' && grid[r][c] <= 'f' {
				keyCount++
			}
		}
	}

	// BFS: [row, col, keys_mask, steps]
	type state struct {
		r, c, keys, steps int
	}

	queue := []state{{startR, startC, 0, 0}}
	visited := make(map[[3]int]bool)
	visited[[3]int{startR, startC, 0}] = true
	allKeys := (1 << keyCount) - 1

	directions := [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		// Check if we have all keys
		if curr.keys == allKeys {
			return curr.steps
		}

		// Explore 4 directions
		for _, dir := range directions {
			nr, nc := curr.r+dir[0], curr.c+dir[1]

			// Check bounds
			if nr < 0 || nr >= m || nc < 0 || nc >= n {
				continue
			}

			cell := grid[nr][nc]

			// Check if wall
			if cell == '#' {
				continue
			}

			newKeys := curr.keys

			// Check if it's a key
			if cell >= 'a' && cell <= 'f' {
				newKeys = curr.keys | (1 << (cell - 'a'))
			}

			// Check if it's a lock
			if cell >= 'A' && cell <= 'F' {
				// Check if we have the key
				if (curr.keys & (1 << (cell - 'A'))) == 0 {
					continue
				}
			}

			// Check if state already visited
			stateKey := [3]int{nr, nc, newKeys}
			if visited[stateKey] {
				continue
			}

			visited[stateKey] = true

			queue = append(queue, state{nr, nc, newKeys, curr.steps + 1})
		}
	}

	return -1
}
