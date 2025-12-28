package trees

// https://leetcode.com/problems/cycle-length-queries-in-a-tree/

// Approach: Optimized - Simultaneous traversal to LCA
// Move both pointers up the tree simultaneously until they meet
// In a complete binary tree: parent of node x is x / 2
// Time: O(M * log N) - for each query, traverse up to log(N) levels
// Space: O(1) - no extra space needed (result array doesn't count)
func cycleLengthQueries(n int, queries [][]int) []int {
	_ = n
	result := make([]int, len(queries))

	for i, query := range queries {
		a, b := query[0], query[1]
		distance := 0

		// Move both nodes up until they meet (find LCA)
		for a != b {
			if a > b {
				a /= 2 // Move to parent
			} else {
				b /= 2 // Move to parent
			}

			distance++
		}

		// Add 1 for the edge we're adding
		result[i] = distance + 1
	}

	return result
}
