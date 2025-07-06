from typing import List
import unittest

# https://leetcode.com/problems/maximum-count-of-positive-integer-and-negative-integer/

# python3 -m unittest binarysearch/2529-maximum-count-of-positive-integer-and-negative-integer.py


class Solution(unittest.TestCase):
    def maximumCount(self, nums: List[int]) -> int:
        # return max(bisect_left(nums, 0), len(nums)-bisect_right(nums, 0))
        n = len(nums)
        left, right = 0, n - 1
        while left <= right:
            mid = (left + right) // 2
            if nums[mid] >= 0:
                right = mid - 1
            else:
                left = mid + 1
        neg = right + 1
        left, right = 0, n - 1
        while left <= right:
            mid = (left + right) // 2
            if nums[mid] <= 0:
                left = mid + 1
            else:
                right = mid - 1
        pos = n - left
        return max(neg, pos)

    def test(self):
        for nums, expected in [
            ([-2, -1, -1, 1, 2, 3], 3),
            ([-3, -2, -1, 0, 0, 1, 2], 3),
            ([5, 20, 66, 1314], 4),
        ]:
            output = self.maximumCount(nums)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
