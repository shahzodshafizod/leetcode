from typing import List
import unittest

# https://leetcode.com/problems/number-of-zero-filled-subarrays/

# python3 -m unittest arrays/2348-number-of-zero-filled-subarrays.py

# Formula for finding the number of subarrays: n*(n+1) / 2


class Solution(unittest.TestCase):
    # def zeroFilledSubarray(self, nums: List[int]) -> int:
    #     res, cnt = 0, 0
    #     for num in nums:
    #         if num == 0:
    #             cnt += 1
    #             res += cnt
    #         else:
    #             cnt = 0
    #     return res

    # Approach: Two-Pointers
    def zeroFilledSubarray(self, nums: List[int]) -> int:
        res, start = 0, 0
        for end in range(len(nums)):
            if nums[end] != 0:
                start = end + 1
            res += end - start + 1
        return res

    def test(self):
        for nums, expected in [
            ([1, 3, 0, 0, 2, 0, 0, 4], 6),
            ([0, 0, 0, 2, 0, 0], 9),
            ([2, 10, 2019], 0),
        ]:
            output = self.zeroFilledSubarray(nums)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
