from typing import List
import unittest

# https://leetcode.com/problems/max-sum-of-a-pair-with-equal-sum-of-digits/

# python3 -m unittest hashes/2342-max-sum-of-a-pair-with-equal-sum-of-digits.py

class Solution(unittest.TestCase):
    def maximumSum(self, nums: List[int]) -> int:
        pairs = {}
        max_sum = -1
        for num in nums:
            key = sum(int(c) for c in str(num))
            if key in pairs:
                if num > pairs[key][0]:
                    pairs[key][1] = pairs[key][0]
                    pairs[key][0] = num
                elif num > pairs[key][1]:
                    pairs[key][1] = num
                max_sum = max(max_sum, pairs[key][0] + pairs[key][1])
            else:
                pairs[key] = [num, 0]
        return max_sum

    def test(self):
        for nums, expected in [
            ([18,43,36,13,7], 54),
            ([10,12,19,14], -1),
            ([4], -1),
            ([5,1,6], -1),
            ([4,6,10,6], 12),
            ([2,1,5,5,2,4], 10),
        ]:
            output = self.maximumSum(nums)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
