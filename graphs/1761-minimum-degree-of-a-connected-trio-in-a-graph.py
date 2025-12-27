import unittest
from typing import List

# https://leetcode.com/problems/minimum-degree-of-a-connected-trio-in-a-graph/

# python3 -m unittest graphs/1761-minimum-degree-of-a-connected-trio-in-a-graph.py


class Solution(unittest.TestCase):
    # # Approach #1: Brute-Force (Check all possible trios)
    # # Time: O(n^3) - three nested loops to check all node combinations
    # # Space: O(n^2) - adjacency matrix
    # def minTrioDegree(self, n: int, edges: List[List[int]]) -> int:
    #     # Build adjacency list and count degrees
    #     adj: List[Set[int]] = [set() for _ in range(n + 1)]
    #     degree = [0] * (n + 1)
    
    #     for u, v in edges:
    #         adj[u].add(v)
    #         adj[v].add(u)
    #         degree[u] += 1
    #         degree[v] += 1
    
    #     MAX = 1 << 31
    #     min_degree = MAX
    
    #     # Check all possible trios (i, j, k) where i < j < k
    #     for i in range(1, n + 1):
    #         for j in range(i + 1, n + 1):
    #             if j not in adj[i]:
    #                 continue
    #             for k in range(j + 1, n + 1):
    #                 # Check if (i, j, k) forms a trio
    #                 if k in adj[i] and k in adj[j]:
    #                     # Calculate degree of this trio
    #                     # degree = (deg(i) + deg(j) + deg(k)) - 6
    #                     # The -6 removes the 3 internal edges (counted twice each)
    #                     trio_degree = degree[i] + degree[j] + degree[k] - 6
    #                     min_degree = min(min_degree, trio_degree)
    
    #     return min_degree if min_degree != MAX else -1

    # Approach #2: Optimized with Adjacency Matrix
    # Time: O(m * sqrt(m)) where m = number of edges
    # Space: O(n^2) - adjacency matrix
    def minTrioDegree(self, n: int, edges: List[List[int]]) -> int:
        # Build adjacency matrix for O(1) edge lookup
        adj_matrix = [[False] * (n + 1) for _ in range(n + 1)]
        degree = [0] * (n + 1)

        for u, v in edges:
            adj_matrix[u][v] = True
            adj_matrix[v][u] = True
            degree[u] += 1
            degree[v] += 1

        MAX = 1 << 31
        min_degree = MAX

        # For each edge (u, v), find common neighbors to form trios
        for u, v in edges:
            # Ensure u < v for consistent ordering
            if u > v:
                u, v = v, u

            # Find all nodes w that are connected to both u and v
            for w in range(1, n + 1):
                if w != u and w != v and adj_matrix[u][w] and adj_matrix[v][w]:
                    # Found a trio (u, v, w)
                    # Degree = sum of degrees - 6 (3 internal edges * 2)
                    trio_degree = degree[u] + degree[v] + degree[w] - 6
                    min_degree = min(min_degree, trio_degree)

        return min_degree if min_degree != MAX else -1

    def test(self):
        for n, edges, expected in [
            (6, [[1, 2], [1, 3], [3, 2], [4, 1], [5, 2], [3, 6]], 3),
            (7, [[1, 3], [4, 1], [4, 3], [2, 5], [5, 6], [6, 7], [7, 5], [2, 6]], 0),
            (4, [[1, 2], [1, 3], [2, 3], [2, 4], [3, 4]], 2),
        ]:
            output = self.minTrioDegree(n, edges)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
