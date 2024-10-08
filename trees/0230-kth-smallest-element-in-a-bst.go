package trees

import "github.com/shahzodshafizod/leetcode/pkg"

// https://leetcode.com/problems/kth-smallest-element-in-a-bst/

func kthSmallest(root *pkg.TreeNode, k int) int {
	leftCount := kthSmallestCountNodes(root.Left)
	if k <= leftCount {
		return kthSmallest(root.Left, k)
	}
	if k > leftCount+1 { // 1 is counted as the current node
		return kthSmallest(root.Right, k-leftCount-1)
	}
	return root.Val
}

func kthSmallestCountNodes(curr *pkg.TreeNode) int {
	if curr == nil {
		return 0
	}
	return 1 + kthSmallestCountNodes(curr.Left) + kthSmallestCountNodes(curr.Right)
}

// func kthSmallest(root *pkg.TreeNode, k int) int {
// 	var idx = 0
// 	return kthSmallestInOrder(root, &idx, k).Val
// }

// func kthSmallestInOrder(curr *pkg.TreeNode, idx *int, k int) *pkg.TreeNode {
// 	if curr == nil {
// 		return nil
// 	}
// 	if left := kthSmallestInOrder(curr.Left, idx, k); left != nil {
// 		return left
// 	}
// 	(*idx)++
// 	if *idx == k {
// 		return curr
// 	}
// 	return kthSmallestInOrder(curr.Right, idx, k)
// }

/*
Follow up: If the BST is modified often (i.e., we can do insert and delete operations)
and you need to find the kth smallest frequently, how would you optimize?

1. Create a map of counts: map[*pkg.TreeNode]int.
	It means how many nodes there are in the subtree of the node.
	Whatever you update or delete, you change the subtree until the min element,
	at the same time change count of those nodes by +1/-1.

2. Then finding kth element takes O(Log(N)):
*/

// func kthSmallest(root *pkg.TreeNode, k int) int {
// 	var counts = make(map[*pkg.TreeNode]int)
// 	kthSmallestCount(root, counts)
// 	return kthSmallestFind(root, k, counts).Val
// }

// func kthSmallestFind(curr *pkg.TreeNode, k int, counts map[*pkg.TreeNode]int) *pkg.TreeNode {
// 	if curr == nil || k > counts[curr] {
// 		return nil
// 	}
// 	var leftCount = 0
// 	if curr.Left != nil {
// 		leftCount = counts[curr.Left]
// 	}
// 	if k <= leftCount {
// 		return kthSmallestFind(curr.Left, k, counts)
// 	}
// 	if k > leftCount+1 {
// 		return kthSmallestFind(curr.Right, k-leftCount-1, counts)
// 	}
// 	return curr
// }

// func kthSmallestCount(curr *pkg.TreeNode, counts map[*pkg.TreeNode]int) {
// 	if curr == nil {
// 		return
// 	}
// 	kthSmallestCount(curr.Left, counts)
// 	kthSmallestCount(curr.Right, counts)
// 	counts[curr] = 1
// 	if curr.Left != nil {
// 		counts[curr] += counts[curr.Left]
// 	}
// 	if curr.Right != nil {
// 		counts[curr] += counts[curr.Right]
// 	}
// }
