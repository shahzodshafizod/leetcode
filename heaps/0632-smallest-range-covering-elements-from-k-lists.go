package heaps

import "sort"

// https://leetcode.com/problems/smallest-range-covering-elements-from-k-lists/

/*
[4,10,15,24,26] | [0,9,12,20] | [5,18,22,30] = [0,5]
[4,10,15,24,26] | [9,12,20]   | [5,18,22,30] = [4,9]
[10,15,24,26]   | [9,12,20]   | [5,18,22,30] = [5,10]
[10,15,24,26]   | [9,12,20]   | [18,22,30]   = [9,18]
[10,15,24,26]   | [12,20]     | [18,22,30]   = [10,18]
[15,24,26]      | [12,20]     | [18,22,30]   = [12,18]
[24,26]         | [20]        | [18,22,30]   = [18,24]
[24,26]         | [20]        | [22,30]      = [20,24]
[24,26]         | []          | [22,30]      = stop
*/

// Approach: Two Pointers
// Time: O(N+KogN), N=len(nums), K=min(len(nums[i]))
// Space: O(N)
func smallestRange(nums [][]int) []int {
	var merged = make([][2]int, 0)
	for idx := range nums {
		for _, num := range nums[idx] {
			merged = append(merged, [2]int{num, idx})
		}
	}
	sort.Slice(merged, func(i, j int) bool {
		return merged[i][0] < merged[j][0]
	})
	var n = len(nums)
	var freq = make([]int, n)
	var left, count = 0, 0
	var length int
	var start, end = merged[0][0], merged[len(merged)-1][0]
	for right := range merged {
		freq[merged[right][1]]++
		if freq[merged[right][1]] == 1 {
			count++
		}
		for count == n {
			length = merged[right][0] - merged[left][0]
			if length < end-start {
				start = merged[left][0]
				end = merged[right][0]
			}
			freq[merged[left][1]]--
			if freq[merged[left][1]] == 0 {
				count--
			}
			left++
		}
	}
	return []int{start, end}
}

// // Approach: Heap (Priority Queue)
// // Time: O(N+KogN), N=len(nums), K=min(len(nums[i]))
// // Space: O(N)
// func smallestRange(nums [][]int) []int {
// 	var n = len(nums)
// 	var data = make([][2]int, n) // [row,col]
// 	var end int = -1e5
// 	for idx := range nums { // O(N)
// 		data[idx] = [2]int{idx, 0}
// 		end = max(end, nums[idx][0])
// 	}
// 	var minHeap = pkg.NewHeap(data, func(x, y [2]int) bool {
// 		return nums[x[0]][x[1]] < nums[y[0]][y[1]]
// 	})
// 	heap.Init(minHeap)
// 	var start int
// 	var interval []int
// 	var length = 2 * int(1e5)
// 	for { // O(min(len(nums[i])))
// 		start = nums[data[0][0]][data[0][1]]
// 		if end-start < length {
// 			length = end - start
// 			interval = []int{start, end}
// 		}
// 		data[0][1]++
// 		if data[0][1] == len(nums[data[0][0]]) {
// 			break
// 		}
// 		end = max(end, nums[data[0][0]][data[0][1]])
// 		heap.Fix(minHeap, 0) // O(LogN)
// 	}
// 	return interval
// }
