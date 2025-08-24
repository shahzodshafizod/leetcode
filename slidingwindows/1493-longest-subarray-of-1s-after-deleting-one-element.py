from typing import List
import unittest

# https://leetcode.com/problems/longest-subarray-of-1s-after-deleting-one-element/

# python3 -m unittest slidingwindows/1493-longest-subarray-of-1s-after-deleting-one-element.py


class Solution(unittest.TestCase):
    # # Approach #1: Dynamic Programming
    # # Time: O(nn)
    # # Space: O(nn)
    # def longestSubarray(self, nums: List[int]) -> int:
    #     n = len(nums)
    #     memo = [[[-1] * 2 for _ in range(n)] for _ in range(n)]

    #     def dp(idx: int, size: int, tries: int) -> int:
    #         if idx == n:
    #             return size - tries
    #         if memo[idx][size][tries] != -1:
    #             return memo[idx][size][tries]
    #         res = size
    #         if nums[idx] == 1:
    #             res = max(res, dp(idx + 1, size + 1, tries))
    #         elif tries > 0:
    #             res = max(res, dp(idx + 1, size, 0), dp(idx + 1, 0, 1))
    #         memo[idx][size][tries] = res
    #         return res

    #     return dp(0, 0, 1)

    # Approach #2: Sliding Window
    # Time: O(n)
    # Space: O(1)
    def longestSubarray(self, nums: List[int]) -> int:
        size, start, tries = 0, 0, 1
        for end in range(len(nums)):
            if nums[end] == 0:
                tries -= 1
            while tries < 0:
                if nums[start] == 0:
                    tries += 1
                start += 1
            size = max(size, end - start)
        return size

    def test(self):
        for nums, expected in [
            ([1, 1, 0, 1], 3),
            ([0, 1, 1, 1, 0, 1, 1, 0, 1], 5),
            ([1, 1, 1], 2),
        ]:
            output = self.longestSubarray(nums)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
