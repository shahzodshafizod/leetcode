from typing import List
import unittest

# https://leetcode.com/problems/count-hills-and-valleys-in-an-array/

# python3 -m unittest arrays/2210-count-hills-and-valleys-in-an-array.py


class Solution(unittest.TestCase):
    def countHillValley(self, nums: List[int]) -> int:
        prv, cnt = 0, 0
        for cur in range(1, len(nums) - 1):
            if nums[prv] < nums[cur] > nums[cur + 1] or nums[prv] > nums[cur] < nums[cur + 1]:
                cnt += 1
                prv = cur
        return cnt

    def test(self):
        for nums, expected in [
            ([2, 4, 1, 1, 6, 5], 3),
            ([6, 6, 5, 5, 4, 1], 0),
        ]:
            output = self.countHillValley(nums)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
