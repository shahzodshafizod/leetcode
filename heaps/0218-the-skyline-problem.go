package heaps

import (
	"container/heap"
	"sort"

	"github.com/shahzodshafizod/leetcode/pkg"
)

// https://leetcode.com/problems/the-skyline-problem/

// Approach: Sweep Line + Max Heap
// Create events for each building's start and end
// Use max heap to track active building heights at each x-coordinate
// When max height changes, add key point to result
// N = number of buildings
// Time: O(N log N) - N events, each heap operation is O(log N)
// Space: O(N) - heap and events storage

func getSkyline(buildings [][]int) [][]int {
	type Event struct {
		x      int
		isEnd  bool
		height int
	}

	// Create events
	events := make([]Event, 0, len(buildings)*2)
	for _, b := range buildings {
		left, right, height := b[0], b[1], b[2]
		// Start event: not end, with height
		events = append(events, Event{x: left, isEnd: false, height: height})
		// End event: is end, with height
		events = append(events, Event{x: right, isEnd: true, height: height})
	}

	// Sort events
	sort.Slice(events, func(i, j int) bool {
		if events[i].x != events[j].x {
			return events[i].x < events[j].x
		}
		// Same x: start before end
		if events[i].isEnd != events[j].isEnd {
			return !events[i].isEnd // false (start) before true (end)
		}
		// Same x and type: for starts, taller first; for ends, shorter first
		if !events[i].isEnd {
			return events[i].height > events[j].height
		}

		return events[i].height < events[j].height
	})

	result := [][]int{}
	maxHeap := pkg.NewHeap(make([]int, 0), func(x, y int) bool { return x > y })
	heap.Init(maxHeap)

	heightCount := make(map[int]int)
	heightCount[0] = 1

	i := 0
	for i < len(events) {
		currX := events[i].x

		// Process all events at the same x-coordinate
		for i < len(events) && events[i].x == currX {
			event := events[i]

			if !event.isEnd { // Start event
				heap.Push(maxHeap, event.height)
				heightCount[event.height]++
			} else { // End event
				heightCount[event.height]--
				if heightCount[event.height] == 0 {
					delete(heightCount, event.height)
				}
			}

			i++
		}

		// Remove invalid heights from top of heap
		for maxHeap.Len() > 0 {
			top := maxHeap.Peak()
			if _, exists := heightCount[top]; exists {
				break
			}

			heap.Pop(maxHeap)
		}

		// Get current max height
		maxHeight := 0
		if maxHeap.Len() > 0 {
			maxHeight = maxHeap.Peak()
		}

		// If height changed, add to result
		if len(result) == 0 || result[len(result)-1][1] != maxHeight {
			result = append(result, []int{currX, maxHeight})
		}
	}

	return result
}
