from typing import List
import unittest

# https://leetcode.com/problems/find-number-of-coins-to-place-in-tree-nodes/

# python3 -m unittest heaps/2973-find-number-of-coins-to-place-in-tree-nodes.py


class Solution(unittest.TestCase):
    # Approach: Depth-First Search + Sorting
    # Time: O(n)
    # Space: O(n)
    def placedCoins(self, edges: List[List[int]], cost: List[int]) -> List[int]:
        n = len(cost)
        graph = [[] for _ in range(n)]
        for a, b in edges:
            graph[a].append(b)
            graph[b].append(a)
        coin = [0] * n

        def dfs(parent: int, node: int) -> List[int]:
            cs = [cost[node]]
            for child in graph[node]:
                if child != parent:
                    cs.extend(dfs(node, child))
            cs.sort()
            if len(cs) > 5:
                # [neg, neg, pos, pos, pos]
                cs = [cs[0], cs[1], cs[-3], cs[-2], cs[-1]]
            coin[node] = 1
            if len(cs) >= 3:
                coin[node] = max(0, cs[0] * cs[1] * cs[-1], cs[-3] * cs[-2] * cs[-1])
            return cs

        dfs(-1, 0)
        return coin

    def test(self):
        for edges, cost, expected in [
            ([[0, 1], [0, 2], [0, 3], [0, 4], [0, 5]], [1, 2, 3, 4, 5, 6], [120, 1, 1, 1, 1, 1]),
            (
                [[0, 1], [0, 2], [1, 3], [1, 4], [1, 5], [2, 6], [2, 7], [2, 8]],
                [1, 4, 2, 3, 5, 7, 8, -4, 2],
                [280, 140, 32, 1, 1, 1, 1, 1, 1],
            ),
            ([[0, 1], [0, 2]], [1, 2, -2], [0, 1, 1]),
        ]:
            output = self.placedCoins(edges, cost)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
