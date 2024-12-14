import bisect
from typing import List
import unittest

# https://leetcode.com/problems/search-insert-position/

# python3 -m unittest binarysearch/0035-search-insert-position.py

class Solution(unittest.TestCase):
    def searchInsert(self, nums: List[int], target: int) -> int:
        # return bisect.bisect_left(nums, target)
        left, right = 0, len(nums)-1
        while left <= right:
            mid = (left+right) // 2
            if nums[mid] >= target:
                right = mid-1
            else:
                left = mid+1
        return left

    def testSearchInsert(self) -> None:
        for nums, target, expected in [
            ([1,3,5,6], 5, 2),
            ([1,3,5,6], 2, 1),
            ([1,3,5,6], 7, 4),
        ]:
            output = self.searchInsert(nums, target)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
