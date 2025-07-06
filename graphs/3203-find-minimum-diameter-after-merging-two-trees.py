from collections import deque
from typing import List
import unittest

# https://leetcode.com/problems/find-minimum-diameter-after-merging-two-trees/

# python3 -m unittest graphs/3203-find-minimum-diameter-after-merging-two-trees.py


class Solution(unittest.TestCase):
    # # Approach: Depth-First Search
    # # Time: O(n+m), n=len(edges1), m=len(edges2)
    # # Space: O(n+m)
    # def minimumDiameterAfterMerge(self, edges1: List[List[int]], edges2: List[List[int]]) -> int:
    #     def dfs(graph: List[List[int]], parent: int, node: int) -> tuple[int, int]: # endpoint, long_path
    #         endpoint, distance = node, 0
    #         for next in graph[node]:
    #             if next == parent: continue
    #             point, dist = dfs(graph, node, next)
    #             dist += 1 # include the edge between node and next
    #             if dist > distance:
    #                 distance = dist
    #                 endpoint = point
    #         return endpoint, distance
    #     def find_diameter(edges: List[List[int]]) -> int:
    #         n = len(edges)+1
    #         graph = [[] for _ in range(n)]
    #         for a, b in edges:
    #             graph[a].append(b)
    #             graph[b].append(a)
    #         endpoint, _ = dfs(graph, -1, 0)
    #         _, diameter = dfs(graph, -1, endpoint)
    #         return diameter
    #     d1 = find_diameter(edges1)
    #     d2 = find_diameter(edges2)
    #     return max(d1, d2, (d1+1)//2 + (d2+1)//2 + 1)

    # Approach: Breadth-First Search
    # Time: O(n+m), n=len(edges1), m=len(edges2)
    # Space: O(n+m)
    def minimumDiameterAfterMerge(self, edges1: List[List[int]], edges2: List[List[int]]) -> int:
        def bfs(graph: List[List[int]], node: int) -> tuple[int, int]:  # endpoint, long_path
            queue = deque([(-1, node, 0)])
            while queue:
                parent, node, distance = queue.popleft()
                for nei in graph[node]:
                    if nei == parent:
                        continue
                    queue.append((node, nei, distance + 1))
            return node, distance

        def find_diameter(edges: List[List[int]]) -> int:
            n = len(edges) + 1
            graph = [[] for _ in range(n)]
            for a, b in edges:
                graph[a].append(b)
                graph[b].append(a)
            endpoint, _ = bfs(graph, 0)
            _, diameter = bfs(graph, endpoint)
            return diameter

        d1 = find_diameter(edges1)
        d2 = find_diameter(edges2)
        return max(d1, d2, (d1 + 1) // 2 + (d2 + 1) // 2 + 1)

    def test(self):
        for edges1, edges2, expected in [
            ([], [], 1),
            ([], [[0, 1]], 2),
            ([[0, 1], [0, 2], [0, 3]], [[0, 1]], 3),
            ([], [[0, 1], [1, 2]], 2),
            ([[0, 1], [2, 0], [3, 2], [3, 6], [8, 7], [4, 8], [5, 4], [3, 5], [3, 9]], [[0, 1], [0, 2], [0, 3]], 7),
            (
                [[0, 1], [0, 2], [0, 3], [2, 4], [2, 5], [3, 6], [2, 7]],
                [[0, 1], [0, 2], [0, 3], [2, 4], [2, 5], [3, 6], [2, 7]],
                5,
            ),
        ]:
            output = self.minimumDiameterAfterMerge(edges1, edges2)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
