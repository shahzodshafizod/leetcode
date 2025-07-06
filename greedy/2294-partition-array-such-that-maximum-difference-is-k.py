from typing import List
import unittest

# https://leetcode.com/problems/partition-array-such-that-maximum-difference-is-k/

# python3 -m unittest greedy/2294-partition-array-such-that-maximum-difference-is-k.py


class Solution(unittest.TestCase):
    # Approach: Sort + Greedy
    # Time: O(nlogn)
    # Space: O(n)
    def partitionArray(self, nums: List[int], k: int) -> int:
        nums.sort()
        paritions, mn = 1, nums[0]
        for mx in nums:
            if mx - mn > k:
                mn = mx
                paritions += 1
        return paritions

    def test(self):
        for nums, k, expected in [
            ([3, 6, 1, 2, 5], 2, 2),
            ([1, 2, 3], 1, 2),
            ([2, 2, 4, 5], 0, 3),
        ]:
            output = self.partitionArray(nums, k)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
