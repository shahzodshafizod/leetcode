import unittest
from typing import List, Dict
from collections import defaultdict

# https://leetcode.com/problems/maximum-score-of-a-node-sequence/

# python3 -m unittest graphs/2242-maximum-score-of-a-node-sequence.py


class Solution(unittest.TestCase):
    # Approach: Graph + Greedy (Store top 3 neighbors for each node)
    # Time: O(E * 9) = O(E) - for each edge, check up to 9 combinations
    # Space: O(n * 3) = O(n) - store top 3 neighbors per node
    def maximumScore(self, scores: List[int], edges: List[List[int]]) -> int:
        n = len(scores)

        # Build adjacency list with scores
        # For each node, store neighbors sorted by score (descending)
        neighbors: Dict[int, List[int]] = defaultdict(list)

        for u, v in edges:
            neighbors[u].append(v)
            neighbors[v].append(u)

        # Keep only top 3 neighbors for each node (sorted by score descending)
        top_neighbors: Dict[int, List[int]] = {}
        for node in range(n):
            # Sort neighbors by their scores in descending order
            sorted_neighbors = sorted(neighbors[node], key=lambda x: scores[x], reverse=True)
            # Keep only top 3
            top_neighbors[node] = sorted_neighbors[:3]

        max_score = -1

        # For each edge (b, c), try to form sequence: a - b - c - d
        for b, c in edges:
            # Try all combinations of top neighbors
            for a in top_neighbors[b]:
                # a must not be c
                if a == c:
                    continue

                for d in top_neighbors[c]:
                    # d must not be b or a
                    if d == b or d == a:
                        continue

                    # Valid sequence: a - b - c - d
                    current_score = scores[a] + scores[b] + scores[c] + scores[d]
                    max_score = max(max_score, current_score)

        return max_score

    def test(self):
        for scores, edges, expected in [
            ([5, 2, 9, 8, 4], [[0, 1], [1, 2], [2, 3], [0, 2], [1, 3], [2, 4]], 24),
            ([9, 20, 6, 4, 11, 12], [[0, 3], [5, 3], [2, 4], [1, 3]], -1),
            ([1, 1, 1, 1], [[0, 1], [1, 2], [2, 3]], 4),
            ([5, 10, 15, 20], [[0, 1], [1, 2], [2, 3], [0, 2]], 50),
        ]:
            output = self.maximumScore(scores, edges)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
