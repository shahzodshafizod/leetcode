from typing import List
import unittest

# https://leetcode.com/problems/most-profitable-path-in-a-tree/

# python3 -m unittest graphs/2467-most-profitable-path-in-a-tree.py


class Solution(unittest.TestCase):
    # Approach: Depth-First Search
    # Time: O(n)
    # Space: O(n)
    def mostProfitablePath(self, edges: List[List[int]], bob: int, amount: List[int]) -> int:
        n = len(edges) + 1
        graph = [[] for _ in range(n)]
        for a, b in edges:
            graph[a].append(b)
            graph[b].append(a)

        def dfs(parent: int, node: int, adist: int) -> tuple[int, int]:
            profit = float("-inf")
            bdist = -1
            for nei in graph[node]:
                if nei != parent:
                    p, d = dfs(node, nei, adist + 1)
                    profit = max(profit, p + amount[node])
                    bdist = max(bdist, d)
            if profit == float("-inf"):
                profit = amount[node]
            if node == bob:
                bdist = 0
            if bdist != -1:
                if bdist == adist:
                    profit -= amount[node] // 2
                elif bdist < adist:
                    profit -= amount[node]
                bdist += 1
            return profit, bdist

        profit, _ = dfs(-1, 0, 0)
        return profit

    def test(self):
        for edges, bob, amount, expected in [
            ([[0, 1]], 1, [-7280, 2350], -7280),
            ([[0, 1], [0, 2]], 2, [-3360, -5394, -1146], -3360),
            ([[0, 1], [1, 2], [1, 3], [3, 4]], 3, [-2, 4, 2, -4, 6], 6),
            ([[0, 1], [1, 2], [2, 3]], 3, [-5644, -6018, 1188, -8502], -11662),
            ([[0, 2], [0, 4], [1, 3], [1, 2]], 1, [3958, -9854, -8334, -9388, 3410], 7368),
            ([[0, 2], [0, 5], [1, 3], [1, 5], [2, 4]], 4, [5018, 8388, 6224, 3466, 3808, 3456], 20328),
            ([[0, 2], [0, 6], [1, 3], [1, 5], [2, 5], [4, 6]], 6, [8838, -6396, -5940, 2694, -1366, 4616, 2966], 7472),
            (
                [[0, 2], [1, 4], [1, 6], [2, 4], [3, 6], [3, 7], [5, 7]],
                4,
                [-6896, -1216, -1208, -1108, 1606, -7704, -9212, -8258],
                -34998,
            ),
            (
                [[0, 2], [1, 4], [1, 6], [2, 3], [2, 8], [3, 7], [4, 5], [6, 7]],
                1,
                [-1410, -9440, 5536, -774, 3044, 7924, -2122, -1192, 9166],
                14320,
            ),
        ]:
            output = self.mostProfitablePath(edges, bob, amount)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
