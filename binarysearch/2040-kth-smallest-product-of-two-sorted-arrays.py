from math import ceil
from bisect import bisect_left, bisect_right
from typing import List
import unittest

# https://leetcode.com/problems/kth-smallest-product-of-two-sorted-arrays/

# python3 -m unittest binarysearch/2040-kth-smallest-product-of-two-sorted-arrays.py

class Solution(unittest.TestCase):
    def kthSmallestProduct(self, nums1: List[int], nums2: List[int], k: int) -> int:
        if len(nums1) > len(nums2):
            nums1, nums2 = nums2, nums1
        m = len(nums2)
        def calc_less_than(target: int) -> int:
            count = 0
            for n1 in nums1:
                if n1 > 0:
                    count += bisect_right(nums2, target // n1)
                elif n1 < 0:
                    count += m - bisect_left(nums2, ceil(target / n1))
                elif target >= 0:
                    count += m
            return count
        left, right = -10**10, 10**10
        while left <= right:
            mid = (left + right) // 2
            if calc_less_than(mid) >= k:
                right = mid - 1
            else:
                left = mid + 1
        return left

    def test(self):
        for nums1, nums2, k, expected in [
            ([2,5], [3,4], 2, 8),
            ([-4,-2,0,3], [2,4], 6, 0),
            ([-2,-1,0,1,2], [-3,-1,2,4,5], 3, -6),
            ([-6], [-9], 1, 54),
        ]:
            output = self.kthSmallestProduct(nums1, nums2, k)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
