import heapq
from typing import List
import unittest


# https://leetcode.com/problems/maximize-sum-of-weights-after-edge-removals/

# python3 -m unittest trees/3367-maximize-sum-of-weights-after-edge-removals.py


class Solution(unittest.TestCase):
    # # Approach #1: Brute-Force (Time Limit Exceeded)
    # # Time: O(nn)
    # # Space: O(n)
    # def maximizeSumOfWeights(self, edges: List[List[int]], k: int) -> int:
    #     n = len(edges) + 1
    #     degrees = [0] * n  # for nodes
    #     for u, v, _ in edges:
    #         degrees[u] += 1
    #         degrees[v] += 1
    #     deleted = [False] * (n - 1)  # for edges

    #     def dp(eid: int) -> int:
    #         weights = 0
    #         for idx, (u, v, w) in enumerate(edges):
    #             if deleted[idx]:
    #                 continue
    #             if degrees[u] > k or degrees[v] > k:
    #                 weights = 0
    #                 break
    #             weights += w

    #         if eid == n - 1:
    #             return weights

    #         u, v, w = edges[eid]

    #         # 1. skip
    #         weights = max(weights, dp(eid + 1))

    #         # 2. delete
    #         deleted[eid] = True
    #         degrees[u] -= 1
    #         degrees[v] -= 1
    #         weights = max(weights, dp(eid + 1))
    #         deleted[eid] = False
    #         degrees[u] += 1
    #         degrees[v] += 1

    #         return weights

    #     return dp(0)

    # Approach #2: Depth-First Search
    # Time: O(n)
    # Space: O(n)
    def maximizeSumOfWeights(self, edges: List[List[int]], k: int) -> int:
        n = len(edges) + 1
        graph: List[List[tuple[int, int]]] = [[] for _ in range(n)]
        for u, v, w in edges:
            graph[u].append((v, w))
            graph[v].append((u, w))

        def dfs(node: int, parent: int = -1) -> tuple[int, int]:
            res = 0
            diff: List[int] = []
            for nei, weight in graph[node]:
                if nei == parent:
                    continue
                # w1 is the weight sum of node i with atmost k - 1 children.
                # w2 is the weight sum of node i with atmost k children.
                w1, w2 = dfs(nei, node)
                res += w2
                diff.append(max(0, w1 + weight - w2))
            return res + sum(heapq.nlargest(k - 1, diff)), res + sum(heapq.nlargest(k, diff))

        return dfs(0)[1]

    def test(self):
        for edges, k, expected in [
            ([[0, 1, 4], [0, 2, 2], [2, 3, 12], [2, 4, 6]], 2, 22),
            ([[0, 1, 5], [1, 2, 10], [0, 3, 15], [3, 4, 20], [3, 5, 5], [0, 6, 10]], 3, 65),
            ([[0, 1, 34], [0, 2, 17]], 1, 34),
            ([[0, 2, 10], [0, 1, 23]], 1, 23),
        ]:
            output = self.maximizeSumOfWeights(edges, k)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
