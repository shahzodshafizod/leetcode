import heapq
import unittest
from typing import List

# https://leetcode.com/problems/modify-graph-edge-weights/

# python3 -m unittest graphs/2699-modify-graph-edge-weights.py


# Time: O((E+V)logV)
# Space: O(E+V)
class Solution(unittest.TestCase):
    def modifiedGraphEdges(
        self, n: int, edges: List[List[int]], source: int, destination: int, target: int
    ) -> List[List[int]]:
        # 1. Graph Representation
        adjacency_list = [[] for _ in range(n)]
        for idx, (u, v, _) in enumerate(edges):
            adjacency_list[u].append((v, idx))
            adjacency_list[v].append((u, idx))

        # 2. Distance Arrays
        distances = [[float("inf")] * 2 for _ in range(n)]

        def runDijkstra(index: int, difference: int) -> None:
            priority_queue = [(0, source)]
            distances[source][index] = 0
            while priority_queue:
                curr_dist, curr_node = heapq.heappop(priority_queue)
                if curr_dist > distances[curr_node][index]:
                    continue
                for next_node, edge_index in adjacency_list[curr_node]:
                    next_weight = edges[edge_index][2]
                    if next_weight == -1:
                        next_weight = 1
                    if index == 1 and edges[edge_index][2] == -1:
                        new_weight = difference + distances[next_node][0] - distances[curr_node][1]
                        if new_weight > next_weight:
                            next_weight = new_weight
                            edges[edge_index][2] = new_weight
                    if distances[curr_node][index] + next_weight < distances[next_node][index]:
                        distances[next_node][index] = distances[curr_node][index] + next_weight
                        heapq.heappush(priority_queue, (distances[next_node][index], next_node))

        # 3. First Dijkstra Run
        runDijkstra(0, 0)

        # 4. Check Feasibility
        different = target - distances[destination][0]
        if different < 0:
            return []

        # 5. Second Dijkstra Run
        runDijkstra(1, different)

        # 6. Final Check
        if distances[destination][1] < target:
            return []

        # 7. Update Edges
        for edge in edges:
            if edge[2] == -1:
                edge[2] = 1

        return edges

    def test(self):
        for n, edges, source, destination, target, expected in [
            (
                5,
                [[4, 1, -1], [2, 0, -1], [0, 3, -1], [4, 3, -1]],
                0,
                1,
                5,
                [[4, 1, 1], [2, 0, 3], [0, 3, 3], [4, 3, 1]],
            ),
            (3, [[0, 1, -1], [0, 2, 5]], 0, 2, 6, []),
            (4, [[1, 0, 4], [1, 2, 3], [2, 3, 5], [0, 3, -1]], 0, 2, 6, [[1, 0, 4], [1, 2, 3], [2, 3, 5], [0, 3, 1]]),
            (4, [[0, 1, -1], [1, 2, -1], [3, 1, -1], [3, 0, 2], [0, 2, 5]], 2, 3, 8, []),
            (
                4,
                [[0, 1, -1], [2, 0, 2], [3, 2, 6], [2, 1, 10], [3, 0, -1]],
                1,
                3,
                12,
                [[0, 1, 11], [2, 0, 2], [3, 2, 6], [2, 1, 10], [3, 0, 1]],
            ),
            (
                5,
                [[1, 4, 1], [2, 4, -1], [3, 0, 2], [0, 4, -1], [1, 3, 10], [1, 0, 10]],
                0,
                2,
                15,
                [[1, 4, 1], [2, 4, 4], [3, 0, 2], [0, 4, 14], [1, 3, 10], [1, 0, 10]],
            ),
            (
                5,
                [[0, 2, 5], [2, 1, -1], [2, 4, 3], [3, 4, 5], [4, 0, 1], [0, 3, -1], [2, 3, -1]],
                0,
                1,
                9,
                [[0, 2, 5], [2, 1, 5], [2, 4, 3], [3, 4, 5], [4, 0, 1], [0, 3, 7], [2, 3, 3]],
            ),
        ]:
            output = self.modifiedGraphEdges(n, edges, source, destination, target)
            self.assertListEqual(expected, output, f"expected: {expected}, output: {output}")
