from typing import List
import unittest

# https://leetcode.com/problems/special-array-i/

# python3 -m unittest arrays/3151-special-array-i.py

class Solution(unittest.TestCase):
    def isArraySpecial(self, nums: List[int]) -> bool:
        return all((nums[i-1]+nums[i])&1!=0 for i in range(1,len(nums)))

    def test(self):
        for nums, expected in [
            ([1], True),
            ([2,1,4], True),
            ([4,3,1,6], False),
            ([2], True),
            ([3], True),
            ([1,5,7], False),
            ([4,8,6,2], False),
            ([6,4,8,8,2,4,2,4,8,2,2,4,4,6,5], False),
        ]:
            output = self.isArraySpecial(nums)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
