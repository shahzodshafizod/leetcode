package arrays

// https://leetcode.com/problems/minimum-common-value/

// Two Pointers
// time: O(n + m)
// space: O(1)
func getCommon(nums1 []int, nums2 []int) int {
	n1, n2 := len(nums1), len(nums2)
	idx1, idx2 := 0, 0

	for idx1 < n1 && idx2 < n2 {
		if nums1[idx1] < nums2[idx2] {
			idx1++
		} else if nums1[idx1] > nums2[idx2] {
			idx2++
		} else {
			return nums1[idx1]
		}
	}

	return -1
}

// // Binary Search
// // time: O(n log m)
// // space: O(1)
// func getCommon(nums1 []int, nums2 []int) int {
// 	var mid int
// 	var len = len(nums2)
// 	for _, num := range nums1 {
// 		for left, right := 0, len-1; left <= right; {
// 			mid = (left + right) / 2
// 			if nums2[mid] < num {
// 				left = mid + 1
// 			} else if nums2[mid] > num {
// 				right = mid - 1
// 			} else {
// 				return num
// 			}
// 		}
// 	}
// 	return -1
// }
