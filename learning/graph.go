package learning

import (
	"container/heap"
	"math"

	"github.com/shahzodshafizod/leetcode/pkg"
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

type Graph interface {
	DFS(adjList [][]int) []int
	BFS(adjList [][]int) []int
	TopologicalSortBFS(adjList map[int][]*Edge, n int) []int
	TopologicalSortDFS(adjList map[int][]*Edge, n int) []int
	Dungeon(grid [][]byte) int
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
				newDist = dist[index] + edge.Weight
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
				newDist = dist[index] + edge.Weight
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
				newDist = dist[index] + edge.Weight
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
				newDist = dist[index] + edge.Weight
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

Indexed Priority Queue

Insertion Pseudo Code
```
	# Inserts a value into the min indexed binary
	# heap. The key index must not already be in
	# the heap and the value must not be null.
	function insert(ki, value):
		values[ki] = value
		# 'sz' is the current size of the heap
		pm[ki] = sz
		im[sz] = ki
		swim(sz)
		sz = sz + 1
```

Swim Pseudo Code
```
	# Swims up node i (zero based) until heap
	# invariant is satisfied.
	function swim(i):
		for (p = (i-1)/2; i > 0 and less(i, p)):
			swap(i, p)
			i = p
			p = (i-1)/2

	function swap(i, j):
		pm[im[j]] = i
		pm[im[i]] = j
		tmp = im[i]
		im[i] = im[j]
		im[j] = tmp

	function less(i, j):
		return values[im[i]] < values[im[j]]
```

Polling & Removals
Polling and removing in a min indexed binary heap are both O(log(n)) operations
rather than O(n) for a traditional PQ since node position lookups are O(1) but
repositioning is still O(log(n))

Eager Dijkstra's

```
	# Runs Dijkstra's algorithm and returns an array that contains
	# the shortest distance to every node from the start node s.
	# g - adjacency list of weighted graph
	# n - the number of nodes in the graph
	# s - the index of the starting node (0 <= s < n)
	function dijkstra(g, n, s):
		vis = [false, false, ..., false] # size n
		dist = [+inf, +inf, ..., +inf, +inf] # size
		dist[s] = 0
		ipq = empty indexed priority queue
		ipq.insert(s, 0)
		while ipq.size() != 0:
			index, minValue = ipq.poll()
			vis[index] = true
			if dist[index] < minValue: continue
			for (edge : g[index]):
				if vis[edge.to]: continue
				newDist = dist[index] + edge.Weight
				if newDist < dist[edge.to]:
					dist[edge.to] = newDist
					if !ipq.contains(edge.to):
						ipq.insert(edge.to, newDist)
					else:
						ipq.decreaseKey(edge.to, newDist)
		return dist
```
The main advantage of using decreaseKey is to prevent duplicate node indexes to
be present in the PQ.

D-ary Heap Optimization

When executing Dijkstra's algorithm, especially on dense graphs, there are a
lot more updates (i.e. decreaseKey operations) to key-value pairs than there are
deqeue (poll) operations.

A D-ary heap is a heap variant in which each node has D children, This speeds up
decrease key operations at the expense of more costly removals.

Optimal D-ary Heap degree

Q: What is the optimal D-ary heap degree to maximize performance of Dijkstra's
algorithm?

A: In general D = E/V is the best degree to use to balance removals against
decreaseKey operations improving Dijkstra's time complexity to O(E*log E/V (V))
which is much better especially for dense graphs which have lots of decreaseKey
operations.

The state of the art

The current state of the art as of now is the Fibonacci heap which gives
Dijkstra's algorithm a time complexity of O(E + Vlog(V))

However, in practice, Fibonacci heaps are very difficult to implement and have
a large enough constant amortized overhead to make them impractical unless your
graph is quite large.

=== Dijkstra's Shortest Path Algorithm Source Code ===

```
	// An edge class to represent a directed edge
	// between two nodes with a certain cost.
	type Edge struct {
		to int
		cost float64
	}

	var n int
	var edgeCount int
	var prev []*int
	var graph [][]*Edge

	// n - The number of nodes in the graph
	func DijkstrasShortestPathAdjacencyListWithDHeap(n int) {
		this.n = n
		createEmptyGraph(n)
	}

	// Construct an empty graph with n nodes including the source and sink nodes.
	func createEmptyGraph(n int) {
		graph = make([][]*Edge, n)
		for idx := range graph {
			graph[i] = make([]*Edge, 0)
		}
	}

	// Adds a directed edge to the graph
	// from - The index of the node the directed edge starts at.
	// to - The index of the node the directed edge ends at.
	// cost - The cost of the edge.
	func addEdge(from int, to int, cost int) {
		edgeCount++
		graph[from] = append(graph[from], &Edge{to, cost})
	}

	// Run Dijkstra's algorithm on a directed graph to find the shortest path
	// from a starting node to an ending node. If there is no path between the
	// starting node and the destination node the returned value is set to be
	// Double.POSITIVE_INFINITY.
	func dijkstra(start int, end int) float64 {

		// Keep an Indexed Priority Queue (ipq) of the next most promising node
		// to visit.
		var degree int = edgeCount / n
		var ipq *MinIndexedDHeap = &MinIndexedDHeap{degree, n}
		ipq.insert(start, 0.0)

		// Maintain an array of the minimum distance to each node.
		dist = make([]float64, n)
		for idx := range dist {
			dist[idx] = Double.POSITIVE_INFINITY
		}
		dist[start] = 0.0

		var visited = make([]bool, n)
		prev = make([]*int, n)

		for !ipq.isEmpty() {
			var nodeId = ipq.peekMinKeyIndex()

			visited[nodeId] = true
			var minValue = ipq.pollMinValue()

			// We already found a better path before we got to
			// processing this node so we can ignore it.
			if minValue > dist[nodeId] { continue }

			for _, edge := range graph[nodeId] {

				// We cannot get a shorter path by revisiting
				// a node we have already visited before.
				if visited[edge.to] { continue }

				// Relax edge by updating minimum cost if applicable.
				var newDist = dist[nodeId] + edge.Weight
				if newDist < dist[edge.to] {
					prev[edge.to] = nodeId
					dist[edge.to] = newDist
					// Insert the cost of going to a node for the first time in
					// the PQ, or try and update it to a better value by calling
					// decrease.
					if !ipq.contains(edge.to) {
						ipq.insert(edge.to, newDist)
					} else {
						ipq.decrease(edge.to, newDist)
					}
				}
			}
			// Once we've processed the end node we can return early (without
			// necessarily visiting the whole graph) because we know we cannot
			// get a shorter path by routing through any other nodes since
			// Dijkstra's is greedy and there are no negative edge weights.
			if nodeId == end {
				return return dist[end]
			}
		}
		// End node is unreachable.
		return Double.POSITIVE_INFINITY
	}

	func reconstructPath(start int, end int) {
		var path = make([]int, 0)
		var dist = dijkstra(start, end)
		if dist == Double.POSITIVE_INFINITY {
			return path
		}
		for at := end; at != null; at = prev[at] {
			path = append(path, at)
		}
		path.reverse()
		return path
	}
```

=== Bellman-Ford Algorithm ===

BF algorithm overview

In graph theory, the Bellman-Ford (BF) algorithm is a Single Source Shortest
Path (SSSP) algorithm. This means it can find the shortest path from one node
to any other node.

However, BF is not ideal for most SSSP problems because it has a time complexity
of O(EV). It is better to use Dijkstra's algorithm which is much faster. It is
on the order of O((E+V)log(V)) when using a binary heap priority queue.

However, Dijkstra's algorithm can fail when the graph has negative edge weights.
This is when BF becomes really handy because it can be used to detect negative
cycles and determine where they occur.

Finding negative cycles can be useful in many types of applications. One
particularly neat application arises in finance when performing an arbitrage
between two or more markets.

Negative Cycles

Negative Cycles can manifest themselves in many ways...

BF Algorithm Steps

Let's define a few variables...
	- Let E be the number of edges.
	- Let V be the number of vertices.
	- Let S be the id of the starting node.
	- Let D be an array of size V that tracks the best distance from S to each
		node.

1) Set every entry in D to +inf
2) Set D[S] = 0
3) Relax eah edge V-1 times:
```
	for (i := 0; i < V-1; i = i + 1):
		for edge in graph.edges:
			// Relax edge (update D with shorter path)
			if (D[edge.from] + edge.Weight < D[edge.to])
				D[edge.to] = D[edge.from] + edge.Weight

	// Repeat to find nodes caught in a negative cycle
	for (i = 0; i < V-1; i = i + 1):
		for edge in graph.edges:
			if (D[edge.from] + edge.Weight < D[edge.to])
				D[edge.to] = -inf
```

