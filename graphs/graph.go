package graphs

import (
	"math"

	"github.com/shahzodshafizod/alkhwarizmi/design"
)

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

/*
Graph Basics
Graph Algorithms Course for Technical Interviews
https://www.youtube.com/watch?v=2_Uuixtc5i0
(from Alvin Zablan)

graph = nodes(vertices) + edges
nodes: things, edges: relationships
edges are connections between nodes

directed & undirected graphs

(0)---->(2)
 |       |
 |       |
 v       v
(1)<----(4)
 |
 |
 v
(3)---->(5)

how to represent graph information:
- adjacency list
	{
		0: [1, 2],
		1: [3],
		2: [4],
		3: [5],
		4: [1],
		5: []
	}
- matrix

Traversal Approaches (Algorithms):
- Depth-First (uses a Stack)
	0, 1, 3, 5, 2, 4
- Breadth-First (uses a Queue)
	0, 1, 2, 3, 4, 5

asyclic = no cycles
*/

type Structy interface {
	DFS(adjList [][]int) []int
	BFS(adjList [][]int) []int
	/*
		https://structy.net/problems/has-path

		Write a function, has_path, that takes in a dictionary representing
		the adjacency list of a directed acyclic graph and two nodes (src, dst).
		The function should return a boolean indicating whether or not
		there exists a directed path between the source and destination nodes.
	*/
	HasPath(graph map[byte][]byte, src byte, dts byte) bool
	/*
		https://structy.net/problems/undirected-path

		Write a function, undirected_path, that takes in a list of edges
		for an undirected graph and two nodes (node_A, node_B).
		The function should return a boolean indicating whether or not
		there exists a path between node_A and node_B.
	*/
	UndirectedPath(edges [][]byte, nodeA byte, nodeB byte) bool
	/*
		https://structy.net/problems/connected-components-count

		Write a function, connected_components_count, that takes in
		the adjacency list of an undirected graph. The function should return
		the number of connected components within the graph.
	*/
	ConnectedComponentsCount(graph map[int][]int) int
	/*
		https://structy.net/problems/largest-component

		Write a function, largest_component, that takes in
		the adjacency list of an undirected graph. The function should
		return the size of the largest connected component in the graph.
	*/
	LargestComponent(graph map[int][]int) int
	/*
		https://structy.net/problems/shortest-path

		Write a function, shortest_path, that takes in a list of edges
		for an undirected graph and two nodes (node_A, node_B).
		The function should return the length of the shortest path between A and B.
		Consider the length as the number of edges in the path, not the number of nodes.
		If there is no path between A and B, then return -1.
		You can assume that A and B exist as nodes in the graph.
	*/
	ShortestPath(edges [][]byte, nodeA byte, nodeB byte) int
	/*
		https://structy.net/problems/island-count

		Write a function, island_count, that takes in a grid containing Ws and Ls.
		W represents water and L represents land.
		The function should return the number of islands on the grid.
		An island is a vertically or horizontally connected region of land.
	*/
	IslandCount(grid [][]byte) int
	/*
		https://structy.net/problems/minimum-island

		Write a function, minimum_island, that takes in a grid containing Ws and Ls.
		W represents water and L represents land.
		The function should return the size of the smallest island.
		An island is a vertically or horizontally connected region of land.
		You may assume that the grid contains at least one island.
	*/
	MinimumIsland(grid [][]byte) int
}

