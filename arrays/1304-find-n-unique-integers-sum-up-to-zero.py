from typing import List
import unittest

# https://leetcode.com/problems/find-n-unique-integers-sum-up-to-zero/

# python3 -m unittest arrays/1304-find-n-unique-integers-sum-up-to-zero.py


class Solution(unittest.TestCase):
    def sumZero(self, n: int) -> List[int]:
        nums: List[int] = [0] * n
        diff = 1 if n & 1 == 0 else 0
        for idx in range(1, n // 2 + 1):
            nums[idx - diff] = -idx
            nums[n - idx] = idx
        return nums

    def test(self):
        for n in [
            (5),
            (3),
            (1),
            (4),
        ]:
            nums = self.sumZero(n)
            if len(set(nums)) != n or sum(nums) != 0:
                self.assertEqual(0, sum(nums), f"sum(expected): {sum(nums)}, nums: {nums}")
