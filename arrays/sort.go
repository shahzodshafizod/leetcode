package arrays

import "github.com/shahzodshafizod/alkhwarizmi/design"

// Sorting Algorithms Dance
// https://www.youtube.com/user/AlgoRythmics/videos

// 1. Elementary Sort

// time: O(N^2), if array isn't sorted; else O(N)
// space: O(1)
func bubleSort(array []int) []int {
	for i, len := 0, len(array); i < len-1; i++ {
		sorted := true
		for j := 1; j < len-i; j++ {
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
	for i, len := 0, len(array); i < len-1; i++ {
		var min = i
		for j := i + 1; j < len; j++ {
			if array[j] < array[min] {
				min = j
			}
		}
		array[i], array[min] = array[min], array[i]
	}
	return array
}

// time: O(N^2), if array isn't sorted (not ASC, nor DEC); else O(N)
// space: O(1)
func insertionSort(array []int) []int {
	for i, len := 0, len(array); i < len; i++ {
		if array[i] < array[0] {
			array = append(append([]int{array[i]}, array[:i]...), array[i+1:]...)
		} else {
			for j := i; j > 0; j-- {
				if array[j-1] > array[j] {
					array[j-1], array[j] = array[j], array[j-1]
				}
			}
		}
	}
	return array
}

// 2. Divide & Conquer

// time: O(N*Log(N))
// space: O(N)
func mergeSort(array []int) []int {
	var len = len(array)
	if len <= 1 {
		return array
	}
	var middle = len / 2
	var left, right = array[:middle], array[middle:]
	return merge(mergeSort(left), mergeSort(right))
}

func merge(array1 []int, array2 []int) []int {
	var len1, len2 = len(array1), len(array2)
	var array = make([]int, 0, len1+len2)
	var idx1, idx2 = 0, 0
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
	quickSortHelper(array, 0, len(array)-1)
	return array
}

func quickSortHelper(array []int, left int, right int) {
	if left < right {
		mid := partition(array, left, right)
		quickSortHelper(array, left, mid-1)
		quickSortHelper(array, mid+1, right)
	}
}

func partition(array []int, left int, right int) int {
	var pivot = array[right]
	var mid = left
	for i := left; i < right; i++ {
		if array[i] < pivot {
			array[i], array[mid] = array[mid], array[i]
			mid++
		}
	}
	array[right], array[mid] = array[mid], array[right]
	return mid
}

// 3. Tree Sort

func heapSort(array []int) []int {
	var heap = design.NewPriorityQueue[int](func(item1, item2 int) bool { return item1 > item2 })
	for _, item := range array {
		heap.Push(item)
	}
	for i, len := 0, len(array); i < len; i++ {
		array[i] = heap.Pop()
	}
	return array
}
