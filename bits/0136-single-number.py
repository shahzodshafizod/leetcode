from typing import List
import unittest

# https://leetcode.com/problems/single-number/

# python3 -m unittest bits/0136-single-number.py

class Solution(unittest.TestCase):
    def singleNumber(self, nums: List[int]) -> int:
        single = 0
        for num in nums:
            single ^= num
        return single

    def test(self) -> None:
        for nums, expected in [
            ([2,2,1], 1),
            ([4,1,2,1,2], 4),
            ([1], 1),
        ]:
            output = self.singleNumber(nums)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
