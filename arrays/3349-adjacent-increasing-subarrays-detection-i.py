from typing import List
import unittest

# https://leetcode.com/problems/adjacent-increasing-subarrays-detection-i/

# python3 -m unittest arrays/3349-adjacent-increasing-subarrays-detection-i.py


class Solution(unittest.TestCase):
    def hasIncreasingSubarrays(self, nums: List[int], k: int) -> bool:
        n = len(nums)
        prv, cur = 0, 1
        for i in range(1, n):
            if nums[i - 1] < nums[i]:
                cur += 1
            else:
                prv, cur = cur, 1
            # cur//2 means: Both subarrays are part of the same increasing segment
            if max(cur // 2, min(prv, cur)) >= k:
                return True
        return False

    def test(self):
        for nums, k, expected in [
            ([2, 5, 7, 8, 9, 2, 3, 4, 3, 1], 3, True),
            ([1, 2, 3, 4, 4, 4, 4, 5, 6, 7], 5, False),
        ]:
            output = self.hasIncreasingSubarrays(nums, k)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
