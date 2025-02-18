package stacks

// https://leetcode.com/problems/construct-smallest-number-from-di-string/

// Approach #2: Stack
// Time: O(n)
// Space: O(n)
func smallestNumber(pattern string) string {
	var n = len(pattern)
	var stack = make([]int, n+1)
	var size = 0
	var num = make([]byte, 0, n+1)
	for idx := 0; idx <= n; idx++ {
		stack[size] = idx + 1
		size++
		if idx == n || pattern[idx] == 'I' {
			for size > 0 {
				num = append(num, byte('0'+stack[size-1]))
				size--
			}
		}
	}
	return string(num)
}

// // Approach #1: Backtracking
// // Time: O(n^n)
// // Space: O(n)
// func smallestNumber(pattern string) string {
// 	var n = len(pattern)
// 	var nums = make([]byte, n+1)
// 	var used = make(map[byte]bool, 10)
// 	var backtrack func(int) bool
// 	backtrack = func(idx int) bool {
// 		if idx == n {
// 			return true
// 		}
// 		var num byte
// 		for num = '1'; num <= '9'; num++ {
// 			if used[num] {
// 				continue
// 			}
// 			if pattern[idx] == 'I' && nums[idx] > num {
// 				continue
// 			}
// 			if pattern[idx] == 'D' && nums[idx] < num {
// 				continue
// 			}
// 			nums[idx+1] = num
// 			used[num] = true
// 			if backtrack(idx + 1) {
// 				return true
// 			}
// 			used[num] = false
// 		}
// 		return false
// 	}
// 	var num byte
// 	for num = '1'; num <= '9'; num++ {
// 		nums[0] = num
// 		used[num] = true
// 		if backtrack(0) {
// 			break
// 		}
// 		used[num] = false
// 	}
// 	return string(nums)
// }
