from typing import List
import unittest

# https://leetcode.com/problems/count-the-number-of-good-subarrays/

# python3 -m unittest slidingwindows/2537-count-the-number-of-good-subarrays.py


class Solution(unittest.TestCase):
    def countGood(self, nums: List[int], k: int) -> int:
        count = {}
        goods, pairs = 0, 0
        n, start = len(nums), 0
        for end in range(n):
            count[nums[end]] = count.get(nums[end], 0) + 1
            pairs += count[nums[end]] - 1
            while pairs >= k:
                count[nums[start]] -= 1
                pairs -= count[nums[start]]
                start += 1
                goods += n - end
        return goods

    def test(self):
        for nums, k, expected in [
            ([1, 1, 1, 1, 1], 10, 1),
            ([3, 1, 4, 3, 2, 2, 4], 2, 4),
        ]:
            output = self.countGood(nums, k)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
