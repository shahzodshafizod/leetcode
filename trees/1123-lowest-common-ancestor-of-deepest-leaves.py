from typing import Optional
from pkg.tree import TreeNode, create_tree
import unittest

# https://leetcode.com/problems/lowest-common-ancestor-of-deepest-leaves/

# python3 -m unittest trees/1123-lowest-common-ancestor-of-deepest-leaves.py

class Solution(unittest.TestCase):
    def lcaDeepestLeaves(self, root: Optional[TreeNode]) -> Optional[TreeNode]:
        def dfs(node: TreeNode):
            if not node:
                return [None,0]
            left = dfs(node.left)
            right = dfs(node.right)
            if left[1] == right[1]:
                return [node,left[1]+1]
            if left[1] > right[1]:
                return left[0], left[1]+1
            return right[0], right[1]+1
        return dfs(root)[0]

    def test(self):
        for root, expected in [
            ([3,5,1,6,2,0,8,None,None,7,4], [2,7,4]),
            ([1], [1]),
            ([0,1,3,None,2], [2]),
        ]:
            root = create_tree(root)
            output = self.lcaDeepestLeaves(root)
            expected = create_tree(expected)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
