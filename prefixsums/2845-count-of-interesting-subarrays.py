from typing import List
import unittest

# https://leetcode.com/problems/count-of-interesting-subarrays/

# python3 -m unittest prefixsums/2845-count-of-interesting-subarrays.py

class Solution(unittest.TestCase):
    # Approach: Prefix Sum + Hash Table
    # Time: O(n)
    # Space: O(n)
    def countInterestingSubarrays(self, nums: List[int], modulo: int, k: int) -> int:
        total, precount = 0, 0
        freq = {0:1}
        for idx in range(len(nums)):
            precount += 1 if nums[idx] % modulo == k else 0
            total += freq.get((precount - k + modulo) % modulo, 0)
            freq[precount % modulo] = freq.get(precount % modulo, 0) + 1
        return total

    def test(self):
        for nums, module, k, expected in [
            ([3,2,4], 2, 1, 3),
            ([3,1,9,6], 3, 0, 2),
        ]:
            output = self.countInterestingSubarrays(nums, module, k)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
