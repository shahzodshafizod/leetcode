from typing import List
import unittest

# https://leetcode.com/problems/longest-subarray-with-maximum-bitwise-and/

# python3 -m unittest bits/2419-longest-subarray-with-maximum-bitwise-and.py


class Solution(unittest.TestCase):
    # # Approach #1: Two-Pass
    # # Time: O(n)
    # # Space: O(1)
    # def longestSubarray(self, nums: List[int]) -> int:
    #     target, length = max(nums), 0
    #     maxlen = 1
    #     for num in nums:
    #         if num == target:
    #             length += 1
    #         else:
    #             length = 0
    #         maxlen = max(maxlen, length)
    #     return maxlen

    # Approach #2: One-Pass
    # Time: O(n)
    # Space: O(1)
    def longestSubarray(self, nums: List[int]) -> int:
        # 1. if n < cur_max, n & cur_max < cur_max
        # 2. if n == cur_max, n & cur_max = cur_max
        # 3. if n > cur_max, n & cur_max < cur_max
        cur_max, length = 0, 0
        maxlen = 0
        for num in nums:
            if num > cur_max:
                cur_max = num
                length = 1
                maxlen = 0
            elif num == cur_max:
                length += 1
            else:
                length = 0
            maxlen = max(maxlen, length)
        return maxlen

    def test(self):
        for nums, expected in [
            ([1, 2, 3, 3, 2, 2], 2),
            ([1, 2, 3, 4], 1),
        ]:
            output = self.longestSubarray(nums)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
