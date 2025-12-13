from typing import List
import unittest

# https://leetcode.com/problems/max-consecutive-ones/

# python3 -m unittest arrays/0485-max-consecutive-ones.py


class Solution(unittest.TestCase):
    # Approach: Single pass tracking
    # Track current consecutive count and max.
    # Time: O(n) - single pass
    # Space: O(1) - constant space
    def findMaxConsecutiveOnes(self, nums: List[int]) -> int:
        max_count = 0
        current_count = 0

        for num in nums:
            if num == 1:
                current_count += 1
                max_count = max(max_count, current_count)
            else:
                current_count = 0

        return max_count

    def test(self):
        for nums, expected in [
            ([1, 1, 0, 1, 1, 1], 3),
            ([1, 0, 1, 1, 0, 1], 2),
            ([0, 0, 0], 0),
            ([1, 1, 1, 1], 4),
        ]:
            output = self.findMaxConsecutiveOnes(nums)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
