from typing import List
import unittest

# https://leetcode.com/problems/critical-connections-in-a-network/

# python3 -m unittest graphs/1192-critical-connections-in-a-network.py


class Solution(unittest.TestCase):
    # Approach: Tarjan's Algorithm
    # Time: O(V + E), Space: O(V + E)
    def criticalConnections(self, n: int, connections: List[List[int]]) -> List[List[int]]:
        graph: List[List[int]] = [[] for _ in range(n)]
        for u, v in connections:
            graph[u].append(v)
            graph[v].append(u)

        dfn = [-1] * n
        low = [-1] * n
        bridges: List[List[int]] = []
        timer = 0

        def dfs(node: int, parent: int):
            nonlocal timer, bridges

            dfn[node] = timer
            low[node] = timer
            timer += 1

            for neighbor in graph[node]:
                if neighbor == parent:
                    continue

                if dfn[neighbor] == -1:
                    dfs(neighbor, node)
                    low[node] = min(low[node], low[neighbor])

                    if low[neighbor] > dfn[node]:
                        bridges.append([node, neighbor])
                else:
                    low[node] = min(low[node], dfn[neighbor])

        dfs(0, -1)
        return bridges

    def test(self):
        for n, connections, expected in [
            (4, [[0, 1], [1, 2], [2, 0], [1, 3]], [[1, 3]]),
            (2, [[0, 1]], [[0, 1]]),
        ]:
            output = self.criticalConnections(n, connections)
            output.sort()
            expected.sort()
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
