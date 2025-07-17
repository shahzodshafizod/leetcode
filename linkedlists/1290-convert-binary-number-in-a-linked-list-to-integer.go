package linkedlists

import "github.com/shahzodshafizod/leetcode/pkg"

// https://leetcode.com/problems/convert-binary-number-in-a-linked-list-to-integer/

func getDecimalValue(head *pkg.ListNode) int {
	var value = 0
	for node := head; node != nil; node = node.Next {
		// value = value * 2 + node.Val
		value = (value << 1) + node.Val
	}
	return value
}
