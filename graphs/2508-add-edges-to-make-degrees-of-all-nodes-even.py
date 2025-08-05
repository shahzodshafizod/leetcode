from collections import defaultdict
from typing import List, Set
import unittest

# https://leetcode.com/problems/add-edges-to-make-degrees-of-all-nodes-even/

# python3 -m unittest graphs/2508-add-edges-to-make-degrees-of-all-nodes-even.py


class Solution(unittest.TestCase):
    def isPossible(self, n: int, edges: List[List[int]]) -> bool:
        adj_set: dict[int, Set[int]] = defaultdict(set)
        for a, b in edges:
            adj_set[a].add(b)
            adj_set[b].add(a)

        # can_connect
        def cc(a: int, b: int) -> bool:
            return a not in adj_set[b]

        odds = [n for n in adj_set if len(adj_set[n]) & 1 == 1]
        if len(odds) == 2:
            return any(cc(odds[0], x) and cc(odds[1], x) for x in range(1, n + 1))
        if len(odds) == 4:
            a, b, c, d = odds
            return cc(a, b) and cc(c, d) or cc(a, c) and cc(b, d) or cc(a, d) and cc(b, c)
        return len(odds) == 0

    def test(self):
        for n, edges, expected in [
            (5, [[1, 2], [2, 3], [3, 4], [4, 2], [1, 4], [2, 5]], True),
            (4, [[1, 2], [3, 4]], True),
            (4, [[1, 2], [1, 3], [1, 4]], False),
        ]:
            output = self.isPossible(n, edges)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
