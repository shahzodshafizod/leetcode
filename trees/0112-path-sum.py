import unittest
from typing import Optional
from design.tree import TreeNode, create_tree

# https://leetcode.com/problems/path-sum/

# python3 -m unittest trees/0112-path-sum.py

class Solution(unittest.TestCase):
    def hasPathSum(self, root: Optional[TreeNode], targetSum: int) -> bool:
        def dfs(node: TreeNode, currSum: int) -> bool:
            if not node:
                return False
            currSum += node.val
            if not node.left and not node.right:
                return currSum == targetSum
            return dfs(node.left, currSum) or dfs(node.right, currSum)
        return dfs(root, 0)

    def test(self) -> None:
        for root, targetSum, expected in [
            ([5,4,8,11,None,13,4,7,2,None,None,None,1], 22, True),
            ([1,2,3], 5, False),
            ([], 0, False),
            ([1,2], 1, False),
            ([1,2,None,3,None,4,None,5], 6, False),
            ([1,2,None,3,None,4,None,5], 15, True),
        ]:
            root = create_tree(root)
            output = self.hasPathSum(root, targetSum)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
