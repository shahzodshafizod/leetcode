package arrays

// https://leetcode.com/problems/shuffle-the-array/

func shuffle(nums []int, n int) []int {
	var shuffled = make([]int, 0, 2*n)
	for i := 0; i < n; i++ {
		shuffled = append(shuffled, nums[i], nums[n+i])
	}
	return shuffled
}

// // in-place, works for: 0<=nums[i]<=10^3(1000)
// // (1000)10 = (1111101000)2
// func shuffle(nums []int, n int) []int {

// 	// step 1: copy left part to right part
// 	// was:
// 	// 1, 2, 3, 4, 5, 6, 7, 8
// 	// became:
// 	// 0, 0, 0, 0, 5, 6, 7, 8
// 	// 1, 2, 3, 4, 1, 2, 3, 4

// 	// step 2:
// 	// 1, 5, 2, 6, 3, 7, 4, 8

// 	// bits it shouldn't be less than 10, becase the maximal value of the array 1000 has 10 bits
// 	var bits = 10
// 	var length = 2 * n
// 	for i := n; i < length; i++ {
// 		nums[i] = nums[i]<<bits | nums[i-n]
// 	}
// 	for i, j := 0, n; j < length; i, j = i+2, j+1 {
// 		nums[i] = nums[j] & (1<<bits - 1)
// 		nums[i+1] = nums[j] >> bits
// 	}
// 	return nums
// }
