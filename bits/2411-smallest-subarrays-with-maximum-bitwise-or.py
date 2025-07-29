from typing import List
import unittest

# https://leetcode.com/problems/smallest-subarrays-with-maximum-bitwise-or/

# python3 -m unittest bits/2411-smallest-subarrays-with-maximum-bitwise-or.py


class Solution(unittest.TestCase):
    def smallestSubarrays(self, nums: List[int]) -> List[int]:
        pos, res = [0] * 32, [0] * len(nums)
        for idx in range(len(nums) - 1, -1, -1):
            bit, num = 0, nums[idx]
            while num:
                if num & 1:
                    pos[bit] = idx
                num >>= 1
                bit += 1
            res[idx] = max(1, max(pos) - idx + 1)
        return res

    def test(self):
        for nums, expected in [
            ([1, 0, 2, 1, 3], [3, 3, 2, 2, 1]),
            ([1, 2], [2, 1]),
        ]:
            output = self.smallestSubarrays(nums)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
