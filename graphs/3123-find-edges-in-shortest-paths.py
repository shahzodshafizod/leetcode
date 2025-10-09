from typing import List, Tuple
import unittest
import heapq

# https://leetcode.com/problems/find-edges-in-shortest-paths/

# python3 -m unittest graphs/3123-find-edges-in-shortest-paths.py


class Solution(unittest.TestCase):
    def findAnswer(self, n: int, edges: List[List[int]]) -> List[bool]:
        graph: List[List[Tuple[int, int, int]]] = [[] for _ in range(n)]
        for i, (a, b, w) in enumerate(edges):
            graph[a].append((b, w, i))
            graph[b].append((a, w, i))
        prev: List[List[Tuple[int, int]]] = [[] for _ in range(n)]
        best = [(1 << 32) - 1] * n
        best[0] = 0
        pq: List[Tuple[int, int]] = [(0, 0)]  # cost, node
        while pq:  # Dijkstra's algorithm -> Time: O(E log E)
            cost, node = heapq.heappop(pq)
            if cost > best[node]:
                continue
            for nei, weight, eidx in graph[node]:
                weight += best[node]
                if weight < best[nei]:
                    best[nei] = weight
                    prev[nei] = [(node, eidx)]
                    heapq.heappush(pq, (weight, nei))
                elif weight == best[nei]:
                    prev[nei].append((node, eidx))

        answer = [False] * len(edges)
        seen = [False] * n

        def track(node: int):  # DFS -> Time: O(V + E)
            if seen[node]:
                return
            seen[node] = True
            for prv, eidx in prev[node]:
                answer[eidx] = True
                track(prv)

        track(n - 1)

        return answer

    def test(self):
        for n, edges, expected in [
            (6, [[0, 1, 4], [0, 2, 1], [1, 3, 2], [1, 4, 3], [1, 5, 1], [2, 3, 1], [3, 5, 3], [4, 5, 2]], [True, True, True, False, True, True, True, False]),
            (4, [[2, 0, 1], [0, 1, 1], [0, 3, 4], [3, 2, 2]], [True, False, False, True]),
        ]:
            output = self.findAnswer(n, edges)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
