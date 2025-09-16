from math import gcd
from typing import List
import unittest

# https://leetcode.com/problems/replace-non-coprime-numbers-in-array/

# python3 -m unittest stacks/2197-replace-non-coprime-numbers-in-array.py


class Solution(unittest.TestCase):
    def replaceNonCoprimes(self, nums: List[int]) -> List[int]:
        stack: List[int] = []
        for lcm in nums:
            while len(stack) > 0:
                g = gcd(stack[-1], lcm)
                if g == 1:
                    break
                lcm = stack[-1] * lcm // g
                stack.pop()
            stack.append(lcm)
        return stack

    def test(self):
        for nums, expected in [
            ([6, 4, 3, 2, 7, 6, 2], [12, 7, 6]),
            ([2, 2, 1, 1, 3, 3, 3], [2, 1, 1, 3]),
        ]:
            output = self.replaceNonCoprimes(nums)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
