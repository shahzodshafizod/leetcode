import unittest
from typing import Optional
from pkg.tree import TreeNode, create_tree

# https://leetcode.com/problems/binary-tree-tilt/

# python3 -m unittest trees/0563-binary-tree-tilt.py


class Solution(unittest.TestCase):
    # Approach: Postorder DFS Traversal
    # Key insight: For each node, we need the sum of its left and right subtrees.
    # Postorder traversal (left, right, root) naturally provides this.
    #
    # Strategy:
    # 1. Use DFS to traverse tree in postorder
    # 2. For each node, calculate left subtree sum and right subtree sum
    # 3. Tilt = |left_sum - right_sum|
    # 4. Accumulate total tilt
    # 5. Return current subtree sum to parent
    #
    # Time Complexity: O(n) where n is number of nodes
    # Space Complexity: O(h) where h is height (recursion stack)
    def findTilt(self, root: Optional[TreeNode]) -> int:
        total_tilt = 0

        def dfs(node: Optional[TreeNode]) -> int:
            nonlocal total_tilt

            if not node:
                return 0

            # Postorder: process left and right first
            left_sum = dfs(node.left)
            right_sum = dfs(node.right)

            # Calculate tilt for current node
            tilt = abs(left_sum - right_sum)

            # Accumulate total tilt
            total_tilt += tilt

            # Return sum of current subtree (for parent)
            return left_sum + right_sum + node.val

        dfs(root)
        return total_tilt

    def test(self):
        for root, expected in [
            (create_tree([1, 2, 3]), 1),
            (create_tree([4, 2, 9, 3, 5, None, 7]), 15),
            (create_tree([21, 7, 14, 1, 1, 2, 2, 3, 3]), 9),
            (create_tree([]), 0),
            (create_tree([1]), 0),
        ]:
            output = self.findTilt(root)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
