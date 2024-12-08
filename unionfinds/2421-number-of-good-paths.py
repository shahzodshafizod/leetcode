from collections import defaultdict
from typing import List
import unittest

# https://leetcode.com/problems/number-of-good-paths/

# python3 -m unittest unionfinds/2421-number-of-good-paths.py

class Solution(unittest.TestCase):
    # # Approach: Brute-Force with Union Find (Time Limit Exceeded)
    # # Time: O(?)
    # # Space: O(?)
    # def numberOfGoodPaths(self, vals: List[int], edges: List[List[int]]) -> int:
    #     n = len(vals)
    #     parent = list(range(n))
    #     def find(x: int) -> int:
    #         while parent[x] != x:
    #             parent[x] = parent[parent[x]]
    #             x = parent[x]
    #         return x
    #     def union(x: int, y: int) -> None:
    #         px = find(x)
    #         py = find(y)
    #         parent[px] = py
    #     graph = defaultdict(list)
    #     for node1, node2 in edges:
    #         graph[node1].append(node2)
    #         graph[node2].append(node1)
    #     same = defaultdict(set)
    #     result = n
    #     for val, node in sorted([(vals[node],node) for node in range(len(vals))]):
    #         for neighbor in graph[node]:
    #             if vals[neighbor] <= val:
    #                 union(neighbor, node)
    #         for twin in same[vals[node]]:
    #             if find(twin) == find(node):
    #                 result += 1
    #         same[vals[node]].add(node)
    #     return result

    # Approach: Union Find
    # Time: O(nlogn)
    # Space: O(n)
    def numberOfGoodPaths(self, vals: List[int], edges: List[List[int]]) -> int:
        n = len(vals)
        parent = list(range(n))
        def find(x: int) -> int: # O(logn)
            while parent[x] != x:
                parent[x] = parent[parent[x]]
                x = parent[x]
            return x
        count = [1]*n
        result = n
        edges.sort(key=lambda edge: max(vals[edge[0]],vals[edge[1]]))
        for x, y in edges: # O(n)
            px, py = find(x), find(y)
            if vals[px] < vals[py]:
                parent[px] = py
            elif vals[px] > vals[py]:
                parent[py] = px
            else:
                parent[py] = px
                result += count[px] * count[py]
                count[px] += count[py]
        return result

    def test(self) -> None:
        for vals, edges, expected in [
            ([1], [], 1),
            ([2,1,1], [[0,1],[2,0]], 3),
            ([1,3,2,1,3], [[0,1],[0,2],[2,3],[2,4]], 6),
            ([1,1,2,2,3], [[0,1],[1,2],[2,3],[2,4]], 7),
        ]:
            output = self.numberOfGoodPaths(vals, edges)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
