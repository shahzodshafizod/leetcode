package linkedlists

import "github.com/shahzodshafizod/leetcode/design"

/*
Doubly Linked Lists

Problem:
Given a doubly linked list, list nodes also have a child property that
can point to a separate doubly linked list. These child lists can also
have one or more child doubly linked lists of their own, and so on.
Return the list as a single level flattened doubly linked list.

Step 1: Verify the constraints
	- Can a doubly linked list have multiple child list nodes?
		: Yes, every list node can have multiple levels of children.
	- What do we do with child properties after flattening?
		: Just set the child property value to null.
Step 2: Write out some test cases
	- null <- (1) <-> (2) <-> (3) <-> (4) <-> (5) <-> (6) -> null
	                   |                       |
	                   v                       v
	                  (7) <-> (8) <-> (9)     (12)<-> (13)
	                           |
	                           v
	                          (10)<-> (11)
		: 1, 2, [7, 8, [10, 11], 9], 3, 4, 5, [12, 13], 6
*/

// https://leetcode.com/problems/flatten-a-multilevel-doubly-linked-list/

func flatten(root *design.DListNode) *design.DListNode {
	for node := root; node != nil; {
		nextNode := node.Next
		if node.Child != nil {
			tail := node.Child
			nextNode = tail
			for ; tail.Next != nil; tail = tail.Next {
				if tail.Child == nil && tail == nextNode {
					nextNode = tail
				}
			}
			tail.Next = node.Next
			if tail.Next != nil {
				tail.Next.Prev = tail
			}
			node.Next = node.Child
			node.Child.Prev = node
			node.Child = nil
		}
		node = nextNode
	}
	return root
}

// func flatten(root *design.DListNode) *design.DListNode {
// 	flattenChild(root, nil, nil)
// 	return root
// }

// func flattenChild(root *design.DListNode, prev *design.DListNode, next *design.DListNode) *design.DListNode {
// 	if root == nil {
// 		return nil
// 	}
// 	node := root
// 	for node != nil {
// 		if node.Child != nil {
// 			nextNode := flattenChild(node.Child, node, node.Next)
// 			node.Child = nil
// 			node = nextNode
// 		}
// 		if node.Next == nil {
// 			break
// 		}
// 		node = node.Next
// 	}
// 	root.Prev = prev
// 	if prev != nil {
// 		prev.Next = root
// 	}
// 	node.Next = next
// 	if next != nil {
// 		next.Prev = node
// 	}
// 	return node
// }
