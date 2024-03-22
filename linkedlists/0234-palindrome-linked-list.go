package linkedlists

// https://leetcode.com/problems/palindrome-linked-list/

// time: O(n)
// space: O(1)
// modifying
func isPalindrome(head *ListNode) bool {
	tortoise, hare := head, head
	// 1. find middle (prev,next) and reverse the first half
	var prev, next *ListNode
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
// // space: O(n/2) = O(n)
// // NOT modifying
// func isPalindrome(head *ListNode) bool {
// 	var recur func(tortoise *ListNode, hare *ListNode) (*ListNode, bool)
// 	recur = func(tortoise *ListNode, hare *ListNode) (*ListNode, bool) {
// 		if hare == nil || hare.Next == nil {
// 			if hare != nil {
// 				tortoise = tortoise.Next
// 			}
// 			return tortoise, true
// 		}
// 		next, _ := recur(tortoise.Next, hare.Next.Next)
// 		if next == nil || next.Val != tortoise.Val {
// 			return nil, false
// 		}
// 		return next.Next, true
// 	}
// 	_, ok := recur(head, head)
// 	return ok
// }
