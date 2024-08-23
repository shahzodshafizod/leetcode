package maths

// https://leetcode.com/problems/permutation-sequence/

// Time: O(n)
// Space: O(n)
func getPermutation(n int, k int) string {
	var nums = make([]int, n)
	var factorial = 1
	for idx := 1; idx <= n; idx++ {
		nums[idx-1] = idx
		factorial *= idx
	}
	k--
	var permutation = make([]byte, 0, n)
	for n > 0 {
		factorial /= n
		idx := k / factorial
		k %= factorial
		permutation = append(permutation, byte('0'+nums[idx]))
		nums = append(nums[:idx], nums[idx+1:]...)
		n--
	}
	return string(permutation)
}

// // Time: O(n + k)
// // Space: O(n)
// func getPermutation(n int, k int) string {
// 	var nums = list.New()
// 	var factorial = 1
// 	for num := 1; num <= n; num++ {
// 		nums.PushBack(num)
// 		factorial *= num
// 	}
// 	var permutation = make([]byte, 0, n)
// 	for num := n; num > 0; num-- {
// 		factorial /= num
// 		e := nums.Front()
// 		for k > factorial {
// 			e = e.Next()
// 			k -= factorial
// 		}
// 		permutation = append(permutation, byte('0'+e.Value.(int)))
// 		nums.Remove(e)
// 	}
// 	return string(permutation)
// }

// // Solved with the help of this problem:
// // 31. Next Permutation (https://leetcode.com/problems/next-permutation/)
// // Time: O(n x k)
// // Space: O(n)
// func getPermutation(n int, k int) string {
// 	var nextPermutation = func(nums []byte) {
// 		// find the starting point
// 		var pivot = n - 2
// 		for pivot >= 0 && nums[pivot] > nums[pivot+1] {
// 			pivot--
// 		}
// 		// swap pivot with the next element
// 		if pivot >= 0 {
// 			successor := n - 1
// 			for nums[successor] < nums[pivot] {
// 				successor--
// 			}
// 			nums[successor], nums[pivot] = nums[pivot], nums[successor]
// 		}
// 		// reverse the tail
// 		var left, right = pivot + 1, n - 1
// 		for left < right && nums[left] > nums[right] {
// 			nums[left], nums[right] = nums[right], nums[left]
// 			left++
// 			right--
// 		}
// 	}

// 	var nums = make([]byte, n)
// 	for idx := range nums {
// 		nums[idx] = byte('0' + idx + 1)
// 	}
// 	for k > 1 {
// 		k--
// 		nextPermutation(nums)
// 	}
// 	return string(nums)
// }
