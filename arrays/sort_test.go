package arrays

import (
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
		array:  []int{24, 7, 64, 53, 31, 28, -78, -9, 54, 28, -56, -70, 95, 43, 42, -64, 18, 13, 40, 81},
		sorted: []int{-78, -70, -64, -56, -9, 7, 13, 18, 24, 28, 28, 31, 40, 42, 43, 53, 54, 64, 81, 95},
	},
	{
		array:  []int{1, 5, 10, 5, 0, 2, 52, 14, 95},
		sorted: []int{0, 1, 2, 5, 5, 10, 14, 52, 95},
	},
	{
		array:  []int{24, 7, 64, 53, 31, 28, 54, 28, 82, 43, 42, 18, 13, 40, 81},
		sorted: []int{7, 13, 18, 24, 28, 28, 31, 40, 42, 43, 53, 54, 64, 81, 82},
	},
}

// go test -v -count=1 ./arrays/ -run ^TestBubleSort$
func TestBubleSort(t *testing.T) {
	for _, tc := range sortingTestCases {
		sorted := bubleSort(tc.array)
		assert.Equal(t, tc.sorted, sorted)
	}
}

// go test -v -count=1 ./arrays/ -run ^TestSelectionSort$
func TestSelectionSort(t *testing.T) {
	for _, tc := range sortingTestCases {
		sorted := selectionSort(tc.array)
		assert.Equal(t, tc.sorted, sorted)
	}
}

// go test -v -count=1 ./arrays/ -run ^TestInsertionSort$
func TestInsertionSort(t *testing.T) {
	for _, tc := range sortingTestCases {
		sorted := insertionSort(tc.array)
		assert.Equal(t, tc.sorted, sorted)
	}
}

// go test -v -count=1 ./arrays/ -run ^TestBucketSort$
func TestBucketSort(t *testing.T) {
	for _, tc := range []struct {
		array  []int
		k      int
		sorted []int
	}{
		{
			array:  []int{1, 5, 10, 5, 0, 2, 52, 14, 95},
			k:      95,
			sorted: []int{0, 1, 2, 5, 5, 10, 14, 52, 95},
		},
		{
			array:  []int{24, 7, 64, 53, 31, 28, 54, 28, 82, 43, 42, 18, 13, 40, 81},
			k:      82,
			sorted: []int{7, 13, 18, 24, 28, 28, 31, 40, 42, 43, 53, 54, 64, 81, 82},
		},
		{
			array:  []int{2, 1, 2, 0, 0, 2},
			k:      2,
			sorted: []int{0, 0, 1, 2, 2, 2},
		},
	} {
		sorted := bucketSort(tc.array, tc.k)
		assert.Equal(t, tc.sorted, sorted)
	}
}

// go test -v -count=1 ./arrays/ -run ^TestMergeSort$
func TestMergeSort(t *testing.T) {
	for _, tc := range sortingTestCases {
		sorted := mergeSort(tc.array)
		assert.Equal(t, tc.sorted, sorted)
	}
}

// go test -v -count=1 ./arrays/ -run ^TestQuickSort$
func TestQuickSort(t *testing.T) {
	for _, tc := range sortingTestCases {
		sorted := quickSort(tc.array)
		assert.Equal(t, tc.sorted, sorted)
	}
}

// go test -v -count=1 ./arrays/ -run ^TestHeapSort$
func TestHeapSort(t *testing.T) {
	for _, tc := range sortingTestCases {
		sorted := heapSort(tc.array)
		assert.Equal(t, tc.sorted, sorted)
	}
}
