package fenwicks

// https://leetcode.com/problems/count-good-triplets-in-an-array/

// Approach: Fenwick Tree (Binary Indexed Tree)
// Time: O(nlogn)
// Space: O(n)
func goodTriplets(nums1 []int, nums2 []int) int64 {
	n := len(nums1)
	tree := make([]int, n+2) // 1-indexed
	update := func(num int) {
		for num < n+2 {
			tree[num]++
			num += num & -num
		}
	}
	query := func(num int) int {
		total := 0
		for num > 0 {
			total += tree[num]
			num -= num & -num
		}

		return total
	}

	indices2 := make([]int, n)
	for idx := range n {
		indices2[nums2[idx]] = idx
	}

	var total int64

	var idx2, left, right int
	for idx1 := range n {
		idx2 = indices2[nums1[idx1]]
		left = query(idx2 + 1)
		right = (n - idx2 - 1) - (idx1 - left)
		total += int64(left * right)

		update(idx2 + 1)
	}

	return total
}
