package dp

import "sort"

// https://leetcode.com/problems/russian-doll-envelopes/

// Approach: Sort + LIS (Longest Increasing Subsequence)
// Sort by width ascending, height descending
// Find LIS on heights using binary search
// Time: O(n log n)
// Space: O(n)
func maxEnvelopes(envelopes [][]int) int {
	binarySearch := func(arr []int, target int) int {
		left, right := 0, len(arr)
		for left < right {
			mid := left + (right-left)/2
			if arr[mid] < target {
				left = mid + 1
			} else {
				right = mid
			}
		}

		return left
	}

	n := len(envelopes)
	if n == 0 {
		return 0
	}

	// Sort: width ascending, height descending
	sort.Slice(envelopes, func(i, j int) bool {
		if envelopes[i][0] != envelopes[j][0] {
			return envelopes[i][0] < envelopes[j][0]
		}

		return envelopes[i][1] > envelopes[j][1]
	})

	// Find LIS on heights
	lis := []int{}

	for _, env := range envelopes {
		h := env[1]

		pos := binarySearch(lis, h)
		if pos == len(lis) {
			lis = append(lis, h)
		} else {
			lis[pos] = h
		}
	}

	return len(lis)
}
