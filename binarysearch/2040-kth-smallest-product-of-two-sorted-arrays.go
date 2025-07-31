package binarysearch

// https://leetcode.com/problems/kth-smallest-product-of-two-sorted-arrays/

func kthSmallestProduct(nums1 []int, nums2 []int, k int64) int64 {
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}

	m := len(nums2)
	calcLessThan := func(target int64) int64 {
		var count int64

		var left, right, mid int

		var product int64

		for _, n1 := range nums1 {
			left, right = 0, m-1
			for left <= right {
				mid = left + (right-left)/2
				product = int64(n1 * nums2[mid])

				if n1 >= 0 && product <= target || n1 < 0 && product > target {
					left = mid + 1
				} else {
					right = mid - 1
				}
			}

			if n1 >= 0 {
				count += int64(left)
			} else {
				count += int64(m - left)
			}
		}

		return count
	}

	left, right := -int64(1e10), int64(1e10)
	for left <= right {
		mid := left + (right-left)/2
		if calcLessThan(mid) >= k {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return left
}
