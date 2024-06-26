package trees

// https://leetcode.com/problems/balance-a-binary-search-tree/

// // Does not work due to math helper functions
// // Approach 3: Day-Stout-Warren Algorithm / In-Place Balancing
// // time: O(n)
// // space: O(n) for recursion stack
// func balanceBST(root *TreeNode) *TreeNode {
// 	if root == nil {
// 		return nil
// 	}
// 	// Function to perform a right rotation
// 	var rightRotate = func(parent *TreeNode, right *TreeNode) {
// 		var tmp = right.Left
// 		right.Left = tmp.Right
// 		tmp.Right = right
// 		parent.Right = tmp
// 	}
// 	// Function to perform a left rotation
// 	var leftRotate = func(parent *TreeNode, right *TreeNode) {
// 		var tmp = right.Right
// 		right.Right = tmp.Left
// 		tmp.Left = right
// 		parent.Right = tmp
// 	}
// 	// Function to perform a series of left rotations to balance the vine
// 	var makeRotations = func(vineHead *TreeNode, count int) {
// 		var curr = vineHead
// 		for count > 0 {
// 			count--
// 			var tmp = curr.Right
// 			leftRotate(curr, tmp)
// 			curr = curr.Right
// 		}
// 	}
// 	// Step 1: Create the backbone (vine) - DLL (doubly linked list)
// 	// Temporary dummy node
// 	var vineHead = &TreeNode{Right: root}
// 	var curr = vineHead
// 	for curr.Right != nil {
// 		if curr.Right.Left != nil {
// 			rightRotate(curr, curr.Right)
// 		} else {
// 			curr = curr.Right
// 		}
// 	}
// 	// Step 2: Count the nodes
// 	var nodeCount = 0
// 	curr = vineHead.Right
// 	for curr != nil {
// 		nodeCount++
// 		curr = curr.Right
// 	}
// 	// Step 3: Create a balanced BST
// 	var m = int(math.Pow(2, math.Floor(math.Log2(float64(nodeCount+1))))) - 1
// 	makeRotations(vineHead, nodeCount-m)
// 	for m > 1 {
// 		m /= 2
// 		makeRotations(vineHead, m)
// 	}
// 	var balancedRoot = vineHead.Right
// 	// Delete the temporary dummy node
// 	vineHead = nil
// 	return balancedRoot
// }

// Approach #2: Morris Traversal + Recursive Construction
// time: O(2n) = O(n)
// space: O(2n) = O(n): 1*recursion+nodes
func balanceBST(root *TreeNode) *TreeNode {
	var getNodes = func(root *TreeNode) []*TreeNode {
		var nodes = make([]*TreeNode, 0)
		var curr = root
		for curr != nil {
			if curr.Left == nil {
				nodes = append(nodes, curr)
				curr = curr.Right
			} else {
				prev := curr.Left
				for prev.Right != nil && prev.Right != curr {
					prev = prev.Right
				}
				if prev.Right == curr {
					prev.Right = nil
					nodes = append(nodes, curr)
					curr = curr.Right
				} else {
					prev.Right = curr
					curr = curr.Left
				}
			}
		}
		return nodes
	}
	var balance func(nodes []*TreeNode, left int, right int) *TreeNode
	balance = func(nodes []*TreeNode, left int, right int) *TreeNode {
		if left > right {
			return nil
		}
		var mid = (left + right) / 2
		var node = nodes[mid]
		node.Left = balance(nodes, left, mid-1)
		node.Right = balance(nodes, mid+1, right)
		return node
	}
	var nodes = getNodes(root)
	root = balance(nodes, 0, len(nodes)-1)
	return root
}

// // Approach #1: Inorder Traversal + Recursive Construction
// // time: O(2n) = O(n)
// // space: O(3n) = O(n): 2*recursions+nodes
// func balanceBST(root *TreeNode) *TreeNode {
// 	var getNodes func(node *TreeNode) []*TreeNode
// 	getNodes = func(node *TreeNode) []*TreeNode {
// 		if node == nil {
// 			return []*TreeNode{}
// 		}
// 		var nodes = getNodes(node.Left)
// 		nodes = append(nodes, node)
// 		nodes = append(nodes, getNodes(node.Right)...)
// 		return nodes
// 	}
// 	var balance func(nodes []*TreeNode, left int, right int) *TreeNode
// 	balance = func(nodes []*TreeNode, left int, right int) *TreeNode {
// 		if left > right {
// 			return nil
// 		}
// 		var mid = (left + right) / 2
// 		var node = nodes[mid]
// 		node.Left = balance(nodes, left, mid-1)
// 		node.Right = balance(nodes, mid+1, right)
// 		return node
// 	}
// 	var nodes = getNodes(root)
// 	root = balance(nodes, 0, len(nodes)-1)
// 	return root
// }
