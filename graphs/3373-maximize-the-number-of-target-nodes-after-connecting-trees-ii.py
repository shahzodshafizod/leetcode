from typing import List
import unittest

# https://leetcode.com/problems/maximize-the-number-of-target-nodes-after-connecting-trees-ii/

# python3 -m unittest graphs/3373-maximize-the-number-of-target-nodes-after-connecting-trees-ii.py

class Solution(unittest.TestCase):
    # Approach: Depth-First Search
    # Time: O(n+m)
    # Space: O(n+m)
    def maxTargetNodes(self, edges1: List[List[int]], edges2: List[List[int]]) -> List[int]:
        def dfs(parent: int, node: int, adj: List[List[int]], color: List[int], col: int = 1) -> int:
            targets = col
            color[node] = col
            for nei in adj[node]:
                if nei != parent:
                    targets += dfs(node, nei, adj, color, col^1)
            return targets

        def make_adj_list(edges: List[List[int]]) -> List[List[int]]:
            n = len(edges) + 1
            adj = [[] for _ in range(n)]
            for a, b in edges:
                adj[a].append(b)
                adj[b].append(a)
            return adj

        m = len(edges2)+1
        evens2 = dfs(-1, 0, make_adj_list(edges2), [0]*m)
        max2 = max(evens2, m - evens2)

        n = len(edges1)+1
        color1 = [0] * n
        evens1 = dfs(-1, 0, make_adj_list(edges1), color1)
        count1 = [n-evens1, evens1]

        return [count1[color1[node]] + max2 for node in range(n)]

    def test(self):
        for edges1, edges2, expected in [
            ([[0,1],[0,2],[2,3],[2,4]], [[0,1],[0,2],[0,3],[2,7],[1,4],[4,5],[4,6]], [8,7,7,8,8]),
            ([[0,1],[0,2],[0,3],[0,4]], [[0,1],[1,2],[2,3]], [3,6,6,6,6]),
        ]:
            output = self.maxTargetNodes(edges1, edges2)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
