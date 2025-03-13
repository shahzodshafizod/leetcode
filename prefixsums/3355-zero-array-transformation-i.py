from typing import List
import unittest

# https://leetcode.com/problems/zero-array-transformation-i/

# python3 -m unittest prefixsums/3355-zero-array-transformation-i.py

class Solution(unittest.TestCase):
    # Approach: Line Sweep
    # Time: O(n+m), n=len(nums), m=len(queries)
    # Space: O(n)
    def isZeroArray(self, nums: List[int], queries: List[List[int]]) -> bool:
        n = len(nums)
        line = [0] * (n+1)
        for l, r in queries:
            line[l] += 1
            line[r+1] -= 1
        presum = 0
        for idx in range(n):
            presum += line[idx]
            if nums[idx] > presum:
                return False
        return True

    def test(self):
        for nums, queries, expected in [
            ([1,0,1], [[0,2]], True),
            ([4,3,2,1], [[1,3],[0,2]], False),
            ([2], [[0,0],[0,0],[0,0]], True),
        ]:
            output = self.isZeroArray(nums, queries)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