/*
Graph Theory
Intro & Overview
https://www.youtube.com/watch?v=09_LlHjoEiY
https://www.youtube.com/watch?v=DgXR2OWQnLc&list=PLDV1Zeh2NRsDGO4--qE8yH72HFL1Km93P
(from William Fiset)

=== Brief introduction ===

Graph theory is the mathematical theory of the properties and applications
of graphs (networks).

The goal of this series is to gain an understanding of how to apply graph
theory to real world applications.

A graph theory problem might be:
- How many different sets of clothing can I make by choosing an article from each category?
- Another canonical graph theory example if a social network of friends.

=== Types of Graphs ===

- Undirected Graph
	An undirected graph is a graph in which edges have no orientation.
	The edge (u, v) is identical to the edge (v, u). - Wiki
	Example: cities and roads connecting them

- Directed Graph (Digraph)
	A directed graph or digraph is a graph in which edges have orientations.
	For example, the edge (u, v) is the edge from node u to node v.

- Weighted Graphs
	Many graphs (directed and undirected) can have edges that contain a certain
	weight to represent an arbitrary value such as cost, distance, quantity, etc...

=== Special Graphs ===

- Trees
	A tree is an undirected graph with no cycles.
	Equivalently, it is a connected graph with N nodes and N-1 edges.

- Rooted Trees
	A rooted tree is a tree with a designated root node where every edge either
	points away from or towards the root node.
	When edges point away from the root the graph is called an arborescence
	(out-tree) and anti-arborescence (in-tree) otherwise.

- Directed Asyclic Graphs (DAGs)
	DAGs are directed graphs with no cycles. These graphs play an important
	role in representing structures with dependencies.
	Several efficient algorithms exist to operate on DAGs.

	Cool fact: All out-trees are DAGs but not all DAGs are out-trees.

- Bipartite Graph
	A bipartite graph is one whose vertices can be split into two independent
	groups U, V such that every edge connects between U and V.

	Other definitions exist such as: The graph is two colourable or there is no
	odd length cycle.

- Complete Graphs
	A complete graph is one where there is a unique edge between every pair of nodes.
	A complete graph with n vertices is denoted as the graph Kn.

=== Representing Graphs ===

- Adjacency Matrix
	An adjacency matric m is a very simple way to represent a graph.
	The idea is that the cell m[i][j] represents the edge weight of going
	from node i to node j.

	Note: It is often assumed that the edge of going from a node to itself
	has a cost of zero.

	Pros:
		- Space efficient for representing dense graphs
		- Edge weight lookup is O(1)
		- Simplest graph representation
	Cons:
		- Requires O(V^2) space
		- Iterating over all edges O(V^2) time

- Adjacency List
	An adjacency list is a way to represent a graph as a map from nodes to
	lists of edges.

	Pros:
		- Space efficient for representing sparse graphs
		- Iterating over all edges is efficient
	Cons:
		- Less space efficient for denser graphs
		- Edge weight lookup is O(E)
		- Slightly more complex graph representation

- Edge List
	An edge list is a way to represent a graph simply as an unordered
	list of edges. Assume the notation for any triplet (u,v,w) means:
	"the cost from node u to node v is w"

	This representation is seldomly used because of its lack of structure.
	However, it is conceptually simple and practical in a handful of algorithms.

	Pros:
		- Space efficient for representing sparse graphs
		- Iterating over all edges is efficient
		- Very simple structure
	Cons:
		- Less space efficient for denser graphs
		- Edge weight lookup is O(E)

=== Common Graph Theory Problems ===

	For the upcoming problems ask yourself:
		- Is the graph directed or undirected?
		- Are the edges of the graph weighted?
		- Is the graph I will encounter likely to be sparse or dense with edges?
		- Should I use an adjacency matrix, adjacency list, an edge list or other
			structure to represent the graph efficiently?

Problem #1: Shortest path problem
	Given a weighted graph, find the shortest path of edges from node A to node B.

	Algorithms: BFS (unweighted graph), Dijkstra's, Bellman-Ford, Floyd-Warshall,
	A*, and many more.

Problem #2: Connectivity
	Does there exist a path between node A and node B?

	Tipical solution: Use union find data structure or any search algorithm (e.g. DFS).

Problem #3: Negative cycles
	Does my weighted digraph have any negative cycles? If so, where?

	Algorithms: Bellman-Ford and Floyd-Warshall

Problem #4: Strongly Connected Components
	Strongly Connected Components (SCCs) can be thought of as self-contained
	cycles within a directed graph where every vertex in a given cycle can reach
	every other vertex in the same cycle.

	Algorithms: Tarjan's and Kosaraju's algorithm

Problem #5: Traveling Salesman Problem (TSP)
	"Given a list of cities and the distances between each pair of cities, what
	is the shortest possible route that visits each city exactly once and
	returns to the origin city?" - Wiki

	Algorithms: Held-Karp, branch and bound and many approximation algorithms

	The TSP problem is NP-Hard meaning it's a very computationally challanging problem.
	This is unfortunate because the TSP has several very important applications.

Problem #6: Bridges
	A bridge / cut edge is any edge in a graph whose removal increases the
	number of connected components.

	Bridges are important in graph theory because they often hint at weak
	points, bottlenecks or vulnerabilities in a graph.

Problem #7: Articulation points
	An articulation point / cut vertex is any node in a graph whose removal
	increases the number of connected components.

	Articulation points are important in graph theory because they often hint
	at weak points, bottlenecks or vulnerabilities in a graph.

Problem #8: Minimum Spanning Tree (MST)
	A minimum spanning tree (MST) is a subset of the edges of a connected,
	edge-weighted graph that connects all the vertices together, without any
	cycles and with the minimum possible total edge weight. - Wiki

	Algorithms: Kruskal's, Prim's & Boruvka's algorithm

	MSTs are seen in many applications including: Designing a least cost network,
	circuit design, transportation networks, and etc...

Problem #9: Network flow: max flow
	Q: With an infinite input source how much "flow" can we push through the network?
	Suppose the edges are roads with cars, pipes with water or hallways with packed with people.
	Flow represents the volume of water allowed to flow through the pipes, the
	number of cars the roads can sustain in traffic and the maximum amount of
	people that can navigate through the hallways.

	Algrithms: Ford-Fulkerson, Edmonds-Karp & Dinic's algorithm

=== Depth First Search ===

The Depth First Search (DFS) is the most fundamental search algorithm used to
explore nodes and edges of a graph. It runs with a time complexity of O(V+E)
and is often used as a building block in other algorithms.

By itself the DFS isn't all that useful, but when augmented to perform other
tasks such as count connected components, determine connectivity, or find
bridges/articulation point then DFS really shines.

Basic DFS

As the name suggests, a DFS plunges depth first into a graph without regard for
which edge it takes next until it cannot go any further at which point it
backtracks and continues.

```
	# Global or class scope variables
	n = number of nodes in the graph
	graph = adjacency list representing graph
	visited = [false, ..., false] # size n

	function dfs(at):
		if visited[at]:
			return
		visited[at] = true

		neighbors = graph[at]
		for next in neighbors:
			dfs(next)

	# Start DFS at node zero
	start_node = 0
	dfs(start_node)
```

=== Conneted Components ===

Sometimes a graph is split into multiple components. It's useful to be able to
identify and count these components.

We can use a DFS to identify components. First, make sure all the nodes are
labeled from [0, n) where n is the number of nodes.

Algorithm: Start a DFS at every node (except if it's already been visited) and
mark all reachable nodes as being part of the same component.

```
	# Global or class scope variables
	n = number of nodes in the graph
	g = adjacency list representing graph
	count = 0
	components = empty integer array # size n
	visited = [false, ..., false] # size n

	function findComponents():
		for (i = 0; i < n; i++):
			if !visited[i]:
				count++
				dfs(i)
		return (count, components)

	function dfs(at):
		visited[at] = true
		components[at] = count
		for (next : g[at]):
			if !visited[next]:
				dfs(next)
```

What else can DFS do?

We can augment the DFS algorithm to:
	- Compute a graph's minimum spanning tree.
	- Detect and find cycles in a graph.
	- Check if a graph is bipartite.
	- Find stringly connected components.
	- Topologically sort the nodes of a graph.
	- Find bridges and articulation points.
	- Find augmenting paths in a flow network.
	- Generate mazes.

=== Breadth First Search ===

BFS overview

The Breadth First Search (BFS) is another fundamental search algorithm used to
explore nodes and edges of a graph. It runs with a time complexity of O(V+E) and
is often used as a building block in other algorithms.

The BFS algorithm is particularly useful for one thing: finding the shortest
path on unweighted graphs.

A BFS starts at some arbitrary node of a graph and explores the neighbor
nodes first, before moving to the next level neighbors.

Using a Queue

The BFS algorithm uses a queue data structure to track which node to visit next.
Upon reaching a new node the algorithm adds it to the queue to visit it later.
The queue data structure works just like a real world queue such as a waiting
line at a restaurant. People can either enter the waiting line (enqueue) or get
seated (dequeue).

```
	# Global/class scope variables
	n = number of nodes in the graph
	g = adjacency list representing unweighted graph

	# s = start node, e = end node, and 0 <= e,s < n
	function bfs(s, e):

		# Do a BFS starting at node s
		prev = solve(s)

		# Return reconstructed path from s -> e
		return reconstructPath(s, e, prev)

	function solve(s):
		q = queue data structure with enqueue and dequeue
		q.enqueue(s)

		visited = [false, ..., false] # size n
		visited[s] = true

		prev = [null, ..., null] # size n
		while !q.isEmpty():
			node = q.dequeue()
			neighbors = g.get(node)

			for (next : neighbors):
				if !visited[next]:
					q.enqueue(next)
					visited[next] = true
					prev[next] = node
		return prev

	function reconstructPath(s, e, prev):

		# Reconstruct path going backward from e
		path = []
		for (at = e; at != null; at = prev[at]):
			path.add(at)

		path.reverse()

		# If s and e are connected return the path
		if path[0] == s:
			return path
		return []
```

=== BFS Shortest Path on a Grid ===

Motivation

Many problems in graph theory can be represented using a grid. Grids are a form
of implicit graph because we can determine a node's neighbors based on our
location within the grid.

A type of problem that involves finding a path through a grid is solving a maze.
Another example could be routing through obstacles (trees, rivers, rocks, etc)
to get to a location.

Graph Theory on Grids

A common approach to solving graph theory problems on grids is to first convert
the grid to a familiar format such as an adjacency list/matrix.
IMPORTANT: Assume the grid is unweighted and cells connect left, right, up and
down.

Once we have an adjacency list/matrix we can run whatever specialized graph
algorithm to solve our problem such as: shortest path, connected components,
etc...

Whatever, transformations between graph representations can usually be avoided
due to the structure of a grid.

Direction Vectors

Due to the structure of a grid, if we are at the red ball in the middle we know
we can move left, right, up and down to reach adjacent cells.

Mathematically, if the red ball is at the row-column coordinate (r, c) we can
add the row vectors [-1, 0], [1, 0], [0, 1], and [0, -1] to reach adjacent cells.

If the problem you are trying to solve allows moving diagonally then you can
also include the row vectors: [-1, -1], [-1, 1], [1, 1], [1, -1].

This makes it very easy to access neighboring cells from the current row-column
position:

```
	# Define the direction vectors for
	# north, south, east and west.
	dr = [-1, 1, 0,  0]
	dc = [ 0, 0, 1, -1]

	for (i = 0; i < 4; i++):
		rr = r + dr[i]
		cc = c + dc[i]
		# Skip invalid cells. Assume R and
		# C for the number of rows and columns
		if rr < 0 or cc < 0: continue
		if rr >= R or cc >= C: continue
		# (rr, cc) is a beighboring cell for (r, c)
```

Dungeon Problem Statement

You are tripped in a 2D dungeon and need to find the quickest way out! The
dungeon is composed of unit cubes which may or may not be filled with rock. It
takes one minute to move one unit north, south, east, west. You cannot move
diagonally and the maze is surrounded by solid rock on all sides.

Is an escape possible? If yes, how long will it take?

This is an easier version of the "Dungeon Master" problem on Kattis:
open.kattis.com/problems/dungeon

The dungeon has a size of R x C and you start at cell 'S' and there's an exit at
cell 'E'. A cell full of rock is indicated by a '#' and empty cells are
represented by a '.'

|S|.|.|#|.|.|.|
|.|#|.|.|.|#|.|
|.|#|.|.|.|.|.|
|.|.|#|#|.|.|.|
|#|.|#|E|.|#|.|

Alternative State representation

So far we have been storing the next x-y position in the queue as an (x, y) pair.
This works well but requires either an array or an object wrapper to store the
coordinate values. In practice, this requires a lot of packing and unpacking of
values to and from the queue.

Let's take a look at an alternative approach which also scales well in higher
dimensions and (IMHO) requires less setup effort...

An alternative approach is to use one queue for each dimension, so in a 3D grid
you would have one queue for each of the x, y, and z dimensions.

```
	# Global/class scope variables
	R, C = ...   # R = number of rows, C = number of columns
	m = ...      # Input character matrix of size R x C
	sr, sc = ... # 'S' symbol row and column values
	rq, cq = ... # Empty Row Queue (RQ) and Column Queue (CQ)

	# Variables used to track the number of steps taken.
	move_count = 0
	nodes_left_in_layer = 1
	nodes_in_next_layer = 0

	# Variable used to track whether the 'E' character
	# ever gets reached during the BFS.
	reached_end = false

	# R x C matrix of false values used to track whether
	# the node at position (i, j) has been visited.
	visited = ...

	# North, south, east, west direction vectors.
	dr = [-1, 1, 0,  0]
	dc = [ 0, 0, 1, -1]

	function solve():
		rq.enqueue(sr)
		cq.enqueue(sc)
		visited[sr][sc] = true
		while rq.size() > 0: # or cq.size() > 0
			r = rq.dequeue()
			c = cq.dequeue()
			if m[r][c] == 'E':
				reached_end = true
				break
			explore_neighbors(r, c)
			nodes_left_in_layer--
			if nodes_left_in_layer == 0:
				nodes_left_in_layer = nodes_in_next_layer
				nodes_in_next_layer = 0
				move_count++
		if reached_end:
			return move_count
		return -1

	function explore_neighbors(r, c):
		for (i = 0; i < 4; i++):
			rr = r + dr[i]
			cc = c + dc[i]

			# Skip out of bounds locations
			if rr < 0 or cc < 0: continue
			if rr >= R or cc >= C: continue

			# Skip visited locations or blocked cells
			if visited[rr][cc]: continue
			if m[rr][cc] == '#: continue

			rq.enqueue(rr)
			cq.enqueue(cc)
			visited[rr][cc] = true
			nodes_in_next_layer++
```

=== Topological Sort ===

Many real world situations can be modelled as a graph directed edges where some
events must occur before others.
	- School class prerequisites
	- Program dependencies
	- Event scheduling
	- Assembly instructions
	- Etc...

Suppose you're a student at university X and you want to take Class H, then you
must take classes A, B, D and E as prerequisites. In this sense there is an
ordering on the nodes of the graph.

                   +=========+
                   | Class C |
+=========+------->+=========+------->+=========+
| Class A |                           | Class J |
+=========+------->+=========+------->+=========+
                   | Class D |
              +--->+=========+----+
              |                   |
              |    +=========+    +-->+=========+
              |    | Class E |        | Class H |
+=========+---+    +=========+------->+=========+
| Class B |
+=========+------->+=========+        +=========+
                   | Class F |        | Class I |
                   +=========+        +=========+

Another canonical example where an ordering on the nodes of the graph matters is
for program build dependencies. A program cannot be built unless its dependencies
are first built.

A topological ordering is an ordering of the nodes in a directed graph where for
each directed edge from node A to node B, node A appears before node B in the
ordering.

The topological sort algorithm can find a topological ordering in O(V+E) time!

NOTE: Topological orderings are NOT unique.

=== Directed Acyclic Graphs (DAG) ===

Not every graph can have a topoligical ordering. A graph which contains a cycle
cannot have a valid ordering.

The only type of graph which has a valid topological ordering is a Directed
Acyclic Graph (DAG). These are graphs with directed edges and no cycles.

Q: How do I verify that my graph does not contain a directed cycle?
A: One method is to use Tarjan's strongly connected component algorithm which
can be used to find these cycles.

By definition, all rooted trees have a topological ordering since they do not
contain any cycles.

Topological Sort Algorithm

- Pick an unvisited node
- Beginning with the selected node, do a Depth First Search (DFS) exploring only
	unvisited nodes.
- On the recurive callback of the DFS, add the current node to the topological
	ordering in reverse order.

Topsort pseudocode

```
	# Assumption: graph is sorted as adjacency list
	function topsort(graph):
		N = graph.numberOfNodes()
		V = [false, ..., false] # Length N
		ordering = [0, ..., 0] # Length N
		i = N - 1 # Index for ordering array

		for (at = 0; at < N; at++):
			if V[at] == false:
				visitedNodes = []
				dfs(at, V, visitedNodes, graph)
				for nodeId in visitedNodes:
					ordering[i] = nodeId
					i = i - 1
			return ordering

	# Execute Depth First Search (DFS)
	function dfs(at, V, visitedNodes, graph):

		V[at] = true

		edges = graph.getEdgesOutFromNode(at)
		for edge in edges:
			if V[edge.to] == false:
				dfs(edge.to, V, visitedNodes, graph)

		visitedNodes.add(at)
```

Topsort Optimization

```
	# Assumption: graph is stored as adjacency list
	function topsort(graph):

		N = graph.numberOfNodes()
		V = [false, ..., false] # Length N
		ordering = [0, ..., 0] # Length N
		i = N - 1 # Index for ordering array

		for (at = 0; at < N; at++):
			if V[at] == false:
				i = dfs(i, at, V, ordering, graph)

		return ordering

	# Execute Depth First Search (DFS)
	function dfs(i, at, V, ordering, graph):

		V[at] = true

		edges = graph.getEdgesOutFromNode(at)
		for edge in edges:
			if V[edge.to] == false:
				i = dfs(i, edge.to, V, ordering, graph)

		ordering[i] = at
		return i - 1
```

=== Shortest and longest paths on DAGs ===

Recall that a Directed Acyclic Graph (DAG) is a graph with directed edges and no
cycles. By definition this means all trees are automatically DAGs since they do
not contain cycles.

SSSP on DAG

The Single Source Shortest Path (SSSP) problem can be solved efficiently on a
DAG in O(V+E) time. This is due to the fact that the nodes can be ordered in a
topological ordering via topsort and processed sequentially.

Longest path on DAG

What about the longest path? On a general graph this problem is NP-Hard, but on
a DAG this problem is solvable in O(V+E)!

The trick is to multiply all edge values by -1 then find the shortest path and
then multiply the edge values by -1 again!

Source Code Link
Implementation source code can be found at the following link:
	github.com/williamfiset/algorithms

A useful application of the topological sort is to find the shortest path
between two nodes in a Directed Acyclic Graph (DAG). Given an adjacency list
this method finds the shortest path to all nodes starting at 'start'

NOTE: 'numNodes' is not necessarily the number of nodes currently present
in the adjacency list since you can have singleton nodes with no edges which
wouldn't be present in the adjacency list but are still part of the graph!

```
	func dagShortestPath(graph map[int][]int, start int, numNodes int) []int {
		var topsort = topologicalSort(graph, numNodes)
		var dist = make([]any, numNodes)
		dist[0] = 0

		for i := 0; i < numNodes; i++ {

			nodeIndex = topsort[i]
			if dist[nodeIndex] != nil {
				var adjacentEdges = graph.get(nodeIndex)
				for _, edge := range adjacentEdges {
					var newDist = dist[nodeIndex].(int) + edge.weight
					if dist[edge.to] == nil || newDist < dist[edge.to].(int) {
						dist[edge.to] = newDist
					}
				}
			}
		}

		return dist
	}
```

=== Dijkstra's Shortest Path Algorithm ===

What is Dijkstra's algorithm?

Dijkstra's algorithm is a Single Source Shortest Path (SSSP) algorithm for
graphs with non-negative edge weights.

Depending on how the algorithm is implemented and what data structures are used
the time complexity is typically O(E*log(V)) which is competitive against other
shortest path algorithms.

Algorithm prerequisites

One constraint for Dijkstra's algorithm is that the graph must only contain
non-negative edge weights. This constraint is imposed to ensure that once a node
has been visited its optimal distance cannot be improved.

This property is especially important because it enables Dijkstra's algorithm
to act in a greedy manner by always selecting the next most promising node.

Outline

The goal of this slide deck is for you to understand how to implement Dijkstra's
algorithm and implement it efficiently.

- Lazy Dijkstra's animation
- Lazy Dijkstra's pseudo-code
- Finding SP + stopping early optimazation
- Using indexed priority queue + decreaseKey to reduce soace and increase
	performance.
- Eager Dijkstra's animation
- Eager Dijkstra's pseudo-code
- Heap optimization with D-ary heap

Quick Algorithm Overview

Maintain a 'dist' array where the distance to every node is positive infinity.
Mark the distance to the shortest start node 's' to be 0.

Maintain a PQ of key-value pairs of (node index, distance) pairs which tell you
which node to visit next based on sorted min value.

Insert (s, 0) into the PQ and loop while PQ is not empty pulling out the next
most promising (node index, distance) pair.

Iterate over all edges outwards from the current node and relax each edge
appending a new (node index, distance) key-value pair to the PQ for every
relaxation.

Lazy Dijkstra's

```
	# Runs Dijkstra's algorithm and returns an array that contains
	# the shortest distance to every node from the start node s.
	# g - adjacency list of weighted graph
	# n - the number of nodes in the graph
	# s - the index of the starting node (0 <= s < n)
	function dijkstra(g, n, s):
		vis = [false, false, ..., false] # size n
		dist = [+inf, +inf, ..., +inf, +inf] # size n
		dist[s] = 0
		pq = empty priority queue
		pq.insert((s, 0))
		while pq.size() != 0:
			index, minValue = pq.poll()
			vis[index] = true
			for (edge : g[index]):
				if vis[edge.to]: continue
				newDist = dist[index] + edge.cost
				if newDist < dist[edge.to]:
					dist[edge.to] = newDist
					pq.insert((edge.to, newDist))
		return dist
```

In practice most standard libraries do not support the decrease key operation
for PQs. A way to get around this is to add a new (node index, best distance)
pair every time we update the distance to a node.

As a result, it is possible to have duplicate node indices in the PQ. Ths is not
ideal, but inserting a new key-value pair in O(log(n)) is much faster than
searching for the key in the PQ which takes O(n)

A neat optimization we can do which ignores stale (index, dist) pairs in our PQ
is to skip nodes where we already found a better path routing through other
nodes before we got to processing this node.

```
	# Runs Dijkstra's algorithm and returns an array that contains
	# the shortest distance to every node from the start node s.
	# g - adjacency list of weighted graph
	# n - the number of nodes in the graph
	# s - the index of the starting node (0 <= s < n)
	function dijkstra(g, n, s):
		vis = [false, false, ..., false] # size n
		dist = [+inf, +inf, ..., +inf, +inf] # size n
		dist[s] = 0
		pq = empty priority queue
		pq.insert((s, 0))
		while pq.size() != 0:
			index, minValue = pq.poll()
			vis[index] = true
			if dist[index] < minValue: continue
			for (edge : g[index]):
				if vis[edge.to]: continue
				newDist = dist[index] + edge.cost
				if newDist < dist[edge.to]:
					dist[edge.to] = newDist
					pq.insert((edge.to, newDist))
		return dist
```

Finding the optimal path

If you wish to not only find the optimal distance to a particular node but also
what sequence of nodes were taken to get there you need to track some additional
information.

```
	# Runs Dijkstra's algorithm and returns an array that contains
	# the shortest distance to every node from the start node s and
	# the prev array to reconstruct the shortest path itself
	# g = adjacency list of weighted graph
	# n - the number of nodes in the graph
	# s - the index of the starting node (0 <= s < n)
	function dijkstra(g, n, s):
		vis = [false, false, ..., false] # size n
		prev = [null, null, ..., null] # size n
		dist = [+inf, +inf, ..., +inf, +inf] # size n
		dist[s] = 0
		pq = empty priority queue
		pq.insert((s, 0))
		while pq.size() != 0:
			index, minValue = pq.poll()
			vis[index] = true
			if dist[index] < minValue: continue
			for (edge : g[index]):
				if vis[edge.to]: continue
				newDist = dist[index] + edge.cost
				if newDist < dist[edge.to]:
					prev[edge.to] = index
					dist[edge.to] = newDist
					pq.insert((edge.to, newDist))
		return (dist, prev)

	# Finds the shortest path between two nodes.
	# g - adjacency list of weighted graph
	# n - the number of nodes in the graph
	# s - the index of the starting node (0 <= e < n)
	# e - the index of the end node (0 <= e < n)
	function findShortestPath(g, n, s, e):
		dist, prev = dijkstra(g, n, s)
		path = []
		if (dist[e] == +inf) return path
		for (at = e; at != null; at = prev[at]):
			path.add(at)
		path.reverse()
		return path
```

Stopping Early

Q: Suppose you know the destination node you're trying to reach is 'e' and you
start at node 's' do you still have to visit every node in the graph?

A: Yes, in the worst case. However, it is possible to stop early once you have
finished visiting the destination node.

The main idea for stopping early is that Dijkstra's algorithm processes each
next most promising node in order. So if the destination node has been visited,
its shortest distance will not change as more future nodes are visited.

```
	# Runs Dijkstra's algorithm and returns the shortest distance
	# between nodes 's' and 'e'. If there is no path, +inf is returned.
	# g - adjacency list of weighted graph
	# n - the number of nodes in the graph
	# s - the index of the starting node (0 <= s < n)
	# e - the index of the end node (0 <= e < n)
	function dijkstra(g, n, s, e):
		vis = [false, false, ..., false] # size n
		dist = [+inf, +inf, ..., +inf, +inf] # size n
		dist[s] = 0
		pq = empty priority queue
		pq.insert((s, 0))
		while pq.size() != 0:
			index, minValue = pq.poll()
			vis[index] = true
			if dist[index] < minValue: continue
			for (edge : g[index]):
				if vis[edge.to]: continue
				newDist = dist[index] + edge.cost
				if newDist < dist[edge.to]:
					dist[edge.to] = newDist
					pq.insert((edge.to, newDist))
			if index == e:
				return dist[e]
		return +inf
```

Eager Dijkstra's using an Indexed Priority Queue

Our current lazy implementation of Dijkstra's inserts duplicate key-value pairs
(keys being the node index and the value being the shortest distance to get to
that node) in our PQ because it's more efficient to insert a new key-value pair
in O(log(n)) than it is to update an existing key's value in O(n)

This approach is inefficient for dense graphs because we end up with several
stale outdated key-value pairs in our PQ. The eager version of Dijkstra's avoids
duplicate key-value pairs and supports efficient value updates in O(log(n))
by using an Indexed Priority Queue (IPQ)

continue from 1:34:44

*/

