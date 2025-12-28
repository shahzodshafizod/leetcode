package graphs

import (
	"strings"
)

/*
This is an Eulerian Circuit problem. The key insight:
- Nodes represent (n-1)-digit sequences
- Edges represent n-digit passwords
- Each node has exactly k outgoing edges (digits 0 to k-1)
- Total edges = k^n (all possible passwords)
- We need an Eulerian path visiting every edge exactly once

Graph Properties:
- For n=2, k=2: nodes are "0", "1"
  - Node "0" has edges "00" and "01"
  - Node "1" has edges "10" and "11"
- Each node has degree k (both in-degree and out-degree)
- This guarantees an Eulerian circuit exists

Example for n=2, k=2:
Start at "0" → visit edges: "00" → "01" → "11" → "10" → back to "0"
Result: "0" + "0" + "1" + "1" + "0" = "01100"
All passwords appear: "01", "11", "10", "00"
*/

// https://leetcode.com/problems/cracking-the-safe/

// Approach: Hierholzer's Algorithm (Optimal - Eulerian Circuit)
// Time: O(k^n) - visit each edge exactly once
// Space: O(k^n) - for graph and result
//
// Hierholzer's Algorithm for finding Eulerian Circuit:
// 1. Build graph with (n-1)-digit nodes
// 2. Each node has k outgoing edges
// 3. Use DFS with a stack to find the circuit
// 4. Visit edges greedily, backtrack when stuck
// 5. Build result in reverse order
//
// Why it works:
// - Graph is balanced (in-degree = out-degree for all nodes)
// - Eulerian circuit is guaranteed to exist
// - We traverse each edge (password) exactly once
// - Result length = n-1 (start) + k^n (edges) = k^n + n - 1
func crackSafe(n int, k int) string {
	// Special case: n=1, just concatenate all digits
	if n == 1 {
		var result strings.Builder
		for i := range k {
			_ = result.WriteByte(byte('0' + i))
		}

		return result.String()
	}

	// Helper: power function
	pow := func(base, exp int) int {
		result := 1
		for range exp {
			result *= base
		}

		return result
	}

	// Helper: convert number to base-k representation with fixed length
	toBaseK := func(num, k, length int) string {
		if num == 0 {
			return strings.Repeat("0", length)
		}

		var result strings.Builder
		for range length {
			_ = result.WriteByte(byte('0' + (num % k)))
			num /= k
		}

		// Reverse the string
		s := result.String()

		runes := []rune(s)
		for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
			runes[i], runes[j] = runes[j], runes[i]
		}

		return string(runes)
	}

	// Build adjacency list
	// Key: (n-1)-digit node, Value: available digits to append
	graph := make(map[string][]int)

	// Initialize all nodes with all possible edges
	totalNodes := pow(k, n-1)
	for i := range totalNodes {
		node := toBaseK(i, k, n-1)

		graph[node] = make([]int, k)
		for digit := range k {
			graph[node][digit] = digit
		}
	}

	// Start from all zeros
	start := strings.Repeat("0", n-1)
	stack := []string{start}

	var path []string

	// Hierholzer's algorithm
	for len(stack) > 0 {
		node := stack[len(stack)-1]

		if len(graph[node]) > 0 {
			// Take an edge (digit)
			digit := graph[node][len(graph[node])-1]
			graph[node] = graph[node][:len(graph[node])-1]

			// Next node is current shifted + new digit
			nextNode := node[1:] + string(byte('0'+digit))
			stack = append(stack, nextNode)
		} else {
			// No more edges, add node to path
			path = append(path, stack[len(stack)-1])
			stack = stack[:len(stack)-1]
		}
	}

	// Reverse path (built backwards)
	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}

	// Build result: first node + last char of each subsequent node
	var result strings.Builder

	_, _ = result.WriteString(path[0])

	for i := 1; i < len(path); i++ {
		_ = result.WriteByte(path[i][len(path[i])-1])
	}

	return result.String()
}
