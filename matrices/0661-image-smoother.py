from typing import List
import unittest

# https://leetcode.com/problems/image-smoother/

# python3 -m unittest matrices/0661-image-smoother.py


class Solution(unittest.TestCase):
    # # Approach #1: Brute Force - Check all 8 directions for each cell
    # # For each cell, iterate through all 9 possible cells (including itself)
    # # and calculate the average of valid cells
    # # Time: O(m*n) - visit each cell and check up to 9 neighbors
    # # Space: O(m*n) - create a new result matrix
    # def imageSmoother(self, img: List[List[int]]) -> List[List[int]]:
    #     m, n = len(img), len(img[0])
    #     result = [[0] * n for _ in range(m)]
    
    #     # 8 directions + current cell
    #     directions = [(-1, -1), (-1, 0), (-1, 1), (0, -1), (0, 0), (0, 1), (1, -1), (1, 0), (1, 1)]
    
    #     for i in range(m):
    #         for j in range(n):
    #             total_sum = 0
    #             count = 0
    
    #             # Check all 9 cells (current + 8 neighbors)
    #             for di, dj in directions:
    #                 ni, nj = i + di, j + dj
    #                 # Check if the neighbor is within bounds
    #                 if 0 <= ni < m and 0 <= nj < n:
    #                     total_sum += img[ni][nj]
    #                     count += 1
    
    #             # Calculate average and floor it (integer division does this)
    #             result[i][j] = total_sum // count
    
    #     return result

    # Approach #2: Optimized - Same algorithm but cleaner implementation
    # Instead of using direction arrays, use nested loops with range
    # Time: O(m*n) - visit each cell once, check constant 9 neighbors
    # Space: O(m*n) - result matrix (required by problem)
    def imageSmoother(self, img: List[List[int]]) -> List[List[int]]:
        m, n = len(img), len(img[0])
        result = [[0] * n for _ in range(m)]

        for i in range(m):
            for j in range(n):
                total_sum = 0
                count = 0

                # Check 3x3 grid centered at (i, j)
                for di in range(-1, 2):  # -1, 0, 1
                    for dj in range(-1, 2):  # -1, 0, 1
                        ni, nj = i + di, j + dj
                        if 0 <= ni < m and 0 <= nj < n:
                            total_sum += img[ni][nj]
                            count += 1

                result[i][j] = total_sum // count

        return result

    def test(self):
        for img, expected in [
            ([[1, 1, 1], [1, 0, 1], [1, 1, 1]], [[0, 0, 0], [0, 0, 0], [0, 0, 0]]),
            ([[100, 200, 100], [200, 50, 200], [100, 200, 100]], [[137, 141, 137], [141, 138, 141], [137, 141, 137]]),
            ([[1]], [[1]]),
            ([[1, 2], [3, 4]], [[2, 2], [2, 2]]),
        ]:
            output = self.imageSmoother(img)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
