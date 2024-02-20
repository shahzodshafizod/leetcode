package graphs

// https://leetcode.com/problems/clone-graph/

// DFS
func cloneGraph(node *Node) *Node {
	var seen [101]*Node
	var newNode = func(val int) *Node { return &Node{Val: val, Neighbors: make([]*Node, 0)} }
	var dfs func(curr *Node) *Node
	dfs = func(curr *Node) *Node {
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
// func cloneGraph(node *Node) *Node {
// 	if node == nil {
// 		return nil
// 	}
// 	var newNode = func(val int) *Node { return &Node{Val: val, Neighbors: make([]*Node, 0)} }
// 	var queue = []*Node{node}
// 	var seen [101]*Node
// 	var root *Node
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
