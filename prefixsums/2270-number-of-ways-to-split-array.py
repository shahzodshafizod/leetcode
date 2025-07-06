from typing import List
import unittest

# https://leetcode.com/problems/number-of-ways-to-split-array/

# python3 -m unittest prefixsums/2270-number-of-ways-to-split-array.py


class Solution(unittest.TestCase):
    def waysToSplitArray(self, nums: List[int]) -> int:
        left, right = sum(nums), 0
        count = 0
        for idx in range(len(nums) - 1, -1, -1):
            left -= nums[idx]
            right += nums[idx]
            if left >= right:
                count += 1
        return count

    def test(self):
        for nums, expected in [
            ([10, 4, -8, 7], 2),
            ([2, 3, 1, 0], 2),
        ]:
            output = self.waysToSplitArray(nums)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
