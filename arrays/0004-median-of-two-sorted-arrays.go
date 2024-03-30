package arrays

// https://leetcode.com/problems/median-of-two-sorted-arrays/

// time: O((m+n)/2) = O(m+n)
// space: O(1)
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var len1, len2 = len(nums1), len(nums2)
	var total = len1 + len2
	var half = total / 2
	var idx, idx1, idx2 = 0, 0, 0
	var median, prev int
	for idx <= half && (idx1 < len1 || idx2 < len2) {
		idx++
		prev = median
		switch {
		case idx1 == len1:
			median = nums2[idx2]
			idx2++
		case idx2 == len2:
			median = nums1[idx1]
			idx1++
		case nums1[idx1] < nums2[idx2]:
			median = nums1[idx1]
			idx1++
		default:
			median = nums2[idx2]
			idx2++
		}
	}
	if total&1 == 0 {
		return float64(prev+median) / 2
	}
	return float64(median)
}

// // time: O(log(min(m,n)))
// // space: O(log(min(m,n))) calling a function
// func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
// 	var getMiddle = func(nums []int, mid int, len int) (int, int) {
// 		var curr, next int
// 		if mid >= 0 {
// 			curr = nums[mid]
// 		} else {
// 			curr = math.MinInt
// 		}
// 		if mid+1 < len {
// 			next = nums[mid+1]
// 		} else {
// 			next = math.MaxInt
// 		}
// 		return curr, next
// 	}
// 	var len1, len2 = len(nums1), len(nums2)
// 	if len1 > len2 {
// 		nums1, nums2 = nums2, nums1
// 		len1, len2 = len2, len1
// 	}
// 	var total = len1 + len2
// 	var half = total/2 + total&1
// 	var left, right = 1, len1
// 	var mid1, mid2, curr1, next1, curr2, next2 int
// 	var curr, next int
// 	for {
// 		mid1 = (left + right) / 2
// 		mid2 = half - mid1
// 		curr1, next1 = getMiddle(nums1, mid1-1, len1)
// 		curr2, next2 = getMiddle(nums2, mid2-1, len2)
// 		if curr1 <= next2 && curr2 <= next1 {
// 			if total&1 == 1 {
// 				curr = max(curr1, curr2)
// 				next = curr
// 			} else {
// 				curr, next = max(curr1, curr2), min(next1, next2)
// 			}
// 			break
// 		}
// 		if curr1 > next2 {
// 			right = mid1 - 1
// 		} else {
// 			left = mid1 + 1
// 		}
// 	}
// 	return float64(curr+next) / 2
// }
