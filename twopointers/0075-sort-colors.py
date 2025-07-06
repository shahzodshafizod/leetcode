from typing import List
import unittest

# https://leetcode.com/problems/sort-colors/

# python3 -m unittest twopointers/0075-sort-colors.py

# DNF algorithm 🇳🇱
# The Dutch National Flag Algorithm, also known as the DNF algorithm or
# the Three-Way Partitioning Algorithm, is a simple and efficient approach
# to sorting an array containing three distinct elements.


class Solution(unittest.TestCase):
    # # Approach: Count Sort
    # # Time: O(n)
    # # Space: O(3)
    # def sortColors(self, nums: List[int]) -> None:
    #     """
    #     Do not return anything, modify nums in-place instead.
    #     """
    #     buckets = [0] * 3
    #     for num in nums:
    #         buckets[num] += 1
    #     c = 0
    #     for idx in range(len(nums)):
    #         if buckets[c] == 0:
    #             c += 1
    #         nums[idx] = c
    #         buckets[c] -= 1

    # Approach: Three Pointers (DNF Algorithm)
    # Time: O(n)
    # Space: O(1)
    def sortColors(self, nums: List[int]) -> None:
        """
        Do not return anything, modify nums in-place instead.
        """
        idx, left, right = 0, 0, len(nums) - 1
        while idx <= right:
            match nums[idx]:
                case 0:
                    nums[idx], nums[left] = nums[left], nums[idx]
                    left += 1
                case 2:
                    nums[idx], nums[right] = nums[right], nums[idx]
                    right -= 1
                    idx -= 1
            idx += 1

    def test(self):
        for nums, expected in [
            ([2, 0, 2, 1, 1, 0], [0, 0, 1, 1, 2, 2]),
            ([2, 0, 1], [0, 1, 2]),
        ]:
            self.sortColors(nums)
            output = nums
            self.assertListEqual(output, expected, f"output: {output}, expected: {expected}")