=== Floyd-Warshall Algorithm ===

FW algorithm overview

In graph theory, the Floyd-Warshall (FW) algorithm is an All-Pairs Shortest Path
(APSP) algorithm. This means it can find the shortest path between all pairs of
nodes.

The time complexity to run FW is O(V^3) which is ideal for graphs no larger
than a couple hundred nodes.

Graph setup

With FW, the optimal way to represent our graph is with a 2D adjacency matrix m
where cell m[i][j] represents the edge weight of going from node i to node j

A->B: 4
B->A: 3
A->C: 1
C->A: 4
A->D: 9
D->A: 6
B->C: 6
C->B: 1
B->D: 11
D->B: 5
C->D: 2
D->C: -4

=>

+---+---+---+----+----+
|   | A | B | C  | D  |
+---+---+---+----+----+
| A | 0 | 4 |  1 |  9 |
+---+---+---+----+----+
| B | 3 | 0 |  6 | 11 |
+---+---+---+----+----+
| C | 4 | 1 |  0 |  2 |
+---+---+---+----+----+
| D | 6 | 5 | -4 |  0 |
+---+---+---+----+----+

NOTE: In the graph above, it is assumed that the distance from a node to itself
is zero. This is why the diagonal is all zeros.

If there is no edge from node i to node j then set the edge value for m[i][j]
to be positive infinity.

