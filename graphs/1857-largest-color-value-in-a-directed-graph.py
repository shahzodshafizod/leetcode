from collections import deque
from typing import List
import unittest

# https://leetcode.com/problems/largest-color-value-in-a-directed-graph/

# python3 -m unittest graphs/1857-largest-color-value-in-a-directed-graph.py


class Solution(unittest.TestCase):
    # # Approach: Depth-First Search
    # # Time: O(26n) = O(n)
    # # Space: O(26n) = O(n)
    # def largestPathValue(self, colors: str, edges: List[List[int]]) -> int:
    #     n = len(colors)
    #     adj = [[] for _ in range(n)]
    #     for src, dst in edges:
    #         adj[src].append(dst)
    #     colors = [ord(c) - ord('a') for c in colors]
    #     count = [[0] * 26 for _ in range(n)]
    #     memo = [-1] * n
    #     def dfs(node: int) -> int:
    #         if memo[node] != -1:
    #             return memo[node]
    #         color = colors[node]
    #         memo[node] = n+1
    #         for nei in adj[node]:
    #             if dfs(nei) > n:
    #                 return n+1
    #             for c in range(26):
    #                 count[node][c] = max(
    #                     count[node][c],
    #                     count[nei][c],
    #                 )
    #         count[node][color] += 1
    #         memo[node] = max(count[node])
    #         return memo[node]
    #     value = max(dfs(node) for node in range(n))
    #     return -1 if value > n else value

    # Approach: Breadth-First Search
    # Time: O(n+m)
    # Space: O(n)
    def largestPathValue(self, colors: str, edges: List[List[int]]) -> int:
        n = len(colors)
        adj = [[] for _ in range(n)]
        indegrees = [0] * n
        for src, dst in edges:
            adj[src].append(dst)
            indegrees[dst] += 1
        queue = deque([n for n in range(n) if indegrees[n] == 0])
        colors = [ord(c) - ord('a') for c in colors]
        dp = [[0] * 26 for _ in range(n)]
        value = 0
        while queue:
            node = queue.popleft()
            n -= 1
            dp[node][colors[node]] += 1
            value = max(dp[node] + [value])
            for nei in adj[node]:
                for c in range(26):
                    dp[nei][c] = max(
                        dp[nei][c],
                        dp[node][c],
                    )
                indegrees[nei] -= 1
                if indegrees[nei] == 0:
                    queue.append(nei)
        return value if n == 0 else -1

    def test(self):
        for colors, edges, expected in [
            ("a", [[0, 0]], -1),
            ("ab", [[0, 1], [1, 1]], -1),
            ("abaca", [[0, 1], [0, 2], [2, 3], [3, 4]], 3),
            ("aaaca", [[0, 1], [0, 2], [2, 3], [3, 4], [4, 0]], -1),
        ]:
            output = self.largestPathValue(colors, edges)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
