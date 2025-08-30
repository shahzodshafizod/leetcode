package graphs

// https://leetcode.com/problems/build-a-matrix-with-conditions/

// Time: O(K^2)
// Space: O(K^2)
func buildMatrix(k int, rowConditions [][]int, colConditions [][]int) [][]int {
	buildGraph := func(conditions [][]int, n int) ([][]int, []int) {
		adjList := make([][]int, n+1)
		for idx := 1; idx <= n; idx++ {
			adjList[idx] = make([]int, 0)
		}

		inDegrees := make([]int, n+1)

		var before, after int
		for _, condition := range conditions {
			before, after = condition[0], condition[1]
			adjList[before] = append(adjList[before], after)
			inDegrees[after]++
		}

		return adjList, inDegrees
	}

	topSort := func(conditions [][]int, n int) []int {
		adjList, inDegrees := buildGraph(conditions, n)
		queue := make([]int, 0)

		for node := 1; node <= n; node++ {
			if inDegrees[node] == 0 {
				queue = append(queue, node)
			}
		}

		ordering := make([]int, 0, n)

		var node int
		for len(queue) > 0 {
			node = queue[0]
			queue = queue[1:]

			ordering = append(ordering, node)
			n--

			for _, next := range adjList[node] {
				inDegrees[next]--
				if inDegrees[next] == 0 {
					queue = append(queue, next)
				}
			}
		}

		if n != 0 {
			return nil
		}

		return ordering
	}

	rowOrdering := topSort(rowConditions, k)
	if len(rowOrdering) == 0 {
		return nil
	}

	colOrdering := topSort(colConditions, k)
	if len(colOrdering) == 0 {
		return nil
	}

	matrix := make([][]int, k)
	for row := range k {
		matrix[row] = make([]int, k)
		matrix[row][0] = rowOrdering[row]
	}

	for col := 1; col < k; col++ {
		for row := range k {
			if matrix[row][col-1] != colOrdering[col-1] {
				matrix[row][col] = matrix[row][col-1]
				matrix[row][col-1] = 0
			}
		}
	}

	return matrix
}
