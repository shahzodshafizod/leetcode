from typing import List
import unittest

# https://leetcode.com/problems/check-if-array-is-sorted-and-rotated/

# python3 -m unittest arrays/1752-check-if-array-is-sorted-and-rotated.py

class Solution(unittest.TestCase):
    # # Approach: Sliding Window
    # # Time: O(n)
    # # Space: O(1)
    # def check(self, nums: List[int]) -> bool:
    #     n = len(nums)
    #     count = 1
    #     for idx in range(1, 2*n):
    #         if nums[(idx-1)%n] <= nums[idx%n]:
    #             count += 1
    #         else:
    #             count = 1
    #         if count == n:
    #             return True
    #     return n == 1

    # Approach: Finding Smallest Elements
    # Time: O(n)
    # Space: O(1)
    def check(self, nums: List[int]) -> bool:
        breaks = 0
        for idx in range(1, len(nums)):
            if nums[idx-1] > nums[idx]:
                breaks += 1
        if nums[-1] > nums[0]:
            breaks += 1
        return breaks <= 1

    def test(self):
        for nums, expected in [
            ([3,4,5,1,2], True),
            ([2,1,3,4], False),
            ([1,2,3], True),
            ([1,3,2], False),
            ([2,1], True),
            ([2,4,1,3], False),
            ([2,1,2,2,2], True),
            ([6,7,2,7,5], False),
            ([7,9,1,1,1,3], True),
            ([10,10,1,4,5,10], True),
        ]:
            output = self.check(nums)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
