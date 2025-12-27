import unittest
from typing import List, Tuple

# https://leetcode.com/problems/range-addition-ii/

# python3 -m unittest maths/0598-range-addition-ii.py


class Solution(unittest.TestCase):
    # # Approach #1: Brute-Force (Simulate the matrix)
    # # Time: O(k * m * n) - k operations, each updating up to m*n cells
    # # Space: O(m * n) - storing the entire matrix
    # def maxCount(self, m: int, n: int, ops: List[List[int]]) -> int:
    #     # Create m x n matrix initialized with zeros
    #     matrix = [[0] * n for _ in range(m)]
    
    #     # Apply each operation
    #     for a, b in ops:
    #         for i in range(a):
    #             for j in range(b):
    #                 matrix[i][j] += 1
    
    #     # Find maximum value
    #     max_val = 0
    #     for i in range(m):
    #         for j in range(n):
    #             max_val = max(max_val, matrix[i][j])
    
    #     # Count cells with maximum value
    #     count = 0
    #     for i in range(m):
    #         for j in range(n):
    #             if matrix[i][j] == max_val:
    #                 count += 1
    
    #     return count

    # Approach #2: Math (Find intersection of all ranges)
    # Time: O(k) - k operations
    # Space: O(1) - constant space
    def maxCount(self, m: int, n: int, ops: List[List[int]]) -> int:
        # If no operations, all cells are 0 (maximum), return total cells
        if not ops:
            return m * n

        # The cells with maximum value are those incremented by ALL operations
        # This is the intersection of all operation ranges: [0, a) x [0, b)
        # The intersection is [0, min(all a)) x [0, min(all b))

        min_a = m  # Initialize to matrix dimensions
        min_b = n

        for a, b in ops:
            min_a = min(min_a, a)
            min_b = min(min_b, b)

        # The number of cells in the intersection
        return min_a * min_b

    def test(self):
        test_cases: List[Tuple[int, int, List[List[int]], int]] = [
            (3, 3, [[2, 2], [3, 3]], 4),
            (3, 3, [[2, 2], [3, 3], [3, 3], [3, 3], [2, 2], [3, 3], [3, 3], [3, 3], [2, 2], [3, 3], [3, 3], [3, 3]], 4),
            (3, 3, [], 9),
            (40000, 40000, [], 1600000000),
            (3, 3, [[1, 1], [2, 2], [3, 3]], 1),
        ]
        for m, n, ops, expected in test_cases:
            output = self.maxCount(m, n, ops)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
