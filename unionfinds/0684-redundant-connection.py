from typing import List
import unittest

# https://leetcode.com/problems/redundant-connection/

# python3 -m unittest unionfinds/0684-redundant-connection.py


class Solution(unittest.TestCase):
    # Approach: Union Find
    # Time: O(N⋅α(N))
    # Space: O(n)
    def findRedundantConnection(self, edges: List[List[int]]) -> List[int]:
        n = len(edges)
        parent = list(range(n + 1))

        def find(x: int) -> int:
            if parent[x] != x:
                parent[x] = find(parent[x])
            return parent[x]

        def union(x: int, y: int) -> bool:
            px = find(x)
            py = find(y)
            if px == py:
                return False
            parent[py] = px
            return True

        edge = []
        for a, b in edges:
            if not union(a, b):
                edge = [a, b]
        return edge

    def test(self):
        for edges, expected in [
            ([[1, 2], [1, 3], [2, 3]], [2, 3]),
            ([[1, 2], [2, 3], [3, 4], [1, 4], [1, 5]], [1, 4]),
        ]:
            output = self.findRedundantConnection(edges)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
