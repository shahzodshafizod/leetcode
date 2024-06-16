package greedy

/*
1: 1
2: 1,2,3
3: 1,2,3,4,5,6
4: 1,2,3,4,5,6,7,8,9,10
5: 1,2,3,4,5,6,7,8,9,10,11,12,13,14,15
6: 1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21
*/

// https://leetcode.com/problems/patching-array/

// time: O(LogN)
// space: O(1)
func minPatches(nums []int, n int) int {
	var patches, maxPatch = 0, 0
	var patch int
	for idx := 0; maxPatch < n; { // O(log n)
		patch = maxPatch + 1
		if idx < len(nums) && patch >= nums[idx] {
			maxPatch += nums[idx]
			idx++
		} else {
			maxPatch += patch
			patches++
		}
	}
	return patches
}

// // time: O(N * N * Len)
// // space: O(1)
// func minPatches(nums []int, n int) int {
// 	var num int
// 	var patches = 0
// 	for n > 0 { // O(N)
// 		num = n
// 		n--
// 		for idx := len(nums) - 1; idx >= 0 && num > 0; idx-- { // O(Len)
// 			if num >= nums[idx] {
// 				num -= nums[idx]
// 			}
// 		}
// 		if num > 0 {
// 			nums = append(nums, num)                   // O(N)
// 			for idx := len(nums) - 1; idx > 1; idx-- { // O(N)
// 				if nums[idx-1] > nums[idx] {
// 					nums[idx-1], nums[idx] = nums[idx], nums[idx-1]
// 				}
// 			}
// 			patches++
// 		}
// 	}
// 	return patches
// }
