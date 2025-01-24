from typing import List
import unittest

# https://leetcode.com/problems/find-eventual-safe-states/

# python3 -m unittest graphs/0802-find-eventual-safe-states.py

class Solution(unittest.TestCase):
    # Approach: Depth-First Search
    # Time: O(v+e), v=# of nodes, e=# of edges
    # Space: O(v)
    def eventualSafeNodes(self, graph: List[List[int]]) -> List[int]:
        n = len(graph)
        safe = [None] * n
        seen = [False] * n
        def dfs(node: int) -> bool:
            if seen[node]:
                return False
            if safe[node] != None:
                return safe[node]
            seen[node] = True
            safe[node] = True
            for next in graph[node]:
                safe[node] = safe[node] and dfs(next)
            seen[node] = False
            return safe[node]
        return [node for node in range(n) if dfs(node)]

    def test(self):
        for graph, expected in [
            ([[1,2],[2,3],[5],[0],[5],[],[]], [2,4,5,6]),
            ([[1,2,3,4],[1,2],[3,4],[0,4],[]], [4]),
        ]:
            output = self.eventualSafeNodes(graph)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
