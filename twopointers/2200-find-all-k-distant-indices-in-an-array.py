from typing import List
import unittest

# https://leetcode.com/problems/find-all-k-distant-indices-in-an-array/

# python3 -m unittest twopointers/2200-find-all-k-distant-indices-in-an-array.py

class Solution(unittest.TestCase):
    # Approach: Two Pointers
    # Time: O(n)
    # Space: O(1)
    def findKDistantIndices(self, nums: List[int], key: int, k: int) -> List[int]:
        indices = []
        n, right = len(nums), 0
        for idx in range(n):
            if nums[idx] == key:
                left = max(right, idx-k)
                right = min(n-1, idx+k)+1
                while left < right:
                    indices.append(left)
                    left += 1
        return indices

    def test(self):
        for nums, key, k, expected in [
            ([3,4,9,1,3,9,5], 9, 1, [1,2,3,4,5,6]),
            ([2,2,2,2,2], 2, 2, [0,1,2,3,4]),
        ]:
            output = self.findKDistantIndices(nums, key, k)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
