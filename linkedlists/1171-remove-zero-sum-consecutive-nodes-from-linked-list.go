package linkedlists

import "github.com/shahzodshafizod/leetcode/pkg"

// https://leetcode.com/problems/remove-zero-sum-consecutive-nodes-from-linked-list/

// Prefix Sum + Hash Table
func removeZeroSumSublists(head *pkg.ListNode) *pkg.ListNode {
	dummy := &pkg.ListNode{Next: head}
	sum := 0
	hashset := make(map[int]*pkg.ListNode)
	var start, node *pkg.ListNode
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
