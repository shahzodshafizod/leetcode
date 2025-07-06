from typing import List, Optional
import unittest
from pkg.tree import TreeNode, create_tree

# https://leetcode.com/problems/binary-tree-postorder-traversal/

# python3 -m unittest trees/0145-binary-tree-postorder-traversal.py


class Solution(unittest.TestCase):
    def postorderTraversal(self, root: Optional[TreeNode]) -> List[int]:
        if root is None:
            return []
        order = self.postorderTraversal(root.left)
        order += self.postorderTraversal(root.right)
        order.append(root.val)
        return order

    def test(self):
        for root, expected in [
            ([1, None, 2, 3], [3, 2, 1]),
            ([1, 2, 3, 4, 5, None, 8, None, None, 6, 7, 9], [4, 6, 7, 5, 2, 9, 8, 3, 1]),
            ([], []),
            ([1], [1]),
        ]:
            root = create_tree(root)
            output = self.postorderTraversal(root)
            self.assertListEqual(expected, output, f"expected: {expected}, output: {output}")
