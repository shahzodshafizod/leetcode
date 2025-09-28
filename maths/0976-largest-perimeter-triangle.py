from typing import List
import unittest

# https://leetcode.com/problems/largest-perimeter-triangle/

# python3 -m unittest maths/0976-largest-perimeter-triangle.py


class Solution(unittest.TestCase):
    def largestPerimeter(self, nums: List[int]) -> int:
        nums.sort()
        for i in range(len(nums) - 3, -1, -1):
            if nums[i] + nums[i + 1] > nums[i + 2]:
                return nums[i] + nums[i + 1] + nums[i + 2]
        return 0

    def test(self):
        for nums, expected in [
            ([2, 1, 2], 5),
            ([1, 2, 1, 10], 0),
        ]:
            output = self.largestPerimeter(nums)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
