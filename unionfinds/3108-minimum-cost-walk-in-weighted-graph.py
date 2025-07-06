from typing import List
import unittest
from pkg.unionfind import UnionFind

# https://leetcode.com/problems/minimum-cost-walk-in-weighted-graph/

# python3 -m unittest unionfinds/3108-minimum-cost-walk-in-weighted-graph.py


class Solution(unittest.TestCase):
    # Time: O(v+e+q), v=n, e=len(edges), q=len(query)
    # Space: O(v)
    def minimumCost(self, n: int, edges: List[List[int]], query: List[List[int]]) -> List[int]:
        uf = UnionFind(n)
        for u, v, _ in edges:
            uf.Union(u, v)
        walk = [-1] * n
        for u, _, w in edges:
            walk[uf.Find(u)] &= w
        answer = [-1] * len(query)
        for idx, (s, t) in enumerate(query):
            ps, pt = uf.Find(s), uf.Find(t)
            if ps == pt:
                answer[idx] = walk[ps]
        return answer

    def test(self):
        for n, edges, query, expected in [
            (5, [[0, 1, 7], [1, 3, 7], [1, 2, 1]], [[0, 3], [3, 4]], [1, -1]),
            (3, [[0, 2, 7], [0, 1, 15], [1, 2, 6], [1, 2, 1]], [[1, 2]], [0]),
            (3, [[1, 0, 4], [0, 2, 5], [0, 2, 3], [0, 2, 14], [0, 2, 12], [2, 0, 14], [0, 2, 4]], [[2, 1]], [0]),
            (
                6,
                [
                    [2, 5, 3],
                    [0, 3, 1],
                    [1, 4, 0],
                    [0, 3, 0],
                    [0, 2, 5],
                    [2, 0, 2],
                    [5, 1, 2],
                    [1, 3, 1],
                    [2, 1, 4],
                    [3, 2, 3],
                ],
                [[1, 5], [0, 1], [1, 5], [3, 1], [1, 2]],
                [0, 0, 0, 0, 0],
            ),
            (
                9,
                [[1, 7, 7], [5, 6, 2], [3, 8, 5], [3, 6, 3]],
                [[3, 8], [1, 4], [5, 2], [5, 2], [1, 3]],
                [0, -1, -1, -1, -1],
            ),
        ]:
            output = self.minimumCost(n, edges, query)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
