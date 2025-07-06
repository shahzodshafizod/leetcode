from typing import List
import unittest

# https://leetcode.com/problems/count-subarrays-with-fixed-bounds/

# python3 -m unittest slidingwindows/2444-count-subarrays-with-fixed-bounds.py


class Solution(unittest.TestCase):
    def countSubarrays(self, nums: List[int], minK: int, maxK: int) -> int:
        total = 0
        bad_id, min_id, max_id = -1, -1, -1
        for idx, num in enumerate(nums):
            if num == minK:
                min_id = idx
            if num == maxK:
                max_id = idx
            if minK > num or num > maxK:
                bad_id, min_id, max_id = idx, idx, idx
            total += min(min_id, max_id) - bad_id
        return total

    def test(self):
        for nums, minK, maxK, expected in [
            ([1, 1, 1, 1], 1, 1, 10),
            ([1, 3, 5, 2, 7, 5], 1, 5, 2),
            ([1, 2, 1, 4, 1, 3, 4], 1, 4, 16),
            ([1, 3, 2, 2, 1, 3, 2, 2], 1, 3, 20),
        ]:
            output = self.countSubarrays(nums, minK, maxK)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
