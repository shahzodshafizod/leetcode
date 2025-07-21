package hashes

// https://leetcode.com/problems/intersection-of-two-arrays/

// time: O(n+m)
// space: O(n)
func intersection(nums1 []int, nums2 []int) []int {
	hashset := make(map[int]bool)
	for _, num := range nums1 {
		hashset[num] = true
	}
	common := make([]int, 0)
	for _, num := range nums2 {
		if hashset[num] {
			common = append(common, num)
			hashset[num] = false
		}
	}
	return common
}
