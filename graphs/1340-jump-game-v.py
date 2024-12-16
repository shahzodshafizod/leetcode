from typing import List
import unittest

# https://leetcode.com/problems/jump-game-v/

# python3 -m unittest graphs/1340-jump-game-v.py

class Solution(unittest.TestCase):
    # # Top-Down Dynamic Programming
    # # Time: O(n x d)
    # # Space: O(n)
    # def maxJumps(self, arr: List[int], d: int) -> int:
    #     n = len(arr)
    #     dp = [0] * n
    #     def dfs(idx: int) -> int:
    #         if dp[idx]:
    #             return dp[idx]
    #         dp[idx] = 1
    #         # move left
    #         for j in range(1, d+1):
    #             next = idx-j
    #             if next < 0 or arr[next] >= arr[idx]:
    #                 break
    #             dp[idx] = max(dp[idx], 1 + dfs(next))
    #         # move right
    #         for j in range(1, d+1):
    #             next = idx+j
    #             if next >= n or arr[next] >= arr[idx]:
    #                 break
    #             dp[idx] = max(dp[idx], 1 + dfs(next))
    #         return dp[idx]
    #     steps = 0
    #     for idx in range(n):
    #         steps = max(steps, dfs(idx))
    #     return steps

    # Approach: Monotonic Stack + DFS on Graph
    # Time: O(3n) = O(n)
    # Space: O(3n) = O(n) (graph, stack, dp)
    def maxJumps(self, arr: List[int], d: int) -> int:
        n = len(arr)
        graph = [[] for _ in range(n)] # adjacency list
        def make_graph(start: int, end: int, step: int) -> None:
            stack = []
            for curr in range(start, end, step):
                while stack and arr[stack[-1]] < arr[curr]:
                    neighbor = stack.pop()
                    if abs(curr-neighbor) <= d:
                      graph[curr].append(neighbor)
                stack.append(curr)
        make_graph(0, n, 1) # add prev neighbors
        make_graph(n-1, -1, -1) # add next neighbors
        dp = [1] * n
        def dfs(idx: int) -> int:
            if dp[idx] != 1:
                return dp[idx]
            for next in graph[idx]:
                dp[idx] = max(dp[idx], 1+dfs(next))
            return dp[idx]
        steps = 0
        for idx in range(n):
            steps = max(steps, dfs(idx))
        return steps

    def test(self):
        for arr, d, expected in [
            ([6,4,14,6,8,13,9,7,10,6,12], 2, 4),
            ([3,3,3,3,3], 3, 1),
            ([7,6,5,4,3,2,1], 1, 7),
            ([1,1,1,1,1,2,2,2,2,2], 1, 2),
            # ([39,1,1,19,40,34,87,44,30,3,89,55,81,97,84,52,10,8,96,69,17,48,93,84,10,48,1,93,65,24,100,26,24,33,52,17,15,26,8,87,69,47,61,58,78,52,2,72,23,9,3,27,36,38,88,20,21,79,5,67,22,24,39,7,17,29,3,97,36,51,91,53,98,48,83,52,14,71,91,46,42,88,44,52,74,8,39,11,48,59,98,34,43,94,46,20,26,62,6,36,59,77,23,93,6,93,64,18,33,69,56,48,54,98,98,53,14,97,47,50,33,87,10,51,92,1,14,27,19,34,83,65,48,44,82,51,81,83,23,8,63,70,76,83,46,84,20,7,37,4,69,63,84,71,91,78,58,25,63,85,98,78,21], 62, 13),
        ]:
            output = self.maxJumps(arr, d)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
