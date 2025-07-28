package binarysearch

import (
	"math"
	"slices"
)

// https://leetcode.com/problems/minimum-time-to-repair-cars/

func repairCars(ranks []int, cars int) int64 {
	n := len(ranks)

	var low, high int64 = 1, int64(slices.Max(ranks) * cars * cars)

	var mid int64

	var idx, remained int

	for low < high {
		mid = low + (high-low)/2
		remained = cars

		for idx = 0; idx < n && remained > 0; idx++ {
			// rank * n*n <= mid
			remained -= int(math.Sqrt(float64(mid / int64(ranks[idx]))))
		}

		if remained <= 0 {
			high = mid
		} else {
			low = mid + 1
		}
	}

	return high
}
