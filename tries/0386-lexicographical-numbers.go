package tries

// https://leetcode.com/problems/lexicographical-numbers/

// Approach #2: Iterative
// Time: O(n)
// Space: O(1)
func lexicalOrder(n int) []int {
	nums := make([]int, n)
	num := 1
	for idx := 0; idx < n; idx++ {
		nums[idx] = num
		if num*10 <= n {
			num *= 10
		} else {
			for num == n || num%10 == 9 {
				num /= 10
			}
			num++
		}
	}
	return nums
}

// // Approach #1: Depth-First Search
// // Time: O(n)
// // Space: O(log(10)n), for recursion stack
// func lexicalOrder(n int) []int {
// 	var nums = make([]int, 0, n)
// 	var dfs func(num int)
// 	dfs = func(num int) {
// 		if num > n {
// 			return
// 		}
// 		nums = append(nums, num)
// 		for k := 0; k < 10; k++ {
// 			dfs(num*10 + k)
// 		}
// 	}
// 	for num := 1; num <= 9; num++ {
// 		dfs(num)
// 	}
// 	return nums
// }
