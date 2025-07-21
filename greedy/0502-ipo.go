package greedy

import (
	"container/heap"

	"github.com/shahzodshafizod/leetcode/pkg"
)

/*
0. Initially, you have capital in size of w.
1. Divide your capital:profits into two groups:
    1.1. The 1st group contains capital you can afford to finance (spend): capital[i] <= w
    1.2. The 2nd group contains capital you cannot afford to finance (spend): capital[i] > w
2. Pick from the first group of capital the one with the highest profit and add that profit to your capital.
3. After step (2), your capital grew, and there may appear capital from the second group that are affordable (<= w). Move them to the first group.
4. Repeat steps (2) and (4), k times.
*/

// https://leetcode.com/problems/ipo/

// time: O(k log n)
// space: O(n)
func findMaximizedCapital(k int, w int, profits []int, capital []int) int {
	maxHeap := pkg.NewHeap(make([]int, 0), func(x, y int) bool { return x > y })
	minHeap := pkg.NewHeap(
		make([][2]int, 0),
		func(x, y [2]int) bool {
			if x[0] == y[0] {
				return x[1] < y[1]
			}
			return x[0] < y[0]
		},
	)
	for idx := range capital { // O(n)
		if capital[idx] <= w {
			heap.Push(maxHeap, profits[idx])
		} else {
			heap.Push(minHeap, [2]int{capital[idx], profits[idx]})
		}
	}
	for k > 0 && maxHeap.Len() > 0 { // O(k)
		k--
		w += heap.Pop(maxHeap).(int)
		for minHeap.Len() > 0 && minHeap.Peak()[0] <= w {
			heap.Push(maxHeap, heap.Pop(minHeap).([2]int)[1]) // O(log n)
		}
	}
	return w
}
