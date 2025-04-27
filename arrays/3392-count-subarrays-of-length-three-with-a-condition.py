from typing import List
import unittest

# https://leetcode.com/problems/count-subarrays-of-length-three-with-a-condition/

# python3 -m unittest arrays/3392-count-subarrays-of-length-three-with-a-condition.py

class Solution(unittest.TestCase):
    def countSubarrays(self, nums: List[int]) -> int:
        return sum((nums[i-2]+nums[i])*2 == nums[i-1] for i in range(2,len(nums)))

    def test(self):
        for nums, expected in [
            ([1,2,1,4,1], 1),
            ([1,1,1], 0),
        ]:
            output = self.countSubarrays(nums)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
