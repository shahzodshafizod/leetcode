from collections import deque
from design.tree import TreeNode, create_tree
from typing import Optional
import unittest

# https://leetcode.com/problems/same-tree/

# python3 -m unittest trees/0100-same-tree.py

class Solution(unittest.TestCase):
    # # Approach: DFS
    # def isSameTree(self, p: Optional[TreeNode], q: Optional[TreeNode]) -> bool:
    #     if not p or not q:
    #         return p == q
    #     return all([
    #         p.val == q.val,
    #         self.isSameTree(p.left, q.left),
    #         self.isSameTree(p.right, q.right),
    #     ])

    # Approach: BFS
    def isSameTree(self, p: Optional[TreeNode], q: Optional[TreeNode]) -> bool:
        queue = deque()
        queue.append((p, q))
        while queue:
            p, q = queue.popleft()
            if not p or not q:
                if p != q:
                    return False
                continue
            if p.val != q.val:
                return False
            queue.append((p.left, q.left))
            queue.append((p.right, q.right))
        return True

    def test(self) -> None:
        for p, q, expected in [
            ([1,2,3], [1,2,3], True),
            ([1,2], [1,None,2], False),
            ([1,2,1], [1,1,2], False),
        ]:
            p = create_tree(p)
            q = create_tree(q)
            output = self.isSameTree(p, q)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
