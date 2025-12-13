from typing import List
import unittest

# https://leetcode.com/problems/check-if-all-1s-are-at-least-length-k-places-away/

# python3 -m unittest arrays/1437-check-if-all-1s-are-at-least-length-k-places-away.py


class Solution(unittest.TestCase):
    # Approach: Single Pass with Distance Tracking
    # Time: O(n)
    # Space: O(1)
    def kLengthApart(self, nums: List[int], k: int) -> bool:
        distance = k
        for num in nums:
            if num == 0:
                distance += 1
            else:
                if distance < k:
                    return False
                distance = 0
        return True

    def test(self):
        for nums, k, expected in [
            ([1, 0, 0, 0, 1, 0, 0, 1], 2, True),
            ([1, 0, 0, 1, 0, 1], 2, False),
            ([1, 1, 1, 1, 1], 0, True),
            ([0, 1, 0, 1], 1, True),
            ([0], 1, True),
            ([1], 0, True),
            ([1, 0, 1], 1, True),
            ([1, 0, 1], 2, False),
        ]:
            output = self.kLengthApart(nums, k)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
