from typing import List
import unittest

# https://leetcode.com/problems/find-the-number-of-subarrays-where-boundary-elements-are-maximum/

# python3 -m unittest stacks/monotonic/3113-find-the-number-of-subarrays-where-boundary-elements-are-maximum.py

class Solution(unittest.TestCase):
    def numberOfSubarrays(self, nums: List[int]) -> int:
        total = 0
        stack = []
        for num in nums:
            while stack and stack[-1][0] < num:
                stack.pop()
            if not stack or stack[-1][0] != num:
                stack.append([num,0])
            stack[-1][1] += 1
            total += stack[-1][1]
        return total

    def test(self):
        for nums, expected in [
            ([1,4,3,3,2], 6),
            ([3,3,3], 6),
            ([1], 1),
        ]:
            output = self.numberOfSubarrays(nums)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
