from typing import List
import unittest

# https://leetcode.com/problems/checking-existence-of-edge-length-limited-paths/

# python3 -m unittest unionfinds/1697-checking-existence-of-edge-length-limited-paths.py


class Solution(unittest.TestCase):
    # Approach: Union-Find + Sorting
    # Time: O((E + Q) log(E + Q)) - sorting edges and queries
    # Space: O(N + Q) - parent array and result array
    def distanceLimitedPathsExist(self, n: int, edgeList: List[List[int]], queries: List[List[int]]) -> List[bool]:
        # Union-Find data structure
        parent = list(range(n))

        def find(x: int) -> int:
            if parent[x] != x:
                parent[x] = find(parent[x])  # Path compression
            return parent[x]

        def union(x: int, y: int) -> bool:
            px, py = find(x), find(y)
            if px == py:
                return False
            parent[py] = px
            return True

        # Sort edges by distance (ascending)
        edgeList.sort(key=lambda e: e[2])

        # Create indexed queries and sort by limit (ascending)
        indexed_queries = [(p, q, limit, i) for i, (p, q, limit) in enumerate(queries)]
        indexed_queries.sort(key=lambda q: q[2])

        result = [False] * len(queries)
        edge_idx = 0

        # Process queries in order of increasing limit
        for p, q, limit, original_idx in indexed_queries:
            # Add all edges with distance < limit to Union-Find
            while edge_idx < len(edgeList) and edgeList[edge_idx][2] < limit:
                u, v, _ = edgeList[edge_idx]
                union(u, v)
                edge_idx += 1

            # Check if p and q are connected
            result[original_idx] = find(p) == find(q)

        return result

    def test(self):
        for n, edgeList, queries, expected in [
            (
                3,
                [[0, 1, 2], [1, 2, 4], [2, 0, 8], [1, 0, 16]],
                [[0, 1, 2], [0, 2, 5]],
                [False, True],
            ),
            (
                5,
                [[0, 1, 10], [1, 2, 5], [2, 3, 9], [3, 4, 13]],
                [[0, 4, 14], [1, 4, 13]],
                [True, False],
            ),
        ]:
            output = self.distanceLimitedPathsExist(n, edgeList, queries)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
