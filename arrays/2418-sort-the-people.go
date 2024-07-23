package arrays

// https://leetcode.com/problems/sort-the-people/

// Approach #4: Space Optimized Priority Queue
// Time: O(N Log N)
// Space: O(1)
func sortPeople(names []string, heights []int) []string {
	var left = func(parent int) int { // right: 2*idx + 2
		return 2*parent + 1
	}
	var compare = func(i int, j int) bool {
		return heights[i] > heights[j]
	}
	var swap = func(i int, j int) {
		heights[i], heights[j] = heights[j], heights[i]
		names[i], names[j] = names[j], names[i]
	}
	var shiftDown = func(parent int, n int) { // O(Log N)
		child := left(parent)
		for child < n {
			if child+1 < n && compare(child, child+1) { // right
				child++
			}
			if !compare(parent, child) {
				break
			}
			swap(parent, child)
			parent = child
			child = left(parent)
		}
	}
	var heapify = func(n int) {
		for parent := n/2 - 1; parent >= 0; parent-- { // O(N/2) = O(N)
			shiftDown(parent, n)
		}
	}
	var sort = func(n int) {
		for n > 0 { // O(N)
			n--
			swap(0, n)
			shiftDown(0, n)
		}
	}

	var n = len(names)
	heapify(n) // O(N Log N)
	sort(n)    // O(N Log N)
	return names
}

// // Approach #3: Quick Sort
// // Time: average: O(N Log N); worst: O(N^2)
// // Space: O(N)
// func sortPeople(names []string, heights []int) []string {
// 	var swap = func(i int, j int) {
// 		heights[i], heights[j] = heights[j], heights[i]
// 		names[i], names[j] = names[j], names[i]
// 	}
// 	var partition = func(left int, right int) int {
// 		var pivot = heights[right]
// 		var mid = left
// 		for i := left; i < right; i++ {
// 			if heights[i] > pivot {
// 				swap(i, mid)
// 				mid++
// 			}
// 		}
// 		swap(mid, right)
// 		return mid
// 	}
// 	var quickSort func(left int, right int)
// 	quickSort = func(left int, right int) {
// 		if left >= right {
// 			return
// 		}
// 		var mid = partition(left, right)
// 		quickSort(left, mid-1)
// 		quickSort(mid+1, right)
// 	}

// 	quickSort(0, len(heights)-1)
// 	return names
// }

// // Approach #2: Index Sorting
// // Time: O(N Log N)
// // Space: O(N)
// func sortPeople(names []string, heights []int) []string {
// 	var n = len(names)
// 	var indices = make([]int, n)
// 	for idx := range indices {
// 		indices[idx] = idx
// 	}
// 	sort.Slice(indices,
// 		func(i, j int) bool {
// 			return heights[indices[i]] > heights[indices[j]]
// 		},
// 	)
// 	for idx := 0; idx < n; idx++ {
// 		names[idx], names[indices[idx]] = names[indices[idx]], names[idx]
// 		indices[idx], indices[indices[idx]] = indices[indices[idx]], indices[idx]
// 	}
// 	return names
// }

// // Approach #1: Sorting
// // Time: O(N Log N)
// // Space: O(N)
// func sortPeople(names []string, heights []int) []string {
// 	type people struct {
// 		name   string
// 		height int
// 	}
// 	var n = len(names)
// 	var p = make([]*people, 0, n)
// 	for idx := 0; idx < n; idx++ {
// 		p = append(p, &people{names[idx], heights[idx]})
// 	}
// 	sort.Slice(p, func(i, j int) bool { return p[i].height > p[j].height })
// 	for idx := 0; idx < n; idx++ {
// 		names[idx] = p[idx].name
// 	}
// 	return names
// }
