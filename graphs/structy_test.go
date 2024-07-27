package graphs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./graphs/ -run ^TestStructyHasPath$
func TestStructyHasPath(t *testing.T) {
	for _, tc := range []struct {
		adjList map[byte][]byte
		src     byte
		dst     byte
		hasPath bool
	}{
		{
			adjList: map[byte][]byte{
				'f': {'g', 'i'},
				'g': {'h'},
				'h': {},
				'i': {'g', 'k'},
				'j': {'i'},
				'k': {},
			},
			src:     'f',
			dst:     'k',
			hasPath: true,
		},
		{
			adjList: map[byte][]byte{
				'f': {'g', 'i'},
				'g': {'h'},
				'h': {},
				'i': {'g', 'k'},
				'j': {'i'},
				'k': {},
			},
			src:     'f',
			dst:     'j',
			hasPath: false,
		},
		{
			adjList: map[byte][]byte{
				'f': {'g', 'i'},
				'g': {'h'},
				'h': {},
				'i': {'g', 'k'},
				'j': {'i'},
				'k': {},
			},
			src:     'i',
			dst:     'h',
			hasPath: true,
		},
		{
			adjList: map[byte][]byte{
				'v': {'x', 'w'},
				'w': {},
				'x': {},
				'y': {'z'},
				'z': {},
			},
			src:     'v',
			dst:     'w',
			hasPath: true,
		},
		{
			adjList: map[byte][]byte{
				'v': {'x', 'w'},
				'w': {},
				'x': {},
				'y': {'z'},
				'z': {},
			},
			src:     'v',
			dst:     'z',
			hasPath: false,
		},
	} {
		var structy Structy = &structy{}
		hasPath := structy.HasPath(tc.adjList, tc.src, tc.dst)
		assert.Equal(t, tc.hasPath, hasPath)
	}
}

// go test -v -count=1 ./graphs/ -run ^TestStructyUndirectedPath$
func TestStructyUndirectedPath(t *testing.T) {
	for _, tc := range []struct {
		edges   [][]byte
		nodeA   byte
		nodeB   byte
		hasPath bool
	}{
		{
			edges: [][]byte{
				{'i', 'j'},
				{'k', 'i'},
				{'m', 'k'},
				{'k', 'l'},
				{'o', 'n'},
			},
			nodeA:   'j',
			nodeB:   'm',
			hasPath: true,
		},
		{
			edges: [][]byte{
				{'i', 'j'},
				{'k', 'i'},
				{'m', 'k'},
				{'k', 'l'},
				{'o', 'n'},
			},
			nodeA:   'm',
			nodeB:   'j',
			hasPath: true,
		},
		{
			edges: [][]byte{
				{'i', 'j'},
				{'k', 'i'},
				{'m', 'k'},
				{'k', 'l'},
				{'o', 'n'},
			},
			nodeA:   'l',
			nodeB:   'j',
			hasPath: true,
		},
		{
			edges: [][]byte{
				{'i', 'j'},
				{'k', 'i'},
				{'m', 'k'},
				{'k', 'l'},
				{'o', 'n'},
			},
			nodeA:   'k',
			nodeB:   'o',
			hasPath: false,
		},
		{
			edges: [][]byte{
				{'i', 'j'},
				{'k', 'i'},
				{'m', 'k'},
				{'k', 'l'},
				{'o', 'n'},
			},
			nodeA:   'i',
			nodeB:   'o',
			hasPath: false,
		},
		{
			edges: [][]byte{
				{'b', 'a'},
				{'c', 'a'},
				{'b', 'c'},
				{'q', 'r'},
				{'q', 's'},
				{'q', 'u'},
				{'q', 't'},
			},
			nodeA:   'a',
			nodeB:   'b',
			hasPath: true,
		},
		{
			edges: [][]byte{
				{'b', 'a'},
				{'c', 'a'},
				{'b', 'c'},
				{'q', 'r'},
				{'q', 's'},
				{'q', 'u'},
				{'q', 't'},
			},
			nodeA:   'a',
			nodeB:   'c',
			hasPath: true,
		},
		{
			edges: [][]byte{
				{'b', 'a'},
				{'c', 'a'},
				{'b', 'c'},
				{'q', 'r'},
				{'q', 's'},
				{'q', 'u'},
				{'q', 't'},
			},
			nodeA:   'r',
			nodeB:   't',
			hasPath: true,
		},
		{
			edges: [][]byte{
				{'b', 'a'},
				{'c', 'a'},
				{'b', 'c'},
				{'q', 'r'},
				{'q', 's'},
				{'q', 'u'},
				{'q', 't'},
			},
			nodeA:   'r',
			nodeB:   'b',
			hasPath: false,
		},
		{
			edges: [][]byte{
				{'s', 'r'},
				{'t', 'q'},
				{'q', 'r'},
			},
			nodeA:   'r',
			nodeB:   't',
			hasPath: true,
		},
	} {
		var structy Structy = &structy{}
		hasPath := structy.UndirectedPath(tc.edges, tc.nodeA, tc.nodeB)
		assert.Equal(t, tc.hasPath, hasPath)
	}
}

