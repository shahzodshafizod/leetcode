from typing import List
import unittest

# https://leetcode.com/problems/count-complete-subarrays-in-an-array/

# python3 -m unittest slidingwindows/2799-count-complete-subarrays-in-an-array.py


class Solution(unittest.TestCase):
    def countCompleteSubarrays(self, nums: List[int]) -> int:
        k = len(set(nums))
        total = start = 0
        count = {}
        for num in nums:
            count[num] = count.get(num, 0) + 1
            if count[num] == 1:
                k -= 1
            while k == 0:
                if count[nums[start]] == 1:
                    k += 1
                count[nums[start]] -= 1
                start += 1
            total += start
        return total

    def test(self):
        for nums, expected in [
            ([1, 3, 1, 2, 2], 4),
            ([5, 5, 5, 5], 10),
            ([1898, 370, 822, 1659, 1360, 128, 370, 360, 261, 1898], 4),
        ]:
            output = self.countCompleteSubarrays(nums)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
