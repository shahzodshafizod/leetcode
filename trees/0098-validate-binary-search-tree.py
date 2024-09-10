from typing import Optional
from design.tree import TreeNode, create_tree
import unittest

# https://leetcode.com/problems/validate-binary-search-tree/

# python3 -m unittest trees/0098-validate-binary-search-tree.py

class Solution(unittest.TestCase):
    # # Approach: Recursive
    # # N = # of tree nodes
    # # H = tree height
    # # Time: O(N)
    # # Space: O(H)
    # def isValidBST(self, root: Optional[TreeNode]) -> bool:
    #     def check(node: TreeNode, min_val: int, max_val: int) -> bool:
    #         if not node:
    #             return True
    #         if node.val <= min_val or node.val >= max_val:
    #             return False
    #         return check(node.left, min_val, node.val) and \
    #             check(node.right, node.val, max_val)
    #     return check(root, float('-inf'), float('inf'))

    # Approach: Iterative
    # N = # of tree nodes
    # H = tree height
    # Time: O(N)
    # Space: O(H)
    def isValidBST(self, root: Optional[TreeNode]) -> bool:
        stack = []
        prev, node = None, root
        while node or stack:
            while node:
                stack.append(node)
                node = node.left
            node = stack.pop()
            if prev and node.val <= prev.val:
                return False
            prev, node = node, node.right
        return True

    def testIsValidBST(self) -> None:
        for root, expected in [
            ([2, 1, 3], True),
            ([5, 1, 4, None, None, 3, 6], False),
            ([10, 5, 15, None, None, 6, 20], False),
            ([12, 8, 18, 5, 10, 14, 25], True),
            ([16, 8, 22, 9, None, 19, 25], False),
            ([13, 6, 17, 2, None, 10, 22], False),
            ([], True),
            ([12, 7, 18, 5, 9, 16, 25], True),
            ([1], True),
        ]:
            root = create_tree(root)
            output = self.isValidBST(root)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
