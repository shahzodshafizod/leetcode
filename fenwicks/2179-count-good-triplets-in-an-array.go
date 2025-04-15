package fenwicks

// https://leetcode.com/problems/count-good-triplets-in-an-array/

// Approach: Fenwick Tree (Binary Indexed Tree)
// Time: O(nlogn)
// Space: O(n)
func goodTriplets(nums1 []int, nums2 []int) int64 {
	var n = len(nums1)
	var tree = make([]int, n+2) // 1-indexed
	var update = func(num int) {
		for num < n+2 {
			tree[num]++
			num += num & -num
		}
	}
	var query = func(num int) int {
		var total = 0
		for num > 0 {
			total += tree[num]
			num -= num & -num
		}
		return total
	}
	var indices2 = make([]int, n)
	for idx := 0; idx < n; idx++ {
		indices2[nums2[idx]] = idx
	}
	var total int64 = 0
	var idx2, left, right int
	for idx1 := 0; idx1 < n; idx1++ {
		idx2 = indices2[nums1[idx1]]
		left = query(idx2 + 1)
		right = (n - idx2 - 1) - (idx1 - left)
		total += int64(left * right)
		update(idx2 + 1)
	}
	return total
}
