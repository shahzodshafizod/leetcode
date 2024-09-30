from typing import Optional
from design.tree import TreeNode, create_tree
import unittest

# https://leetcode.com/problems/maximum-depth-of-binary-tree/

# python3 -m unittest trees/0104-maximum-depth-of-binary-tree.py

class Solution(unittest.TestCase):
    def maxDepth(self, root: Optional[TreeNode]) -> int:
        if not root:
            return 0
        return 1 + max(
            self.maxDepth(root.left),
            self.maxDepth(root.right),
        )

    def test(self) -> None:
        for root, expected in [
            ([3,9,20,None,None,15,7], 3),
            ([1,None,2], 2),
            ([], 0),
            ([1], 1),
            ([1,2,3,4,5,None,None,None,3,None,None,None,7], 5),
            ([1,None,2,None,3,None,4,None,5], 5),
        ]:
            root = create_tree(root)
            output = self.maxDepth(root)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
