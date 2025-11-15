import statistics
from typing import List
import unittest

# https://leetcode.com/problems/majority-element/

# python3 -m unittest hashes/0169-majority-element.py


class Solution(unittest.TestCase):
    def majorityElement(self, nums: List[int]) -> int:
        # return Counter(nums).most_common(1)[0][0]
        # return sorted(nums)[len(nums) // 2]
        # return sorted((cnt, num) for num, cnt in Counter(nums).items())[-1][1]
        return statistics.mode(nums)

    def test(self):
        for nums, expected in [
            ([3, 2, 3], 3),
            ([2, 2, 1, 1, 1, 2, 2], 2),
        ]:
            output = self.majorityElement(nums)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
