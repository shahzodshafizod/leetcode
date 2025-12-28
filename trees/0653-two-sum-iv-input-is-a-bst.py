from typing import Optional, List, Tuple, Any, Set
import unittest
from pkg.tree import TreeNode, create_tree

# https://leetcode.com/problems/two-sum-iv-input-is-a-bst/

# python3 -m unittest trees/0653-two-sum-iv-input-is-a-bst.py


class Solution(unittest.TestCase):
    # # Approach: Brute Force - DFS with Search
    # # For each node, search if complement exists in tree
    # # N = number of nodes, H = height of tree
    # # Time: O(N*H) - visit each node and search for complement
    # # Space: O(H) - recursion stack
    # def findTarget(self, root: Optional[TreeNode], k: int) -> bool:
    #     def search(node: Optional[TreeNode], target: int, source: TreeNode) -> bool:
    #         # Search for target value in tree, excluding source node
    #         if not node:
    #             return False
    #         # Found target but make sure it's not the same node
    #         if node.val == target and node is not source:
    #             return True
    #         if target < node.val:
    #             return search(node.left, target, source)
    #         return search(node.right, target, source)
    
    #     def dfs(node: Optional[TreeNode]) -> bool:
    #         if not node:
    #             return False
    #         # Check if complement exists
    #         complement = k - node.val
    #         if search(root, complement, node):
    #             return True
    #         # Check left and right subtrees
    #         return dfs(node.left) or dfs(node.right)
    
    #     return dfs(root)

    # Approach: Hash Set - One Pass DFS
    # Store seen values and check for complement
    # N = number of nodes
    # Time: O(N) - visit each node once
    # Space: O(N) - hash set stores all values
    def findTarget(self, root: Optional[TreeNode], k: int) -> bool:
        seen: Set[int] = set()

        def dfs(node: Optional[TreeNode]) -> bool:
            if not node:
                return False
            # Check if complement exists in seen values
            complement = k - node.val
            if complement in seen:
                return True
            # Add current value to seen set
            seen.add(node.val)
            # Check left and right subtrees
            return dfs(node.left) or dfs(node.right)

        return dfs(root)

    def test(self):
        test_cases: List[Tuple[List[Any], int, bool]] = [
            ([5, 3, 6, 2, 4, None, 7], 9, True),  # 3 + 6 = 9
            ([5, 3, 6, 2, 4, None, 7], 28, False),  # No pair sums to 28
            ([2, 1, 3], 4, True),  # 1 + 3 = 4
            ([2, 1, 3], 1, False),  # No pair sums to 1
            ([2, 1, 3], 3, True),  # 1 + 2 = 3
            ([1], 2, False),  # Single node, no pair
        ]
        for root, k, expected in test_cases:
            root = create_tree(root)
            output = self.findTarget(root, k)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
