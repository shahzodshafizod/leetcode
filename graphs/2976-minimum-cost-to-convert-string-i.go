package graphs

import "math"

// https://leetcode.com/problems/minimum-cost-to-convert-string-i/

// Approach 2: Floyd-Warshall Algorithm
// N = Len(source)
// M = Len(original)
// Time: O(M + 26^3 + N) = O(M+N)
// Space: O(26x26) = O(1)
func minimumCost(source string, target string, original []byte, changed []byte, cost []int) int64 {
	var minCosts [26][26]int64
	for src := range minCosts {
		for dst := range minCosts[src] {
			if src == dst {
				minCosts[src][dst] = 0
			} else {
				minCosts[src][dst] = math.MaxInt32
			}
		}
	}
	var src, dst int
	for idx := range original {
		src, dst = int(original[idx]-'a'), int(changed[idx]-'a')
		minCosts[src][dst] = min(minCosts[src][dst], int64(cost[idx]))
	}
	for through := 0; through < 26; through++ {
		for src := 0; src < 26; src++ {
			for dst := 0; dst < 26; dst++ {
				minCosts[src][dst] = min(
					minCosts[src][dst],
					minCosts[src][through]+minCosts[through][dst], // !!!dangerous section: MAX+MAX becomes negative!!!
				)
			}
		}
	}
	var totalCost int64 = 0
	for idx := range source {
		src, dst = int(source[idx]-'a'), int(target[idx]-'a')
		if minCosts[src][dst] >= math.MaxInt32 {
			return -1
		}
		totalCost += minCosts[src][dst]
	}

	return totalCost
}

// // Approach 1: Dijkstra's Algorithm
// // N = Len(source)
// // M = Len(original)
// // Time: O(26*M+N)) = O(M+N)
// // Space: O(M+N)
// func minimumCost(source string, target string, original []byte, changed []byte, cost []int) int64 {
// 	type Edge struct {
// 		To   byte
// 		Cost int
// 	}
// 	var adjList = make(map[byte][]*Edge)
// 	for idx := range original {
// 		from, to, cost := original[idx], changed[idx], cost[idx]
// 		adjList[from] = append(adjList[from], &Edge{to, cost})
// 	}
// 	var dijkstra = func(src byte) map[byte]int {
// 		var minCostMap = make(map[byte]int)
// 		var pq = design.NewPQ(make([]*Edge, 0), func(x, y *Edge) bool { return x.Cost > y.Cost })
// 		pq.Push(&Edge{src, 0})
// 		for pq.Len() > 0 {
// 			curr := pq.Pop()
// 			if _, exists := minCostMap[curr.To]; exists {
// 				continue
// 			}
// 			minCostMap[curr.To] = curr.Cost
// 			for _, edge := range adjList[curr.To] {
// 				pq.Push(&Edge{edge.To, curr.Cost + edge.Cost})
// 			}
// 		}
// 		return minCostMap
// 	}
// 	var minCostMaps = make(map[byte]map[byte]int)
// 	for idx := range source {
// 		if _, exists := minCostMaps[source[idx]]; !exists {
// 			minCostMaps[source[idx]] = dijkstra(source[idx])
// 		}
// 	}
// 	var totalCost int64 = 0
// 	for idx := range source {
// 		source, target := source[idx], target[idx]
// 		if minCostMaps[source] == nil {
// 			return -1
// 		}
// 		if cost, exists := minCostMaps[source][target]; !exists {
// 			return -1
// 		} else {
// 			totalCost += int64(cost)
// 		}
// 	}
// 	return totalCost
// }