A->B: 4
A->C: 1
C->A: 4
B->C: 6
C->B: 1
C->D: 2

+---+------+------+------+------+
|   |   A  |   B  |   C  |   D  |
+---+------+------+------+------+
| A |   0  |   4  |   1  | +inf |
+---+------+------+------+------+
| B | +inf |   0  |   6  | +inf |
+---+------+------+------+------+
| C |   4  |   1  |   0  |   2  |
+---+------+------+------+------+
| D | +inf | +inf | +inf |   0  |
+---+------+------+------+------+

IMPORTANT: If your programming language does not support a special constant for
+inf such that inf + inf = inf and x + inf = inf then avoid using (2^31)-1 as
infinity! This will cause integer overflow; prefer to use a large constant such
as 10^7 instead.

The main idea behind the Floyd-Warshall algorithm is to gradually build up all
intermediate routes between nodes i and j to find the optimal path.

Suppose our adjacency matrix tells us that the distance from a to b is:
m[a][b] = 11.
Suppose there exists a third node, c. m[a][c] = 5, m[c][b] = 5
If m[a][c] + m[c][b] < m[a][b] then it's better to route through c!

The goal of Floyd-Warshall is to eventually consider going through all possible
intermediate nodes on paths of different lengths.

The Memo Table

Let 'dp' (short for Dynamic Programming) be a 3D matrix of size n x n x n that
acts as a memo table.

dp[k][i][j] = shortest path from i to j routing through nodes {0,1,...,k-1,k}

Start with k = 0, then k = 1, then k = 2, ... This gradually builds up the
optimal solution routing through 0, then all optimal solutions routing through
0 and 1, then all optimal solutions routing through 0, 1, 2, etc... up until
n-1 which stores to APSP solution.
Specially dp[n-1] is the 2D matrix solution we're after.

