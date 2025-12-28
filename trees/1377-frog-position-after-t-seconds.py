from typing import List, Tuple
import unittest

# https://leetcode.com/problems/frog-position-after-t-seconds/

# python3 -m unittest trees/1377-frog-position-after-t-seconds.py


class Solution(unittest.TestCase):
    # # Approach #1: Brute Force - BFS with full probability tracking
    # # Use BFS to explore all paths, track probability for each node at each time
    # # Store (node, probability, time) in queue and visited set
    # # N = number of nodes, E = number of edges
    # # Time: O(N + E) - visit each node and edge once
    # # Space: O(N) - queue and visited set store up to N nodes
    # def frogPosition(self, n: int, edges: List[List[int]], t: int, target: int) -> float:
    #     # Build adjacency list
    #     graph: List[List[int]] = [[] for _ in range(n + 1)]
    #     for a, b in edges:
    #         graph[a].append(b)
    #         graph[b].append(a)
    
    #     # BFS: (current_node, probability, time, parent)
    #     from collections import deque
    #     queue: Deque[Tuple[int, float, int, int]] = deque([(1, 1.0, 0, -1)])
    
    #     while queue:
    #         node, prob, time, parent = queue.popleft()
    
    #         # Count unvisited neighbors (exclude parent)
    #         neighbors = [neighbor for neighbor in graph[node] if neighbor != parent]
    #         num_neighbors = len(neighbors)
    
    #         # If we reached target at exactly time t, or if we're at target
    #         # and can't move anymore (no unvisited neighbors)
    #         if node == target:
    #             if time == t or (time < t and num_neighbors == 0):
    #                 return prob
    #             if time > t:
    #                 return 0.0
    
    #         # If we've exhausted time, continue searching
    #         if time >= t:
    #             continue
    
    #         # Move to unvisited neighbors with equal probability
    #         if num_neighbors > 0:
    #             for neighbor in neighbors:
    #                 queue.append((neighbor, prob / num_neighbors, time + 1, node))
    
    #     return 0.0

    # Approach #2: Optimized - DFS with early termination
    # Use DFS to find target path, calculate probability along the way
    # Key insight: probability = product of (1/choices) at each step
    # Can only be at target if: (1) exactly at time t, or (2) stuck at target before t
    # N = number of nodes, E = number of edges
    # Time: O(N + E) - worst case visit all nodes, but often terminates early
    # Space: O(N) - recursion stack depth (tree height)
    def frogPosition(self, n: int, edges: List[List[int]], t: int, target: int) -> float:
        # Build adjacency list
        graph: List[List[int]] = [[] for _ in range(n + 1)]
        for a, b in edges:
            graph[a].append(b)
            graph[b].append(a)

        def dfs(node: int, parent: int, time: int, prob: float) -> float:
            # Count unvisited neighbors (exclude parent)
            neighbors = [neighbor for neighbor in graph[node] if neighbor != parent]
            num_neighbors = len(neighbors)

            # If we found the target
            if node == target:
                # Case 1: We're exactly at time t
                if time == t:
                    return prob
                # Case 2: We reached target before t, but can't move (stuck here)
                if num_neighbors == 0:
                    return prob
                # Case 3: We reached target but will have to move away
                return 0.0

            # If time is up or no neighbors to explore, can't reach target
            if time >= t or num_neighbors == 0:
                return 0.0

            # Explore neighbors with equal probability distribution
            for neighbor in neighbors:
                result = dfs(neighbor, node, time + 1, prob / num_neighbors)
                if result > 0:
                    return result

            return 0.0

        return dfs(1, -1, 0, 1.0)

    def test(self):
        test_cases: List[Tuple[int, List[List[int]], int, int, float]] = [
            (7, [[1, 2], [1, 3], [1, 7], [2, 4], [2, 6], [3, 5]], 2, 4, 0.16666666666666666),
            (7, [[1, 2], [1, 3], [1, 7], [2, 4], [2, 6], [3, 5]], 1, 7, 0.3333333333333333),
            (7, [[1, 2], [1, 3], [1, 7], [2, 4], [2, 6], [3, 5]], 20, 6, 0.16666666666666666),
            (3, [[2, 1], [3, 2]], 1, 2, 1.0),
            (1, [], 1, 1, 1.0),
            (8, [[2, 1], [3, 2], [4, 1], [5, 1], [6, 4], [7, 1], [8, 7]], 7, 7, 0.0),
        ]
        for n, edges, t, target, expected in test_cases:
            output = self.frogPosition(n, edges, t, target)
            self.assertAlmostEqual(expected, output, places=5, msg=f"expected: {expected}, output: {output}")
