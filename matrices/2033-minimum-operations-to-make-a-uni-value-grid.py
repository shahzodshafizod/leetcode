from typing import List
import unittest

# https://leetcode.com/problems/minimum-operations-to-make-a-uni-value-grid/

# python3 -m unittest matrices/2033-minimum-operations-to-make-a-uni-value-grid.py

class Solution(unittest.TestCase):
    def minOperations(self, grid: List[List[int]], x: int) -> int:
        m, n = len(grid), len(grid[0])
        total = 0
        nums = [None] * (m*n)
        idx = 0
        mod = grid[0][0] % x
        for row in range(m):
            for col in range(n):
                if grid[row][col] % x != mod:
                    return -1
                total += grid[row][col]
                nums[idx] = grid[row][col]
                idx += 1
        nums.sort()
        mid = nums[m*n//2]
        return sum(abs(num - mid) // x for num in nums)

    def test(self):
        for grid, x, expected in [
            ([[2,4],[6,8]], 2, 4),
            ([[1,5],[2,3]], 1, 5),
            ([[1,2],[3,4]], 2, -1),
        ]:
            output = self.minOperations(grid, x)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
