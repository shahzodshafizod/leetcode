from typing import List
import unittest

# https://leetcode.com/problems/maximum-path-quality-of-a-graph/

# python3 -m unittest graphs/2065-maximum-path-quality-of-a-graph.py


class Solution(unittest.TestCase):
    # Approach: Backtracking
    # Time: O(n∗(4^n))
    # Space: O(n)
    def maximalPathQuality(self, values: List[int], edges: List[List[int]], maxTime: int) -> int:
        n = len(values)
        adj: List[List[tuple[int, int]]] = [[] for _ in range(n)]
        for u, v, t in edges:
            adj[u].append((v, t))
            adj[v].append((u, t))

        def dfs(node: int, value: int, time: int) -> int:
            best = value if node == 0 else 0
            old_value = values[node]
            values[node] = 0
            for nei, wei in adj[node]:
                if time >= wei:
                    best = max(best, dfs(nei, value + values[nei], time - wei))
            values[node] = old_value
            return best

        return dfs(0, values[0], maxTime)

    def test(self):
        for values, edges, maxTime, expected in [
            ([0, 32, 10, 43], [[0, 1, 10], [1, 2, 15], [0, 3, 10]], 49, 75),
            ([5, 10, 15, 20], [[0, 1, 10], [1, 2, 10], [0, 3, 10]], 30, 25),
            ([1, 2, 3, 4], [[0, 1, 10], [1, 2, 11], [2, 3, 12], [1, 3, 13]], 50, 7),
        ]:
            output = self.maximalPathQuality(values, edges, maxTime)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