type Graph interface {
	DFS(adjList [][]int) []int
	BFS(adjList [][]int) []int
	Dungeon(grid [][]byte) int
}

type graph struct{}

var _ Structy = &graph{}
var _ Graph = &graph{}

func (g *graph) DFS(adjList [][]int) []int {
	var dfs func(curr int, visited map[int]bool) []int
	dfs = func(curr int, visited map[int]bool) []int {
		var nodes = []int{curr}
		visited[curr] = true
		for _, neighbor := range adjList[curr] {
			if !visited[neighbor] {
				nodes = append(nodes, dfs(neighbor, visited)...)
			}
		}
		return nodes
	}
	return dfs(0, make(map[int]bool))
}

func (g *graph) BFS(adjList [][]int) []int {
	var nodes = make([]int, 0)
	var queue = design.NewQueue[int]()
	queue.Enqueue(0)
	var visited = make(map[int]bool)
	for !queue.Empty() {
		curr := queue.Dequeue()
		nodes = append(nodes, curr)
		visited[curr] = true
		for _, neighbor := range adjList[curr] {
			if !visited[neighbor] {
				queue.Enqueue(neighbor)
			}
		}
	}
	return nodes
}

// N = # nodes
// E = # edges
// Time: O(E) = O(N^2)
// Space: O(N)
func (g *graph) HasPath(graph map[byte][]byte, src byte, dst byte) bool {
	var dfs func(src byte, dst byte, visited map[byte]bool) bool
	dfs = func(src byte, dst byte, visited map[byte]bool) bool {
		if src == dst {
			return true
		}
		visited[src] = true
		for _, neighbor := range graph[src] {
			if !visited[neighbor] && dfs(neighbor, dst, visited) {
				return true
			}
		}
		return false
	}
	return dfs(src, dst, make(map[byte]bool))
}

