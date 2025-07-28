package hashes

// https://leetcode.com/problems/make-two-arrays-equal-by-reversing-subarrays/

// Approach #2: Array
// Time: O(N)
// Space: O(N)
func canBeEqual(target []int, arr []int) bool {
	var count [1001]int
	for idx := range arr {
		count[arr[idx]]++
		count[target[idx]]--
	}

	for idx := range arr {
		if count[target[idx]] != 0 || count[arr[idx]] != 0 {
			return false
		}
	}

	return true
}

// // Approach #2: Hash Table
// // Time: O(N)
// // Space: O(N)
// func canBeEqual(target []int, arr []int) bool {
// 	var count = make(map[int]int)
// 	for idx := range arr {
// 		count[arr[idx]]++
// 		count[target[idx]]--
// 	}
// 	for _, count := range count {
// 		if count != 0 {
// 			return false
// 		}
// 	}
// 	return true
// }

// // Approach #1: Sorting
// // Time: O(N Log N)
// // Space: O(1)
// func canBeEqual(target []int, arr []int) bool {
// 	sort.Ints(target)
// 	sort.Ints(arr)
// 	for idx := range arr {
// 		if arr[idx] != target[idx] {
// 			return false
// 		}
// 	}
// 	return true
// }
