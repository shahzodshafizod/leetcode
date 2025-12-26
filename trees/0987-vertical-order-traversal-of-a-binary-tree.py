from typing import List, Optional, Tuple, Any
import unittest
from pkg.tree import TreeNode, create_tree

# https://leetcode.com/problems/vertical-order-traversal-of-a-binary-tree/

# python3 -m unittest trees/0987-vertical-order-traversal-of-a-binary-tree.py


class Solution(unittest.TestCase):
    # Approach: DFS + Sorting
    # Track (column, row, value) for each node
    # Group by column, sort by row then value
    # Time Complexity: O(n log n)
    # Space Complexity: O(n)
    def verticalTraversal(self, root: Optional[TreeNode]) -> List[List[int]]:
        nodes: List[Tuple[int, int, int]] = []

        def dfs(node: Optional[TreeNode], row: int, col: int):
            if not node:
                return
            nodes.append((col, row, node.val))
            dfs(node.left, row + 1, col - 1)
            dfs(node.right, row + 1, col + 1)

        dfs(root, 0, 0)

        # Sort by column, then row, then value
        nodes.sort()

        # Group by column
        result: List[List[int]] = []
        if not nodes:
            return result

        current_col = nodes[0][0]
        current_group = [nodes[0][2]]

        for i in range(1, len(nodes)):
            col, _, val = nodes[i]
            if col == current_col:
                current_group.append(val)
            else:
                result.append(current_group)
                current_col = col
                current_group = [val]

        result.append(current_group)
        return result

    def test(self):
        test_cases: List[Tuple[List[Any], List[List[int]]]] = [
            ([3, 9, 20, None, None, 15, 7], [[9], [3, 15], [20], [7]]),
            ([1, 2, 3, 4, 5, 6, 7], [[4], [2], [1, 5, 6], [3], [7]]),
        ]
        for root_vals, expected in test_cases:
            root = create_tree(root_vals)
            output = self.verticalTraversal(root)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
