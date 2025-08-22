from typing import List
import unittest

# https://leetcode.com/problems/count-square-submatrices-with-all-ones/

# python3 -m unittest dp/1277-count-square-submatrices-with-all-ones.py

class Solution(unittest.TestCase):
    # # Approach #1: Top-Down Dynamic Programming
    # # Time: O(mn)
    # # Space: O(mn)
    # def countSquares(self, matrix: List[List[int]]) -> int:
    #     m, n = len(matrix), len(matrix[0])
    #     cache = [[None] * n for _ in range(n)]
    #     def dp(row: int, col: int) -> int:
    #         if row == m or col == n or not matrix[row][col]:
    #             return 0
    #         if cache[row][col] != None:
    #             return cache[row][col]
    #         cache[row][col] = 1 + min(
    #             dp(row+1, col),
    #             dp(row, col+1),
    #             dp(row+1, col+1),
    #         )
    #         return cache[row][col]
    #     count = 0
    #     for row in range(m):
    #         for col in range(n):
    #             count += dp(row, col)
    #     return count

    # # Approach #2: Bottom-Up Dynamic Programming
    # # Time: O(mn)
    # # Space: O(mn)
    # def countSquares(self, matrix: List[List[int]]) -> int:
    #     m, n = len(matrix), len(matrix[0])
    #     dp = [[0] * (n+1) for _ in range(m+1)]
    #     count = 0
    #     for row in range(1, m+1):
    #         for col in range(1, n+1):
    #             if matrix[row-1][col-1]:
    #                 dp[row][col] = 1 + min(
    #                     dp[row-1][col],
    #                     dp[row][col-1],
    #                     dp[row-1][col-1],
    #                 )
    #             count += dp[row][col]
    #     return count

    # # Approach #3: Bottom-Up Dynamic Programming (Space Optimized)
    # # Time: O(mn)
    # # Space: O(n)
    # def countSquares(self, matrix: List[List[int]]) -> int:
    #     m, n = len(matrix), len(matrix[0])
    #     prev, curr = [0]*(n+1), [0]*(n+1)
    #     count = 0
    #     for row in range(1, m+1):
    #         prev, curr = curr, prev
    #         for col in range(1, n+1):
    #             curr[col] = matrix[row-1][col-1]
    #             if matrix[row-1][col-1]:
    #                 curr[col] += min(
    #                     prev[col],
    #                     curr[col-1],
    #                     prev[col-1],
    #                 )
    #             count += curr[col]
    #     return count

    # Approach #4: Bottom-Up Dynamic Programming (in-place)
    # Time: O(mn)
    # Space: O(1)
    def countSquares(self, matrix: List[List[int]]) -> int:
        m, n = len(matrix), len(matrix[0])
        count = 0
        for row in range(m):
            for col in range(n):
                if row and col and matrix[row][col]:
                    matrix[row][col] += min(
                        matrix[row-1][col],
                        matrix[row][col-1],
                        matrix[row-1][col-1],
                    )
                count += matrix[row][col]
        return count

    def test(self) -> None:
        for matrix, expected in [
            ([[0,1,1,1],[1,1,1,1],[0,1,1,1]], 15),
            ([[1,0,1],[1,1,0],[1,1,0]], 7),
            ([[1,0,0],[1,1,0]], 3),
            ([[1,1,1,1],[1,0,0,1],[1,1,0,0],[1,0,0,0]], 9),
            ([[0,1,1,1],[1,1,0,1],[1,1,1,1],[1,0,1,0]], 13),
            ([[1,0,1,0,1],[1,0,0,1,1],[0,1,0,1,1],[1,0,0,1,1]], 14),
            ([[1,1,0,0,1],[1,0,1,1,1],[1,1,1,1,1],[1,0,1,0,1],[0,0,1,0,1]], 19),
        ]:
            output = self.countSquares(matrix)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
