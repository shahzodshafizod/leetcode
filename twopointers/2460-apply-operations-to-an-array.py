from typing import List
import unittest

# https://leetcode.com/problems/apply-operations-to-an-array/

# python3 -m unittest twopointers/2460-apply-operations-to-an-array.py

class Solution(unittest.TestCase):
    # Approach: Brute-Force
    # Time: O(n)
    # Space: O(1)
    def applyOperations(self, nums: List[int]) -> List[int]:
        n, left = len(nums), 0
        for idx in range(n):
            if nums[idx] == 0:
                continue
            if idx < n-1 and nums[idx] == nums[idx+1]:
                nums[left] = nums[idx] * 2
                nums[idx+1] = 0
            else:
                nums[left] = nums[idx]
            left += 1
        while left < n:
            nums[left] = 0
            left += 1
        return nums

    def test(self):
        for nums, expected in [
            ([1,2,2,1,1,0], [1,4,2,0,0,0]),
            ([0,1], [1,0]),
            ([1,4,0,2,0,0], [1,4,2,0,0,0]),
            ([847,847,0,0,0,399,416,416,879,879,206,206,206,272], [1694,399,832,1758,412,206,272,0,0,0,0,0,0,0]),
        ]:
            output = self.applyOperations(nums)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