// go test -v -count=1 ./graphs/ -run ^TestStructyConnectedComponentsCount$
func TestStructyConnectedComponentsCount(t *testing.T) {
	for _, tc := range []struct {
		adjList map[int][]int
		count   int
	}{
		{
			adjList: map[int][]int{
				0: {8, 1, 5},
				1: {0},
				5: {0, 8},
				8: {0, 5},
				2: {3, 4},
				3: {2, 4},
				4: {3, 2},
			},
			count: 2,
		},
		{
			adjList: map[int][]int{
				1: {2},
				2: {1, 8},
				6: {7},
				9: {8},
				7: {6, 8},
				8: {9, 7, 2},
			},
			count: 1,
		},
		{
			adjList: map[int][]int{
				3: {},
				4: {6},
				6: {4, 5, 7, 8},
				8: {6},
				7: {6},
				5: {6},
				1: {2},
				2: {1},
			},
			count: 3,
		},
		{
			adjList: map[int][]int{},
			count:   0,
		},
		{
			adjList: map[int][]int{
				0: {4, 7},
				1: {},
				2: {},
				3: {6},
				4: {0},
				6: {3},
				7: {0},
				8: {},
			},
			count: 5,
		},
	} {
		var structy Structy = &structy{}
		count := structy.ConnectedComponentsCount(tc.adjList)
		assert.Equal(t, tc.count, count)
	}
}

// go test -v -count=1 ./graphs/ -run ^TestStructyLargestComponent$
func TestStructyLargestComponent(t *testing.T) {
	for _, tc := range []struct {
		adjList map[int][]int
		maxsize int
	}{
		{
			adjList: map[int][]int{
				0: {8, 1, 5},
				1: {0},
				5: {0, 8},
				8: {0, 5},
				2: {3, 4},
				3: {2, 4},
				4: {3, 2},
			},
			maxsize: 4,
		},
		{
			adjList: map[int][]int{
				1: {2},
				2: {1, 8},
				6: {7},
				9: {8},
				7: {6, 8},
				8: {9, 7, 2},
			},
			maxsize: 6,
		},
		{
			adjList: map[int][]int{
				3: {},
				4: {6},
				6: {4, 5, 7, 8},
				8: {6},
				7: {6},
				5: {6},
				1: {2},
				2: {1},
			},
			maxsize: 5,
		},
		{
			adjList: map[int][]int{},
			maxsize: 0,
		},
		{
			adjList: map[int][]int{
				0: {4, 7},
				1: {},
				2: {},
				3: {6},
				4: {0},
				6: {3},
				7: {0},
				8: {},
			},
			maxsize: 3,
		},
	} {
		var structy Structy = &structy{}
		maxsize := structy.LargestComponent(tc.adjList)
		assert.Equal(t, tc.maxsize, maxsize)
	}
}

