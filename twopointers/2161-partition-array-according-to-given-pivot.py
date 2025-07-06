from typing import List
import unittest

# https://leetcode.com/problems/partition-array-according-to-given-pivot/

# python3 -m unittest twopointers/2161-partition-array-according-to-given-pivot.py


class Solution(unittest.TestCase):
    # Approach: Additional space
    # Time: O(n)
    # Space: O(n)
    def pivotArray(self, nums: List[int], pivot: int) -> List[int]:
        n = len(nums)
        ans = [pivot] * n
        left, right = 0, n - 1
        for idx in range(n):
            if nums[idx] < pivot:
                ans[left] = nums[idx]
                left += 1
            if nums[n - idx - 1] > pivot:
                ans[right] = nums[n - idx - 1]
                right -= 1
        return ans

    def test(self):
        for nums, pivot, expected in [
            ([9, 12, 5, 10, 14, 3, 10], 10, [9, 5, 3, 10, 10, 12, 14]),
            ([-3, 4, 3, 2], 2, [-3, 2, 4, 3]),
        ]:
            output = self.pivotArray(nums, pivot)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
