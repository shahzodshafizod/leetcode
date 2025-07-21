package graphs

// https://leetcode.com/problems/sum-of-distances-in-tree/

func sumOfDistancesInTree(n int, edges [][]int) []int {
	adjList := make([][]int, n)
	for idx := range adjList {
		adjList[idx] = make([]int, 0)
	}
	var a, b int
	for _, edge := range edges {
		a, b = edge[0], edge[1]
		adjList[a] = append(adjList[a], b)
		adjList[b] = append(adjList[b], a)
	}

	answer := make([]int, n)
	count := make([]int, n)

	var calcCount func(curr int, prev int, depth int)
	calcCount = func(curr int, prev int, depth int) {
		count[curr] = 1
		depth++
		for _, next := range adjList[curr] {
			if next != prev {
				calcCount(next, curr, depth)
				count[curr] += count[next]
				answer[0] += depth
			}
		}
	}
	calcCount(0, -1, 0)

	var calcSum func(curr int, prev int)
	calcSum = func(curr int, prev int) {
		for _, next := range adjList[curr] {
			if next != prev {
				// when we move from parent to child
				// we are moving 1 step closer to all the child nodes
				// and 1 step away from the parent nodes
				answer[next] = answer[curr] - count[next] + (n - count[next])
				calcSum(next, curr)
			}
		}
	}
	calcSum(0, -1)

	return answer
}

// // Time Limit Exceeded
// // time: O(n^2)
// // space: O(n^2)
// func sumOfDistancesInTree(n int, edges [][]int) []int {
// 	var adjList = make([][]int, n)
// 	for idx := range adjList {
// 		adjList[idx] = make([]int, 0)
// 	}
// 	var a, b int
// 	for _, edge := range edges {
// 		a, b = edge[0], edge[1]
// 		adjList[a] = append(adjList[a], b)
// 		adjList[b] = append(adjList[b], a)
// 	}
// 	var dfs func(curr int, prev int, depth int) int
// 	dfs = func(curr int, prev int, depth int) int {
// 		var sum = depth
// 		for _, next := range adjList[curr] {
// 			if next != prev {
// 				sum += dfs(next, curr, depth+1)
// 			}
// 		}
// 		return sum
// 	}
// 	var answer = make([]int, n)
// 	for idx := range answer {
// 		answer[idx] = dfs(idx, -1, 0)
// 	}
// 	return answer
// }
