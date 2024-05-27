package arrays

// https://leetcode.com/problems/special-array-with-x-elements-greater-than-or-equal-x/

// time: O(n)
// space: O(n)
func specialArray(nums []int) int {
	var n = len(nums)
	var freq = make([]int, n+1)
	for _, num := range nums {
		freq[min(n, num)]++
	}
	var count = 0
	for x := n; x >= 0; x-- {
		count += freq[x]
		if x == count {
			return x
		}
	}
	return -1
}

// // time: O(1001)
// // space: O(1001)
// func specialArray(nums []int) int {
// 	const limit = 1001
// 	var freq [limit]int
// 	for _, num := range nums {
// 		freq[num]++
// 	}
// 	var count = len(nums)
// 	for x := 0; x < limit; x++ {
// 		if count == x {
// 			return x
// 		}
// 		count -= freq[x]
// 	}
// 	return -1
// }

// // brute-force
// // time: O(n ^ 2)
// func specialArray(nums []int) int {
// 	var max = -1
// 	for _, num := range nums {
// 		if num > max {
// 			max = num
// 		}
// 	}
// 	for x := 1; x <= max; x++ {
// 		var count = 0
// 		for _, num := range nums {
// 			if num >= x {
// 				count++
// 			}
// 		}
// 		if count == x {
// 			return x
// 		}
// 	}
// 	return -1
// }
