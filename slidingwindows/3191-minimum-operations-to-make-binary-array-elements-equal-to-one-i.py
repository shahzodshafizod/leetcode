from typing import List
import unittest

# https://leetcode.com/problems/minimum-operations-to-make-binary-array-elements-equal-to-one-i/

# python3 -m unittest slidingwindows/3191-minimum-operations-to-make-binary-array-elements-equal-to-one-i.py


class Solution(unittest.TestCase):
    def minOperations(self, nums: List[int]) -> int:
        flips = 0
        curr, nxt = nums[0], nums[1]
        prev = -1
        for idx in range(2, len(nums)):
            prev, curr = curr, nxt
            nxt = nums[idx]
            if prev == 0:
                curr ^= 1
                nxt ^= 1
                flips += 1
        return -1 if curr == 0 or nxt == 0 else flips

    def test(self):
        for nums, expected in [
            ([0, 1, 1, 1, 0, 0], 3),
            ([0, 1, 1, 1], -1),
        ]:
            output = self.minOperations(nums)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
