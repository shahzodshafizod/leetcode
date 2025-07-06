from typing import List
import unittest

# https://leetcode.com/problems/number-of-sub-arrays-with-odd-sum/

# python3 -m unittest prefixsums/1524-number-of-sub-arrays-with-odd-sum.py


class Solution(unittest.TestCase):
    # Approach: Prefix Sum
    # Time: O(n)
    # Space: O(1)
    def numOfSubarrays(self, arr: List[int]) -> int:
        odd_cnt, evn_cnt = 0, 0
        ttl_cnt, presum = 0, 0
        MOD = 10**9 + 7
        for num in arr:
            presum += num
            if presum & 1:
                ttl_cnt = (ttl_cnt + evn_cnt + 1) % MOD
                odd_cnt += 1
            else:
                ttl_cnt = (ttl_cnt + odd_cnt) % MOD
                evn_cnt += 1
        return ttl_cnt

    def test(self):
        for arr, expected in [
            ([1, 3, 5], 4),
            ([2, 4, 6], 0),
            ([1, 2, 3, 4, 5, 6, 7], 16),
        ]:
            output = self.numOfSubarrays(arr)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
