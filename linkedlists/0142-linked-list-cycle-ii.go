package linkedlists

import "github.com/shahzodshafizod/alkhwarizmi/design"

/*
Linked Lists: Cycle Detection.

Floyd's Tortoise and Hare (Sangpusht va Xargushi Floyd)
tortoise=1_step, hare=2_steps.
*/

// https://leetcode.com/problems/linked-list-cycle-ii/

func detectCycle(head *design.ListNode) *design.ListNode {
	tortoise, hare := head, head
	for hare != nil && hare.Next != nil {
		tortoise = tortoise.Next
		hare = hare.Next.Next
		if tortoise == hare {
			tortoise = head
			for tortoise != hare {
				tortoise = tortoise.Next
				hare = hare.Next
			}
			return tortoise
		}
	}
	return nil
}

// func detectCycle(head *design.ListNode) *design.ListNode {
// 	if head == nil {
// 		return nil
// 	}
// 	var tortoise, hare *design.ListNode = head, head
// 	for {
// 		tortoise = tortoise.Next
// 		hare = hare.Next
// 		if hare == nil || hare.Next == nil {
// 			return nil
// 		}
// 		hare = hare.Next
// 		if hare == tortoise {
// 			break
// 		}
// 	}
// 	for node := head; tortoise != nil && node != tortoise; {
// 		node = node.Next
// 		tortoise = tortoise.Next
// 	}
// 	return tortoise
// }

// func detectCycle(head *design.ListNode) *design.ListNode {
// 	var seen = make(map[*design.ListNode]bool)
// 	for node := head; node != nil; node = node.Next {
// 		if _, exists := seen[node]; exists {
// 			return node
// 		}
// 		seen[node] = true
// 	}
// 	return nil
// }

// func detectCycle(head *design.ListNode) *design.ListNode {
// 	var seen = make(map[*design.ListNode]bool)
// 	node := head
// 	for node != nil && !seen[node] {
// 		seen[node] = true
// 		node = node.Next
// 	}
// 	return nil
// }
