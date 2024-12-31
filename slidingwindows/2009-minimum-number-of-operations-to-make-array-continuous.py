from typing import List
import unittest

# https://leetcode.com/problems/minimum-number-of-operations-to-make-array-continuous/

# python3 -m unittest slidingwindows/2009-minimum-number-of-operations-to-make-array-continuous.py

class Solution(unittest.TestCase):
    # # Approach #1: Binary Search
    # # Time: O(n log n)
    # # Space: O(n), for set(nums)
    # def minOperations(self, nums: List[int]) -> int:
    #     origlen = len(nums)
    #     nums = sorted(set(nums))
    #     length = len(nums)
    #     ops = origlen
    #     for start in range(length):
    #         end_num = nums[start] + origlen - 1
    #         left, right = 0, length-1
    #         while left <= right:
    #             mid = (left + right) // 2
    #             if nums[mid] > end_num:
    #                 right = mid-1
    #             else:
    #                 left = mid+1
    #         window_size = right - start + 1
    #         ops = min(ops, origlen - window_size)
    #     return ops

    # Approach #2: Sliding Window
    # Time: O(n log n)
    # Space: O(n), for set(nums)
    def minOperations(self, nums: List[int]) -> int:
        length = len(nums)
        nums = sorted(set(nums))
        operations = length
        end = 0
        for start in range(len(nums)):
            # target window is: [start; start+length-1]
            while end < len(nums) and nums[end] <= nums[start] + length - 1:
                end += 1
            window_size = end - start
            operations = min(operations, length - window_size)
        return operations

    def test(self):
        for nums, expected in [
            ([4,2,5,3], 0),
            ([1,2,3,5,6], 1),
            ([1,10,100,1000], 3),
        ]:
            output = self.minOperations(nums)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
