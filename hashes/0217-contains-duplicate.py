import unittest
from typing import List

# https://leetcode.com/problems/contains-duplicate/

# python3 -m unittest hashes/0217-contains-duplicate.py


class Solution(unittest.TestCase):
    def containsDuplicate(self, nums: List[int]) -> bool:
        return len(nums) != len(set(nums))

    def testContainsDuplicate(self) -> None:
        for nums, expected in [
            ([1, 2, 3, 1], True),
            ([1, 2, 3, 4], False),
            ([1, 1, 1, 3, 3, 4, 3, 2, 4, 2], True),
        ]:
            output = self.containsDuplicate(nums)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
