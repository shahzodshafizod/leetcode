from collections import defaultdict
from typing import List
import unittest
from pkg.unionfind import UnionFind  # pylint: disable=unused-import

# https://leetcode.com/problems/count-the-number-of-complete-components/

# python3 -m unittest graphs/2685-count-the-number-of-complete-components.py


class Solution(unittest.TestCase):
    # # Approach #1: Depth-First Search + Union Find
    # # Time: O(2E+2V) = O(E+V)
    # # Space: O(E+4V) = O(E+V)
    # def countCompleteComponents(self, n: int, edges: List[List[int]]) -> int:
    #     uf = UnionFind(n)
    #     graph = defaultdict(list)
    #     for a,b in edges:
    #         uf.Union(a, b)
    #         graph[a].append(b)
    #         graph[b].append(a)
    #     comp_edges = defaultdict(int)
    #     for a,_ in edges:
    #         comp_edges[uf.Find(a)] += 1
    #     seen = [False] * n
    #     def dfs(node: int) -> int:
    #         if seen[node]:
    #             return 0
    #         seen[node] = True
    #         return 1 + sum(dfs(n) for n in graph[node])
    #     count = 0
    #     for node in range(n):
    #         if not seen[node]:
    #             comp_nodes = dfs(node)
    #             if comp_nodes*(comp_nodes-1)//2 == comp_edges[uf.Find(node)]:
    #                 count += 1
    #     return count

    # Approach #2: Depth-First Search
    # Time: O(E+2V) = O(E+V)
    # Space: O(E+3V) = O(E+V)
    def countCompleteComponents(self, n: int, edges: List[List[int]]) -> int:
        graph = defaultdict(list)
        for a, b in edges:
            graph[a].append(b)
            graph[b].append(a)
        seen = [False] * n

        def dfs(node: int) -> List[int]:
            if seen[node]:
                return []
            seen[node] = True
            vertices = [node]
            for nei in graph[node]:
                vertices.extend(dfs(nei))
            return vertices

        count = 0
        for node in range(n):
            if seen[node]:
                continue
            vertices = dfs(node)
            if all(len(vertices) - 1 == len(graph[node]) for node in vertices):
                count += 1
        return count

    def test(self):
        for n, edges, expected in [
            (6, [[0, 1], [0, 2], [1, 2], [3, 4]], 3),
            (6, [[0, 1], [0, 2], [1, 2], [3, 4], [3, 5]], 1),
            (5, [[1, 2], [3, 4], [1, 4], [2, 3], [1, 3], [2, 4]], 2),
            (3, [[1, 0], [2, 0]], 0),
            (4, [[1, 0], [2, 0], [3, 1], [3, 2]], 0),
            (2, [[1, 0]], 1),
        ]:
            output = self.countCompleteComponents(n, edges)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
