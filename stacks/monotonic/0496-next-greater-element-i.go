package monotonic

// https://leetcode.com/problems/next-greater-element-i/

// time: O(2n) = O(n) // n = len(nums1)
// space: O(m) // m = len(nums2)
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	var idmap = make(map[int]int)
	for idx, num := range nums1 {
		idmap[num] = idx
		nums1[idx] = -1
	}
	var size = 0
	var stack = make([]int, size) // monotonic stack
	var curr, ptrid int
	var exists bool
	for idx := len(nums2) - 1; idx >= 0; idx-- {
		curr = nums2[idx]
		for size > 0 && curr > stack[size-1] { // O(2n)
			stack = stack[:size-1]
			size--
		}
		if ptrid, exists = idmap[curr]; exists && size > 0 {
			nums1[ptrid] = stack[size-1]
		}
		stack = append(stack, curr)
		size++
	}
	return nums1
}

// // time: O(n^2)
// // space: O(1)
// func nextGreaterElement(nums1 []int, nums2 []int) []int {
// 	for idx, num1 := range nums1 {
// 		var found = false
// 		nums1[idx] = -1
// 		for _, num2 := range nums2 {
// 			if num2 == num1 {
// 				found = true
// 			} else if found && num2 > num1 {
// 				nums1[idx] = num2
// 				break
// 			}
// 		}
// 	}
// 	return nums1
// }
