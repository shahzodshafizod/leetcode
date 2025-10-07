from typing import List
import unittest

# https://leetcode.com/problems/longest-cycle-in-a-graph/

# python3 -m unittest graphs/2360-longest-cycle-in-a-graph.py

# Hint: Each node can be part of at most one cycle


class Solution(unittest.TestCase):
    def longestCycle(self, edges: List[int]) -> int:
        n = len(edges)
        dist = [0] * n

        def dfs(node: int, length: int) -> int:
            if node == -1 or dist[node] == -1:
                return -1
            if 0 < dist[node] < length:
                return length - dist[node]
            dist[node] = length
            length = dfs(edges[node], length + 1)
            dist[node] = -1
            return length

        return max(dfs(node, 1) for node in range(n) if dist[node] == 0)

    def test(self):
        for edges, expected in [
            ([3, 3, 4, 2, 3], 3),
            ([2, -1, 3, 1], -1),
            ([-1, 4, -1, 2, 0, 4], -1),
            ([1, 0], 2),
            # ([1,19,30,87,53,91,36,6,95,14,73,2,59,76,49,41,29,28,8,9,96,80,68,10,31,24,0,42,39,4,51,64,25,90,35,71,97,32,16,18,62,22,40,78,55,13,99,93,66,26,98,5,88,74,89,81,43,12,44,57,75,47,34,72,85,77,3,65,46,20,60,33,48,94,84,21,69,54,56,11,70,83,86,79,61,37,67,15,7,38,23,52,58,27,50,63,92,45,17,82], 50),
            # ([49,29,24,24,-1,-1,-1,2,23,-1,44,47,52,49,9,31,40,34,-1,53,33,-1,2,-1,18,31,0,9,47,35,-1,-1,-1,-1,4,12,14,19,-1,4,4,43,25,11,54,-1,25,39,17,-1,22,44,-1,44,29,50,-1,-1], -1),
        ]:
            output = self.longestCycle(edges)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
