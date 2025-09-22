package linkedlists

import "github.com/shahzodshafizod/leetcode/pkg"

/*
Linked Lists: Cycle Detection.

Floyd's Tortoise and Hare (Sangpusht va Xargushi Floyd)
tortoise=1_step, hare=2_steps.
*/

// https://leetcode.com/problems/linked-list-cycle-ii/

func detectCycle(head *pkg.ListNode[int]) *pkg.ListNode[int] {
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

// func detectCycle(head *pkg.ListNode) *pkg.ListNode {
// 	if head == nil {
// 		return nil
// 	}
// 	var tortoise, hare *pkg.ListNode = head, head
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

// func detectCycle(head *pkg.ListNode) *pkg.ListNode {
// 	var seen = make(map[*pkg.ListNode]bool)
// 	for node := head; node != nil; node = node.Next {
// 		if _, exists := seen[node]; exists {
// 			return node
// 		}
// 		seen[node] = true
// 	}
// 	return nil
// }

// func detectCycle(head *pkg.ListNode) *pkg.ListNode {
// 	var seen = make(map[*pkg.ListNode]bool)
// 	node := head
// 	for node != nil && !seen[node] {
// 		seen[node] = true
// 		node = node.Next
// 	}
// 	return nil
// }
