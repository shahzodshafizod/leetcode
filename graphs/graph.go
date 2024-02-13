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
