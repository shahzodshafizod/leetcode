package graphs

// https://leetcode.com/problems/find-closest-node-to-given-two-nodes/

// Approach #2: Depth-First Search (Optimized)
// Time: O(n)
// Space: O(n)
func closestMeetingNode(edges []int, node1 int, node2 int) int {
	type Node struct{ node, mask int }
	var n = len(edges)
	var seen = make([]int, n+1)
	var n1, n2 = &Node{node1, 1}, &Node{node2, 2}
	for n1.node != -1 || n2.node != -1 {
		if n1.node > n2.node {
			n1, n2 = n2, n1
		}
		if seen[n1.node+1]&n1.mask == n1.mask {
			n1.node = -1
		}
		if seen[n2.node+1]&n2.mask == n2.mask {
			n2.node = -1
		}
		seen[n1.node+1] |= n1.mask
		seen[n2.node+1] |= n2.mask
		if seen[n1.node+1] == 3 {
			return n1.node
		}
		if seen[n2.node+1] == 3 {
			return n2.node
		}
		if n1.node != -1 {
			n1.node = edges[n1.node]
		}
		if n2.node != -1 {
			n2.node = edges[n2.node]
		}
	}
	return -1
}

// // Approach #1: Depth-First Search
// // Time: O(3n) = O(n)
// // Space: O(2n) = O(n)
// func closestMeetingNode(edges []int, node1 int, node2 int) int {
// 	var n = len(edges)
// 	var dist1 = make([]int, n)
// 	var dist2 = make([]int, n)
// 	for i := 0; i < n; i++ {
// 		dist1[i] = n + 1
// 		dist2[i] = n + 1
// 	}
// 	for dist := 0; node1 != -1 && dist1[node1] > dist; dist++ {
// 		dist1[node1] = dist
// 		node1 = edges[node1]
// 	}
// 	for dist := 0; node2 != -1 && dist2[node2] > dist; dist++ {
// 		dist2[node2] = dist
// 		node2 = edges[node2]
// 	}
// 	var node, dist = -1, n + 1
// 	var d int
// 	for i := 0; i < n; i++ {
// 		d = max(dist1[i], dist2[i])
// 		if d < dist {
// 			dist = d
// 			node = i
// 		}
// 	}
// 	return node
// }
