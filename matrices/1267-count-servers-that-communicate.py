from typing import List
import unittest

# https://leetcode.com/problems/count-servers-that-communicate/

# python3 -m unittest matrices/1267-count-servers-that-communicate.py

# Approach: Counting
# Time: O(mn)
# Space: O(m+n)
class Solution(unittest.TestCase):
    def countServers(self, grid: List[List[int]]) -> int:
        m, n = len(grid), len(grid[0])
        rcount, ccount = [0] * m, [0] * n
        for row in range(m):
            for col in range(n):
                rcount[row] += grid[row][col]
                ccount[col] += grid[row][col]
        count = 0
        for row in range(m):
            for col in range(n):
                if grid[row][col] and (rcount[row] + ccount[col] > 2):
                    count += 1
        return count

    def test(self):
        for grid, expected in [
            ([[1,0],[0,1]], 0),
            ([[1,0],[1,1]], 3),
            ([[1,1,0,0],[0,0,1,0],[0,0,1,0],[0,0,0,1]], 4),
        ]:
            output = self.countServers(grid)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
