from typing import List, Any, Optional, Tuple
import heapq
import unittest

# https://leetcode.com/problems/design-graph-with-shortest-path-calculator/

# python3 -m unittest design/2642-design-graph-with-shortest-path-calculator.py


class Graph:
    # Time: O(E)
    # Space: O(E)
    def __init__(self, n: int, edges: List[List[int]]):
        self.adj: List[List[Tuple[int, int]]] = [[] for _ in range(n)]
        for edge in edges:
            self.addEdge(edge)

    def addEdge(self, edge: List[int]) -> None:
        frm, to, cost = edge
        self.adj[frm].append((cost, to))

    # Approach: Dijkstra's Algorithm
    # Time: O(E log E)
    # Space: O(E)
    def shortestPath(self, node1: int, node2: int) -> int:
        dist = [(1 << 32) - 1] * len(self.adj)
        pq = [(0, node1)]
        while pq:
            cost, node = heapq.heappop(pq)
            if cost >= dist[node]:
                continue
            if node == node2:
                return cost
            dist[node] = cost
            for nc, nei in self.adj[node]:
                nc += cost
                if nc < dist[nei]:
                    heapq.heappush(pq, (nc, nei))
        return -1


# Your Graph object will be instantiated and called as such:
# obj = Graph(n, edges)
# obj.addEdge(edge)
# param_2 = obj.shortestPath(node1,node2)


class Solution(unittest.TestCase):
    def test(self):
        test_cases: List[tuple[List[str], List[Any], List[Any]]] = [
            (
                ["Graph", "shortestPath", "shortestPath", "addEdge", "shortestPath"],
                [[4, [[0, 2, 5], [0, 1, 2], [1, 2, 1], [3, 0, 3]]], [3, 2], [0, 3], [[1, 3, 4]], [0, 3]],
                [None, 6, -1, None, 6],
            ),
        ]
        for commands, values, expected in test_cases:
            graph: Optional[Graph] = None
            for idx, command in enumerate(commands):
                output = None
                match command:
                    case "Graph":
                        n: int = values[idx][0]
                        edges: List[List[int]] = values[idx][1]
                        graph = Graph(n, edges)
                    case "addEdge":
                        if graph is not None:
                            graph.addEdge(values[idx][0])
                    case "shortestPath":
                        if graph is not None:
                            output = graph.shortestPath(values[idx][0], values[idx][1])
                    case _:
                        pass
                self.assertEqual(expected[idx], output, f"expected: {expected[idx]}, output: {output}")
