import unittest
from typing import Optional, Tuple, List, Any
from pkg.tree import Node, create_n_ary_tree

# https://leetcode.com/problems/maximum-depth-of-n-ary-tree/

# python3 -m unittest trees/0559-maximum-depth-of-n-ary-tree.py


class Solution(unittest.TestCase):
    # # Approach #1: Brute Force - BFS level-order traversal
    # # Count the number of levels by doing BFS.
    # # Time: O(n) - visit each node once
    # # Space: O(w) - queue size where w is maximum width
    # def maxDepth(self, root: Optional[Node]) -> int:
    #     if root is None:
    #         return 0

    #     from collections import deque

    #     queue = deque([root])
    #     depth = 0

    #     while queue:
    #         depth += 1
    #         level_size = len(queue)
    #         for _ in range(level_size):
    #             node = queue.popleft()
    #             for child in node.children:
    #                 queue.append(child)

    #     return depth

    # Approach #2: Recursive DFS
    # Recursively find the maximum depth among all children and add 1 for current node.
    # Time: O(n) - visit each node once
    # Space: O(h) - recursion stack where h is height
    def maxDepth(self, root: Optional[Node]) -> int:
        if root is None:
            return 0
        if not root.children:
            return 1
        return 1 + max(self.maxDepth(child) for child in root.children)

    def test(self):
        test_cases: List[Tuple[List[Any], int]] = [
            ([1, None, 3, 2, 4, None, 5, 6], 3),  # root -> 3/2/4 -> 5/6
            ([1, None, 2, 3, 4, 5, None, None, 6, 7, None, 8, None, 9, 10, None, None, 11, None, 12, None, 13, None, None, 14], 5),
            ([], 0),
            ([1], 1),
        ]
        for vals, expected in test_cases:
            root = create_n_ary_tree(vals)
            output = self.maxDepth(root)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
