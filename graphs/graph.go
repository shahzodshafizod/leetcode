package graphs

/*
	(1)			(4)-----------(6)
	 |    		 |			   |
	(0)---------(3)----(5)	  (7)
				 |
				(2)----(8)

Adjacency List:
	[[1, 3], [0], [3, 8], [4, 5, 2, 0], [6, 3], [3], [7, 4], [6], [2]]

Adjacency Matrix:
	[[0, 1, 0, 1, 0, 0, 0, 0, 0]
	 [1, 0, 0, 0, 0, 0, 0, 0, 0]
	 [0, 0, 0, 1, 0, 0, 0, 0, 1]
	 [1, 0, 1, 0, 1, 1, 0, 0, 0]
	 [0, 0, 0, 1, 0, 0, 1, 0, 0]
	 [0, 0, 0, 1, 0, 0, 0, 0, 0]
	 [0, 0, 0, 0, 1, 0, 0, 1, 0]
	 [0, 0, 0, 0, 0, 0, 1, 0, 0]
	 [0, 0, 1, 0, 0, 0, 0, 0, 0]]
*/

// using adjacency list
func traversalBFS(graph [][]int) []int {
	var values = make([]int, 0)
	length := len(graph)
	if length == 0 {
		return values
	}
	var queue = []int{0}
	var seen = map[int]bool{0: true}
	for length := len(queue); length > 0; length = len(queue) {
		for count := 0; count < length; count++ {
			var vertex = queue[count]
			values = append(values, vertex)
			seen[vertex] = true
			var connections = graph[vertex]
			for _, connection := range connections {
				if !seen[connection] {
					queue = append(queue, connection)
				}
			}
		}
		queue = queue[length:]
	}
	return values
}

func traversalDFS(graph [][]int) []int {
	return traversalDFSHelper(graph, 0, make(map[int]bool))
}

func traversalDFSHelper(graph [][]int, vertex int, seen map[int]bool) []int {
	var values = []int{vertex}
	seen[vertex] = true
	var connections = graph[vertex]
	for _, connection := range connections {
		if !seen[connection] {
			values = append(values, traversalDFSHelper(graph, connection, seen)...)
		}
	}
	return values
}

// Definition for a Node.
type Node struct {
	Val       int
	Neighbors []*Node
}

func makeGraph(adjList [][]int) *Node {
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

// func printGraph(root *Node, seen map[*Node]bool) {
// 	if root == nil {
// 		return
// 	}
// 	seen[root] = true
// 	var neighbors = make([]int, 0)
// 	for _, neighbor := range root.Neighbors {
// 		neighbors = append(neighbors, neighbor.Val)
// 		if seen[neighbor] {
// 			continue
// 		}
// 		printGraph(neighbor, seen)
// 	}
// 	fmt.Printf("%d: %v\n", root.Val, neighbors)
// }
