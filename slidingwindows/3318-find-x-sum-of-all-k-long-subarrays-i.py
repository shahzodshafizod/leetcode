from collections import Counter
from typing import List
import unittest

# https://leetcode.com/problems/find-x-sum-of-all-k-long-subarrays-i/

# python3 -m unittest slidingwindows/3318-find-x-sum-of-all-k-long-subarrays-i.py


class Solution(unittest.TestCase):
    def findXSum(self, nums: List[int], k: int, x: int) -> List[int]:
        return [sum(cnt * num for cnt, num in sorted(((cnt, num) for num, cnt in Counter(nums[i : i + k]).items()), reverse=True)[:x] if cnt > 0) for i in range(len(nums) - k + 1)]

    def test(self):
        for nums, k, x, expected in [
            ([1, 1, 2, 2, 3, 4, 2, 3], 6, 2, [6, 10, 12]),
            ([3, 8, 7, 8, 7, 5], 2, 2, [11, 15, 15, 15, 12]),
        ]:
            output = self.findXSum(nums, k, x)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
