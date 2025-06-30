from collections import Counter
from typing import List
import unittest

# https://leetcode.com/problems/longest-harmonious-subsequence/

# python3 -m unittest slidingwindows/0594-longest-harmonious-subsequence.py

class Solution(unittest.TestCase):
    # # Approach #1: Sliding Window
    # # Time: O(n log n)
    # # Space: O(1)
    # def findLHS(self, nums: List[int]) -> int:
    #     nums.sort()
    #     length, left = 0, 0
    #     for right in range(len(nums)):
    #         while nums[right] - nums[left] > 1:
    #             left += 1
    #         if nums[right] - nums[left] == 1:
    #             length = max(length, right - left + 1)
    #     return length

    # Approach #2: Hash Table
    # Time: O(n)
    # Space: O(n)
    def findLHS(self, nums: List[int]) -> int:
        c = Counter(nums)
        return max([c[num] + c[num+1] for num in c if num+1 in c] or [0])

    def test(self):
        for nums, expected in [
            ([1,3,2,2,5,2,3,7], 5),
            ([1,2,3,4], 2),
            ([1,1,1,1], 0),
        ]:
            output = self.findLHS(nums)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
