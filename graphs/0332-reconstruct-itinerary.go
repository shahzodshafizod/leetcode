package graphs

import (
	"container/heap"
)

/*
The initial vertex v must have indeg(v)+1=outdeg(v) or indeg(v)=outdeg(v) (if an Euler cycle)
The terminal vertex w has indeg(w)=outdeg(w)+1 or indeg(w)=outdeg(w) (if an Euler cycle v=w)
For other vertex vi has indeg(vi)=outdeg(vi).
*/

// https://leetcode.com/problems/reconstruct-itinerary/

type itineraryHeap []string

var _ heap.Interface = &itineraryHeap{}

func (i *itineraryHeap) Len() int { return len(*i) }
func (i *itineraryHeap) Less(x, y int) bool {
	var str1, str2 = (*i)[x], (*i)[y]
	len1, len2 := len(str1), len(str2)
	for idx := 0; idx < len1 && idx < len2; idx++ {
		if str1[idx] == str2[idx] {
			continue
		}
		return str1[idx] < str2[idx]
	}
	return len1 < len2
}
func (i *itineraryHeap) Swap(x, y int) { (*i)[x], (*i)[y] = (*i)[y], (*i)[x] }
func (i *itineraryHeap) Push(x any)    { *i = append(*i, x.(string)) }
func (i *itineraryHeap) Pop() any      { last := (*i)[i.Len()-1]; *i = (*i)[:i.Len()-1]; return last }

func findItinerary(tickets [][]string) []string {
	var adjList = make(map[string]*itineraryHeap)
	for _, ticket := range tickets {
		var src, dst = ticket[0], ticket[1]
		if adjList[src] == nil {
			adjList[src] = &itineraryHeap{}
		}
		heap.Push(adjList[src], dst)
	}
	var length = len(tickets) + 1
	var result = make([]string, length)
	findItineraryDFS(adjList, "JFK", &result, &length)
	return result
}

func findItineraryDFS(adjList map[string]*itineraryHeap, source string, result *[]string, index *int) {
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
