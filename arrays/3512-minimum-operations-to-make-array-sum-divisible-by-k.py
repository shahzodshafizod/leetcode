from typing import List
import unittest

# https://leetcode.com/problems/minimum-operations-to-make-array-sum-divisible-by-k/

# python3 -m unittest arrays/3512-minimum-operations-to-make-array-sum-divisible-by-k.py


class Solution(unittest.TestCase):
    # Approach: Calculate remainder and operations needed
    # Time: O(n) - sum calculation
    # Space: O(1) - constant space
    def minOperations(self, nums: List[int], k: int) -> int:
        return sum(nums) % k

    def test(self):
        for nums, k, expected in [
            ([1, 2, 3, 4], 5, 0),
            ([1, 2, 3], 7, 6),
            ([5, 6, 7], 9, 0),
            ([1], 3, 1),
        ]:
            output = self.minOperations(nums, k)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
