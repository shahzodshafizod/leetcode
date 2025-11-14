package monotonic

import "github.com/shahzodshafizod/leetcode/pkg"

// https://leetcode.com/problems/final-prices-with-a-special-discount-in-a-shop/

// Approach #3: Monotonic Increasing Stack
// Time: O(n)
// Space: O(n)
func finalPrices(prices []int) []int {
	var top *pkg.ListNode // stack top

	n := len(prices)

	answer := make([]int, n)
	for idx, price := range prices {
		answer[idx] = price
		for top != nil && answer[top.Val] >= price {
			answer[top.Val] -= price
			top = top.Next
		}

		top = &pkg.ListNode{Val: idx, Next: top}
	}

	return answer
}

// // Approach #2: Monotonic Increasing Stack + Binary Search
// // Time: O(nlogn)
// // Space: O(n)
// func finalPrices(prices []int) []int {
// 	var stack = make([]int, 0)
// 	var size = 0
// 	var n = len(prices)
// 	var answer = make([]int, n)
// 	var left, right, mid int
// 	for idx := n - 1; idx >= 0; idx-- {
// 		answer[idx] = prices[idx]
// 		left, right = 0, size-1
// 		for left <= right {
// 			mid = left + (right-left)/2
// 			if stack[mid] > prices[idx] {
// 				right = mid - 1
// 			} else {
// 				left = mid + 1
// 			}
// 		}
// 		if right >= 0 {
// 			answer[idx] -= stack[right]
// 		}
// 		for size > 0 && stack[size-1] >= prices[idx] {
// 			stack = stack[:size-1]
// 			size--
// 		}
// 		stack = append(stack, prices[idx])
// 		size++
// 	}
// 	return answer
// }

// // Approach #1: Brute Force
// // Time: O(nn)
// // Space: O(1)
// func finalPrices(prices []int) []int {
// 	var n = len(prices)
// 	var answer = make([]int, n)
// 	for idx := range answer {
// 		answer[idx] = prices[idx]
// 		for j := idx + 1; j < n; j++ {
// 			if prices[j] <= prices[idx] {
// 				answer[idx] -= prices[j]
// 				break
// 			}
// 		}
// 	}
// 	return answer
// }
