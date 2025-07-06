from collections import deque
from typing import List
import unittest

# https://leetcode.com/problems/map-of-highest-peak/

# python3 -m unittest matrices/1765-map-of-highest-peak.py


class Solution(unittest.TestCase):
    # Approach: Breadth-First Search
    # Time: O(mn)
    # Space: O(mn)
    def highestPeak(self, isWater: List[List[int]]) -> List[List[int]]:
        m, n = len(isWater), len(isWater[0])
        height = [[-1] * n for _ in range(m)]
        queue = deque()
        for row in range(m):
            for col in range(n):
                if isWater[row][col] == 1:
                    height[row][col] = 0
                    queue.append((row, col))
        dirs = [-1, 0, 1, 0, -1]
        while queue:
            row, col = queue.popleft()
            for d in range(1, 5):
                r, c = row + dirs[d - 1], col + dirs[d]
                if min(r, c) < 0 or r == m or c == n or height[r][c] != -1:
                    continue
                height[r][c] = height[row][col] + 1
                queue.append((r, c))
        return height

    def test(self):
        for isWater, expected in [
            ([[1]], [[0]]),
            ([[1], [0]], [[0], [1]]),
            ([[0, 1], [0, 0]], [[1, 0], [2, 1]]),
            ([[0, 0, 1], [1, 0, 0], [0, 0, 0]], [[1, 1, 0], [0, 1, 1], [1, 2, 2]]),
            ([[0], [0], [0], [0], [1], [0], [0], [1], [1]], [[4], [3], [2], [1], [0], [1], [1], [0], [0]]),
        ]:
            output = self.highestPeak(isWater)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
