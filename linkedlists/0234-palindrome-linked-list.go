package linkedlists

import "github.com/shahzodshafizod/leetcode/design"

// https://leetcode.com/problems/palindrome-linked-list/

// time: O(n)
// space: O(1)
// modifying
func isPalindrome(head *design.ListNode) bool {
	tortoise, hare := head, head
	// 1. find middle (prev,next) and reverse the first half
	var prev, next *design.ListNode
	for hare != nil && hare.Next != nil {
		hare = hare.Next.Next
		next = tortoise.Next
		tortoise.Next = prev
		prev = tortoise
		tortoise = next
	}
	if hare != nil && next != nil {
		next = next.Next
	}
	// 2. check palindrome (prev: left side, next: right side)
	for prev != nil {
		if prev.Val != next.Val {
			return false
		}
		prev = prev.Next
		next = next.Next
	}
	return true
}

// // time: O(n)
// // space: O(n) (O(n/2))
// // NOT modifying
// func isPalindrome(head *design.ListNode) bool {
// 	var check func(left *design.ListNode, last *design.ListNode) *design.ListNode
// 	check = func(left *design.ListNode, last *design.ListNode) *design.ListNode {
// 		if last == nil || last.Next == nil {
// 			if last != nil && last != left {
// 				left = left.Next
// 			}
// 			return left
// 		}
// 		right := check(left.Next, last.Next.Next)
// 		if right == nil || right.Val != left.Val {
// 			return nil
// 		}
// 		if right.Next != nil {
// 			right = right.Next
// 		}
// 		return right
// 	}
// 	return check(head, head) != nil
// }
