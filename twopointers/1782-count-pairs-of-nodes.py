from collections import defaultdict
from typing import List
import unittest

# https://leetcode.com/problems/count-pairs-of-nodes/

# python3 -m unittest twopointers/1782-count-pairs-of-nodes.py


class Solution(unittest.TestCase):
    # # Approach: Brute-Force
    # # Time: O(q*n^2), q=len(queries)
    # # Space: O(n+e), e=len(edges)
    # def countPairs(self, n: int, edges: List[List[int]], queries: List[int]) -> List[int]:
    #     degree = defaultdict(int)
    #     shared = defaultdict(int)
    #     for u, v in edges:
    #         degree[u] += 1
    #         degree[v] += 1
    #         shared[(min(u,v),max(u,v))] += 1
    #     answer = [0] * len(queries)
    #     for idx, query in enumerate(queries):
    #         for u in range(1, n+1):
    #             for v in range(u+1, n+1):
    #                 if degree[u] + degree[v] - shared[(u,v)] > query:
    #                     answer[idx] += 1
    #     return answer

    # Approach: Two Pointers
    # Time: O(q*(n+e) + nlogn), q=len(queries)
    # Space: O(n+e), e=len(edges)
    def countPairs(self, n: int, edges: List[List[int]], queries: List[int]) -> List[int]:
        degree = [0] * (n + 1)
        shared = defaultdict(int)
        for u, v in edges:  # O(e)
            degree[u] += 1
            degree[v] += 1
            shared[(min(u, v), max(u, v))] += 1
        answer = [0] * len(queries)
        sorted_degree = sorted(degree[1:])  # O(nlogn), [1:] means excluding node 0
        for idx, query in enumerate(queries):  # O(q)
            left, right = 0, n - 1
            while left <= right:  # O(n)
                if sorted_degree[left] + sorted_degree[right] > query:
                    answer[idx] += right - left
                    right -= 1
                else:
                    left += 1
            for (u, v), cnt in shared.items():  # O(e)
                # subtract those edges that were added, but
                # considering shared ones between them breaks condition
                if degree[u] + degree[v] > query >= degree[u] + degree[v] - cnt:
                    answer[idx] -= 1
        return answer

    def test(self):
        for n, edges, queries, expected in [
            (4, [[1, 2], [2, 4], [1, 3], [2, 3], [2, 1]], [2, 3], [6, 5]),
            (5, [[1, 5], [1, 5], [3, 4], [2, 5], [1, 3], [5, 1], [2, 3], [2, 5]], [1, 2, 3, 4, 5], [10, 10, 9, 8, 6]),
        ]:
            output = self.countPairs(n, edges, queries)
            self.assertListEqual(expected, output, f"expected: {expected}, output: {output}")
