package linkedlists

import "github.com/shahzodshafizod/leetcode/pkg"

// https://leetcode.com/problems/merge-two-sorted-lists/

// // recursion
// func mergeTwoLists(list1 *pkg.ListNode, list2 *pkg.ListNode) *pkg.ListNode {
// 	if list1 == nil {
// 		return list2
// 	}
// 	if list2 == nil {
// 		return list1
// 	}
// 	if list1.Val < list2.Val {
// 		list1.Next = mergeTwoLists(list1.Next, list2)
// 		return list1
// 	}
// 	list2.Next = mergeTwoLists(list1, list2.Next)
// 	return list2
// }

// space: O(1)
func mergeTwoLists(list1 *pkg.ListNode, list2 *pkg.ListNode) *pkg.ListNode {
	dummy := &pkg.ListNode{} // to escape "if tail != nil"
	tail := dummy
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			tail.Next = list1
			list1 = list1.Next
		} else {
			tail.Next = list2
			list2 = list2.Next
		}
		tail = tail.Next
	}
	if list1 != nil {
		tail.Next = list1
	}
	if list2 != nil {
		tail.Next = list2
	}
	return dummy.Next
}
