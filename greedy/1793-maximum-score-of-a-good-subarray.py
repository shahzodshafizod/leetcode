from typing import List
import unittest

# https://leetcode.com/problems/maximum-score-of-a-good-subarray/

# python3 -m unittest greedy/1793-maximum-score-of-a-good-subarray.py


class Solution(unittest.TestCase):
    def maximumScore(self, nums: List[int], k: int) -> int:
        n = len(nums)
        minimum = nums[k]
        score = nums[k]
        left = right = k
        while left > 0 or right + 1 < n:
            lnum = nums[left - 1] if left > 0 else 0
            rnum = nums[right + 1] if right + 1 < n else 0
            if lnum > rnum:
                left -= 1
                minimum = min(minimum, lnum)
            else:
                right += 1
                minimum = min(minimum, rnum)
            score = max(score, minimum * (right - left + 1))
        return score

    def test(self):
        for nums, k, expected in [
            ([1, 4, 3, 7, 4, 5], 3, 15),
            ([5, 5, 4, 5, 4, 1, 1, 1], 0, 20),
            ([1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 2], 28, 30),
        ]:
            output = self.maximumScore(nums, k)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
