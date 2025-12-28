package graphs

// https://leetcode.com/problems/sort-items-by-groups-respecting-dependencies/

// Approach: Dual Topological Sort with Global Item Ordering
// Sort all items globally, then sort groups, then merge
// This avoids repeated per-group sorting overhead
// N = number of items, M = number of groups, E = total edges
// Time: O(N + M + E) - single topo sort for items + groups + bucketing
// Space: O(N + M) for graphs, indegrees, and result buckets
func sortItems(n int, m int, group []int, beforeItems [][]int) []int {
	// Step 1: Assign unique group IDs to items with no group (-1)
	for i := range n {
		if group[i] == -1 {
			group[i] = m
			m++
		}
	}

	// Step 2: Build graphs and indegree counts
	itemGraph := make([][]int, n)
	itemIndegree := make([]int, n)

	groupGraph := make([][]int, m)
	groupIndegree := make([]int, m)

	// Build graphs from beforeItems
	for i := range n {
		for _, prev := range beforeItems[i] {
			itemGraph[prev] = append(itemGraph[prev], i)
			itemIndegree[i]++

			// If different groups, add group dependency
			if group[prev] != group[i] {
				groupGraph[group[prev]] = append(groupGraph[group[prev]], group[i])
				groupIndegree[group[i]]++
			}
		}
	}

	// Step 3: Topological sort helper using Kahn's algorithm
	topoSort := func(graph [][]int, indegree []int, total int) []int {
		queue := []int{}

		for i := range total {
			if indegree[i] == 0 {
				queue = append(queue, i)
			}
		}

		result := []int{}

		for len(queue) > 0 {
			x := queue[0]
			queue = queue[1:]

			result = append(result, x)

			for _, y := range graph[x] {
				indegree[y]--
				if indegree[y] == 0 {
					queue = append(queue, y)
				}
			}
		}

		if len(result) != total {
			return []int{}
		}

		return result
	}

	// Step 4: Sort groups
	groupIndegreeCopy := make([]int, len(groupIndegree))
	copy(groupIndegreeCopy, groupIndegree)

	groupOrder := topoSort(groupGraph, groupIndegreeCopy, m)
	if len(groupOrder) == 0 {
		return []int{}
	}

	// Step 5: Sort items globally
	itemIndegreeCopy := make([]int, len(itemIndegree))
	copy(itemIndegreeCopy, itemIndegree)

	itemOrder := topoSort(itemGraph, itemIndegreeCopy, n)
	if len(itemOrder) == 0 {
		return []int{}
	}

	// Step 6: Put items into buckets by group following item order
	buckets := make([][]int, m)
	for _, item := range itemOrder {
		buckets[group[item]] = append(buckets[group[item]], item)
	}

	// Step 7: Follow group ordering to build final answer
	result := []int{}
	for _, g := range groupOrder {
		result = append(result, buckets[g]...)
	}

	return result
}