// N = # nodes
// E = # edges
// Time: O(E) = O(N^2)
// Space: O(N)
func (g *graph) UndirectedPath(edges [][]byte, nodeA byte, nodeB byte) bool {
	var adjList = make(map[byte][]byte)
	var a, b byte
	for _, edge := range edges {
		a, b = edge[0], edge[1]
		adjList[a] = append(adjList[a], b)
		adjList[b] = append(adjList[b], a)
	}
	return g.HasPath(adjList, nodeA, nodeB)
}

// N = # nodes
// E = # edges
// Time: O(E) = O(N^2)
// Space: O(N)
func (g *graph) ConnectedComponentsCount(graph map[int][]int) int {
	var dfs func(node int, visited map[int]bool)
	dfs = func(node int, visited map[int]bool) {
		visited[node] = true
		for _, neighbor := range graph[node] {
			if !visited[neighbor] {
				dfs(neighbor, visited)
			}
		}
	}
	var visited = make(map[int]bool)
	var count = 0
	for src := range graph {
		if !visited[src] {
			count++
			dfs(src, visited)
		}
	}
	return count
}

// N = # nodes
// E = # edges
// Time: O(E) = O(N^2)
// Space: O(N)
func (g *graph) LargestComponent(graph map[int][]int) int {
	var count func(node int, visited map[int]bool) int
	count = func(node int, visited map[int]bool) int {
		if visited[node] {
			return 0
		}
		visited[node] = true
		var size = 1
		for _, neighbor := range graph[node] {
			size += count(neighbor, visited)
		}
		return size
	}
	var visited = make(map[int]bool)
	var maxsize = 0
	for src := range graph {
		maxsize = max(maxsize, count(src, visited))
	}
	return maxsize
}

