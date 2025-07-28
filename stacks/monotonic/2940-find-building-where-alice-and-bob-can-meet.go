package monotonic

import (
	"sort"
)

// https://leetcode.com/problems/find-building-where-alice-and-bob-can-meet/

// Approach #2: Monotonic Decreasing Stack
// Time: O(QlogQ+QlogH), Q=len(queries), H=len(heights)
// Space: O(Q+H)
func leftmostBuildingQueries(heights []int, queries [][]int) []int {
	ans := make([]int, len(queries))

	var a, b int

	remained := make([][3]int, 0)

	for idx := range queries {
		a, b = queries[idx][0], queries[idx][1]
		if a > b {
			a, b = b, a
		}

		if a == b || heights[a] < heights[b] {
			ans[idx] = b
		} else {
			ans[idx] = -1

			remained = append(remained, [3]int{b, idx, heights[a]})
		}
	}

	sort.Slice(remained, func(i, j int) bool { return remained[i][0] > remained[j][0] })

	var qid, aid, height, left, right, mid int

	hid := len(heights) - 1
	stack := make([]int, 0)
	size := 0

	for idx := range remained {
		qid, aid, height = remained[idx][0], remained[idx][1], remained[idx][2]
		for ; hid > qid; hid-- {
			for size > 0 && heights[stack[size-1]] <= heights[hid] {
				stack = stack[:size-1]
				size--
			}

			stack = append(stack, hid)
			size++
		}

		left, right = 0, size-1
		for left <= right {
			mid = left + (right-left)/2
			if heights[stack[mid]] > height {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}

		if right != -1 {
			ans[aid] = stack[right]
		}
	}

	return ans
}

// // Approach #1: Brute-Force
// // Time: O(Q*H), Q=len(queries), H=len(heights)
// // Space: O(1)
// func leftmostBuildingQueries(heights []int, queries [][]int) []int {
// 	var ans = make([]int, len(queries))
// 	var hlen = len(heights)
// 	var a, b int
// 	for idx := range queries {
// 		a, b = queries[idx][0], queries[idx][1]
// 		if a > b {
// 			a, b = b, a
// 		}
// 		if a == b || heights[a] < heights[b] {
// 			ans[idx] = b
// 		} else {
// 			ans[idx] = -1
// 			for next := b + 1; next < hlen; next++ {
// 				if heights[next] > heights[a] {
// 					ans[idx] = next
// 					break
// 				}
// 			}
// 		}
// 	}
// 	return ans
// }
