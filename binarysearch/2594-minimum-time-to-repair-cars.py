from typing import List
import unittest

# https://leetcode.com/problems/minimum-time-to-repair-cars/

# python3 -m unittest binarysearch/2594-minimum-time-to-repair-cars.py


class Solution(unittest.TestCase):
    def repairCars(self, ranks: List[int], cars: int) -> int:
        low, high = 1, max(ranks) * cars * cars
        n = len(ranks)
        while low < high:
            mid = (low + high) // 2
            idx, remained = 0, cars
            while idx < n and remained > 0:
                # rank * n*n <= mid
                remained -= int((mid / ranks[idx]) ** 0.5)
                idx += 1
            if remained <= 0:
                high = mid
            else:
                low = mid + 1
        return high

    def test(self):
        for ranks, cars, expected in [
            ([4, 2, 3, 1], 10, 16),
            ([5, 1, 8], 6, 16),
        ]:
            output = self.repairCars(ranks, cars)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
