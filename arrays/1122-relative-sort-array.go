package arrays

import "slices"

// https://leetcode.com/problems/relative-sort-array/

// n = len(arr1)
// m = len(arr2)
// k = max(arr1)
// time: O(n+m+k)
// space: O(k)
func relativeSortArray(arr1 []int, arr2 []int) []int {
	var count = make([]int, slices.Max(arr1)+1)
	for _, num := range arr1 {
		count[num]++
	}
	var result = make([]int, 0, len(arr1))
	for _, num := range arr2 {
		for count[num] > 0 {
			result = append(result, num)
			count[num]--
		}
	}
	for num, count := range count {
		for count > 0 {
			count--
			result = append(result, num)
		}
	}
	return result
}

// // n = len(arr1)
// // m = len(arr2)
// // time: O(n log n)
// // space: O(m)
// func relativeSortArray(arr1 []int, arr2 []int) []int {
// 	var indexes = make(map[int]int)
// 	for idx, num := range arr2 {
// 		indexes[num] = idx
// 	}
// 	sort.Slice(arr1, func(i, j int) bool { // O(n log n)
// 		_, oki := indexes[arr1[i]]
// 		_, okj := indexes[arr1[j]]
// 		if oki && okj {
// 			return indexes[arr1[i]] < indexes[arr1[j]]
// 		} else if oki {
// 			return true
// 		} else if okj {
// 			return false
// 		} else {
// 			return arr1[i] < arr1[j]
// 		}
// 	})
// 	return arr1
// }

// // n = len(arr1)
// // m = len(arr2)
// // time: O(m + nlogn)
// // space: O(m + n)
// func relativeSortArray(arr1 []int, arr2 []int) []int {
// 	var freq = make(map[int]int)
// 	for _, num := range arr2 {
// 		freq[num] = 0
// 	}
// 	var remaining = make([]int, 0)
// 	for _, num := range arr1 {
// 		if _, exists := freq[num]; exists {
// 			freq[num]++
// 		} else {
// 			remaining = append(remaining, num)
// 		}
// 	}
// 	sort.Ints(remaining)
// 	var result = make([]int, 0, len(arr1))
// 	for _, num := range arr2 {
// 		for freq[num] > 0 {
// 			result = append(result, num)
// 			freq[num]--
// 		}
// 	}
// 	result = append(result, remaining...)
// 	return result
// }