// E = # edges
// Time: O(E)
// Space: O(E)
func (g *graph) ShortestPath(edges [][]byte, nodeA byte, nodeB byte) int {
	var adjList = make(map[byte][]byte)
	var a, b byte
	for _, edge := range edges {
		a, b = edge[0], edge[1]
		adjList[a] = append(adjList[a], b)
		adjList[b] = append(adjList[b], a)
	}
	type node struct {
		val      byte
		distance int
	}
	// BFS
	var queue = design.NewQueue[*node]()
	var visited = make(map[byte]bool)
	queue.Enqueue(&node{val: nodeA, distance: 0})
	for !queue.Empty() {
		current := queue.Dequeue()
		visited[current.val] = true
		if current.val == nodeB {
			return current.distance
		}
		for _, neighbor := range adjList[current.val] {
			if !visited[neighbor] {
				queue.Enqueue(&node{val: neighbor, distance: current.distance + 1})
			}
		}
	}
	return -1
}

// r = # rows
// c = # cols
// Time: O(rc)
// Space: O(rc)
func (g *graph) IslandCount(grid [][]byte) int {
	var m, n = len(grid), len(grid[0])
	var directions = [5]int{-1, 0, 1, 0, -1}
	var dfs func(row int, col int, visited [][]bool) int
	dfs = func(row int, col int, visited [][]bool) int {
		if visited[row][col] || grid[row][col] != 'L' {
			return 0
		}
		visited[row][col] = true
		var r, c int
		for dir := 1; dir < 5; dir++ {
			r = row + directions[dir-1]
			c = col + directions[dir]
			if min(r, c) >= 0 && r < m && c < n {
				dfs(r, c, visited)
			}
		}
		return 1
	}
	var visited = make([][]bool, m)
	for row := 0; row < m; row++ {
		visited[row] = make([]bool, n)
	}
	var islands = 0
	for row := 0; row < m; row++ {
		for col := 0; col < n; col++ {
			islands += dfs(row, col, visited)
		}
	}
	return islands
}

