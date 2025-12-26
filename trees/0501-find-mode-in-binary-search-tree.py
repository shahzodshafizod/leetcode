from typing import List, Optional, Tuple, Any
import unittest
from pkg.tree import TreeNode, create_tree

# https://leetcode.com/problems/find-mode-in-binary-search-tree/

# python3 -m unittest trees/0501-find-mode-in-binary-search-tree.py


class Solution(unittest.TestCase):
    # # Approach 1: HashMap/Frequency Counter
    # # Traverse tree and count all values in dictionary
    # # Find max frequency, collect all values with that frequency
    # # Time Complexity: O(n)
    # # Space Complexity: O(n)
    # def findMode(self, root: Optional[TreeNode]) -> List[int]:
    #     freq: dict[int, int] = {}

    #     def traverse(node: Optional[TreeNode]):
    #         if not node:
    #             return
    #         freq[node.val] = freq.get(node.val, 0) + 1
    #         traverse(node.left)
    #         traverse(node.right)

    #     traverse(root)

    #     max_freq = max(freq.values())
    #     return [val for val, count in freq.items() if count == max_freq]

    # Approach 2: In-order Traversal (Optimal for BST)
    # Use BST property: in-order gives sorted sequence
    # Track current value and count
    # Time Complexity: O(n)
    # Space Complexity: O(1) excluding recursion stack
    def findMode(self, root: Optional[TreeNode]) -> List[int]:
        result: List[int] = []
        prev_val = None
        current_count = 0
        max_count = 0

        def inorder(node: Optional[TreeNode]):
            nonlocal prev_val, current_count, max_count, result

            if not node:
                return

            # Traverse left
            inorder(node.left)

            # Process current
            if prev_val is None or prev_val != node.val:
                current_count = 1
            else:
                current_count += 1

            if current_count > max_count:
                max_count = current_count
                result = [node.val]
            elif current_count == max_count:
                result.append(node.val)

            prev_val = node.val

            # Traverse right
            inorder(node.right)

        inorder(root)
        return result

    def test(self):
        test_cases: List[Tuple[List[Any], List[int]]] = [
            ([1, None, 2, 2], [2]),
            ([1, 1, 2], [1]),
            ([1], [1]),
        ]
        for root_vals, expected in test_cases:
            root = create_tree(root_vals)
            output = self.findMode(root)
            output.sort()
            expected.sort()
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
