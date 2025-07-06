from collections import deque
from typing import List
import unittest

# https://leetcode.com/problems/parallel-courses-iii/

# python3 -m unittest graphs/2050-parallel-courses-iii.py


class Solution(unittest.TestCase):
    # # Approach: Depth-First Search
    # # Time: O(E+V), E=# of edges, V=# of nodes/vertices
    # # Space: O(E+V)
    # def minimumTime(self, n: int, relations: List[List[int]], time: List[int]) -> int:
    #     graph = [[] for _ in range(n+1)]
    #     for src, dst in relations: graph[src].append(dst)
    #     memo = [None] * (n+1)
    #     def dfs(src: int) -> int:
    #         if memo[src] != None: return memo[src]
    #         memo[src] = max((dfs(dst) for dst in graph[src]), default=0) + time[src-1]
    #         return memo[src]
    #     return max(dfs(node) for node in range(1, n+1))

    # Approach: Breadth-First Search
    # Time: O(E+V), E=# of edges, V=# of nodes/vertices
    # Space: O(E+V)
    def minimumTime(self, n: int, relations: List[List[int]], time: List[int]) -> int:
        graph = [[] for _ in range(n)]
        indegree = [0] * n
        for src, dst in relations:
            graph[src - 1].append(dst - 1)
            indegree[dst - 1] += 1
        queue = deque([node for node, degree in enumerate(indegree) if degree == 0])
        dist = time.copy()
        while queue:
            src = queue.popleft()
            for dst in graph[src]:
                dist[dst] = max(dist[dst], dist[src] + time[dst])
                indegree[dst] -= 1
                if indegree[dst] == 0:
                    queue.append(dst)
        return max(dist)

    def test(self):
        for n, relations, time, expected in [
            (1, [], [10], 10),
            (2, [], [3, 5], 5),
            (3, [[1, 3], [2, 3]], [3, 2, 5], 8),
            (2, [[2, 1]], [10000, 10000], 20000),
            (10, [[1, 3]], [5, 95, 21, 41, 87, 9, 60, 67, 36, 96], 96),
            (5, [[1, 5], [2, 5], [3, 5], [3, 4], [4, 5]], [1, 2, 3, 4, 5], 12),
            (7, [[1, 7], [1, 4], [1, 3], [2, 3], [4, 3], [5, 3], [7, 3], [7, 6]], [6, 5, 1, 8, 2, 9, 4], 19),
            (7, [[1, 7], [1, 4], [1, 3], [2, 3], [4, 3], [5, 3], [7, 3], [7, 6]], [6, 5, 1, 8, 2, 9, 4], 19),
        ]:
            output = self.minimumTime(n, relations, time)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
