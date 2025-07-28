package learning

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./learning/ -run ^TestGraphBFS$
func TestGraphBFS(t *testing.T) {
	for _, tc := range []struct {
		adjList [][]int
		values  []int
	}{
		{
			adjList: [][]int{{1, 3}, {0}, {3, 8}, {4, 5, 2, 0}, {6, 3}, {3}, {7, 4}, {6}, {2}},
			values:  []int{0, 1, 3, 4, 5, 2, 6, 8, 7},
		},
		{
			adjList: [][]int{{1, 2}, {3}, {4}, {5}, {1}, {}},
			values:  []int{0, 1, 2, 3, 4, 5},
		},
	} {
		var graph Graph = &graph{}

		values := graph.BFS(tc.adjList)
		assert.Equal(t, tc.values, values)
	}
}

// go test -v -count=1 ./learning/ -run ^TestGraphDFS$
func TestGraphDFS(t *testing.T) {
	for _, tc := range []struct {
		adjList [][]int
		values  []int
	}{
		{
			adjList: [][]int{{1, 3}, {0}, {3, 8}, {4, 5, 2, 0}, {6, 3}, {3}, {7, 4}, {6}, {2}},
			values:  []int{0, 1, 3, 4, 6, 7, 5, 2, 8},
		},
		{
			adjList: [][]int{{1, 2}, {3}, {4}, {5}, {1}, {}},
			values:  []int{0, 1, 3, 5, 2, 4},
		},
	} {
		var graph Graph = &graph{}

		values := graph.DFS(tc.adjList)
		assert.Equal(t, tc.values, values)
	}
}

// go test -v -count=1 ./learning/ -run ^TestTopologicalSort$
func TestTopologicalSort(t *testing.T) {
	for _, tc := range []struct {
		adjList   map[int][]*Edge
		n         int
		sortedBFS []int
		sortedDFS []int
	}{
		{
			adjList: map[int][]*Edge{
				0:  {{To: 3}},
				1:  {{To: 3}},
				2:  {{To: 0}, {To: 1}},
				3:  {{To: 6}, {To: 7}},
				4:  {{To: 0}, {To: 3}, {To: 5}},
				5:  {{To: 9}, {To: 10}},
				6:  {{To: 8}},
				7:  {{To: 8}, {To: 9}},
				8:  {{To: 11}},
				9:  {{To: 11}, {To: 12}},
				10: {{To: 9}},
			},
			n:         13,
			sortedBFS: []int{2, 4, 1, 0, 5, 3, 10, 6, 7, 8, 9, 11, 12},
			sortedDFS: []int{4, 5, 10, 2, 1, 0, 3, 7, 9, 12, 6, 8, 11},
		},
	} {
		var graph Graph = &graph{}

		sorted := graph.TopologicalSortBFS(tc.adjList, tc.n)
		assert.Equal(t, tc.sortedBFS, sorted)

		sorted = graph.TopologicalSortDFS(tc.adjList, tc.n)
		assert.Equal(t, tc.sortedDFS, sorted)
	}
}

// go test -v -count=1 ./learning/ -run ^TestDungeon$
func TestDungeon(t *testing.T) {
	for _, tc := range []struct {
		grid [][]byte
		path int
	}{
		{
			grid: [][]byte{
				{'S', '.', '.', '#', '.', '.', '.'},
				{'.', '#', '.', '.', '.', '#', '.'},
				{'.', '#', '.', '.', '.', '.', '.'},
				{'.', '.', '#', '#', '.', '.', '.'},
				{'#', '.', '#', 'E', '.', '#', '.'},
			},
			path: 10,
		},
	} {
		var graph Graph = &graph{}

		path := graph.Dungeon(tc.grid)
		assert.Equal(t, tc.path, path)
	}
}

// go test -v -count=1 ./learning/ -run ^TestShortestPath$
func TestShortestPath(t *testing.T) {
	for _, tc := range []struct {
		cyclic         bool
		negativeEdges  bool
		negativeCycles bool
		adjList        map[int][]*Edge
		s              int
		n              int
		dist           []int
	}{
		{
			cyclic:         false,
			negativeEdges:  false,
			negativeCycles: false,
			adjList: map[int][]*Edge{
				0: {{To: 1, Weight: 4}, {To: 2, Weight: 1}},
				1: {{To: 3, Weight: 1}},
				2: {{To: 1, Weight: 2}, {To: 3, Weight: 5}},
				3: {{To: 4, Weight: 3}},
			},
			s:    0,
			n:    5,
			dist: []int{0, 3, 1, 4, 7},
		},
		{
			cyclic:         false,
			negativeEdges:  true,
			negativeCycles: false,
			adjList: map[int][]*Edge{
				0: {{To: 1, Weight: 3}, {To: 2, Weight: 6}},
				1: {{To: 2, Weight: 4}, {To: 3, Weight: 4}, {To: 4, Weight: 11}},
				2: {{To: 3, Weight: 8}, {To: 6, Weight: 11}},
				3: {{To: 4, Weight: -4}, {To: 5, Weight: 5}, {To: 6, Weight: 2}},
				4: {{To: 7, Weight: 9}},
				5: {{To: 7, Weight: 1}},
				6: {{To: 7, Weight: 2}},
			},
			s:    0,
			n:    8,
			dist: []int{0, 3, 6, 7, 3, 12, 9, 11},
		},
		{
			cyclic:         true,
			negativeEdges:  false,
			negativeCycles: false,
			adjList: map[int][]*Edge{
				0: {{To: 1, Weight: 5}, {To: 2, Weight: 1}},
				1: {{To: 2, Weight: 2}, {To: 3, Weight: 3}, {To: 4, Weight: 20}},
				2: {{To: 1, Weight: 3}, {To: 4, Weight: 12}},
				3: {{To: 2, Weight: 3}, {To: 4, Weight: 2}, {To: 5, Weight: 6}},
				4: {{To: 5, Weight: 1}},
			},
			s:    0,
			n:    6,
			dist: []int{0, 4, 1, 7, 9, 10},
		},
		{
			cyclic:         true,
			negativeEdges:  true,
			negativeCycles: true,
			adjList: map[int][]*Edge{
				0: {{To: 1, Weight: 5}},
				1: {{To: 2, Weight: 20}, {To: 5, Weight: 30}, {To: 6, Weight: 60}},
				2: {{To: 3, Weight: 10}, {To: 4, Weight: 75}},
				3: {{To: 2, Weight: -15}},
				4: {{To: 9, Weight: 100}},
				5: {{To: 4, Weight: 25}, {To: 6, Weight: 5}, {To: 8, Weight: 50}},
				6: {{To: 7, Weight: -50}},
				7: {{To: 8, Weight: -10}},
			},
			s:    0,
			n:    10,
			dist: []int{},
		},
	} {
		var graph ShortestPath = &graph{}

		var dist []int

		if !tc.cyclic && !tc.negativeEdges && !tc.negativeCycles {
			dist = graph.DAG(tc.adjList, tc.s, tc.n)
			assert.Equal(t, tc.dist, dist)
		}

		if !tc.negativeEdges && !tc.negativeCycles {
			dist = graph.Dijkstra(tc.adjList, tc.s, tc.n)
			assert.Equal(t, tc.dist, dist)
		}

		if !tc.negativeCycles {
			dist = graph.BellmanFord(tc.adjList, tc.s, tc.n)
			assert.Equal(t, tc.dist, dist)
		}
	}
}
