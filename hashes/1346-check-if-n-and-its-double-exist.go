package hashes

// https://leetcode.com/problems/check-if-n-and-its-double-exist/

// Approach #3: Hash Map
// Time: O(n)
// Space: O(n)
func checkIfExist(arr []int) bool {
	seen := make(map[int]bool)
	for _, num := range arr {
		if seen[2*num] || num&1 == 0 && seen[num/2] {
			return true
		}
		seen[num] = true
	}
	return false
}

// // Approach #2: Sorting + Binary Search
// // Time: O(nlogn)
// // Space: O(1)
// func checkIfExist(arr []int) bool {
// 	sort.Ints(arr)
// 	var n = len(arr)
// 	var left, right, mid int
// 	for idx, num := range arr {
// 		num += num
// 		left, right = 0, n-1
// 		for left <= right {
// 			mid = left + (right-left)/2
// 			if arr[mid] < num {
// 				left = mid + 1
// 			} else if arr[mid] > num {
// 				right = mid - 1
// 			} else if mid != idx {
// 				return true
// 			} else {
// 				break
// 			}
// 		}
// 	}
// 	return false
// }

// // Approach #1: Brute-Force
// // Time: O(nn)
// // Space: O(1)
// func checkIfExist(arr []int) bool {
// 	for i := range arr {
// 		for j := range arr {
// 			if j != i && arr[i] == 2*arr[j] {
// 				return true
// 			}
// 		}
// 	}
// 	return false
// }