In the beginning the optimal solution from i to j is simply the distance in the
adjacency matrix.
	dp[k][i][j] = m[i][j] if k = 0
otherwise:
	dp[k][i][j] = min(dp[k-1][i][j], dp[k-1][i][k]+dp[k-1][k][j])

min(
	Reuse the best distance from i to j with values routing through nodes
	{0, 1, ..., k-1}
	,
	Find the best distance from i to j through node k reusing best solutions
	from {0, 1, ..., k-1}.
)

The right side of the min function in english essentially says: "go from i to k"
and the "go from k to j".

Currently we're using O(V^3) memory since our memo table 'dp' has one dimension
for each of k, i, j.

Notice that we will be looping over k starting from 0, then 1, 2, ... and so
fourth. The important thing to note here is that previous result builds off the
last since we need state k-1 to compute state k. With that being said, it is
possible to compute the solution for k in-place saving us a dimension of memory
and reducing the space complexity to O(V^2)!

The new recurrence relation is:
	dp[i][j] = m[i][j] if k = 0
otherwise:
	dp[i][j] = min(dp[i][j], dp[i][k] + dp[k][j])

```
	# Global/class scope variables
	n = size of the adjacency matrix
	dp = the memo table that will contain APSP solution
	next = matrix used to reconstruct shortest paths

	function floydWarshall(m):
		setup(m)

		# Execute FW all pairs shortest path algorithm.
		for (k := 0; k < n; k++):
			for (j := 0; j < n; i++):
				for (j := 0; j < n; j++):
					if (dp[i][k] + dp[k][j] < dp[i][j]):
						dp[i][j] = d[i][k] + dp[k][j]
						next[i][j] = next[i][k]
		# Detect and propagate negative cycles.
		propagateNegativeCycles(dp, n) # Optional

		# Return APSP matrix
		return dp

	function setup(m):
		dp = empty matrix of size n x n

		# Should contain null values by default
		next = empty integer matrix of size n x n

		# Do a deep copy of the input matrix and setup
		# the 'next' matrix for path reconstruction.
		for (i := 0; i < n; i++):
			for (j := 0; j < n; j++):
				dp[i][j] = m[i][j]
				if m[i][j] != +inf:
					next[i][j] = j

	function propagateNegativeCycles(dp, n):
		# Execute FW APSP algorithm a second time but
		# this time if the distance can be improved
		# set the optimal distance to be -inf.
		# Every edge (i, j) marked with -inf is either
		# part of or reaches into a negative cycle.
		for (k := 0; k < n; k++):
			for (i := 0; i < n; i++):
				for (j := 0; j < n; j++):
					if (dp[i][k] + dp[k][j] < d[i][j]):
						dp[i][j] = -inf
						next[i][j] =-1

	# Reconstructs the shortest path between nodes
	# 'start' and 'end'. You must run the
	# floydWarshall solver before calling this method.
	# Returns null if path affected by negative cycle.
	function reconstructPath(start, end):
		path = []
		# Check if there exists a path between
		# the start and the end node.
		if dp[start][end] == +inf: return path

		at := start
		# Reconstruct path from next matrix
		for (; at != -1; at = next[at][end]):
			if at == -1: return null
			path.add(at)

		if next[at][end] == -1: return null
		path.add(end)
		return path
```

Negative Cycles

What do we mean by a negative cycle?

Continue from 2:29:22

*/

type ShortestPath interface {
	DAG(adjList map[int][]*Edge, s int, n int) []int
	Dijkstra(adjList map[int][]*Edge, s int, n int) []int
	BellmanFord(adjList map[int][]*Edge, s int, n int) []int
	FloydWarshall(adjList map[int][]*Edge, s int, n int) []int
}

