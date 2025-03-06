from typing import List
import unittest

# https://leetcode.com/problems/find-missing-and-repeated-values/

# python3 -m unittest matrices/2965-find-missing-and-repeated-values.py

class Solution(unittest.TestCase):
    # Approach: Counting
    # Time: O(nn)
    # Space: O(nn)
    def findMissingAndRepeatedValues(self, grid: List[List[int]]) -> List[int]:
        n = len(grid)
        nn = n * n
        total = nn * (nn + 1) // 2
        seen = [False] * (nn + 1)
        for r in range(n):
            for c in range(n):
                if seen[grid[r][c]]:
                    twice = grid[r][c]
                else:
                    total -= grid[r][c]
                    seen[grid[r][c]] = True
        missing = total
        return [twice, missing]

    def test(self):
        for grid, expected in [
            ([[1,3],[2,2]], [2,4]),
            ([[9,1,7],[8,9,2],[3,4,6]], [9,5]),
        ]:
            output = self.findMissingAndRepeatedValues(grid)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
