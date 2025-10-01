from typing import List
import unittest

# https://leetcode.com/problems/find-triangular-sum-of-an-array/

# python3 -m unittest maths/2221-find-triangular-sum-of-an-array.py


class Solution(unittest.TestCase):
    # Approach: Brute-Force (Simulation)
    # Time: O(nn)
    # Space: O(1)
    def triangularSum(self, nums: List[int]) -> int:
        for n in range(len(nums), 1, -1):
            for i in range(1, n):
                nums[i - 1] = (nums[i - 1] + nums[i]) % 10
        return nums[0]

    def test(self):
        for nums, expected in [
            ([1, 2, 3, 4, 5], 8),
            ([5], 5),
        ]:
            output = self.triangularSum(nums)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
