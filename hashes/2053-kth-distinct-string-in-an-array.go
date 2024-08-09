package hashes

// https://leetcode.com/problems/kth-distinct-string-in-an-array/

// Approach: Hash Map
// Time: O(N)
// Space: O(N)
func kthDistinct(arr []string, k int) string {
	var count = make(map[string]int)
	for _, s := range arr {
		count[s]++
	}
	for _, s := range arr {
		if count[s] == 1 {
			k--
			if k == 0 {
				return s
			}
		}
	}
	return ""
}

// // Approach: Brute Force
// // Time: O(N^2)
// // Space: O(1)
// func kthDistinct(arr []string, k int) string {
// 	var distinct bool
// 	for idx := range arr {
// 		distinct = true
// 		for jdx := range arr {
// 			if idx != jdx && arr[idx] == arr[jdx] {
// 				distinct = false
// 				break
// 			}
// 		}
// 		if distinct {
// 			k--
// 			if k == 0 {
// 				return arr[idx]
// 			}
// 		}
// 	}
// 	return ""
// }
