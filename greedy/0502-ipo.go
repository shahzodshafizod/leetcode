package greedy

import (
	"github.com/shahzodshafizod/alkhwarizmi/design"
)

// https://leetcode.com/problems/ipo/

// time: O(k log n)
// space: O(n)
func findMaximizedCapital(k int, w int, profits []int, capital []int) int {
	var maxHeap = design.NewPQ(make([]int, 0), func(x, y int) bool { return x < y })
	var minHeap = design.NewPQ(
		make([][2]int, 0),
		func(x, y [2]int) bool {
			if x[0] == y[0] {
				return x[1] > y[1]
			}
			return x[0] > y[0]
		},
	)
	for idx := range capital { // O(n)
		if capital[idx] <= w {
			maxHeap.Push(profits[idx])
		} else {
			minHeap.Push([2]int{capital[idx], profits[idx]})
		}
	}
	for k > 0 && maxHeap.Len() > 0 { // O(k)
		k--
		w += maxHeap.Pop()
		for minHeap.Len() > 0 && minHeap.Peek()[0] <= w {
			maxHeap.Push(minHeap.Pop()[1]) // O(log n)
		}
	}
	return w
}
