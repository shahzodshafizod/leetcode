package trees

import (
	"github.com/shahzodshafizod/leetcode/pkg"
)

// https://leetcode.com/problems/convert-sorted-array-to-binary-search-tree/

// Time: O(n), n=len(nums)
// Space: O(h), h=height of tree
func sortedArrayToBST(nums []int) *pkg.TreeNode {
	var createTree func(left int, right int) *pkg.TreeNode
	createTree = func(left int, right int) *pkg.TreeNode {
		if left == right {
			return nil
		}
		if left+1 == right {
			return &pkg.TreeNode{Val: nums[left]}
		}
		mid := left + (right-left)/2
		return &pkg.TreeNode{
			Val:   nums[mid],
			Left:  createTree(left, mid),
			Right: createTree(mid+1, right),
		}
	}
	return createTree(0, len(nums))
}

// // Time: O(n), n=len(nums)
// // Space: O(h), h=height of tree
// func sortedArrayToBST(nums []int) *pkg.TreeNode {
// 	if len(nums) == 0 {
// 		return nil
// 	}
// 	if len(nums) == 1 {
// 		return &pkg.TreeNode{Val: nums[0]}
// 	}
// 	var mid = len(nums) / 2
// 	return &pkg.TreeNode{
// 		Val:   nums[mid],
// 		Left:  sortedArrayToBST(nums[:mid]),
// 		Right: sortedArrayToBST(nums[mid+1:]),
// 	}
// }
