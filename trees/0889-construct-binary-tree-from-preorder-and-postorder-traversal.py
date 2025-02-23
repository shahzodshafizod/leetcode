from pkg.tree import TreeNode, create_tree
from typing import Optional, List
import unittest

# https://leetcode.com/problems/construct-binary-tree-from-preorder-and-postorder-traversal/

# python3 -m unittest trees/0889-construct-binary-tree-from-preorder-and-postorder-traversal.py

class Solution(unittest.TestCase):
    def constructFromPrePost(self, preorder: List[int], postorder: List[int]) -> Optional[TreeNode]:
        preidx, postidx = 0, 0
        def construct() -> TreeNode:
            nonlocal preidx, postidx
            node = TreeNode(preorder[preidx])
            preidx += 1
            if node.val != postorder[postidx]:
                node.left = construct()
            if node.val != postorder[postidx]:
                node.right = construct()
            postidx += 1
            return node
        return construct()

    def test(self):
        for preorder, postorder, expected in [
            ([1,2,4,5,3,6,7], [4,5,2,6,7,3,1], [1,2,3,4,5,6,7]),
            ([1], [1], [1]),
        ]:
            expected = create_tree(expected)
            output = self.constructFromPrePost(preorder, postorder)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
