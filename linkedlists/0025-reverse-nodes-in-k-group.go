package linkedlists

import "github.com/shahzodshafizod/leetcode/pkg"

// https://leetcode.com/problems/reverse-nodes-in-k-group/

func reverseKGroup(head *pkg.ListNode, k int) *pkg.ListNode {
	length := reverseKGroupLLLength(head, 0)
	node := head

	var prevGroup *pkg.ListNode = nil

	times := length / k
	for times > 0 {
		times--

		var localHead, localTrail *pkg.ListNode = nil, node

		count := 1
		for count <= k {
			count++
			// [1]-->[2]-->[3]-->[4]-->[5]-->nil
			// [1]-->nil, [2]-->[3]-->[4]-->[5]-->nil
			// [2]-->[1]-->nil, [3]-->[4]-->[5]-->nil
			// [3]-->[2]-->[1]-->nil, [4]-->[5]-->nil
			// [4]-->[3]-->[2]-->[1]-->nil, [5]-->nil
			current := node
			node = node.Next
			current.Next = localHead
			localHead = current
		}

		if prevGroup != nil {
			prevGroup.Next = localHead
		}

		prevGroup = localTrail

		localTrail.Next = node
		if head == localTrail {
			head = localHead
		}
	}

	return head
}

func reverseKGroupLLLength(node *pkg.ListNode, length int) int {
	if node == nil {
		return length
	}

	return reverseKGroupLLLength(node.Next, length+1)
}

// // If changing NEXT is not allowed, we'll change VAL
// func reverseKGroup(head *pkg.ListNode, k int) *pkg.ListNode {
// 	if reverseKGroupLLLength(head, 0) < k { // time: O(N), space: O(1) because of tail recursion
// 		return head
// 	}
// 	var stack = make([]int, k)
// 	for node := head; node != nil; {
// 		var counter = 0
// 		var pawn = node
// 		for counter < k && pawn != nil {
// 			stack[counter] = pawn.Val
// 			counter++
// 			pawn = pawn.Next
// 		}
// 		if counter < k {
// 			break
// 		}
// 		counter--
// 		for counter >= 0 {
// 			node.Val = stack[counter]
// 			counter--
// 			node = node.Next
// 		}
// 	}
// 	return head
// }

// // Follow-up: Can you solve the problem in O(1) extra memory space?
// // But with time of O(N^2)
// func reverseKGroup(head *pkg.ListNode, k int) *pkg.ListNode {
// 	for node := head; node != nil; {
// 		var next *pkg.ListNode = nil
// 		var counter = 1
// 		for counter < k {
// 			var pawn = node
// 			var i = counter
// 			for i < k && pawn != nil {
// 				i++
// 				pawn = pawn.Next
// 			}
// 			if pawn == nil {
// 				break
// 			}
// 			if next == nil {
// 				next = pawn
// 			}
// 			node.Val, pawn.Val = pawn.Val, node.Val
// 			node = node.Next
// 			counter += 2
// 		}
// 		if next == nil {
// 			next = node
// 		}
// 		node = next.Next
// 	}
// 	return head
// }
