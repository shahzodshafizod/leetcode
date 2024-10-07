package linkedlists

import "github.com/shahzodshafizod/leetcode/design"

// https://leetcode.com/problems/remove-zero-sum-consecutive-nodes-from-linked-list/

// Prefix Sum + Hash Table
func removeZeroSumSublists(head *design.ListNode) *design.ListNode {
	var dummy = &design.ListNode{Next: head}
	var sum = 0
	var hashset = make(map[int]*design.ListNode)
	var start, node *design.ListNode
	for end := dummy; end != nil; end = end.Next {
		sum += end.Val
		start = hashset[sum]
		if start != nil {
			for node = start.Next; node != end; node = node.Next {
				sum += node.Val
				delete(hashset, sum)
			}
			sum += end.Val
			start.Next = end.Next
			end = start
		}
		hashset[sum] = end
	}
	return dummy.Next
}
