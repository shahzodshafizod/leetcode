package binarysearch

import (
	"sort"
)

// https://leetcode.com/problems/maximum-number-of-tasks-you-can-assign/

// Approach: Binary Search + Greedy Selection of Workers
// Time: O(nlogn + mlogm + klogk), k=Min(n,m)
// Space: O(m)
func maxTaskAssign(tasks []int, workers []int, pills int, strength int) int {
	n, m := len(tasks), len(workers)
	sort.Ints(tasks)   // O(n log n)
	sort.Ints(workers) // O(m log m)
	deque := make([]int, m)
	greedy := func(mid int) bool { // O(min(n,m))
		p := pills
		first, last := m-1, m-1
		wid := m - 1
		// enumerate each task from largest to smallest
		for tid := mid - 1; tid >= 0; tid-- {
			for wid >= m-mid && workers[wid]+strength >= tasks[tid] {
				deque[first] = workers[wid]
				first--
				wid--
			}
			if first == last { // not enough strong workers, even with pills
				return false
			}
			// if the strongest worker can manage the task without a pill, use that worker
			if deque[last] >= tasks[tid] {
				last--
			} else if p > 0 {
				first++
				p--
			} else { // not enough pills
				return false
			}
		}
		return true
	}
	left, right := 0, min(n, m) // left=0 is important
	var mid int
	for left < right { // O(log min(n,m))
		mid = left + (right-left+1)/2
		if greedy(mid) {
			left = mid
		} else {
			right = mid - 1
		}
	}
	return left
}
