package learning

import "slices"

// Sorting Algorithms Dance
// https://www.youtube.com/user/AlgoRythmics/videos

// Sorting Humor
// https://www.reddit.com/r/ProgrammerHumor/comments/9s9kgn/nononsense_sorting_algorithm/

// 1. Elementary Sort

// time: O(N^2), if array isn't sorted; else O(N)
// space: O(1)
func bubleSort(array []int) []int {
	for i, n := 0, len(array); i < n-1; i++ {
		sorted := true

		for j := 1; j < n-i; j++ {
			if array[j-1] > array[j] {
				array[j-1], array[j] = array[j], array[j-1]
				sorted = false
			}
		}

		if sorted {
			break
		}
	}

	return array
}

// time: O(N^2)
// space: O(1)
func selectionSort(array []int) []int {
	var minimum int
	for i, n := 0, len(array); i < n-1; i++ {
		minimum = i
		for j := i + 1; j < n; j++ {
			if array[j] < array[minimum] {
				minimum = j
			}
		}

		array[i], array[minimum] = array[minimum], array[i]
	}

	return array
}

// time: O(N^2), if array isn't sorted (not ASC, nor DEC); else O(N)
// space: O(1)
func insertionSort(array []int) []int {
	for i, n := 1, len(array); i < n; i++ {
		if array[i] < array[0] { // space: O(N)
			array = append(append([]int{array[i]}, array[:i]...), array[i+1:]...)
			continue
		}

		for j := i; j > 0 && array[j-1] > array[j]; j-- {
			array[j-1], array[j] = array[j], array[j-1]
		}
	}

	return array
}

// Counting Sort
// time: O(N+N) -> O(N)
// space: O(k) -> O(1)
// works for [0;k]
func bucketSort(array []int) []int {
	k := slices.Max(array)

	buckets := make([]int, k+1)
	for _, num := range array { // O(N)
		buckets[num]++
	}

	idx := 0

	for num, count := range buckets { // O(k)
		for count > 0 { // O(N/k)
			count--
			array[idx] = num
			idx++
		}
	}

	return array
}

// 2. Divide & Conquer

// time: O(N*Log(N))
// space: O(N)
func mergeSort(array []int) []int {
	n := len(array)
	if n <= 1 {
		return array
	}

	middle := n / 2

	return merge(
		mergeSort(array[:middle]), // [0-middle)
		mergeSort(array[middle:]), // [middle-len)
	)
}

func merge(array1 []int, array2 []int) []int {
	len1, len2 := len(array1), len(array2)
	array := make([]int, 0, len1+len2)

	idx1, idx2 := 0, 0
	for idx1 < len1 && idx2 < len2 {
		if array1[idx1] < array2[idx2] {
			array = append(array, array1[idx1])
			idx1++
		} else {
			array = append(array, array2[idx2])
			idx2++
		}
	}

	if idx1 < len1 {
		array = append(array, array1[idx1:]...)
	} else if idx2 < len2 {
		array = append(array, array2[idx2:]...)
	}

	return array
}

// time: O(N^2), is sorted; else O(N*Log(N))
// space: O(Log(N))
func quickSort(array []int) []int {
	partition := func(array []int, left int, right int) int {
		pivot := array[right]
		mid := left

		for i := left; i < right; i++ {
			if array[i] < pivot {
				array[i], array[mid] = array[mid], array[i]
				mid++
			}
		}

		array[right], array[mid] = array[mid], array[right]

		return mid
	}

	var helper func(left int, right int)

	helper = func(left int, right int) {
		if left < right {
			mid := partition(array, left, right)
			helper(left, mid-1)
			helper(mid+1, right)
		}
	}
	helper(0, len(array)-1)

	return array
}

// 3. Tree Sort

func heapSort(array []int) []int {
	var (
		getParent = func(child int) int { return (child - 1) / 2 }
		getLeft   = func(parent int) int { return 2*parent + 1 }
		compare   = func(i int, j int) bool { return array[i] < array[j] }
		swap      = func(i int, j int) { array[i], array[j] = array[j], array[i] }
	)
	// 1. create a max heap
	n := len(array)
	for i := 1; i < n; i++ {
		// sift up
		child := i
		parent := getParent(child)

		for parent >= 0 && compare(parent, child) {
			swap(parent, child)
			child = parent
			parent = getParent(child)
		}
	}
	// 2. sort ascending
	// move max to the end, and continue with the array[0:len-1]
	for n > 0 {
		n--
		swap(0, n)
		// sift down
		parent := 0

		child := getLeft(parent)
		for child < n {
			if child+1 < n && compare(child, child+1) {
				child++
			}

			if !compare(parent, child) {
				break
			}

			swap(parent, child)
			parent = child
			child = getLeft(parent)
		}
	}

	return array
}
