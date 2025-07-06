from typing import List
import unittest

# https://leetcode.com/problems/type-of-triangle/

# python3 -m unittest maths/3024-type-of-triangle.py


class Solution(unittest.TestCase):
    def triangleType(self, nums: List[int]) -> str:
        nums.sort()
        return (
            "none"
            if nums[0] + nums[1] <= nums[2]
            else (
                "equilateral"
                if nums[0] == nums[2]
                else "isosceles" if nums[0] == nums[1] or nums[1] == nums[2] else "scalene"
            )
        )

    def test(self):
        for nums, expected in [
            ([3, 3, 3], "equilateral"),
            ([3, 4, 5], "scalene"),
            ([1, 1, 1], "equilateral"),
            ([3, 1, 2], "none"),
            ([2, 5, 2], "none"),
            ([3, 3, 1], "isosceles"),
            ([1, 1, 2], "none"),
            ([2, 3, 4], "scalene"),
            ([8, 4, 2], "none"),
            ([8, 4, 4], "none"),
        ]:
            output = self.triangleType(nums)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
