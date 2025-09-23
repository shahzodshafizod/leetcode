package arrays

import "sort"

// https://leetcode.com/problems/merge-sorted-array/

func merge(nums1 []int, m int, nums2 []int, n int) {
	_ = n

	copy(nums1[m:], nums2)
	sort.Ints(nums1)
}
