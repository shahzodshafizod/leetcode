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

// // time: O(n log n)
// // space: O(1)
// func specialArray(nums []int) int {
// 	sort.Ints(nums)
// 	var n = len(nums)
// 	var prev = -1
// 	var x int
// 	for i := 0; i < n; i++ {
// 		x = n - i
// 		if nums[i] >= x && prev < x {
// 			return x
// 		}
// 		// skip duplicates, ex: [2,2,2,3,3,3,3,3]
// 		for i+1 < n && nums[i] == nums[i+1] {
// 			i++
// 		}
// 		prev = nums[i]
// 	}
// 	return -1
// }

// // time: O(2 * n log n) = O(n log n)
// // space: O(1)
// func specialArray(nums []int) int {
// 	sort.Ints(nums) // O(n log n)
// 	var left, right = 1, len(nums)
// 	var mid, count int
// 	for left <= right { // O(log n)
// 		mid = (left + right) / 2
// 		count = 0
// 		for _, num := range nums { // O(n)
// 			if num >= mid {
// 				count++
// 			}
// 		}
// 		if count == mid {
// 			return mid
// 		} else if count > mid {
// 			left = mid + 1
// 		} else {
// 			right = mid - 1
// 		}
// 	}
// 	return -1
// }
