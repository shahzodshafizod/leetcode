from typing import List
import unittest

# https://leetcode.com/problems/the-two-sneaky-numbers-of-digitville/

# python3 -m unittest bits/3289-the-two-sneaky-numbers-of-digitville.py


class Solution(unittest.TestCase):
    # # Approach #1: Counting
    # # Time: O(n)
    # # Space: O(n)
    # def getSneakyNumbers(self, nums: List[int]) -> List[int]:
    #     cnt: List[int] = [0] * (len(nums) - 2)
    #     res: List[int] = []
    #     for num in nums:
    #         cnt[num] += 1
    #         if cnt[num] > 1:
    #             res.append(num)
    #     return res

    # # Approach #2: Sorting
    # # Time: O(n log n)
    # # Space: O(n)
    # def getSneakyNumbers(self, nums: List[int]) -> List[int]:
    #     nums.sort()
    #     res: List[int] = []
    #     for i in range(1, len(nums)):
    #         if nums[i - 1] == nums[i]:
    #             res.append(nums[i])
    #     return res

    # Approach #3: Bit Manipulations
    # Time: O(n)
    # Space: O(1)
    def getSneakyNumbers(self, nums: List[int]) -> List[int]:
        xor, n = 0, len(nums) - 2
        for num in nums:
            xor ^= num
        for idx in range(n):
            xor ^= idx
        last_bit = xor & -xor
        num1, num2 = 0, 0
        for num in nums:
            if num & last_bit == 0:
                num1 ^= num
            else:
                num2 ^= num
        for idx in range(n):
            if idx & last_bit == 0:
                num1 ^= idx
            else:
                num2 ^= idx
        return [num1, num2]

    def test(self):
        for nums, expected in [
            ([0, 1, 1, 0], [0, 1]),
            ([0, 3, 2, 1, 3, 2], [2, 3]),
            ([7, 1, 5, 4, 3, 4, 6, 0, 9, 5, 8, 2], [4, 5]),
        ]:
            output = self.getSneakyNumbers(nums)
            output.sort()
            expected.sort()
            self.assertListEqual(expected, output, f"expected: {expected}, output: {output}")
