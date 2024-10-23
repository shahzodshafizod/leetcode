package pkg

// Definition for a Node.
type Node struct {
	Val       int
	Neighbors []*Node
}

func MakeGraph(adjList [][]int) *Node {
	var root *Node
	var ptrs = make(map[int]*Node)
	for node, neighbors := range adjList {
		node++
		if ptrs[node] == nil {
			ptrs[node] = &Node{Val: node, Neighbors: make([]*Node, 0)}
		}
		for _, neighbor := range neighbors {
			if ptrs[neighbor] == nil {
				ptrs[neighbor] = &Node{Val: neighbor, Neighbors: make([]*Node, 0)}
			}
			ptrs[node].Neighbors = append(ptrs[node].Neighbors, ptrs[neighbor])
			if root == nil && ptrs[node].Val == 1 {
				root = ptrs[node]
			}
		}
	}
	return root
}
