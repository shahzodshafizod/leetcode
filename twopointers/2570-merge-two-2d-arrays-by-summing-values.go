package twopointers

// https://leetcode.com/problems/merge-two-2d-arrays-by-summing-values/

func mergeArrays(nums1 [][]int, nums2 [][]int) [][]int {
	var n1, n2 = len(nums1), len(nums2)
	var res = make([][]int, 0, n1+n2)
	var idx1, idx2 = 0, 0
	for idx1 < n1 && idx2 < n2 {
		if nums1[idx1][0] < nums2[idx2][0] {
			res = append(res, nums1[idx1])
			idx1++
		} else if nums2[idx2][0] < nums1[idx1][0] {
			res = append(res, nums2[idx2])
			idx2++
		} else {
			res = append(res, []int{
				nums1[idx1][0],
				nums1[idx1][1] + nums2[idx2][1],
			})
			idx1++
			idx2++
		}
	}
	for ; idx1 < n1; idx1++ {
		res = append(res, nums1[idx1])
	}
	for ; idx2 < n2; idx2++ {
		res = append(res, nums2[idx2])
	}
	return res
}
