from collections import defaultdict
from typing import List
import unittest

# https://leetcode.com/problems/longest-special-path/

# python3 -m unittest graphs/3425-longest-special-path.py


class Solution(unittest.TestCase):
    # # Approach 1: Brute Force
    # # Time: O(nn)
    # # Space: O(n)
    # def longestSpecialPath(self, edges: List[List[int]], nums: List[int]) -> List[int]:
    #     n = len(nums)
    #     graph: List[List[tuple[int, int]]] = [[] for _ in range(n)]
    #     for u, v, length in edges:
    #         graph[u].append((v, length))
    #         graph[v].append((u, length))

    #     downward: List[List[tuple[int, int]]] = [[] for _ in range(n)]
    #     seen = [False] * n

    #     def build_tree(parent: int, node: int):
    #         seen[node] = True
    #         for nei, length in graph[node]:
    #             if nei != parent and not seen[nei]:
    #                 downward[node].append((nei, length))
    #                 build_tree(node, nei)

    #     build_tree(-1, 0)

    #     max_length, min_cnt = 0, 1
    #     visited: dict[int, bool] = defaultdict(bool)

    #     def dfs(node: int, length: int, cnt: int) -> tuple[int, int]:
    #         visited[nums[node]] = True

    #         nonlocal max_length, min_cnt
    #         if length > max_length:
    #             max_length = length
    #             min_cnt = cnt
    #         elif length == max_length and cnt < min_cnt:
    #             min_cnt = cnt

    #         for nei, edge_len in downward[node]:
    #             if not visited[nums[nei]]:
    #                 dfs(nei, length + edge_len, cnt + 1)

    #         visited[nums[node]] = False
    #         return max_length, min_cnt

    #     for start in range(n):
    #         dfs(start, 0, 1)

    #     return [max_length, min_cnt]

    # Approach 2: Depth-First Search & Sliding Window
    # Time: O(n)
    # Space: O(n)
    def longestSpecialPath(self, edges: List[List[int]], nums: List[int]) -> List[int]:
        n = len(nums)
        graph: List[List[tuple[int, int]]] = [[] for _ in range(n)]
        for u, v, length in edges:
            graph[u].append((v, length))
            graph[v].append((u, length))

        result = [0, 1]
        costs: List[int] = []
        last: dict[int, int] = defaultdict(lambda: -1)

        def dfs(parent: int, node: int, cost: int, start: int):
            color = last[nums[node]]
            last[nums[node]] = len(costs)
            costs.append(cost)

            length = cost - costs[start]
            cnt = len(costs) - start

            if length > result[0]:
                result[0] = length
                result[1] = cnt
            elif length == result[0] and cnt < result[1]:
                result[1] = cnt

            for nei, ncost in graph[node]:
                if nei == parent:
                    continue
                nstart = last[nums[nei]] + 1 if start <= last[nums[nei]] else start
                dfs(node, nei, cost + ncost, nstart)
            last[nums[node]] = color
            costs.pop()

        dfs(-1, 0, 0, 0)
        return result

    def test(self):
        for edges, nums, expected in [
            (
                [[0, 1, 2], [1, 2, 3], [1, 3, 5], [1, 4, 4], [2, 5, 6]],
                [2, 1, 2, 1, 3, 1],
                [6, 2],
            ),
            (
                [[1, 0, 8]],
                [2, 2],
                [0, 1],
            ),
            (
                [[0, 1, 1], [1, 2, 1], [2, 3, 1]],
                [1, 2, 3, 4],
                [3, 4],
            ),
        ]:
            output = self.longestSpecialPath(edges, nums)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
