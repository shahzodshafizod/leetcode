package graphs

import "github.com/shahzodshafizod/leetcode/pkg"

// https://leetcode.com/problems/clone-graph/

// DFS
func cloneGraph(node *pkg.Node) *pkg.Node {
	var seen [101]*pkg.Node
	newNode := func(val int) *pkg.Node { return &pkg.Node{Val: val, Neighbors: make([]*pkg.Node, 0)} }
	var dfs func(curr *pkg.Node) *pkg.Node
	dfs = func(curr *pkg.Node) *pkg.Node {
		if curr == nil {
			return nil
		}
		seen[curr.Val] = newNode(curr.Val)
		for _, neighbor := range curr.Neighbors {
			if seen[neighbor.Val] == nil {
				seen[neighbor.Val] = dfs(neighbor)
			}
			seen[curr.Val].Neighbors = append(seen[curr.Val].Neighbors, seen[neighbor.Val])
		}
		return seen[curr.Val]
	}
	return dfs(node)
}

// // BFS
// func cloneGraph(node *pkg.Node) *pkg.Node {
// 	if node == nil {
// 		return nil
// 	}
// 	var newNode = func(val int) *pkg.Node { return &pkg.Node{Val: val, Neighbors: make([]*pkg.Node, 0)} }
// 	var queue = []*pkg.Node{node}
// 	var seen [101]*pkg.Node
// 	var root *pkg.Node
// 	for length := len(queue); length > 0; length = len(queue) {
// 		curr := queue[0]
// 		queue = queue[1:]
// 		if seen[curr.Val] == nil {
// 			seen[curr.Val] = newNode(curr.Val)
// 		}
// 		if root == nil {
// 			root = seen[curr.Val]
// 		}
// 		for _, neighbor := range curr.Neighbors {
// 			if seen[neighbor.Val] == nil {
// 				seen[neighbor.Val] = newNode(neighbor.Val)
// 				queue = append(queue, neighbor)
// 			}
// 			seen[curr.Val].Neighbors = append(seen[curr.Val].Neighbors, seen[neighbor.Val])
// 		}
// 	}
// 	return root
// }
