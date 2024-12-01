package graphs

// https://leetcode.com/problems/shortest-distance-after-road-addition-queries-ii/

// Approach: Linked List
// Time: O(n+q)
// Space: O(n)
func shortestDistanceAfterQueries(n int, queries [][]int) []int {
	var next = make([]int, n)
	for idx := range next {
		next[idx] = idx + 1
	}
	var distance = n - 1
	var answer = make([]int, len(queries))
	var src, dst int
	for idx := range queries {
		src, dst = queries[idx][0], queries[idx][1]
		for next[src] < dst {
			src, next[src] = next[src], dst
			distance--
		}
		answer[idx] = distance
	}
	return answer
}

// // Approach: Union Find
// // Time: O(n+q)
// // Space: O(n)
// func shortestDistanceAfterQueries(n int, queries [][]int) []int {
// 	type UnionFind struct {
// 		Root []int
// 		Size int
// 	}
// 	var find func(u *UnionFind, x int) int
// 	find = func(u *UnionFind, x int) int {
// 		if u.Root[x] != x {
// 			u.Root[x] = find(u, u.Root[x])
// 		}
// 		return u.Root[x]
// 	}
// 	var union = func(u *UnionFind, x int, y int) {
// 		var px = find(u, x)
// 		var py = find(u, y)
// 		if px != py {
// 			u.Root[py] = px
// 			u.Size--
// 		}
// 	}
// 	var uf = &UnionFind{Root: make([]int, n), Size: n}
// 	for x := range uf.Root {
// 		uf.Root[x] = x
// 	}
// 	var answer = make([]int, len(queries))
// 	var src, dst int
// 	for idx := range queries {
// 		src, dst = queries[idx][0]+1, queries[idx][1]
// 		for dst > src {
// 			union(uf, dst-1, dst)
// 			dst = find(uf, dst)
// 		}
// 		answer[idx] = uf.Size - 1
// 	}
// 	return answer
// }
