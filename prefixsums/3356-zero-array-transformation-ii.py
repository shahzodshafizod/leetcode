from typing import List
import unittest

# https://leetcode.com/problems/zero-array-transformation-ii/

# python3 -m unittest prefixsums/3356-zero-array-transformation-ii.py

class Solution(unittest.TestCase):
    # Approach: Line Sweep
    # Time: O(n+m), n=len(nums), m=len(queries)
    # Space: O(n)
    def minZeroArray(self, nums: List[int], queries: List[List[int]]) -> int:
        n = len(nums)
        line = [0] * (n+1)
        k, qlen = 0, len(queries)
        idx, presum = 0, 0
        while idx < n:
            if nums[idx] <= presum+line[idx]:
                presum += line[idx]
                idx += 1
            elif k < qlen:
                l, r, val = queries[k]
                if r >= idx:
                    if l <= idx:
                        presum += val
                    else:
                        line[l] += val
                    line[r+1] -= val
                k += 1
            else:
                break
        return k if idx == n else -1

    def test(self):
        for nums, queries, expected in [
            ([3,6,4], [[2,2,4]], -1),
            ([0,2,0,4], [[3,3,3],[0,0,5]], -1),
            ([5], [[0,0,5],[0,0,1],[0,0,3],[0,0,2]], 1),
            ([0], [[0,0,2],[0,0,4],[0,0,4],[0,0,3],[0,0,5]], 0),
            ([7,6,8], [[0,0,2],[0,1,5],[2,2,5],[0,2,4]], 4),
            ([1,0,6], [[1,2,1],[0,0,4],[1,1,5],[0,0,5],[1,2,4],[0,2,2],[2,2,4],[1,2,2],[1,2,4],[0,1,3]], 6),
        ]:
            output = self.minZeroArray(nums, queries)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
