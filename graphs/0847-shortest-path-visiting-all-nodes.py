from collections import deque
from typing import List, Tuple, Set
import unittest

# https://leetcode.com/problems/shortest-path-visiting-all-nodes/

# python3 -m unittest graphs/0847-shortest-path-visiting-all-nodes.py


class Solution(unittest.TestCase):
    # Approach: BFS with Bitmask to track visited nodes
    # Use state = (current_node, visited_mask) to represent current position
    # visited_mask is a bitmask where bit i is 1 if node i has been visited
    # Start BFS from all nodes simultaneously since we can start anywhere
    # N = number of nodes (1 <= N <= 12, so bitmask fits in int)
    # Time: O(N * 2^N) - at most N * 2^N states (node × visited combinations)
    # Space: O(N * 2^N) - queue and visited set
    def shortestPathLength(self, graph: List[List[int]]) -> int:
        n = len(graph)

        # Special case: single node
        if n == 1:
            return 0

        # Target: all nodes visited (all bits set)
        target = (1 << n) - 1

        # Queue: (current_node, visited_mask, distance)
        queue: deque[Tuple[int, int, int]] = deque()

        # Visited states: set of (node, mask) to avoid revisiting same state
        visited: Set[Tuple[int, int]] = set()

        # Start from all nodes (we can choose any starting point)
        for i in range(n):
            mask = 1 << i  # Only node i visited
            queue.append((i, mask, 0))
            visited.add((i, mask))

        # BFS to find shortest path
        while queue:
            node, mask, dist = queue.popleft()

            # Check if all nodes visited
            if mask == target:
                return dist

            # Explore neighbors
            for neighbor in graph[node]:
                new_mask = mask | (1 << neighbor)  # Mark neighbor as visited
                state = (neighbor, new_mask)

                if state not in visited:
                    visited.add(state)
                    queue.append((neighbor, new_mask, dist + 1))

        return -1  # Should never reach here for connected graph

    def test(self):
        test_cases: List[Tuple[List[List[int]], int]] = [
            ([[1, 2, 3], [0], [0], [0]], 4),
            ([[1], [0, 2, 4], [1, 3, 4], [2], [1, 2]], 4),
            ([[]], 0),
            ([[1], [0]], 1),
            ([[1, 2], [0, 2], [0, 1]], 2),
        ]
        for graph, expected in test_cases:
            output = self.shortestPathLength(graph)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
