package dp

import (
	"container/heap"
	"math"

	"github.com/shahzodshafizod/leetcode/pkg"
)

// https://leetcode.com/problems/minimum-cost-to-reach-destination-in-time/

// Approach #2: Dynamic Programming (Dijkstra-like with time constraint)
// Time: O((n * maxTime) * log(n * maxTime)) - priority queue operations
// Space: O(n * maxTime) - DP table and priority queue
func minCost1928(maxTime int, edges [][]int, passingFees []int) int {
	type State struct {
		cost int
		city int
		time int
	}

	n := len(passingFees)

	// Build adjacency list
	graph := make([][][2]int, n)

	for _, edge := range edges {
		u, v, time := edge[0], edge[1], edge[2]
		graph[u] = append(graph[u], [2]int{v, time})
		graph[v] = append(graph[v], [2]int{u, time})
	}

	// dp[city][time] = minimum cost to reach city with time spent
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, maxTime+1)
		for j := range dp[i] {
			dp[i][j] = math.MaxInt
		}
	}

	dp[0][0] = passingFees[0]

	// Priority queue: (cost, city, time)
	pq := pkg.NewHeap(make([]State, 0), func(x, y State) bool { return x.cost < y.cost })
	heap.Push(pq, State{cost: passingFees[0], city: 0, time: 0})

	for pq.Len() > 0 {
		state, ok := heap.Pop(pq).(State)
		_ = ok
		cost, city, time := state.cost, state.city, state.time

		// If we reached destination
		if city == n-1 {
			return cost
		}

		// Skip if we found a better path already
		if cost > dp[city][time] {
			continue
		}

		// Explore neighbors
		for _, neighbor := range graph[city] {
			nextCity, travelTime := neighbor[0], neighbor[1]
			newTime := time + travelTime

			if newTime <= maxTime {
				newCost := cost + passingFees[nextCity]

				// Only proceed if this is better than previous best for this state
				if newCost < dp[nextCity][newTime] {
					dp[nextCity][newTime] = newCost
					heap.Push(pq, State{newCost, nextCity, newTime})
				}
			}
		}
	}

	return -1
}

// // Approach #1: Brute-Force (DFS with memoization)
// // Time: O(n * maxTime * E) where E is number of edges
// // Space: O(n * maxTime) - memoization cache
// func minCost1928(maxTime int, edges [][]int, passingFees []int) int {
// 	n := len(passingFees)

// 	// Build adjacency list
// 	graph := make([][][2]int, n)
// 	for _, edge := range edges {
// 		u, v, time := edge[0], edge[1], edge[2]
// 		graph[u] = append(graph[u], [2]int{v, time})
// 		graph[v] = append(graph[v], [2]int{u, time})
// 	}

// 	// memo[city][time] = minimum cost to reach city with time spent
// 	memo := make(map[[2]int]int)

// 	var dfs func(city, timeSpent int) int
// 	dfs = func(city, timeSpent int) int {
// 		// Base cases
// 		if city == n-1 {
// 			return passingFees[city]
// 		}

// 		if timeSpent > maxTime {
// 			return math.MaxInt
// 		}

// 		key := [2]int{city, timeSpent}
// 		if val, exists := memo[key]; exists {
// 			return val
// 		}

// 		minCost := math.MaxInt
// 		for _, neighbor := range graph[city] {
// 			nextCity, travelTime := neighbor[0], neighbor[1]
// 			if timeSpent+travelTime <= maxTime {
// 				cost := dfs(nextCity, timeSpent+travelTime)
// 				if cost != math.MaxInt {
// 					if passingFees[city]+cost < minCost {
// 						minCost = passingFees[city] + cost
// 					}
// 				}
// 			}
// 		}

// 		memo[key] = minCost
// 		return minCost
// 	}

// 	result := dfs(0, 0)
// 	if result == math.MaxInt {
// 		return -1
// 	}
// 	return result
// }
