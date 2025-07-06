from typing import Optional
import unittest
from pkg.tree import TreeNode, create_tree

# https://leetcode.com/problems/symmetric-tree/

# python3 -m unittest trees/0101-symmetric-tree.py


class Solution(unittest.TestCase):
    # # Approach: Recursively
    # # Time: O(n), n=# of tree nodes
    # # Space: O(h), h=height of tree
    # def isSymmetric(self, root: Optional[TreeNode]) -> bool:
    #     def isMirror(lnode: TreeNode, rnode: TreeNode) -> bool:
    #         if not lnode or not rnode:
    #             return lnode == rnode
    #         if lnode.val != rnode.val:
    #             return False
    #         return (
    #             isMirror(lnode.left, rnode.right) and
    #             isMirror(lnode.right, rnode.left)
    #         )
    #     return isMirror(root.left, root.right)

    # Approach: Iteratively
    # Time: O(n), n=# of tree nodes
    # Space: O(w), w=width of tree
    def isSymmetric(self, root: Optional[TreeNode]) -> bool:
        lstack, rstack = [root.left], [root.right]
        while lstack and rstack:
            lnode, rnode = lstack.pop(), rstack.pop()
            if not lnode and not rnode:
                continue
            if not lnode or not rnode or lnode.val != rnode.val:
                return False
            lstack.append(lnode.right)
            lstack.append(lnode.left)
            rstack.append(rnode.left)
            rstack.append(rnode.right)
        return not lstack and not rstack

    def testIsSymmetric(self) -> None:
        for root, expected in [
            ([1, 2, 2, 3, 4, 4, 3], True),
            ([1, 2, 2, None, 3, None, 3], False),
            ([1, 2, 2, 5, None, None, 5, 6, None, None, 6], True),
            ([1, 2, 2, 2, None, 2], False),
            ([1, 2, 3], False),
            ([1], True),
            ([2, 3, 3, 4, 5, 5, 4, None, None, 8, 9, 9, 8], True),
            ([1, 2, 2, None, 3, 2], False),
        ]:
            root = create_tree(root)
            output = self.isSymmetric(root)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
