from typing import List
import unittest

# https://leetcode.com/problems/find-the-maximum-number-of-fruits-collected/

# python3 -m unittest dp/3363-find-the-maximum-number-of-fruits-collected.py


class Solution(unittest.TestCase):
    # Approach: Top-Down Dynamic Programming
    # Time: O(nn)
    # Space: O(nn)
    def maxCollectedFruits(self, fruits: List[List[int]]) -> int:
        n = len(fruits)
        first = sum(fruits[i][i] for i in range(n))
        memo = [[-1] * n for _ in range(n)]

        def dp(row: int, col: int, dirs: List[List[int]], f: int) -> int:
            if min(row, col) < 0 or max(row, col) == n or f * row <= f * col:
                return 0
            if memo[row][col] != -1:
                return memo[row][col]
            memo[row][col] = max(dp(row + dr, col + dc, dirs, f) for dr, dc in dirs) + fruits[row][col]
            return memo[row][col]

        second = dp(n - 1, 0, [[1, 1], [0, 1], [-1, 1]], 1)
        third = dp(0, n - 1, [[1, -1], [1, 0], [1, 1]], -1)
        return first + second + third

    def test(self):
        for fruits, expected in [
            ([[1, 2, 3, 4], [5, 6, 8, 7], [9, 10, 11, 12], [13, 14, 15, 16]], 100),
            ([[1, 1], [1, 1]], 4),
        ]:
            output = self.maxCollectedFruits(fruits)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