/*
#######################################################################################################################

Shortest Path (SP) Algorithms
|------------------+------------+--------------+--------------+----------------|
|                  |    BFS     |  Dijkstra's  | Bellman-Ford | Floyd-Warshall |
|------------------+------------+--------------+--------------+----------------|
| Complexity       |   O(V+E)   | O((V+E)logV) |    O(VE)     |     O(V^3)     |
|------------------+------------+--------------+--------------+----------------|
| Recommended      |   Large    | Large/Medium | Medium/Small |     Small      |
| graph size       |            |              |              |                |
|------------------+------------+--------------+--------------+----------------|
| Good for         |  Only for  |              |              |                |
| APSP?            | unweighted |      Ok      |     Bad      |      Yes       |
|                  |   graphs   |              |              |                |
|------------------+------------+--------------+--------------+----------------|
| Can detect       |     No     |      No      |     Yes      |      Yes       |
| negative cycles? |            |              |              |                |
|------------------+------------+--------------+--------------+----------------|
| SP on graph with | Incorrect  |     Best     |    Works     | Bad in general |
| weighted edges   | SP answer  |  algorithm   |              |                |
|------------------+------------+--------------+--------------+----------------|
| SP on graph with |    Best    |      Ok      |     Bad      | Bad in general |
| unweighted edges | algorithm  |              |              |                |
|------------------+------------+--------------+--------------+----------------|

Reference: Competitive Programming 3, P. 161, Steven & Felix Halim

#######################################################################################################################

                      | Dijkstra | Bellman-Ford | Floyd-Warshall |
Single Source         | SSSP*    | SSSP*        | APSP*          |
Cyclic                | yes      | yes          | yes            |
Negative Edge Weights | no       | yes          | yes            |
Negative Cycles*      | no       | no           | no             |

*Negative Cycles: where the sum of the edges in a cycle is negative
*SSSP: Single Source Shortest Path
*APSP: All Pairs Shortest Path

Dijkstra's algorithm solves the shortest-path problem for any weighted,
directed graph with non-negative weights.

It can handle graphs consisting of cycles, but negative weights will cause
this algorithm to produce incorrect results.

Since Dijkstra follows a Greedy Approach, once a node is marked as visited
it cannot be reconsidered even if there is another path with less cost or
distance. This issue arises only if there exists a negative weight or edge in
the graph.

The Bellman-Ford algorithm is a way to find single source shortest paths
in a graph with negative edge weights (but no negative cycles).
The second for loop in this algorithm also detects negative cycles.
The first for loop relaxes each of the edges in the graph n − 1 times.

What is the best shortest path algorithm for unweighted graph?

The most important fact we should know is that BFS traversal can be used to
find the shortest path in an unweighted graph in O(|V| + |E|) time.

Unlike Dijkstra's algorithm, Bellman-Ford can handle graphs with negative
weight edges.

*/

type Edge struct {
	To     int
	Weight int
}

type graph struct{}

var (
	_ Graph        = &graph{}
	_ ShortestPath = &graph{}
)

