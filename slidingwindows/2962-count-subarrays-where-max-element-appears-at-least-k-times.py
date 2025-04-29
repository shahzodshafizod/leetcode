from typing import List
import unittest

# https://leetcode.com/problems/count-subarrays-where-max-element-appears-at-least-k-times/

# python3 -m unittest slidingwindows/2962-count-subarrays-where-max-element-appears-at-least-k-times.py

class Solution(unittest.TestCase):
    def countSubarrays(self, nums: List[int], k: int) -> int:
        total, start, maximum = 0, 0, max(nums)
        for idx in range(len(nums)):
            if nums[idx] == maximum:
                k -= 1
            while k == 0:
                if nums[start] == maximum:
                    k += 1
                start += 1
            total += start
        return total

    def test(self):
        for nums, k, expected in [
            ([10], 1, 1),
            ([1,4,2,1], 3, 0),
            ([1,1,1,1,1], 5, 1),
            ([1,3,2,3,3], 2, 6),
            ([3,3,3,3,3], 3, 6),
            ([1,2,3,4,5], 1, 5),
            ([5,5,1,5,2,5,5], 4, 3),
            ([1,2,2,2,3,2,2], 3, 0),
            ([4,4,4,1,4,4,4,4], 5, 6),
            ([1,5,5,5,2,5,5,1,5], 5, 7),
            ([2,2,2,2,2,2,2,2,2,2], 8, 6),
            ([1000000,1,1000000,2,1000000], 2, 5),
            # ([28,5,58,91,24,91,53,9,48,85,16,70,91,91,47,91,61,4,54,61,49], 1, 187),
            # ([1,3,2,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3], 2, 1221),
            # ([1,3,2,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,3], 1, 1273),
        ]:
            output = self.countSubarrays(nums, k)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
