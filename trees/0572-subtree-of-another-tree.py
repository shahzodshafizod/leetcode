import unittest
from typing import Optional
from pkg.tree import TreeNode, create_tree

# https://leetcode.com/problems/subtree-of-another-tree/

# python3 -m unittest trees/0572-subtree-of-another-tree.py

class Solution(unittest.TestCase):
    # # M = # nodes in root
    # # N = # nodes in subRoot
    # # Time: O(M x N)
    # # Space: O(M + N)
    # def isSubtree(self, root: Optional[TreeNode], subRoot: Optional[TreeNode], check: bool = False) -> bool:
    #     if not root or not subRoot:
    #         return root == subRoot
    #     if check:
    #         return all([
    #             root.val == subRoot.val,
    #             self.isSubtree(root.left, subRoot.left, True),
    #             self.isSubtree(root.right, subRoot.right, True),
    #         ])
    #     return any([
    #         self.isSubtree(root, subRoot, True),
    #         self.isSubtree(root.left, subRoot),
    #         self.isSubtree(root.right, subRoot),
    #     ])

    def isSubtree(self, root: Optional[TreeNode], subRoot: Optional[TreeNode]) -> bool:
        def serialize(node: TreeNode) -> str:
            if not node:
                return ";"
            return f"S{node.val}L{serialize(node.left)}R{serialize(node.right)}E"
        return serialize(subRoot) in serialize(root)

    def test(self) -> None:
        for root, subRoot, expected in [
            ([3,4,5,1,2], [4,1,2], True),
            ([3,4,5,1,2,None,None,None,None,0], [4,1,2], False),
            ([1,1], [1], True),
            ([1,2,3], [1,2], False),
            ([4,None,2], [3,4,5,None,2,None,2], False),
            ([3,4,5,None,2,None,2], [4,None,2], True),
            ([3,1,2,1,None,2], [3,1,2], False),
            ([1,None,1,None,1,None,1,None,1,None,1,None,1,None,1,None,1,None,1,None,1,2], [1,None,1,None,1,None,1,None,1,None,1,2], True),
            ([3,4,5,1,None,2], [3,1,2], False),
            ([4,1], [4,1,2], False),
            ([4,1,2], [4,1], False),
        ]:
            root = create_tree(root)
            subRoot = create_tree(subRoot)
            output = self.isSubtree(root, subRoot)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