func (g *graph) DFS(adjList [][]int) []int {
	var dfs func(curr int, visited map[int]bool) []int

	dfs = func(curr int, visited map[int]bool) []int {
		nodes := []int{curr}

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
	nodes := make([]int, 0)
	queue := pkg.NewQueue[int]()
	queue.Enqueue(0)

	visited := make(map[int]bool)

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

func (g *graph) TopologicalSortBFS(adjList map[int][]*Edge, n int) []int {
	indegrees := make(map[int]int)

	for _, dsts := range adjList {
		for _, dst := range dsts {
			indegrees[dst.To]++
		}
	}

	visited := make(map[int]bool)
	queue := pkg.NewQueue[int]()

	for src := 0; src < n; src++ {
		if indegrees[src] == 0 && !visited[src] {
			queue.Enqueue(src)

			visited[src] = true
		}
	}

	sorted := make([]int, 0)

	var node int
	for !queue.Empty() {
		node = queue.Dequeue()
		sorted = append(sorted, node)

		for _, dst := range adjList[node] {
			if !visited[dst.To] {
				indegrees[dst.To]--
				if indegrees[dst.To] == 0 {
					queue.Enqueue(dst.To)

					visited[dst.To] = true
				}
			}
		}
	}

	return sorted
}

func (g *graph) TopologicalSortDFS(adjList map[int][]*Edge, n int) []int {
	sorted := make([]int, n)

	var dfs func(node int, visited map[int]bool, idx int, sorted []int) int

	dfs = func(node int, visited map[int]bool, idx int, sorted []int) int {
		visited[node] = true
		for _, next := range adjList[node] {
			if !visited[next.To] {
				idx = dfs(next.To, visited, idx, sorted)
			}
		}

		sorted[idx] = node

		return idx - 1
	}
	visited := make(map[int]bool)
	idx := n - 1

	for src := 0; src < n; src++ {
		if !visited[src] {
			visited[src] = true
			idx = dfs(src, visited, idx, sorted)
		}
	}

	return sorted
}

// Approach: BFS
func (g *graph) Dungeon(grid [][]byte) int {
	directions := [5]int{-1, 0, 1, 0, -1}
	m, n := len(grid), len(grid[0])
	visited := make([][]bool, m)

	for idx := range visited {
		visited[idx] = make([]bool, n)
	}

	queue := pkg.NewQueue[[3]int]()
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
		current := queue.Dequeue()

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

func (g *graph) DAG(adjList map[int][]*Edge, s int, n int) []int {
	topsort := g.TopologicalSortBFS(adjList, n)
	dist := make([]int, n)
	hasValue := make([]bool, n)
	dist[s] = 0
	hasValue[s] = true

	for _, node := range topsort {
		if hasValue[node] {
			adjacentEdges := adjList[node]
			for _, edge := range adjacentEdges {
				newDist := dist[node] + edge.Weight
				if !hasValue[edge.To] || newDist < dist[edge.To] {
					dist[edge.To] = newDist
					hasValue[edge.To] = true
				}
			}
		}
	}

	return dist
}

// Lazy Dijkstra's Algorithm
func (g *graph) Dijkstra(adjList map[int][]*Edge, s int, n int) []int {
	visited := make([]bool, n)

	dist := make([]int, n)
	for idx := range dist {
		dist[idx] = math.MaxInt
	}

	dist[s] = 0
	minPQ := pkg.NewHeap(make([]*Edge, 0), func(x, y *Edge) bool { return x.Weight < y.Weight })
	heap.Push(minPQ, &Edge{s, 0})

	for minPQ.Len() > 0 {
		nodeID := heap.Pop(minPQ).(*Edge).To // get the next minimal (promising) distance
		visited[nodeID] = true

		for _, edge := range adjList[nodeID] {
			if visited[edge.To] {
				continue
			}

			newDist := dist[nodeID] + edge.Weight
			if newDist < dist[edge.To] {
				dist[edge.To] = newDist
				heap.Push(minPQ, &Edge{edge.To, newDist})
			}
		}
	}

	return dist
}

func (g *graph) BellmanFord(adjList map[int][]*Edge, s int, n int) []int {
	// step 1: initialization
	dist := make([]int, n)
	for idx := range dist {
		dist[idx] = math.MaxInt
	}

	dist[s] = 0
	// step 2: relax edges
	for i := 0; i < n-1; i++ {
		for src, edges := range adjList {
			for _, edge := range edges {
				if dist[src] != math.MaxInt {
					dist[edge.To] = min(dist[edge.To], dist[src]+edge.Weight)
				}
			}
		}
	}
	// step 3: check for negative-weight cycles
	for i := 0; i < n-1; i++ {
		for src, edges := range adjList {
			if dist[src] == math.MaxInt {
				continue
			}

			for _, edge := range edges {
				if dist[edge.To] > dist[src]+edge.Weight {
					// Graph contains negative weight cycle
					dist[edge.To] = math.MaxInt
				}
			}
		}
	}

	return dist
}

func (g *graph) FloydWarshall(adjList map[int][]*Edge, s int, n int) []int {
	_, _, _ = adjList, s, n
	return nil
}
