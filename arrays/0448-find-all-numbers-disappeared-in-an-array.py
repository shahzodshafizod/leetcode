from typing import List, Tuple
import unittest

# https://leetcode.com/problems/find-all-numbers-disappeared-in-an-array/

# python3 -m unittest arrays/0448-find-all-numbers-disappeared-in-an-array.py


class Solution(unittest.TestCase):
    # # Approach #1: Brute Force - Use Hash Set
    # # Time: O(n) - create set and check each number
    # # Space: O(n) - hash set
    # def findDisappearedNumbers(self, nums: List[int]) -> List[int]:
    #     num_set = set(nums)
    #     absent: List[int] = []
    #     for num in range(1, len(nums) + 1):
    #         if num not in num_set:
    #             absent.append(num)
    #     return absent

    # Approach #2: Mark Visited with Negation
    # Time: O(n) - two passes through array
    # Space: O(1) - modify array in-place
    def findDisappearedNumbers(self, nums: List[int]) -> List[int]:
        # Mark visited numbers by negating values at their indices
        for num in nums:
            idx = abs(num) - 1
            if nums[idx] > 0:
                nums[idx] = -nums[idx]

        # Collect indices with positive values
        absent: List[int] = []
        for idx in range(len(nums)):
            if nums[idx] > 0:
                absent.append(idx + 1)

        return absent

    def test(self):
        test_cases: List[Tuple[List[int], List[int]]] = [
            ([4, 3, 2, 7, 8, 2, 3, 1], [5, 6]),
            ([1, 1], [2]),
            ([1, 2, 3, 4, 5], []),
            ([2, 2, 2], [1, 3]),
        ]
        for nums, expected in test_cases:
            output = self.findDisappearedNumbers(nums)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