// go test -v -count=1 ./graphs/ -run ^TestStructyShortestPath$
func TestStructyShortestPath(t *testing.T) {
	for _, tc := range []struct {
		edges   [][]byte
		nodeA   byte
		nodeB   byte
		minPath int
	}{
		{
			edges: [][]byte{
				{'w', 'x'},
				{'x', 'y'},
				{'z', 'y'},
				{'z', 'v'},
				{'w', 'v'},
			},
			nodeA:   'w',
			nodeB:   'z',
			minPath: 2,
		},
		{
			edges: [][]byte{
				{'w', 'x'},
				{'x', 'y'},
				{'z', 'y'},
				{'z', 'v'},
				{'w', 'v'},
			},
			nodeA:   'y',
			nodeB:   'x',
			minPath: 1,
		},
		{
			edges: [][]byte{
				{'a', 'c'},
				{'a', 'b'},
				{'c', 'b'},
				{'c', 'd'},
				{'b', 'd'},
				{'e', 'd'},
				{'g', 'f'},
			},
			nodeA:   'a',
			nodeB:   'e',
			minPath: 3,
		},
		{
			edges: [][]byte{
				{'a', 'c'},
				{'a', 'b'},
				{'c', 'b'},
				{'c', 'd'},
				{'b', 'd'},
				{'e', 'd'},
				{'g', 'f'},
			},
			nodeA:   'e',
			nodeB:   'c',
			minPath: 2,
		},
		{
			edges: [][]byte{
				{'a', 'c'},
				{'a', 'b'},
				{'c', 'b'},
				{'c', 'd'},
				{'b', 'd'},
				{'e', 'd'},
				{'g', 'f'},
			},
			nodeA:   'b',
			nodeB:   'g',
			minPath: -1,
		},
		{
			edges: [][]byte{
				{'c', 'n'},
				{'c', 'e'},
				{'c', 's'},
				{'c', 'w'},
				{'w', 'e'},
			},
			nodeA:   'w',
			nodeB:   'e',
			minPath: 1,
		},
		{
			edges: [][]byte{
				{'c', 'n'},
				{'c', 'e'},
				{'c', 's'},
				{'c', 'w'},
				{'w', 'e'},
			},
			nodeA:   'n',
			nodeB:   'e',
			minPath: 2,
		},
		{
			edges: [][]byte{
				{'m', 'n'},
				{'n', 'o'},
				{'o', 'p'},
				{'p', 'q'},
				{'t', 'o'},
				{'r', 'q'},
				{'r', 's'},
			},
			nodeA:   'm',
			nodeB:   's',
			minPath: 6,
		},
	} {
		var structy Structy = &structy{}
		minPath := structy.ShortestPath(tc.edges, tc.nodeA, tc.nodeB)
		assert.Equal(t, tc.minPath, minPath)
	}
}

// go test -v -count=1 ./graphs/ -run ^TestStructyIslandCount$
func TestStructyIslandCount(t *testing.T) {
	for _, tc := range []struct {
		grid    [][]byte
		islands int
	}{
		{
			grid: [][]byte{
				{'W', 'L', 'W', 'W', 'W'},
				{'W', 'L', 'W', 'W', 'W'},
				{'W', 'W', 'W', 'L', 'W'},
				{'W', 'W', 'L', 'L', 'W'},
				{'L', 'W', 'W', 'L', 'L'},
				{'L', 'L', 'W', 'W', 'W'},
			},
			islands: 3,
		},
		{
			grid: [][]byte{
				{'L', 'W', 'W', 'L', 'W'},
				{'L', 'W', 'W', 'L', 'L'},
				{'W', 'L', 'W', 'L', 'W'},
				{'W', 'W', 'W', 'W', 'W'},
				{'W', 'W', 'L', 'L', 'L'},
			},
			islands: 4,
		},
		{
			grid: [][]byte{
				{'L', 'L', 'L'},
				{'L', 'L', 'L'},
				{'L', 'L', 'L'},
			},
			islands: 1,
		},
		{
			grid: [][]byte{
				{'W', 'W'},
				{'W', 'W'},
				{'W', 'W'},
			},
			islands: 0,
		},
	} {
		var structy Structy = &structy{}
		islands := structy.IslandCount(tc.grid)
		assert.Equal(t, tc.islands, islands)
	}
}

// go test -v -count=1 ./graphs/ -run ^TestStructyMinimumIsland$
func TestStructyMinimumIsland(t *testing.T) {
	for _, tc := range []struct {
		grid      [][]byte
		minIsland int
	}{
		{
			grid: [][]byte{
				{'W', 'L', 'W', 'W', 'W'},
				{'W', 'L', 'W', 'W', 'W'},
				{'W', 'W', 'W', 'L', 'W'},
				{'W', 'W', 'L', 'L', 'W'},
				{'L', 'W', 'W', 'L', 'L'},
				{'L', 'L', 'W', 'W', 'W'},
			},
			minIsland: 2,
		},
		{
			grid: [][]byte{
				{'L', 'W', 'W', 'L', 'W'},
				{'L', 'W', 'W', 'L', 'L'},
				{'W', 'L', 'W', 'L', 'W'},
				{'W', 'W', 'W', 'W', 'W'},
				{'W', 'W', 'L', 'L', 'L'},
			},
			minIsland: 1,
		},
		{
			grid: [][]byte{
				{'L', 'L', 'L'},
				{'L', 'L', 'L'},
				{'L', 'L', 'L'},
			},
			minIsland: 9,
		},
		{
			grid: [][]byte{
				{'W', 'W'},
				{'L', 'L'},
				{'W', 'W'},
				{'W', 'L'},
			},
			minIsland: 1,
		},
	} {
		var structy Structy = &structy{}
		minIsland := structy.MinimumIsland(tc.grid)
		assert.Equal(t, tc.minIsland, minIsland)
	}
}
