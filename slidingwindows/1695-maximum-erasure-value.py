from collections import defaultdict
from typing import List
import unittest

# https://leetcode.com/problems/maximum-erasure-value/

# python3 -m unittest slidingwindows/1695-maximum-erasure-value.py


class Solution(unittest.TestCase):
    def maximumUniqueSubarray(self, nums: List[int]) -> int:
        score, maxscore, left = 0, 0, 0
        count = defaultdict(int)
        for num in nums:
            score += num
            count[num] += 1
            while count[num] > 1:
                count[nums[left]] -= 1
                score -= nums[left]
                left += 1
            maxscore = max(maxscore, score)
        return maxscore

    def test(self):
        for nums, expected in [
            ([4, 2, 4, 5, 6], 17),
            ([5, 2, 1, 2, 5, 2, 1, 2, 5], 8),
        ]:
            output = self.maximumUniqueSubarray(nums)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
