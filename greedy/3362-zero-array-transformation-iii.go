package greedy

import (
	"container/heap"
	"sort"

	"github.com/shahzodshafizod/leetcode/pkg"
)

// https://leetcode.com/problems/zero-array-transformation-iii/

// Approach: Greedy Line Sweep
// Time: O(n + mlogm); n=len(nums), m=len(queries)
// Space: O(n+m)
func maxRemoval(nums []int, queries [][]int) int {
	sort.Slice(queries, func(i, j int) bool {
		return queries[i][0] < queries[j][0]
	})
	line := make([]int, len(nums)+1)
	ops := 0
	qid, n := 0, len(queries)
	maxHeap := pkg.NewHeap([]int{}, func(x, y int) bool { return x > y })
	for idx, num := range nums {
		ops += line[idx]
		for qid < n && queries[qid][0] == idx {
			heap.Push(maxHeap, queries[qid][1])
			qid++
		}
		for ops < num && maxHeap.Len() > 0 && maxHeap.Peak() >= idx {
			ops++
			line[heap.Pop(maxHeap).(int)+1]--
		}
		if ops < num {
			return -1
		}
	}
	return maxHeap.Len()
}
