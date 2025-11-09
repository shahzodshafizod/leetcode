from typing import Optional, List, Any
import unittest
from pkg.tree import TreeNode, create_tree

# https://leetcode.com/problems/balanced-binary-tree/

# python3 -m unittest trees/0110-balanced-binary-tree.py


class Solution(unittest.TestCase):
    def isBalanced(self, root: Optional[TreeNode]) -> bool:
        def dfs(node: Optional[TreeNode]) -> int:
            if not node:
                return 0
            right = dfs(node.right)
            if right == -1:
                return -1
            left = dfs(node.left)
            if left == -1:
                return -1
            if abs(left - right) > 1:
                return -1
            return 1 + max(left, right)

        return dfs(root) != -1

    def test(self):
        test_cases: List[tuple[List[Any], bool]] = [
            ([3, 9, 20, None, None, 15, 7], True),
            ([1, 2, 2, 3, 3, None, None, 4, 4], False),
            ([], True),
        ]
        for root_vals, expected in test_cases:
            root = create_tree(root_vals)
            output = self.isBalanced(root)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
