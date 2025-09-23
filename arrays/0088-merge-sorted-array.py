from typing import List
import unittest

# https://leetcode.com/problems/merge-sorted-array/

# python3 -m unittest arrays/0088-merge-sorted-array.py


class Solution(unittest.TestCase):
    def merge(self, nums1: List[int], m: int, nums2: List[int], n: int) -> None:
        """
        Do not return anything, modify nums1 in-place instead.
        """
        for i in range(n):
            nums1[m+i] = nums2[i]
        nums1.sort()

    def test(self):
        for nums1, m, nums2, n, expected in [
            ([1,2,3,0,0,0], 3, [2,5,6], 3, [1,2,2,3,5,6]),
            ([1], 1, [], 0, [1]),
            ([0], 0, [1], 1, [1]),
        ]:
            self.merge(nums1, m, nums2, n)
            output = nums1
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
