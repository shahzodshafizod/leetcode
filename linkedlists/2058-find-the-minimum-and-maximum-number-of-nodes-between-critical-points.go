package linkedlists

import "github.com/shahzodshafizod/leetcode/pkg"

// https://leetcode.com/problems/find-the-minimum-and-maximum-number-of-nodes-between-critical-points/

func nodesBetweenCriticalPoints(head *pkg.ListNode) []int {
	first, last := -1, -1
	minima := 100000

	prev := head
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
