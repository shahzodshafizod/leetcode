from typing import List
import unittest

# https://leetcode.com/problems/keep-multiplying-found-values-by-two/

# python3 -m unittest hashes/2154-keep-multiplying-found-values-by-two.py


class Solution(unittest.TestCase):
    # Approach: Hash Set Lookup
    # Time: O(n + log(max_value)) - build set + lookups
    # Space: O(n) - hash set
    def findFinalValue(self, nums: List[int], original: int) -> int:
        snums = set(nums)
        while original in snums:
            original *= 2
        return original

    def test(self):
        for nums, original, expected in [
            ([5, 3, 6, 1, 12], 3, 24),
            ([2, 7, 9], 4, 4),
            ([8, 19, 4, 2, 15, 3], 2, 16),
            ([1], 1, 2),
            ([1, 2], 1, 4),
        ]:
            output = self.findFinalValue(nums, original)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
