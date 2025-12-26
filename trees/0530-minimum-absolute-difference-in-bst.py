from typing import Optional, List, Tuple, Any
import unittest
from pkg.tree import TreeNode, create_tree

# https://leetcode.com/problems/minimum-absolute-difference-in-bst/

# python3 -m unittest trees/0530-minimum-absolute-difference-in-bst.py


class Solution(unittest.TestCase):
    # # Approach 1: In-order with tracking
    # # Time Complexity: O(n)
    # # Space Complexity: O(1) excluding recursion
    # def getMinimumDifference(self, root: Optional[TreeNode]) -> int:
    #     min_diff = 1 << 31
    #     prev = None

    #     def inorder(node: Optional[TreeNode]):
    #         nonlocal min_diff, prev

    #         if not node:
    #             return

    #         inorder(node.left)

    #         if prev is not None:
    #             min_diff = min(min_diff, node.val - prev)

    #         prev = node.val

    #         inorder(node.right)

    #     inorder(root)
    #     return min_diff

    # Approach 2: Collect all values then compute
    # Time Complexity: O(n)
    # Space Complexity: O(n)
    def getMinimumDifference(self, root: Optional[TreeNode]) -> int:
        values: List[int] = []

        def inorder(node: Optional[TreeNode]):
            if not node:
                return
            inorder(node.left)
            values.append(node.val)
            inorder(node.right)

        inorder(root)

        min_diff = 1 << 31
        for i in range(1, len(values)):
            min_diff = min(min_diff, values[i] - values[i - 1])

        return min_diff

    def test(self):
        test_cases: List[Tuple[List[Any], int]] = [
            ([4, 2, 6, 1, 3], 1),
            ([1, 0, 48, None, None, 12, 49], 1),
            ([5, 4, 7], 1),
        ]
        for root_vals, expected in test_cases:
            root = create_tree(root_vals)
            output = self.getMinimumDifference(root)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
