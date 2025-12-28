from typing import Optional, Tuple, List, Any
import unittest
from pkg.tree import TreeNode, create_tree

# https://leetcode.com/problems/second-minimum-node-in-a-binary-tree/

# python3 -m unittest trees/0671-second-minimum-node-in-a-binary-tree.py


class Solution(unittest.TestCase):
    # # Approach #1: Brute Force - Collect all unique values and sort
    # # Traverse tree to collect all unique values into a set
    # # Convert to sorted list and return second smallest
    # # N = number of nodes
    # # Time: O(N log N) - collect N values, sort them
    # # Space: O(N) - set stores unique values
    # def findSecondMinimumValue(self, root: Optional[TreeNode]) -> int:
    #     values: Set[int] = set()
    
    #     def dfs(node: Optional[TreeNode]) -> None:
    #         if not node:
    #             return
    #         values.add(node.val)
    #         dfs(node.left)
    #         dfs(node.right)
    
    #     dfs(root)
    #     sorted_values = sorted(values)
    
    #     return sorted_values[1] if len(sorted_values) >= 2 else -1

    # Approach #2: Optimized DFS - Use tree property
    # Key insight: root.val is always the minimum (by tree property)
    # We only need to find the smallest value > root.val
    # Prune branches where node.val == root.val (can't have second min there)
    # N = number of nodes
    # Time: O(N) - worst case visit all nodes, but often prunes early
    # Space: O(H) - recursion stack depth H (tree height)
    def findSecondMinimumValue(self, root: Optional[TreeNode]) -> int:
        if not root:
            return -1

        min_val = root.val
        MAX = 1 << 31
        second_min = MAX

        def dfs(node: Optional[TreeNode]) -> None:
            nonlocal second_min

            if not node:
                return

            # If current value is greater than min but less than second_min
            if min_val < node.val < second_min:
                second_min = node.val

            # Only explore children if they might contain second minimum
            # If node.val == min_val, children might have larger values
            # If node.val > min_val, we already updated second_min
            if node.val == min_val:
                dfs(node.left)
                dfs(node.right)

        dfs(root)

        return second_min if second_min != MAX else -1

    def test(self):
        test_cases: List[Tuple[List[Any], int]] = [
            ([2, 2, 5, None, None, 5, 7], 5),
            ([2, 2, 2], -1),
            ([1, 1, 3, 1, 1, 3, 4, 3, 1, 1, 1, 3, 8, 4, 8, 3, 3, 1, 6, 2, 1], 2),
            ([2, 2, 2147483647], 2147483647),
            ([5, 8, 5], 8),
        ]
        for root, expected in test_cases:
            root = create_tree(root)
            output = self.findSecondMinimumValue(root)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
