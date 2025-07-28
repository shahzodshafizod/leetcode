package dp

/*
Problem:
For a given staircase, the i-th step is assigned a non-negative cost indicated by a cost array.
Once you pay the cost for a step, you can either climb one or two steps.
Find the minimum cost to teach the top of the staircase.
Your first step can either be the first or second step.

Recognize it's a dynamic programming problem
	- optimization question (min, max)
	- recursion: recurrence relation - Formula for the basis of a recursive solution
*/

// https://leetcode.com/problems/min-cost-climbing-stairs/

// bottom up
func minCostClimbingStairs(cost []int) int {
	n := len(cost)
	first, second := cost[0], cost[1]

	for i := 2; i < n; i++ {
		// third := cost[i] + min(first, second)
		// first = second
		// second = third
		first, second = second, cost[i]+min(first, second)
	}

	return min(first, second)
}

// // top down
// func minCostClimbingStairs(cost []int) int {
// 	var n = len(cost)
// 	var mins = make([]int, n)
// 	return min(minCost(cost, n-1, mins), minCost(cost, n-2, mins))
// }

// func minCost(cost []int, stair int, mins []int) int {
// 	if stair < 0 {
// 		return 0
// 	}
// 	if stair < 2 {
// 		return cost[stair]
// 	}
// 	if mins[stair] == 0 {
// 		mins[stair] = cost[stair] + min(minCost(cost, stair-1, mins), minCost(cost, stair-2, mins))
// 	}
// 	return mins[stair]
// }