// r = # rows
// c = # cols
// Time: O(rc)
// Space: O(rc)
func (g *graph) MinimumIsland(grid [][]byte) int {
	var m, n = len(grid), len(grid[0])
	var directions = [5]int{-1, 0, 1, 0, -1}
	var count func(row int, col int, visited [][]bool) int
	count = func(row int, col int, visited [][]bool) int {
		if visited[row][col] || grid[row][col] != 'L' {
			return 0
		}
		visited[row][col] = true
		var size = 1
		var r, c int
		for dir := 1; dir < 5; dir++ {
			r, c = row+directions[dir-1], col+directions[dir]
			if min(r, c) >= 0 && r < m && c < n {
				size += count(r, c, visited)
			}
		}
		return size
	}
	var visited = make([][]bool, m)
	for row := 0; row < m; row++ {
		visited[row] = make([]bool, n)
	}
	var minSize = math.MaxInt
	for row := 0; row < m; row++ {
		for col := 0; col < n; col++ {
			if count := count(row, col, visited); count > 0 {
				minSize = min(minSize, count)
			}
		}
	}
	return minSize
}

// Approach: BFS
func (g *graph) Dungeon(grid [][]byte) int {
	var directions = [5]int{-1, 0, 1, 0, -1}
	var m, n = len(grid), len(grid[0])
	var visited = make([][]bool, m)
	for idx := range visited {
		visited[idx] = make([]bool, n)
	}
	var queue = design.NewQueue[[3]int]()
	for row := 0; row < m && queue.Empty(); row++ {
		for col := 0; col < n && queue.Empty(); col++ {
			if grid[row][col] == 'S' {
				queue.Enqueue([3]int{row, col, 1})
				visited[row][col] = true
			}
		}
	}
	var row, col, path, r, c int
	for !queue.Empty() {
		var current = queue.Dequeue()
		row, col, path = current[0], current[1], current[2]
		for dir := 1; dir < 5; dir++ {
			r = row + directions[dir-1]
			c = col + directions[dir]
			if min(r, c) >= 0 && r < m && c < n && !visited[r][c] {
				switch grid[r][c] {
				case 'E':
					return path + 1
				case '.':
					queue.Enqueue([3]int{r, c, path + 1})
					visited[r][c] = true
				}
			}
		}
	}
	return -1
}
