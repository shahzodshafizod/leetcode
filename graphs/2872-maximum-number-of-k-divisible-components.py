from collections import deque
from typing import List, Tuple
import unittest

# https://leetcode.com/problems/maximum-number-of-k-divisible-components/

# python3 -m unittest graphs/2872-maximum-number-of-k-divisible-components.py


class Solution(unittest.TestCase):
    # # Approach: Depth-First Search (Graph)
    # # Time: O(n+e), n=# of nodes, e=# of edges
    # # Space: O(n+e)
    # def maxKDivisibleComponents(self, n: int, edges: List[List[int]], values: List[int], k: int) -> int:
    #     graph = [[] for _ in range(n)]
    #     for a, b in edges:
    #         graph[a].append(b)
    #         graph[b].append(a)
    #     count = 0
    #     def dfs(parent: int, node: int) -> int:
    #         total = (values[node] +
    #             sum(dfs(node, next) for next in graph[node] if next != parent))
    #         if total % k == 0:
    #             nonlocal count
    #             count += 1
    #         return total
    #     dfs(-1, 0)
    #     return count

    # Approach: Breadth-First Search (Graph)
    # Time: O(n+e), n=# of nodes, e=# of edges
    # Space: O(n+e)
    def maxKDivisibleComponents(self, n: int, edges: List[List[int]], values: List[int], k: int) -> int:
        if n < 2:
            return 1
        graph: List[List[int]] = [[] for _ in range(n)]
        indegree = [0] * n
        for a, b in edges:
            graph[a].append(b)
            graph[b].append(a)
            indegree[a] += 1
            indegree[b] += 1
        queue = deque(node for node in range(n) if indegree[node] == 1)
        count = 0
        while queue:
            node = queue.popleft()
            indegree[node] -= 1
            total = 0
            if values[node] % k == 0:
                count += 1
            else:
                total = values[node]
            for nei in graph[node]:
                if indegree[nei] == 0:
                    continue
                indegree[nei] -= 1
                values[nei] += total
                if indegree[nei] == 1:
                    queue.append(nei)
        return count

    def test(self):
        test_cases: List[Tuple[int, List[List[int]], List[int], int, int]] = [
            (1, [], [0], 1, 1),
            (1, [], [100], 100, 1),
            (5, [[0, 2], [1, 2], [1, 3], [2, 4]], [1, 8, 1, 4, 4], 6, 2),
            (7, [[0, 1], [0, 2], [1, 3], [1, 4], [2, 5], [2, 6]], [3, 0, 6, 1, 5, 2, 1], 3, 3),
            (4, [[0, 1], [1, 2], [1, 3]], [176968959, 450655404, 922326524, 897145068], 5, 1),
            (4, [[3, 0], [0, 2], [0, 1]], [742593846, 612800589, 813572140, 813559096], 47, 2),
            (5, [[1, 2], [1, 3], [0, 2], [2, 4]], [999999999, 999999999, 999999999, 999999999, 999999999], 5, 1),
            (
                7,
                [[2, 0], [0, 4], [2, 5], [1, 2], [5, 3], [4, 6]],
                [12154284, 649536765, 974051464, 821507385, 392654193, 770357917, 37707285],
                11,
                2,
            ),
            # (20, [[17,4],[4,2],[3,15],[15,19],[19,7],[7,12],[15,16],[19,18],[18,13],[19,9],[19,1],[12,8],[1,11],[7,6],[9,0],[15,10],[13,14],[18,5],[4,5]], [900160,891774,2283737,414736,265741,212260,2983538,280834,17654,221822,996740,1517078,374434,43488,610048,241798,611382,3234714,1156844,93794], 5784192, 2),
        ]
        for n, edges, values, k, expected in test_cases:
            output = self.maxKDivisibleComponents(n, edges, values, k)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
