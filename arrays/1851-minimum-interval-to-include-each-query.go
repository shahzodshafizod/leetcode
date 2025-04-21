package arrays

import (
	"container/heap"
	"sort"

	"github.com/shahzodshafizod/leetcode/pkg"
)

// https://leetcode.com/problems/minimum-interval-to-include-each-query/

// Approach: Heap (Priority Queue)
// Time: O(NLogN + QLogQ), N=len(intervals), Q=len(queries)
// Space: O(N+Q)
func minInterval(intervals [][]int, queries []int) []int {
	sort.Slice(
		intervals,
		func(i, j int) bool {
			return intervals[i][0] < intervals[j][0]
		},
	)

	qlen := len(queries)
	squeries := make([]int, qlen)
	copy(squeries, queries)
	sort.Ints(squeries)

	type interval struct {
		length int
		right  int
	}
	inbound := pkg.NewHeap(
		make([]*interval, 0),
		func(x, y *interval) bool { return x.length < y.length },
	)

	qmap := make(map[int]int)
	var iid, ilen = 0, len(intervals)
	var left, right int
	for _, query := range squeries {
		for iid < ilen && intervals[iid][0] <= query {
			left = intervals[iid][0]
			right = intervals[iid][1]
			heap.Push(inbound, &interval{right - left + 1, right})
			iid++
		}
		for inbound.Len() > 0 && inbound.Peak().right < query {
			heap.Pop(inbound)
		}
		var value = -1
		if inbound.Len() > 0 {
			value = inbound.Peak().length
		}
		qmap[query] = value
	}
	answer := make([]int, qlen)
	for idx, query := range queries {
		answer[idx] = qmap[query]
	}
	return answer
}

// // Approach: Line Sweep + Monotonic Increasing Stack (Time Limit Exceeded)
// // Time: O((N+Q)Log(N+Q)), N=len(intervals), Q=len(queries)
// // Space: O(N+Q)
// func minInterval(intervals [][]int, queries []int) []int {
// 	const (
// 		INTERVAL = iota + 1
// 		QUERY
// 	)
// 	type point struct {
// 		first  int // INTERVAL(left)  QUERY(query) STACK(length)
// 		second int // INTERVAL(right) QUERY(index) STACK(right)
// 		kind   int // INTERVAL=1, QUERY=2
// 	}
// 	var points = design.NewHeap(
// 		make([]*point, 0),
// 		func(x, y *point) bool {
// 			if x.first == y.first {
// 				return x.kind < y.kind
// 			}
// 			return x.first < y.first
// 		},
// 	)
// 	var left, right int
// 	for _, interval := range intervals {
// 		left, right = interval[0], interval[1]
// 		heap.Push(points, &point{left, right, INTERVAL})
// 	}
// 	for idx, query := range queries {
// 		heap.Push(points, &point{query, idx, QUERY})
// 	}
// 	type node struct {
// 		point *point
// 		next  *node
// 	}
// 	var stack *node = nil
// 	var length int
// 	for points.Len() > 0 {
// 		p := heap.Pop(points).(*point)
// 		switch p.kind {
// 		case INTERVAL:
// 			length = p.second - p.first + 1
// 			stack = &node{&point{length, p.second, -1}, stack}
// 			curr := stack
// 			for curr.next != nil && curr.point.first > curr.next.point.first { // sort by length
// 				curr.point, curr.next.point = curr.next.point, curr.point
// 				curr = curr.next
// 			}
// 		case QUERY:
// 			queries[p.second] = -1
// 			for stack != nil && stack.point.second < p.first {
// 				stack = stack.next
// 			}
// 			if stack != nil {
// 				queries[p.second] = stack.point.first
// 			}
// 		}
// 	}
// 	return queries
// }

// // Approach: Brute-Force (Time Limit Exceeded)
// // Time: O(NxQ), N=len(intervals), Q=len(queries)
// // Space: O(1)
// func minInterval(intervals [][]int, queries []int) []int {
// 	var qlen = len(queries)
// 	var answer = make([]int, qlen)
// 	var left, right, length int
// 	for idx, query := range queries {
// 		answer[idx] = -1
// 		for _, interval := range intervals {
// 			left, right = interval[0], interval[1]
// 			if query >= left && query <= right {
// 				length = right - left + 1
// 				if answer[idx] == -1 || length < answer[idx] {
// 					answer[idx] = length
// 				}
// 			}
// 		}
// 	}
// 	return answer
// }
