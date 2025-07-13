package heaps

import (
	"container/heap"
	"slices"
	"sort"

	"github.com/shahzodshafizod/leetcode/pkg"
)

// https://leetcode.com/problems/meeting-rooms-iii/

// Approach: Heap (Priority Queue)
// Time: O(M⋅logM+M⋅logN)
// Space: O(N)
func mostBooked(n int, meetings [][]int) int {
	var count = make([]int, n)
	var held = pkg.NewHeap(
		[][2]int{}, // {end, room}
		func(x, y [2]int) bool {
			return x[0] < y[0] || x[0] == y[0] && x[1] < y[1]
		},
	)
	// sorted array is a valid heap
	var rooms = make([]int, n)
	for room := 0; room < n; room++ {
		rooms[room] = room
	}
	var avai = pkg.NewHeap(rooms, func(x, y int) bool { return x < y })
	sort.Slice(meetings, func(i, j int) bool { return meetings[i][0] < meetings[j][0] })
	var start, end, room int
	var mroom [2]int
	for _, meeting := range meetings {
		start, end = meeting[0], meeting[1]
		for held.Len() > 0 && held.Peak()[0] <= start {
			mroom = heap.Pop(held).([2]int)
			heap.Push(avai, mroom[1])
		}
		if avai.Len() == 0 {
			mroom = heap.Pop(held).([2]int)
			heap.Push(avai, mroom[1])
			end += mroom[0] - start
		}
		room = heap.Pop(avai).(int)
		heap.Push(held, [2]int{end, room})
		count[room]++
	}
	return slices.Index(count, slices.Max(count))
}
