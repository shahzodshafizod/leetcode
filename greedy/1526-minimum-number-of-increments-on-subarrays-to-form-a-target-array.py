from typing import List
import unittest

# https://leetcode.com/problems/minimum-number-of-increments-on-subarrays-to-form-a-target-array/

# python3 -m unittest greedy/1526-minimum-number-of-increments-on-subarrays-to-form-a-target-array.py


class Solution(unittest.TestCase):
    def minNumberOperations(self, target: List[int]) -> int:
        return target[0] + sum(max(target[i] - target[i - 1], 0) for i in range(1, len(target)))

    def test(self):
        for target, expected in [
            ([1, 2, 3, 2, 1], 3),
            ([3, 1, 1, 2], 4),
            ([3, 1, 5, 4, 2], 7),
        ]:
            output = self.minNumberOperations(target)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
