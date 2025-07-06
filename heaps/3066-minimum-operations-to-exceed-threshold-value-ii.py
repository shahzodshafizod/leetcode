import heapq
from typing import List
import unittest

# https://leetcode.com/problems/minimum-operations-to-exceed-threshold-value-ii/

# python3 -m unittest heaps/3066-minimum-operations-to-exceed-threshold-value-ii.py


class Solution(unittest.TestCase):
    def minOperations(self, nums: List[int], k: int) -> int:
        heapq.heapify(nums)
        count = 0
        while len(nums) >= 2 and nums[0] < k:
            first = heapq.heappop(nums)
            second = heapq.heappop(nums)
            heapq.heappush(nums, first * 2 + second)
            count += 1
        return count

    def test(self):
        for nums, k, expected in [
            ([2, 11, 10, 1, 3], 10, 2),
            ([1, 1, 2, 4, 9], 20, 4),
            ([80, 47], 81, 1),
            ([9, 98, 52, 8], 2, 0),
            ([999999999, 999999999, 999999999], 1000000000, 2),
            ([62, 32, 62, 73, 58, 56, 68, 50], 74, 4),
            ([69, 89, 57, 31, 84, 97, 50, 38, 91, 86], 91, 4),
            ([1000000000, 999999999, 1000000000, 999999999, 1000000000, 999999999], 1000000000, 2),
        ]:
            output = self.minOperations(nums, k)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
