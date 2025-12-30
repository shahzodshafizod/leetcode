from typing import List
import unittest

# https://leetcode.com/problems/count-negative-numbers-in-a-sorted-matrix/

# python3 -m unittest matrices/1351-count-negative-numbers-in-a-sorted-matrix.py


class Solution(unittest.TestCase):
    # # Approach 1: Brute Force
    # # Iterate through every element in the matrix and count negatives
    # # Time: O(m * n) where m is rows, n is columns
    # # Space: O(1)
    # def countNegatives(self, grid: List[List[int]]) -> int:
    #     count = 0
    #     for row in grid:
    #         for num in row:
    #             if num < 0:
    #                 count += 1
    #     return count

    # # Approach 2: Binary Search on Each Row
    # # Since each row is sorted in non-increasing order, use binary search
    # # to find the first negative number in each row
    # # Time: O(m * log n) where m is rows, n is columns
    # # Space: O(1)
    # def countNegatives(self, grid: List[List[int]]) -> int:
    #     n = len(grid[0])
    #     count = 0
    
    #     for row in grid:
    #         # Binary search to find first negative (first element < 0)
    #         left, right = 0, n - 1
    #         first_neg = n  # If no negative found, all are non-negative
    
    #         while left <= right:
    #             mid = left + (right - left) // 2
    #             if row[mid] < 0:
    #                 # Found a negative, but there might be an earlier one
    #                 first_neg = mid
    #                 right = mid - 1
    #             else:
    #                 # row[mid] >= 0, search right half
    #                 left = mid + 1
    
    #         # Count negatives from first_neg to end of row
    #         count += n - first_neg
    
    #     return count

    # Approach 3: Staircase/Linear Search (Optimal - Follow-up Solution)
    # Key insight: Start from top-right or bottom-left corner
    # The matrix is sorted both row-wise and column-wise in non-increasing order
    #
    # Starting from bottom-left (grid[m-1][0]):
    # - If current element is negative, all elements to the right are also negative
    #   (because row is non-increasing), so count them and move up
    # - If current element is non-negative, move right to find negatives
    #
    # Time: O(m + n) - at most m+n moves (either move up or move right)
    # Space: O(1)
    def countNegatives(self, grid: List[List[int]]) -> int:
        m, n = len(grid), len(grid[0])
        count = 0
        row, col = m - 1, 0  # Start from bottom-left corner

        while row >= 0 and col < n:
            if grid[row][col] < 0:
                # All elements from col to n-1 in this row are negative
                count += n - col
                # Move up to check previous row
                row -= 1
            else:
                # Current element is non-negative, move right
                col += 1

        return count

    def test(self):
        for grid, expected in [
            ([[4, 3, 2, -1], [3, 2, 1, -1], [1, 1, -1, -2], [-1, -1, -2, -3]], 8),
            ([[3, 2], [1, 0]], 0),
            ([[1, -1], [-1, -1]], 3),
            ([[-1]], 1),
            ([[5, 1, 0], [-5, -5, -5]], 3),
            ([[3, -1, -3, -3, -3], [2, -2, -3, -3, -3], [1, -2, -3, -3, -3], [0, -3, -3, -3, -3]], 16),
        ]:
            output = self.countNegatives([row[:] for row in grid])
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
