from typing import List, Optional
from trees.tree import TreeNode, create_tree
import unittest

# https://leetcode.com/problems/binary-tree-postorder-traversal/

# python3 -m unittest trees/0145-binary-tree-postorder-traversal.py

class Solution(unittest.TestCase):
    def postorderTraversal(self, root: Optional[TreeNode]) -> List[int]:
        if root == None:
            return []
        order = self.postorderTraversal(root.left)
        order += self.postorderTraversal(root.right)
        order.append(root.val)
        return order

    def testPostorderTraversal(self) -> None:
        for root, expected in [
            ([1,None,2,None,None,3], [3,2,1]),
            ([], []),
            ([1], [1]),
        ]:
            root = create_tree(root, 0)
            output = self.postorderTraversal(root)
            self.assertListEqual(expected, output, f"expected: {expected}, output: {output}")
