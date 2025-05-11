from typing import List
import unittest

# https://leetcode.com/problems/minimum-equal-sum-of-two-arrays-after-replacing-zeros/

# python3 -m unittest greedy/2918-minimum-equal-sum-of-two-arrays-after-replacing-zeros.py

class Solution(unittest.TestCase):
    def minSum(self, nums1: List[int], nums2: List[int]) -> int:
        sum1, zeroes1 = sum(nums1), nums1.count(0)
        sum2, zeroes2 = sum(nums2), nums2.count(0)
        sum1 += zeroes1
        sum2 += zeroes2
        if zeroes1 == 0 and sum1 < sum2 or zeroes2 == 0 and sum2 < sum1:
            return -1
        return max(sum1, sum2)

    def test(self):
        for nums1, nums2, expected in [
            ([3,2,0,1,0], [6,5,0], 12),
            ([2,0,2,0], [1,4], -1),
        ]:
            output = self.minSum(nums1, nums2)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
