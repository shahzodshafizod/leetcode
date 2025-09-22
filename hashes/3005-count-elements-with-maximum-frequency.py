from collections import Counter
from typing import List
import unittest

# https://leetcode.com/problems/count-elements-with-maximum-frequency/

# python3 -m unittest hashes/3005-count-elements-with-maximum-frequency.py


class Solution(unittest.TestCase):
    def maxFrequencyElements(self, nums: List[int]) -> int:
        counter = Counter(nums)
        mx_freq = max(counter.values())
        return sum(cnt for cnt in counter.values() if cnt == mx_freq)

    def test(self):
        for nums, expected in [
            ([1, 2, 2, 3, 1, 4], 4),
            ([1, 2, 3, 4, 5], 5),
        ]:
            output = self.maxFrequencyElements(nums)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
