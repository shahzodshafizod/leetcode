package linkedlists

// https://leetcode.com/problems/remove-nodes-from-linked-list/

// Approach: Reverse Twice
// time: O(3n) = O(n)
// space: O(1)
func removeNodes(head *ListNode) *ListNode {
	var reverse = func(head *ListNode) *ListNode {
		var prev, next *ListNode = nil, nil
		var curr = head
		for curr != nil {
			next = curr.Next
			curr.Next = prev
			prev = curr
			head = curr
			curr = next
		}
		return head
	}
	head = reverse(head)
	var max = 0
	var prev *ListNode = nil
	for node := head; node != nil; node = node.Next {
		if node.Val < max {
			if prev == nil {
				head = node.Next
			} else {
				prev.Next = node.Next
			}
		} else {
			max = node.Val
			prev = node
		}
	}
	return reverse(head)
}

// // Approach: Stack
// // time: O(2n)
// // space: O(n)
// func removeNodes(head *ListNode) *ListNode {
// 	var stack = make([]*ListNode, 0)
// 	var node *ListNode
// 	for node = head; node != nil; node = node.Next {
// 		stack = append(stack, node)
// 	}
// 	var size = len(stack)
// 	var max = 0
// 	for size > 0 {
// 		node = stack[size-1]
// 		size--
// 		if node.Val < max {
// 			if size == 0 {
// 				head = node.Next
// 			} else {
// 				stack[size-1].Next = node.Next
// 			}
// 		} else {
// 			max = node.Val
// 		}
// 	}
// 	return head
// }

// // Approach: Monotonic Stack
// // time: O(n)
// // space: O(n) # for recursion stack
// func removeNodes(head *ListNode) *ListNode {
// 	if head == nil {
// 		return nil
// 	}
// 	head.Next = removeNodes(head.Next)
// 	if head.Next != nil && head.Val < head.Next.Val {
// 		return head.Next
// 	}
// 	return head
// }

// // Approach: Monotonic Stack
// // time: O(n)
// // space: O(n) # for recursion stack
// func removeNodes(head *ListNode) *ListNode {
// 	var monostack func(prev *ListNode, node *ListNode) int
// 	monostack = func(prev *ListNode, node *ListNode) int {
// 		if node == nil {
// 			return 0
// 		}
// 		var max = monostack(node, node.Next)
// 		if node.Val < max {
// 			if prev == nil {
// 				head = node.Next
// 			} else {
// 				prev.Next = node.Next
// 			}
// 			return max
// 		}
// 		return node.Val
// 	}
// 	monostack(nil, head)
// 	return head
// }
