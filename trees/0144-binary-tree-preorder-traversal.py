from typing import List, Optional
import unittest
from pkg.tree import TreeNode, create_tree

# https://leetcode.com/problems/binary-tree-preorder-traversal/

# python3 -m unittest trees/0144-binary-tree-preorder-traversal.py


class Solution(unittest.TestCase):
    def preorderTraversal(self, root: Optional[TreeNode]) -> List[int]:
        if not root: return []
        vals = [root.val]
        vals.extend(self.preorderTraversal(root.left))
        vals.extend(self.preorderTraversal(root.right))
        return vals

    def test(self):
        for root, expected in [
            ([1,None,2,3], [1,2,3]),
            ([1,2,3,4,5,None,8,None,None,6,7,9], [1,2,4,5,6,7,3,8,9]),
            ([], []),
            ([1], [1]),
        ]:
            output = self.preorderTraversal(create_tree(root))
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
