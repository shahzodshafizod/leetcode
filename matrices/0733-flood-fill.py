import unittest
from typing import List
from collections import deque

# https://leetcode.com/problems/flood-fill/

# python3 -m unittest matrices/0733-flood-fill.py

class Solution(unittest.TestCase):
    def floodFill(self, image: List[List[int]], sr: int, sc: int, color: int) -> List[List[int]]:
        srcColor = image[sr][sc]
        queue = deque()
        if srcColor != color:
            queue.append((sr, sc))
        directions = [-1, 0, 1, 0, -1]
        m, n = len(image), len(image[0])
        while queue:
            row, col = queue.popleft()
            image[row][col] = color
            for d in range(1, len(directions)):
                r, c = row+directions[d-1], col+directions[d]
                if min(r, c) >= 0 and r < m and c < n and image[r][c] == srcColor:
                    queue.append((r, c))
        return image

    def testFloodFill(self) -> None:
        for image, sr, sc, color, expected in [
            ([[1, 1, 1], [1, 1, 0], [1, 0, 1]], 1, 1, 2, [[2, 2, 2], [2, 2, 0], [2, 0, 1]]),
            ([[0, 0, 0], [0, 0, 0]], 0, 0, 0, [[0, 0, 0], [0, 0, 0]]),
            ([[0, 0, 0], [0, 0, 0]], 0, 0, 2, [[2, 2, 2], [2, 2, 2]]),
            ([[0, 0, 0], [0, 0, 0]], 1, 0, 2, [[2, 2, 2], [2, 2, 2]]),
            ([[1, 1, 1], [1, 1, 0], [1, 0, 1]], 1, 1, 2, [[2, 2, 2], [2, 2, 0], [2, 0, 1]]),
            ([[0, 0, 0], [0, 0, 1]], 1, 0, 2, [[2, 2, 2], [2, 2, 1]]),
        ]:
            output = self.floodFill(image, sr, sc, color)
            self.assertListEqual(expected, output, f"expected: {expected}, output: {output}")
