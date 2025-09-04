import heapq
from typing import List
import unittest

# https://leetcode.com/problems/maximum-average-pass-ratio/

# python3 -m unittest heaps/1792-maximum-average-pass-ratio.py


class Solution(unittest.TestCase):
    def maxAverageRatio(self, classes: List[List[int]], extraStudents: int) -> float:
        pq: List[tuple[float, int, int]] = []
        for p, t in classes:
            heapq.heappush(pq, (p / t - (p + 1) / (t + 1), p, t))
        for _ in range(extraStudents):
            _, p, t = heapq.heappop(pq)
            p, t = p + 1, t + 1
            heapq.heappush(pq, (p / t - (p + 1) / (t + 1), p, t))
        ratio = 0
        for _, p, t in pq:
            ratio += p / t
        return ratio / len(classes)

    def test(self):
        for classes, extraStudents, expected in [
            ([[1, 2], [3, 5], [2, 2]], 2, 0.78333),
            ([[2, 4], [3, 9], [4, 5], [2, 10]], 4, 0.53485),
        ]:
            output = self.maxAverageRatio(classes, extraStudents)
            output = round(output, 5)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
