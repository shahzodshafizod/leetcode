import heapq
from typing import List
import unittest

# https://leetcode.com/problems/trapping-rain-water-ii/

# python3 -m unittest heaps/0407-trapping-rain-water-ii.py

class Solution(unittest.TestCase):
    # Approach: Breadth-First Search + Heaps (Priority Queue)
    # Time: O(mn log(mn))
    # Space: O(mn)
    def trapRainWater(self, heightMap: List[List[int]]) -> int:
        directions = [-1,0,1,0,-1]
        m, n = len(heightMap), len(heightMap[0])
        queue = [] # (height, row, col)
        for row in range(m):
            queue.append((heightMap[row][0], row, 0))
            queue.append((heightMap[row][n-1], row, n-1))
            heightMap[row][0] = -1
            heightMap[row][n-1] = -1
        for col in range(n):
            queue.append((heightMap[0][col], 0, col))
            queue.append((heightMap[m-1][col], m-1, col))
            heightMap[0][col] = -1
            heightMap[m-1][col] = -1
        heapq.heapify(queue)
        volume = 0
        while queue:
            height, row, col = heapq.heappop(queue)
            for dir in range(1,5):
                r = row + directions[dir-1]
                c = col + directions[dir]
                if min(r,c) < 0 or r == m or c == n or heightMap[r][c] < 0:
                    continue
                if heightMap[r][c] < height:
                    volume += height - heightMap[r][c]
                    heightMap[r][c] = height
                heapq.heappush(queue, (heightMap[r][c], r, c))
                heightMap[r][c] = -1
        return volume

    def test(self):
        for heightMap, expected in [
            ([[1,4,3,1,3,2],[3,2,1,3,2,4],[2,3,3,2,3,1]], 4),
            ([[3,3,3,3,3],[3,2,2,2,3],[3,2,1,2,3],[3,2,2,2,3],[3,3,3,3,3]], 10),
        ]:
            output = self.trapRainWater(heightMap)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
