package twopointers

import "sort"

// https://leetcode.com/problems/count-pairs-of-nodes/

// Approach: Two Pointers + Bitmask
// Time: O(q*(n+e) + nlogn), q=len(queries)
// Space: O(n+e), e=len(edges)
func countPairs(n int, edges [][]int, queries []int) []int {
	var degree = make([]int, n+1)
	// max(n) = 2^4, dec(2e5) = bin(0b100111000100000), # of bits = 15
	// so, we'll move the first node by 15 steps to left
	var shared = make(map[int]int)
	var u, v int
	for idx := range edges { // O(e)
		u, v = edges[idx][0], edges[idx][1]
		degree[u]++
		degree[v]++
		if u > v {
			u, v = v, u
		}
		shared[u<<15+v]++
	}
	var sortedDegree = make([]int, n)
	copy(sortedDegree, degree[1:]) // [1:] means excluding node 0
	sort.Ints(sortedDegree)        // O(nlogn)
	var answer = make([]int, len(queries))
	var left, right int
	for idx := range queries { // O(q)
		left, right = 0, n-1
		for left <= right { // O(n)
			if sortedDegree[left]+sortedDegree[right] > queries[idx] {
				answer[idx] += right - left
				right--
			} else {
				left++
			}
		}
		for mask := range shared { // O(e)
			u, v = mask>>15, mask&(1<<15-1)
			if degree[u]+degree[v] > queries[idx] && degree[u]+degree[v]-shared[u<<15+v] <= queries[idx] {
				answer[idx]--
			}
		}
	}
	return answer
}

// // Approach: Brute-Force
// // Time: O(q*n^2), q=len(queries)
// // Space: O(n+e), e=len(edges)
// func countPairs(n int, edges [][]int, queries []int) []int {
// 	var degree = make([]int, n+1)
// 	var shared = make([][]int, n+1)
// 	for idx := range shared { // O(n)
// 		shared[idx] = make([]int, n+1)
// 	}
// 	var u, v int
// 	for idx := range edges { // O(e)
// 		u, v = edges[idx][0], edges[idx][1]
// 		degree[u]++
// 		degree[v]++
// 		if u > v {
// 			u, v = v, u
// 		}
// 		shared[u][v]++
// 	}
// 	var answer = make([]int, len(queries))
// 	for idx := range queries { // O(q)
// 		for u := 1; u <= n; u++ { // O(n)
// 			for v := u + 1; v <= n; v++ { // O(n)
// 				if degree[u]+degree[v]-shared[u][v] > queries[idx] {
// 					answer[idx]++
// 				}
// 			}
// 		}
// 	}
// 	return answer
// }
