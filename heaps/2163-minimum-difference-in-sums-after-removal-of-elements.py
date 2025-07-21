import heapq
from typing import List
import unittest

# https://leetcode.com/problems/minimum-difference-in-sums-after-removal-of-elements/

# python3 -m unittest heaps/2163-minimum-difference-in-sums-after-removal-of-elements.py


class Solution(unittest.TestCase):
    # Time: O(n log n)
    # Space: O(n)
    def minimumDifference(self, nums: List[int]) -> int:
        n = len(nums) // 3
        sumsec = [0] * (n + 1)
        sumsec[n] = sum(nums[2 * n :])
        mnhp = nums[2 * n :][:]
        heapq.heapify(mnhp)
        for i in range(n - 1, -1, -1):
            heapq.heappush(mnhp, nums[n + i])
            sumsec[i] = sumsec[i + 1] + nums[n + i] - heapq.heappop(mnhp)
        sumfir = sum(nums[:n])
        mxhp = [-x for x in nums[:n]]
        heapq.heapify(mxhp)
        diff = sumfir - sumsec[0]
        for i in range(n):
            heapq.heappush(mxhp, -nums[n + i])
            sumfir += nums[n + i] + heapq.heappop(mxhp)
            diff = min(diff, sumfir - sumsec[i + 1])
        return diff

    def test(self):
        for nums, expected in [
            ([3, 1, 2], -1),
            ([7, 9, 5, 8, 1, 3], 1),
            ([7, 9, 5, 8, 1, 3, 1, 1, 10000], -9993),
            ([16, 46, 43, 41, 42, 14, 36, 49, 50, 28, 38, 25, 17, 5, 18, 11, 14, 21, 23, 39, 23], -14),
        ]:
            output = self.minimumDifference(nums)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
