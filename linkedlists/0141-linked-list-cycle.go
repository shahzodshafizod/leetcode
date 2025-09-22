package linkedlists

import "github.com/shahzodshafizod/leetcode/pkg"

// https://leetcode.com/problems/linked-list-cycle/

// Floyd's Tortoise and Hare
func hasCycle(head *pkg.ListNode[int]) bool {
	for tortoise, hare := head, head; hare != nil && hare.Next != nil; {
		tortoise = tortoise.Next
		hare = hare.Next.Next

		if tortoise == hare {
			return true
		}
	}

	return false
}
