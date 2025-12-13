from collections import deque
from typing import Optional, List, Tuple
import unittest
from pkg.tree import TreeNode, create_tree

# https://leetcode.com/problems/invert-binary-tree/

# python3 -m unittest trees/0226-invert-binary-tree.py


class Solution(unittest.TestCase):
    # # Approach #1: Recursive DFS
    # # Time: O(n) - visit each node once
    # # Space: O(h) - recursion stack, where h is height
    # def invertTree(self, root: Optional[TreeNode]) -> Optional[TreeNode]:
    #     if root:
    #         root.left, root.right = root.right, root.left
    #         self.invertTree(root.left)
    #         self.invertTree(root.right)
    #     return root

    # Approach #2: Iterative BFS
    # Time: O(n) - visit each node once
    # Space: O(w) - queue size, where w is maximum width
    def invertTree(self, root: Optional[TreeNode]) -> Optional[TreeNode]:
        if not root:
            return None

        queue = deque([root])

        while queue:
            node = queue.popleft()

            # Swap left and right children
            node.left, node.right = node.right, node.left

            if node.left:
                queue.append(node.left)
            if node.right:
                queue.append(node.right)

        return root

    def test(self):
        test_cases: List[Tuple[List[int], List[int]]] = [
            ([4, 2, 7, 1, 3, 6, 9], [4, 7, 2, 9, 6, 3, 1]),
            ([2, 1, 3], [2, 3, 1]),
            ([], []),
            ([1], [1]),
        ]
        for values, expected in test_cases:
            root = create_tree(values)
            expected = create_tree(expected)
            output = self.invertTree(root)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
