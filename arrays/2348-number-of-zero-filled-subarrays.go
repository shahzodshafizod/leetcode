package arrays

// https://leetcode.com/problems/number-of-zero-filled-subarrays/

// Approach: Two-Pointers
func zeroFilledSubarray(nums []int) int64 {
	var (
		res   int64
		start int
	)

	for end := range nums {
		if nums[end] != 0 {
			start = end + 1
		}

		res += int64(end - start + 1)
	}

	return res
}

// func zeroFilledSubarray(nums []int) int64 {
// 	var res, cnt int64

// 	for _, num := range nums {
// 		if num == 0 {
// 			cnt++
// 			res += cnt
// 		} else {
// 			cnt = 0
// 		}
// 	}

// 	return res
// }
