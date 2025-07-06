from collections import deque
from typing import List
import unittest

# https://leetcode.com/problems/maximum-candies-you-can-get-from-boxes/

# python3 -m unittest graphs/1298-maximum-candies-you-can-get-from-boxes.py


class Solution(unittest.TestCase):
    # Approach: Breadth-First Search
    # Time: O(nn)
    # Space: O(nn)
    def maxCandies(
        self,
        status: List[int],
        candies: List[int],
        keys: List[List[int]],
        containedBoxes: List[List[int]],
        initialBoxes: List[int],
    ) -> int:
        seen = [False] * len(status)
        queue = deque()
        for node in initialBoxes:
            seen[node] = True
            queue.append(node)
        count = 0
        while queue:
            for node in queue:
                for box in keys[node]:
                    if box != node:
                        status[box] = 1
            for _ in range(len(queue)):
                node = queue.popleft()
                if status[node] == 0:
                    continue
                count += candies[node]
                for nei in containedBoxes[node]:
                    if not seen[nei]:
                        seen[nei] = True
                        queue.append(nei)
        return count

    def test(self):
        for status, candies, keys, containedBoxes, initialBoxes, expected in [
            ([1, 1, 1], [100, 1, 100], [[], [0, 2], []], [[], [], []], [1], 1),
            ([1, 0, 1, 0], [7, 5, 4, 100], [[], [], [1], []], [[1, 2], [3], [], []], [0], 16),
            ([1, 0, 1, 0], [7, 5, 4, 100], [[], [], [1], [3]], [[1, 2], [3], [], []], [0], 16),
            (
                [1, 0, 0, 0, 0, 0],
                [1, 1, 1, 1, 1, 1],
                [[1, 2, 3, 4, 5], [], [], [], [], []],
                [[1, 2, 3, 4, 5], [], [], [], [], []],
                [0],
                6,
            ),
        ]:
            output = self.maxCandies(status, candies, keys, containedBoxes, initialBoxes)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
