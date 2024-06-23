package slidingwindows

// https://leetcode.com/problems/longest-continuous-subarray-with-absolute-diff-less-than-or-equal-to-limit/

// Approach #3: Monotonic Queue + Sliding Window
// time: O(5n) = O(n)
// space: O(2n) = O(n)
func longestSubarray(nums []int, limit int) int {
	type monoqueue struct {
		indices []int
		start   int
		end     int
		compare func(a int, b int) bool
	}
	var expand = func(m *monoqueue, idx int) {
		for m.end >= m.start && !m.compare(nums[m.indices[m.end]], nums[idx]) {
			m.end--
		}
		m.end++
		m.indices[m.end] = idx
	}
	var shrink = func(m *monoqueue, idx int) {
		for m.start < m.end && m.indices[m.start] < idx {
			m.start++
		}
	}
	var first = func(m *monoqueue) int {
		return m.indices[m.start]
	}

	var n = len(nums)
	var monoqinc = &monoqueue{make([]int, n), 0, 0, func(a, b int) bool { return a < b }}
	var monoqdec = &monoqueue{make([]int, n), 0, 0, func(a, b int) bool { return a > b }}
	var start = 0
	var size = 0
	for end := range nums { // O(n)
		expand(monoqinc, end) // O(+n)
		expand(monoqdec, end) // O(+n)
		for nums[first(monoqdec)]-nums[first(monoqinc)] > limit {
			start = min(first(monoqdec), first(monoqinc)) + 1
			shrink(monoqinc, start) // O(+n)
			shrink(monoqdec, start) // O(+n)
		}
		size = max(size, end-start+1)
	}
	return size
}

// // Approach #2: Two Heaps
// // time: O(n log n)
// // space: O(2n) = O(n)
// func longestSubarray(nums []int, limit int) int {
// 	type minmax struct {
// 		min design.PQ[int]
// 		max design.PQ[int]
// 	}

// 	var expand = func(mm *minmax, idx int) {
// 		mm.min.Push(idx)
// 		mm.max.Push(idx)
// 	}
// 	var shrink = func(mm *minmax, idx int) {
// 		for mm.min.Peek() < idx {
// 			mm.min.Pop()
// 		}
// 		for mm.max.Peek() < idx {
// 			mm.max.Pop()
// 		}
// 	}
// 	var getMin = func(mm *minmax) int {
// 		return mm.min.Peek()
// 	}
// 	var getMax = func(mm *minmax) int {
// 		return mm.max.Peek()
// 	}

// 	var mm = &minmax{
// 		min: design.NewPQ[int](make([]int, 0), func(x, y int) bool { return nums[x] > nums[y] }),
// 		max: design.NewPQ[int](make([]int, 0), func(x, y int) bool { return nums[x] < nums[y] }),
// 	}

// 	// main logic
// 	var size = 0
// 	var start = 0
// 	for end := range nums {
// 		expand(mm, end)
// 		for nums[getMax(mm)]-nums[getMin(mm)] > limit {
// 			start = min(getMax(mm), getMin(mm)) + 1
// 			shrink(mm, start)
// 		}
// 		size = max(size, end-start+1)
// 	}
// 	return size
// }

// // Approach #1: Brute-Force
// // time: O(n^2)
// // space: O(1)
// func longestSubarray(nums []int, limit int) int {
// 	var size = 0
// 	var start = 0
// 	for end := range nums { // O(n)
// 		for start < end {
// 			var mi, ma = start, start
// 			for idx := start + 1; idx <= end; idx++ { // O(n)
// 				if nums[idx] < nums[mi] {
// 					mi = idx
// 				}
// 				if nums[idx] > nums[ma] {
// 					ma = idx
// 				}
// 			}
// 			if nums[ma]-nums[mi] <= limit {
// 				break
// 			}
// 			start = min(mi, ma) + 1
// 		}
// 		size = max(size, end-start+1)
// 	}
// 	return size
// }
