from typing import List
import unittest

# https://leetcode.com/problems/find-minimum-in-rotated-sorted-array-ii/

# python3 -m unittest binarysearch/0154-find-minimum-in-rotated-sorted-array-ii.py


class Solution(unittest.TestCase):
    # # Approach #1: Brute Force - Linear Scan
    # # Simply iterate through array to find minimum
    # # N = length of array
    # # Time: O(N) - scan entire array
    # # Space: O(1) - only use min variable
    # def findMin(self, nums: List[int]) -> int:
    #     return min(nums)

    # Approach #2: Optimized - Binary Search with duplicate handling
    # Compare mid with right boundary to determine which half to search
    # When duplicates cause ambiguity, shrink search space by decrementing right
    # N = length of array
    # Time: O(log N) average, O(N) worst case (all duplicates)
    # Space: O(1) - only use pointers
    def findMin(self, nums: List[int]) -> int:
        left, right = 0, len(nums) - 1

        while left < right:
            mid = (left + right) // 2

            if nums[mid] < nums[right]:
                # Minimum is in left half (including mid)
                # Example: [4, 5, 1, 2, 3], mid=1, right=3
                # 1 < 3, so minimum is in left half (could be mid itself)
                right = mid
            elif nums[mid] > nums[right]:
                # Minimum is in right half (excluding mid)
                # Example: [4, 5, 6, 7, 0, 1, 2], mid=7, right=2
                # 7 > 2, so minimum must be in right half
                left = mid + 1
            else:
                # nums[mid] == nums[right]
                # Can't determine which half has minimum
                # Example: [2, 2, 2, 0, 1] or [1, 2, 2, 2, 2]
                # Decrement right to shrink search space
                right -= 1

        return nums[left]

    def test(self):
        for nums, expected in [
            ([1, 3, 5], 1),
            ([2, 2, 2, 0, 1], 0),
            ([3, 3, 1, 3], 1),
            ([1, 1], 1),
            ([2, 2, 2, 2, 2], 2),
            ([3, 1, 3], 1),
            ([10, 1, 10, 10, 10], 1),
            ([4, 5, 6, 7, 0, 1, 2], 0),
            ([0, 1, 2, 4, 5, 6, 7], 0),
            ([7, 0, 1, 2, 4, 5, 6], 0),
        ]:
            output = self.findMin(nums)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
