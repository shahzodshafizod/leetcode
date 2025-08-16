package heaps

import (
	"container/heap"

	"github.com/shahzodshafizod/leetcode/pkg"
)

// https://leetcode.com/problems/time-to-cross-a-bridge/

// Approach: Simulation
// Time: O(n log n)
// Space: O(n)
func findCrossingTime(n int, k int, time [][]int) int {
	mxcmp := func(x, y [2]int) bool { return x[0] == y[0] && x[1] > y[1] || x[0] > y[0] }
	mncmp := func(x, y [2]int) bool { return x[0] == y[0] && x[1] < y[1] || x[0] < y[0] }
	l2r := pkg.NewHeap([][2]int{}, mxcmp)  // ready to move from left to right: [left+right, index]
	r2l := pkg.NewHeap([][2]int{}, mxcmp)  // ready to move from right to left: [left+right, index]
	pick := pkg.NewHeap([][2]int{}, mncmp) // ready to pick: [time, index]
	put := pkg.NewHeap([][2]int{}, mncmp)  // ready to put: [time, index]

	// Initially, all k workers are waiting on the left side of the bridge.
	for idx := range k {
		heap.Push(l2r, [2]int{time[idx][0] + time[idx][2], idx})
	}

	var (
		ok             bool
		item           [2]int
		minutes, index int
	)

	for n > 0 {
		// 2. Assign boxes to workers who are waiting to pick
		for pick.Len() > 0 && pick.Peak()[0] <= minutes {
			item, ok = heap.Pop(pick).([2]int)
			index = item[1]
			heap.Push(r2l, [2]int{time[index][0] + time[index][2], index})
		}
		// 3. Empty workers who have carried boxes to the right side
		for put.Len() > 0 && put.Peak()[0] <= minutes {
			item, ok = heap.Pop(put).([2]int)
			index = item[1]
			heap.Push(l2r, [2]int{time[index][0] + time[index][2], index})
		}

		if r2l.Len() > 0 {
			// 4. When the bridge is unused prioritize the least efficient worker
			//    (who have picked up the box) on the right side to cross.
			item, ok = heap.Pop(r2l).([2]int)
			index = item[1]
			minutes += time[index][2]
			heap.Push(put, [2]int{minutes + time[index][3], index})

			n--
		} else if l2r.Len() > 0 && n > pick.Len()+r2l.Len() {
			// 5. If not, prioritize the least efficient worker on the left side to cross.
			item, ok = heap.Pop(l2r).([2]int)
			index = item[1]
			minutes += time[index][0]
			heap.Push(pick, [2]int{minutes + time[index][1], index})
		} else {
			minutes = (1 << 31) - 1
			if put.Len() > 0 && n > pick.Len()+r2l.Len() {
				minutes = min(minutes, put.Peak()[0])
			}

			if pick.Len() > 0 {
				minutes = min(minutes, pick.Peak()[0])
			}
		}

		_ = ok
	}

	return minutes
}
