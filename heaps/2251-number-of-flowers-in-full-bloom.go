package heaps

import (
	"container/heap"
	"sort"

	"github.com/shahzodshafizod/leetcode/pkg"
)

// https://leetcode.com/problems/number-of-flowers-in-full-bloom/

// Approach: Heap (Priority Queue)
// Time: O(nlogn + mlogm), n=len(flowers), m=len(people)
// Space: O(n + m)
func fullBloomFlowers(flowers [][]int, people []int) []int {
	sort.Slice(flowers, func(i, j int) bool {
		return flowers[i][0] < flowers[j][0]
	})
	n, m := len(flowers), len(people)
	answer := make([]int, m)
	peopids := make([][2]int, m)
	for idx, person := range people {
		peopids[idx][0] = person
		peopids[idx][1] = idx
	}
	sort.Slice(peopids, func(i, j int) bool {
		return peopids[i][0] < peopids[j][0]
	})
	ends := pkg.NewHeap(make([]int, 0), func(x, y int) bool { return x < y })
	var person, pid int
	fid := 0
	for idx := range peopids {
		person, pid = peopids[idx][0], peopids[idx][1]
		for fid < n && flowers[fid][0] <= person {
			heap.Push(ends, flowers[fid][1])
			fid++
		}
		for ends.Len() > 0 && ends.Peak() < person {
			heap.Pop(ends)
		}
		answer[pid] = ends.Len()
	}
	return answer
}
