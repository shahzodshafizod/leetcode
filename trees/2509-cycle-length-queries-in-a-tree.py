from typing import List
import unittest

# https://leetcode.com/problems/cycle-length-queries-in-a-tree/

# python3 -m unittest trees/2509-cycle-length-queries-in-a-tree.py


class Solution(unittest.TestCase):
    # # Approach #1: Brute Force - Find LCA by traversing to root
    # # For each query, find path from both nodes to root, then find LCA
    # # N = 2^n - 1 (number of nodes), M = number of queries
    # # Time: O(M * log N) - for each query, traverse up to log(N) levels
    # # Space: O(log N) - store paths to root
    # def cycleLengthQueries(self, n: int, queries: List[List[int]]) -> List[int]:
    #     result: List[int] = []
    
    #     for a, b in queries:
    #         # Get ancestors of both nodes
    #         path_a: List[int] = []
    #         node = a
    #         while node >= 1:
    #             path_a.append(node)
    #             node //= 2
    
    #         path_b: List[int] = []
    #         node = b
    #         while node >= 1:
    #             path_b.append(node)
    #             node //= 2
    
    #         # Find LCA - the first common ancestor
    #         path_a_set = set(path_a)
    #         lca = -1
    #         for node in path_b:
    #             if node in path_a_set:
    #                 lca = node
    #                 break
    
    #         # Calculate distances
    #         dist_a = path_a.index(lca)
    #         dist_b = path_b.index(lca)
    
    #         # Cycle length = distance(a, lca) + distance(b, lca) + 1 (added edge)
    #         result.append(dist_a + dist_b + 1)
    
    #     return result

    # Approach #2: Optimized - Simultaneous traversal to LCA
    # Move both pointers up the tree simultaneously until they meet
    # In a complete binary tree: parent of node x is x // 2
    # Time: O(M * log N) - for each query, traverse up to log(N) levels
    # Space: O(1) - no extra space needed
    def cycleLengthQueries(self, n: int, queries: List[List[int]]) -> List[int]:
        result: List[int] = []

        for a, b in queries:
            distance = 0

            # Move both nodes up until they meet (find LCA)
            while a != b:
                if a > b:
                    a //= 2  # Move to parent
                else:
                    b //= 2  # Move to parent
                distance += 1

            # Add 1 for the edge we're adding
            result.append(distance + 1)

        return result

    def test(self):
        for n, queries, expected in [
            (3, [[5, 3], [4, 7], [2, 3]], [4, 5, 3]),
            (2, [[1, 2]], [2]),
            (5, [[8, 5], [14, 10]], [4, 7]),
            (4, [[7, 2], [15, 1]], [4, 4]),
        ]:
            output = self.cycleLengthQueries(n, queries)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
