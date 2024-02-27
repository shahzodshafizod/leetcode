package trees

// https://leetcode.com/problems/same-tree/

// DFS
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil || q == nil {
		return p == q
	}
	return p.Val == q.Val &&
		isSameTree(p.Left, q.Left) &&
		isSameTree(p.Right, q.Right)
}

// // BFS
// func isSameTree(p *TreeNode, q *TreeNode) bool {
// 	var queue = []*TreeNode{p, q}
// 	for length := len(queue); length > 0; length = len(queue) {
// 		p, q = queue[0], queue[1]
// 		queue = queue[2:]
// 		if p == nil || q == nil {
// 			if p != q {
// 				return false
// 			}
// 			continue
// 		}
// 		if p.Val != q.Val {
// 			return false
// 		}
// 		queue = append(queue, p.Left, q.Left, p.Right, q.Right)
// 	}
// 	return true
// }