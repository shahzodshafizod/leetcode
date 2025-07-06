from typing import Optional
import unittest
from pkg.tree import TreeNode, create_tree

# https://leetcode.com/problems/binary-tree-maximum-path-sum/

# python3 -m unittest trees/0124-binary-tree-maximum-path-sum.py


class Solution(unittest.TestCase):
    def maxPathSum(self, root: Optional[TreeNode]) -> int:
        def dfs(node: TreeNode) -> int:
            if not node:
                return 0
            lpath = max(0, dfs(node.left))
            rpath = max(0, dfs(node.right))
            nonlocal max_path
            max_path = max(max_path, node.val + lpath + rpath)
            curr_path = node.val + max(lpath, rpath)
            return curr_path

        max_path = -1000
        dfs(root)
        return max_path

    def test(self):
        for root, expected in [
            ([1, 2, 3], 6),
            ([-10, 9, 20, None, None, 15, 7], 42),
            ([-10, -9, 20, None, None, -15, -7], 20),
            ([2, -1], 2),
            ([5, 4, 8, 11, None, 13, 4, 7, 2, None, None, None, 1], 48),
            ([9, 6, -3, None, None, -6, 2, None, None, 2, None, -6, -6, -6], 16),
            ([1, -2, -3, 1, 3, -2, None, -1], 3),
            ([0], 0),
            ([1, 1, 1, 1, 1, 1, 1], 5),
            ([1, 2, 3, 4, 5], 11),
            ([1, 2, 1, 10, 10, 1, 1], 22),
            ([-3], -3),
            ([-1, None, 9, -6, 3, None, None, None, -2], 12),
            ([-1, None, 9, -6, 3, None, None, None, 2], 14),
            ([-1, 8, 2], 9),
            ([-1, -2, 10, -6, None, -3, -6], 10),
            ([8, 9, -6, None, None, 5, 9], 20),
            ([-1, 8, 2, None, -9, 0, None, None, None, -3, None, None, -9, None, 2], 9),
            ([2, -1, -2], 2),
            ([10, -20, -30], 10),
            ([1, -2, 3], 4),
            ([1, 2], 3),
        ]:
            root = create_tree(root)
            output = self.maxPathSum(root)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
