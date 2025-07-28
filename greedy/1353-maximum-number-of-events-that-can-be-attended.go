package greedy

import (
	"container/heap"
	"sort"

	"github.com/shahzodshafizod/leetcode/pkg"
)

// https://leetcode.com/problems/maximum-number-of-events-that-can-be-attended/

func maxEvents(events [][]int) int {
	sort.Slice(events, func(i int, j int) bool {
		return events[i][1] > events[j][1]
	})

	pq := pkg.NewHeap([]int{}, func(x, y int) bool { return x > y })
	i, n := 0, len(events)
	count := 0

	for day := events[0][1]; day > 0; day-- {
		for i < n && events[i][1] == day {
			heap.Push(pq, events[i][0])

			i++
		}

		for pq.Len() > 0 && day < pq.Peak() {
			heap.Pop(pq)
		}

		if pq.Len() > 0 {
			count++

			heap.Pop(pq)
		}
	}

	return count
}
