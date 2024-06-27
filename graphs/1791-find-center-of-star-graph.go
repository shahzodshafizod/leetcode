package graphs

// https://leetcode.com/problems/find-center-of-star-graph/

// Approach #2: Greedy
// time: O(1)
// space: O(1)
func findCenter(edges [][]int) int {
	var u1, v1 = edges[0][0], edges[0][1]
	var u2, v2 = edges[1][0], edges[1][1]
	if u1 == u2 || u1 == v2 {
		return u1
	}
	return v1
}

// // Approach #1: Degree Count ("Degree" is the number of connections a node has)
// // time: O(n)
// // space: O(n)
// func findCenter(edges [][]int) int {
// 	var connected = make(map[int]bool)
// 	var u, v, center int
// 	for _, edge := range edges {
// 		u, v = edge[0], edge[1]
// 		if connected[u] {
// 			center = u
// 			break
// 		} else if connected[v] {
// 			center = v
// 			break
// 		} else {
// 			connected[u] = true
// 			connected[v] = true
// 		}
// 	}
// 	return center
// }
