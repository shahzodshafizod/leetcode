package learning

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

var sortingTestCases = []struct {
	array  []int
	sorted []int
}{
	{
		array:  []int{1, 5, -7, 10, 5, -4, 125, 0, 2, 52, 14, 95, -1, 789, -32},
		sorted: []int{-32, -7, -4, -1, 0, 1, 2, 5, 5, 10, 14, 52, 95, 125, 789},
	},
	{
		array:  []int{-32, -7, -4, -1, 0, 1, 2, 5, 5, 10, 14, 52, 95, 125, 789},
		sorted: []int{-32, -7, -4, -1, 0, 1, 2, 5, 5, 10, 14, 52, 95, 125, 789},
	},
	{
		array: []int{
			24,
			7,
			64,
			53,
			31,
			28,
			-78,
			-9,
			54,
			28,
			-56,
			-70,
			95,
			43,
			42,
			-64,
			18,
			13,
			40,
			81,
		},
		sorted: []int{
			-78,
			-70,
			-64,
			-56,
			-9,
			7,
			13,
			18,
			24,
			28,
			28,
			31,
			40,
			42,
			43,
			53,
			54,
			64,
			81,
			95,
		},
	},
	{
		array:  []int{1, 5, 10, 5, 0, 2, 52, 14, 95},
		sorted: []int{0, 1, 2, 5, 5, 10, 14, 52, 95},
	},
	{
		array:  []int{24, 7, 64, 53, 31, 28, 54, 28, 82, 43, 42, 18, 13, 40, 81},
		sorted: []int{7, 13, 18, 24, 28, 28, 31, 40, 42, 43, 53, 54, 64, 81, 82},
	},
	{
		array:  []int{2, 1, 2, 0, 0, 2},
		sorted: []int{0, 0, 1, 2, 2, 2},
	},
}

// go test -v -count=1 ./learning/ -run ^TestBubleSort$
func TestBubleSort(t *testing.T) {
	for _, tc := range sortingTestCases {
		copied := append([]int{}, tc.array...)
		sorted := bubleSort(copied)
		assert.Equal(t, tc.sorted, sorted)
	}
}

// go test -v -count=1 ./learning/ -run ^TestSelectionSort$
func TestSelectionSort(t *testing.T) {
	for _, tc := range sortingTestCases {
		copied := append([]int{}, tc.array...)
		sorted := selectionSort(copied)
		assert.Equal(t, tc.sorted, sorted)
	}
}

// go test -v -count=1 ./learning/ -run ^TestInsertionSort$
func TestInsertionSort(t *testing.T) {
	for _, tc := range sortingTestCases {
		copied := append([]int{}, tc.array...)
		sorted := insertionSort(copied)
		assert.Equal(t, tc.sorted, sorted)
	}
}

// go test -v -count=1 ./learning/ -run ^TestBucketSort$
func TestBucketSort(t *testing.T) {
	for _, tc := range sortingTestCases {
		if slices.Min(tc.array) < 0 {
			continue
		}

		copied := append([]int{}, tc.array...)
		sorted := bucketSort(copied)
		assert.Equal(t, tc.sorted, sorted)
	}
}

// go test -v -count=1 ./learning/ -run ^TestMergeSort$
func TestMergeSort(t *testing.T) {
	for _, tc := range sortingTestCases {
		copied := append([]int{}, tc.array...)
		sorted := mergeSort(copied)
		assert.Equal(t, tc.sorted, sorted)
	}
}

// go test -v -count=1 ./learning/ -run ^TestQuickSort$
func TestQuickSort(t *testing.T) {
	for _, tc := range sortingTestCases {
		copied := append([]int{}, tc.array...)
		sorted := quickSort(copied)
		assert.Equal(t, tc.sorted, sorted)
	}
}

// go test -v -count=1 ./learning/ -run ^TestHeapSort$
func TestHeapSort(t *testing.T) {
	for _, tc := range sortingTestCases {
		copied := append([]int{}, tc.array...)
		sorted := heapSort(copied)
		assert.Equal(t, tc.sorted, sorted)
	}
}
