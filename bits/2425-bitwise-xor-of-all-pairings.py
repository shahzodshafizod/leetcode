from typing import List
import unittest

# https://leetcode.com/problems/bitwise-xor-of-all-pairings/

# python3 -m unittest bits/2425-bitwise-xor-of-all-pairings.py

class Solution(unittest.TestCase):
    # Approach: Bit Manipulation
    # Time: O(n1 + n2)
    # Space: O(1)
    def xorAllNums(self, nums1: List[int], nums2: List[int]) -> int:
        num3 = 0
        if len(nums2)&1:
            for num in nums1:
                num3 ^= num
        if len(nums1)&1:
            for num in nums2:
                num3 ^= num
        return num3

    def test(self):
        for nums1, nums2, expected in [
            ([2,1,3], [10,2,5,0], 13),
            ([1,2], [3,4], 0),
            ([2], [4], 6),
            ([27,6,15], [8,23,15,16,25,5], 28),
            ([10,0,15,25,29,3,2], [12], 12),
            ([16,15,8,6,29,2,96,29], [24,12,13], 115),
            ([29,10,25,6], [24,11,1,17,11,5,1,1,27], 8),
        ]:
            output = self.xorAllNums(nums1, nums2)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
