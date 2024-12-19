from typing import List
import unittest

# https://leetcode.com/problems/find-pivot-index/

# python3 -m unittest prefixsums/0724-find-pivot-index.py

class Solution(unittest.TestCase):
    def pivotIndex(self, nums: List[int]) -> int:
        lsum, rsum = 0, sum(nums)
        pivot = -1
        for idx, num in enumerate(nums):
            rsum -= num
            if lsum == rsum:
                pivot = idx
                break
            lsum += num
        return pivot

    def test(self):
        for nums, expected in [
            ([1,7,3,6,5,6], 3),
            ([1,2,3], -1),
            ([2,1,-1], 0),
            ([-1,-1,-1,0,1,1], 0),
        ]:
            output = self.pivotIndex(nums)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
