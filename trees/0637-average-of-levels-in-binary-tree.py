import unittest
from collections import deque
from typing import List, Optional, Tuple, Any
from pkg.tree import TreeNode, create_tree

# https://leetcode.com/problems/average-of-levels-in-binary-tree/

# python3 -m unittest trees/0637-average-of-levels-in-binary-tree.py


class Solution(unittest.TestCase):
    # Approach: BFS (Breadth-First Search) / Level-Order Traversal
    # Time: O(n) - visit each node once
    # Space: O(w) - where w is maximum width of tree (queue size)
    def averageOfLevels(self, root: Optional[TreeNode]) -> List[float]:
        if not root:
            return []

        result: List[float] = []
        queue = deque([root])

        while queue:
            level_size = len(queue)
            level_sum = 0

            # Process all nodes at current level
            for _ in range(level_size):
                node = queue.popleft()
                level_sum += node.val

                # Add children for next level
                if node.left:
                    queue.append(node.left)
                if node.right:
                    queue.append(node.right)

            # Calculate average for this level
            result.append(level_sum / level_size)

        return result

    def test(self):
        test_cases: List[Tuple[List[Any], List[float]]] = [
            ([3, 9, 20, None, None, 15, 7], [3.00000, 14.50000, 11.00000]),
            ([3, 9, 20, 15, 7], [3.00000, 14.50000, 11.00000]),
            ([1], [1.00000]),
            ([1, 2, 3, 4, 5], [1.00000, 2.50000, 4.50000]),
        ]
        for root, expected in test_cases:
            tree = create_tree(root)
            output = self.averageOfLevels(tree)
            # Compare with small tolerance for floating point
            self.assertEqual(len(expected), len(output))
            for exp, out in zip(expected, output):
                self.assertAlmostEqual(exp, out, places=5, msg=f"expected: {expected}, output: {output}")
