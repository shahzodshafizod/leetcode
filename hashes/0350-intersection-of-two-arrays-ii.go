package hashes

/*
Follow-up Question #1:
	What if the given array is already sorted? How would you optimize your algorithm?
Answer:
	Approach #2 is the best choice since we skip the cost of sorting.
	So time complexity is O(M+N) and the space complexity is O(1).

Follow-up Question #2:
	What if nums1's size is small compared to nums2's size? Which algorithm is better?
Answer:
	Approach #1 is the best choice.
	Time complexity is O(M+N) and the space complexity is O(M),
	where M is length of nums1, N is length of nums2.

Follow-up Question #3:
	What if elements of nums2 are stored on disk, and the memory is limited
	such that you cannot load all elements into the memory at once?
Answer:
	Binary Search???
*/

// https://leetcode.com/problems/intersection-of-two-arrays-ii/

// Approach #1: HashMap
// time: O(max(n1,n2))
// space: O(min(n1,n2))
func intersect(nums1 []int, nums2 []int) []int {
	if len(nums1) > len(nums2) {
		return intersect(nums2, nums1)
	}

	set := make(map[int]int)
	for _, num := range nums1 {
		set[num]++
	}

	intersection := make([]int, 0)

	for _, num := range nums2 {
		if set[num] > 0 {
			intersection = append(intersection, num)
			set[num]--
		}
	}

	return intersection
}

// // Approach #2: Sorting
// // time: O(n log n)
// // space: O(1)
// func intersect(nums1 []int, nums2 []int) []int {
// 	sort.Ints(nums1)
// 	sort.Ints(nums2)
// 	var n1, n2 = len(nums1), len(nums2)
// 	var i1, i2 = 0, 0
// 	var intersection = make([]int, 0)
// 	for i1 < n1 && i2 < n2 {
// 		if nums1[i1] < nums2[i2] {
// 			i1++
// 		} else if nums2[i2] < nums1[i1] {
// 			i2++
// 		} else {
// 			intersection = append(intersection, nums1[i1])
// 			i1++
// 			i2++
// 		}
// 	}
// 	return intersection
// }
