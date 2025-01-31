from collections import deque
from typing import List
import unittest

# https://leetcode.com/problems/is-graph-bipartite/

# python3 -m unittest graphs/0785-is-graph-bipartite.py

class Solution(unittest.TestCase):
    # # Approach #1: Union Find
    # # Time: O(v+e), v=# of vertices, e=# of edges
    # # Space: O(v)
    # def isBipartite(self, graph: List[List[int]]) -> bool:
    #     n = len(graph)
    #     parent = [node for node in range(n)]
    #     def find(x: int) -> int:
    #         if parent[x] != x:
    #             parent[x] = find(parent[x])
    #         return parent[x]
    #     def union(x: int, y: int):
    #         px = find(x)
    #         py = find(y)
    #         if px != py:
    #             parent[py] = px
    #     for node in range(n):
    #         for neighbor in graph[node]:
    #             if find(node) == find(neighbor):
    #                 return False
    #             union(graph[node][0], neighbor)
    #     return True

    # # Approach #2: Breadth-First Search
    # # Time: O(v+e), v=# of vertices (nodes), e=# of edges
    # # Space: O(v+e)
    # def isBipartite(self, graph: List[List[int]]) -> bool:
    #     n = len(graph)
    #     colors = [0] * n
    #     def check(node: int) -> bool:
    #         if colors[node] != 0:
    #             return True
    #         queue = deque([node])
    #         colors[node] = -1
    #         while queue:
    #             for _ in range(len(queue)):
    #                 curr = queue.popleft()
    #                 for next in graph[curr]:
    #                     if colors[next] == colors[curr]:
    #                         return False
    #                     elif colors[next] == 0:
    #                         colors[next] = -1 * colors[curr]
    #                         queue.append(next)
    #         return True
    #     return all(check(node) for node in range(n))

    # Approach #3: Depth-First Search
    # Time: O(v+e), v=# of vertices (nodes), e=# of edges
    # Space: O(v+e)
    def isBipartite(self, graph: List[List[int]]) -> bool:
        n = len(graph)
        colors = [0] * n
        def dfs(node: int, color: int) -> bool:
            colors[node] = color
            for neighbor in graph[node]:
                if colors[neighbor] == -color:
                    continue
                if colors[neighbor] == color or not dfs(neighbor, -color):
                    return False
            return True
        for node in range(n):
            if colors[node] == 0 and not dfs(node, 1):
                return False
        return True

    def test(self):
        for graph, expected in [
            ([[1],[0]], True),
            ([[1,2,3],[0,2],[0,1,3],[0,2]], False),
            ([[1,3],[0,2],[1,3],[0,2]], True),
            ([[4],[],[4],[4],[0,2,3]], True),
            ([[],[2,4,6],[1,4,8,9],[7,8],[1,2,8,9],[6,9],[1,5,7,8,9],[3,6,9],[2,3,4,6,9],[2,4,5,6,7,8]], False),
            ([[3,4,6],[3,6],[3,6],[0,1,2,5],[0,7,8],[3],[0,1,2,7],[4,6],[4],[]], True),
            ([[1,2,3],[0,3,4],[0,3],[0,1,2],[1]], False),
            ([[1,2,3,4],[0,2,3],[0,1,3,4],[0,1,2,4],[0,2,3]], False),
        ]:
            output = self.isBipartite(graph)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
