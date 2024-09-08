package linkedlists

import "github.com/shahzodshafizod/alkhwarizmi/design"

// https://leetcode.com/problems/find-the-minimum-and-maximum-number-of-nodes-between-critical-points/

func nodesBetweenCriticalPoints(head *design.ListNode) []int {
	var first, last = -1, -1
	var minima = 100000
	var prev = head
	for node, idx := head.Next, 1; node.Next != nil; node, idx = node.Next, idx+1 {
		if node.Val > prev.Val && node.Val > node.Next.Val ||
			node.Val < prev.Val && node.Val < node.Next.Val {
			if first == -1 {
				first = idx
			}
			if last != -1 {
				minima = min(minima, idx-last)
			}
			last = idx
		}
		prev = node
	}
	if first == last {
		return []int{-1, -1}
	}
	return []int{minima, last - first}
}
