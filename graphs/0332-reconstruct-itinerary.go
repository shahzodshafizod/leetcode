package graphs

import (
	"container/heap"

	"github.com/shahzodshafizod/leetcode/design"
)

/*
The initial vertex v must have indeg(v)+1=outdeg(v) or indeg(v)=outdeg(v) (if an Euler cycle)
The terminal vertex w has indeg(w)=outdeg(w)+1 or indeg(w)=outdeg(w) (if an Euler cycle v=w)
For other vertex vi has indeg(vi)=outdeg(vi).
*/

// https://leetcode.com/problems/reconstruct-itinerary/

func findItinerary(tickets [][]string) []string {
	var adjList = make(map[string]*design.Heap[string])
	for _, ticket := range tickets {
		var src, dst = ticket[0], ticket[1]
		if adjList[src] == nil {
			adjList[src] = design.NewHeap(make([]string, 0), func(x, y string) bool { return x < y })
		}
		heap.Push(adjList[src], dst)
	}
	var length = len(tickets) + 1
	var result = make([]string, length)
	findItineraryDFS(adjList, "JFK", &result, &length)
	return result
}

func findItineraryDFS(adjList map[string]*design.Heap[string], source string, result *[]string, index *int) {
	for adjList[source] != nil && adjList[source].Len() > 0 {
		next := heap.Pop(adjList[source]).(string)
		findItineraryDFS(adjList, next, result, index)
	}
	(*index)--
	(*result)[*index] = source
}

// func findItinerary(tickets [][]string) []string {
// 	var same = make(map[string]bool)
// 	var adjList = make(map[string][]string)
// 	for _, ticket := range tickets {
// 		var src, dst = ticket[0], ticket[1]
// 		adjList[src] = append(adjList[src], dst)
// 		if _, exists := same[src]; !exists {
// 			same[src] = true
// 		} else if same[src] {
// 			if len := len(adjList[src]); len > 1 && adjList[src][len-2] != adjList[src][len-1] {
// 				same[src] = false
// 			}
// 		}
// 	}
// 	for src := range adjList {
// 		sort.Strings(adjList[src])
// 	}
// 	var result = make([]string, len(tickets)+1)
// 	result[0] = "JFK"
// 	backtracking(adjList, same, "JFK", &result, 1)
// 	return result
// }

// func backtracking(adjList map[string][]string, same map[string]bool, source string, result *[]string, index int) bool {
// 	if len(adjList[source]) == 0 {
// 		for _, dsts := range adjList {
// 			if len(dsts) > 0 {
// 				return false
// 			}
// 		}
// 		return true
// 	}
// 	for i := 0; i < len(adjList[source]); i++ {
// 		value := adjList[source][i]
// 		(*result)[index] = value
// 		adjList[source] = append(adjList[source][:i], adjList[source][i+1:]...)
// 		if backtracking(adjList, same, value, result, index+1) {
// 			return true
// 		}
// 		adjList[source] = append(adjList[source][:i+1], adjList[source][i:]...)
// 		adjList[source][i] = value
// 		if same[source] {
// 			return false
// 		}
// 	}
// 	return false
// }
