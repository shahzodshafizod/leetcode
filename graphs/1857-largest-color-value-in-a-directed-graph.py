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
    #     graph = [[] for _ in range(n)]
    #     for src, dst in edges:
    #         graph[src].append(dst)
    #     colids = [ord(c) - ord('a') for c in colors]
    #     dp = [[0] * 26 for _ in range(n)]
    #     memo = [None] * n
    #     def dfs(node: int) -> int:
    #         if memo[node] == -1:
    #             return float("inf")
    #         if memo[node] != None:
    #             return memo[node]
    #         memo[node] = -1 # sign of detecting cycles
    #         dp[node][colids[node]] = 1
    #         for next in graph[node]:
    #             if dfs(next) == float("inf"):
    #                 return float("inf")
    #             for c in range(26):
    #                 dp[node][c] = max(
    #                     dp[node][c],
    #                     dp[next][c] + int(c == colids[node]),
    #                 )
    #         memo[node] = max(dp[node])
    #         return memo[node]
    #     count = max(dfs(node) for node in range(n))
    #     return count if count < float("inf") else -1

    # Approach: Breadth-First Search
    # Time: O(26n) = O(n)
    # Space: O(26n) = O(n)
    def largestPathValue(self, colors: str, edges: List[List[int]]) -> int:
        n = len(colors)
        graph = [[] for _ in range(n)]
        indegree = [0] * n
        for src, dst in edges:
            graph[src].append(dst)
            indegree[dst] += 1
        queue = deque()
        for node in range(n):
            if indegree[node] == 0:
                queue.append(node)
        colids = [ord(c)-ord('a') for c in colors]
        dp = [[0] * 26 for _ in range(n)]
        result = 0
        while queue:
            curr = queue.popleft()
            n -= 1
            dp[curr][colids[curr]] += 1
            result = max(result, max(dp[curr]))
            for next in graph[curr]:
                for c in range(26):
                    dp[next][c] = max(
                        dp[next][c],
                        dp[curr][c],
                    )
                indegree[next] -= 1
                if indegree[next] == 0:
                    queue.append(next)
        return result if n == 0 else -1

    def test(self):
        for colors, edges, expected in [
            ("a", [[0,0]], -1),
            ("ab", [[0,1],[1,1]], -1),
            ("abaca", [[0,1],[0,2],[2,3],[3,4]], 3),
        ]:
            output = self.largestPathValue(colors, edges)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
