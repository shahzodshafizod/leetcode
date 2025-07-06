from typing import List
import unittest

# https://leetcode.com/problems/house-robber-iv/

# python3 -m unittest binarysearch/2560-house-robber-iv.py


class Solution(unittest.TestCase):
    def minCapability(self, nums: List[int], k: int) -> int:
        n = len(nums)
        low, high = 1, max(nums)
        amount = 0
        while low <= high:
            mid = (low + high) // 2
            idx, count = 0, 0
            while idx < n:
                if nums[idx] <= mid:
                    count += 1
                    idx += 2
                else:
                    idx += 1
            if count >= k:
                amount = mid
                high = mid - 1
            else:
                low = mid + 1
        return amount

    def test(self):
        for nums, k, expected in [
            ([2, 3, 5, 9], 2, 5),
            ([2, 7, 9, 3, 1], 2, 2),
        ]:
            output = self.minCapability(nums, k)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
