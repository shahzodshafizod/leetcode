from typing import List
import unittest

# https://leetcode.com/problems/make-array-elements-equal-to-zero/

# python3 -m unittest prefixsums/3354-make-array-elements-equal-to-zero.py


class Solution(unittest.TestCase):
    def countValidSelections(self, nums: List[int]) -> int:
        ans = 0
        left, right = 0, sum(nums)
        for num in nums:
            left += num
            right -= num
            if num == 0:
                if right == left:
                    ans += 2
                if abs(left - right) == 1:
                    ans += 1
        return ans

    def test(self):
        for nums, expected in [
            ([1, 0, 2, 0, 3], 2),
            ([2, 3, 4, 0, 4, 1, 0], 0),
        ]:
            output = self.countValidSelections(nums)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
