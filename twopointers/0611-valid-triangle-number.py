# from bisect import bisect_left
from typing import List
import unittest

# https://leetcode.com/problems/valid-triangle-number/

# python3 -m unittest twopointers/0611-valid-triangle-number.py


class Solution(unittest.TestCase):
    # # Approach #1: Binary Search
    # # Time: O(nnlogn)
    # # Space: O(1)
    # def triangleNumber(self, nums: List[int]) -> int:
    #     nums.sort()
    #     n = len(nums)
    #     return sum(bisect_left(nums, nums[i] + nums[j], j + 1) - j - 1 for i in range(n - 2) for j in range(i + 1, n - 1))

    # Approach #2: Two Pointers
    # Time: O(nn)
    # Space: O(1)
    def triangleNumber(self, nums: List[int]) -> int:
        nums.sort()
        n, total = len(nums), 0
        for i in range(1, n):
            start, end = 0, i - 1
            while start < end:
                if nums[start] + nums[end] > nums[i]:
                    total += end - start
                    end -= 1
                else:
                    start += 1
        return total

    def test(self):
        for nums, expected in [
            ([2, 2, 3, 4], 3),
            ([4, 2, 3, 4], 4),
        ]:
            output = self.triangleNumber(nums)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
