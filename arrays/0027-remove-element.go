package arrays

// https://leetcode.com/problems/remove-element/

func removeElement(nums []int, val int) int {
	index := -1

	for _, num := range nums {
		if num != val {
			index++
			nums[index] = num
		}
	}

	return index + 1
}

// func removeElement(nums []int, val int) int {
// 	return len(slices.DeleteFunc[[]int](nums, func(i int) bool { return i == val }))
// }

// func removeElement(nums []int, val int) int {
// 	var len = len(nums)
// 	var occurrences = 0
// 	for idx, num := range nums {
// 		if num == val {
// 			occurrences++
// 		} else if occurrences > 0 {
// 			nums[idx-occurrences] = num
// 		}
// 	}
// 	return len - occurrences
// }

// // moving all vals to the end using two pointers
// func removeElement(nums []int, val int) int {
// 	var right = len(nums) - 1
// 	var moveRight = func() {
// 		for right >= 0 && nums[right] == val {
// 			right--
// 		}
// 	}
// 	moveRight()
// 	for left := 0; left < right; left++ {
// 		if nums[left] == val {
// 			nums[left], nums[right] = nums[right], nums[left]
// 			moveRight()
// 		}
// 	}
// 	return right + 1
// }
