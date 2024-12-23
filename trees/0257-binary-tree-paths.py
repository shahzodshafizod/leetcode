from typing import Optional, List
import unittest
from pkg.tree import TreeNode, create_tree

# https://leetcode.com/problems/binary-tree-paths/

# python3 -m unittest trees/0257-binary-tree-paths.py

class Solution(unittest.TestCase):
    # Approach: Depth-First Search + Backtracking
    # Time: O(n), n=# of nodes
    # Space: O(h), h=height of tree
    def binaryTreePaths(self, root: Optional[TreeNode]) -> List[str]:
        paths = []
        def dfs(node: TreeNode, path: str):
            path += str(node.val) + "->"
            if not node.left and not node.right:
                paths.append(path[:-2])
                return
            if node.left: dfs(node.left, path)
            if node.right: dfs(node.right, path)
        dfs(root, "")
        return paths

    def test(self):
        for root, expected in [
            ([1], ["1"]),
            ([1,2,3,None,5], ["1->2->5","1->3"]),
            ([1,2,3,None,None,4], ["1->2","1->3->4"]),
        ]:
            root = create_tree(root)
            output = self.binaryTreePaths(root)
            expected.sort()
            output.sort()
            self.assertListEqual(expected, output, f"expected: {expected}, output: {output}")
