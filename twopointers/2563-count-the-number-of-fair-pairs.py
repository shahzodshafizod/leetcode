from typing import List
import unittest

# https://leetcode.com/problems/count-the-number-of-fair-pairs/

# python3 -m unittest twopointers/2563-count-the-number-of-fair-pairs.py


class Solution(unittest.TestCase):
    # Approach: Two Pointers
    # Time: O(n log n)
    # Space: O(1)
    def countFairPairs(self, nums: List[int], lower: int, upper: int) -> int:
        nums.sort()

        def lessThan(top: int) -> int:
            pairs = 0
            left, right = 0, len(nums) - 1
            while left < right:
                if nums[left] + nums[right] < top:
                    pairs += right - left
                    left += 1
                else:
                    right -= 1
            return pairs

        return lessThan(upper + 1) - lessThan(lower)

    def test(self):
        for nums, lower, upper, expected in [
            ([0, 1, 7, 4, 4, 5], 3, 6, 6),
            ([1, 7, 9, 2, 5], 11, 11, 1),
        ]:
            output = self.countFairPairs(nums, lower, upper)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
