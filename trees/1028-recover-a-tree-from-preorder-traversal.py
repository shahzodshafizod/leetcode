from typing import Optional
from pkg.tree import TreeNode, create_tree
import unittest

# https://leetcode.com/problems/recover-a-tree-from-preorder-traversal/

# python3 -m unittest trees/1028-recover-a-tree-from-preorder-traversal.py

class Solution(unittest.TestCase):
    # Approach: Depth-First Search
    # Time: O(n)
    # Space: O(n)
    def recoverFromPreorder(self, traversal: str) -> Optional[TreeNode]:
        n = len(traversal)
        def recover(idx: int, depth: int) -> tuple[Optional[TreeNode], int]:
            if idx >= n:
                return None, idx
            dashes = 0
            while idx < n and traversal[idx] == '-':
                dashes += 1
                idx += 1
            if dashes != depth:
                return None, idx - dashes
            val = 0
            while idx < n and traversal[idx] != '-':
                val = val * 10 + int(traversal[idx])
                idx += 1
            node = TreeNode(val)
            node.left, idx = recover(idx, depth+1)
            node.right, idx = recover(idx, depth+1)
            return node, idx
        root, _ = recover(0, 0)
        return root

    def test(self):
        for traversal, expected in [
            ("1-2--3--4-5--6--7", [1,2,5,3,4,6,7]),
            ("1-2--3---4-5--6---7", [1,2,5,3,None,6,None,4,None,7]),
            ("1-401--349---90--88", [1,401,None,349,88,90]),
        ]:
            output = self.recoverFromPreorder(traversal)
            expected = create_tree(expected)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
