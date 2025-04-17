from collections import defaultdict
from typing import List
import unittest

# https://leetcode.com/problems/count-equal-and-divisible-pairs-in-an-array/

# python3 -m unittest hashes/2176-count-equal-and-divisible-pairs-in-an-array.py

class Solution(unittest.TestCase):
    # Approach: Hash Table
    # Time: best: O(N); worst: O(NN)
    # Space: O(N)
    def countPairs(self, nums: List[int], k: int) -> int:
        pairs = 0
        indices = defaultdict(list)
        for curr, num in enumerate(nums):
            for prev in indices[num]:
                if prev*curr % k == 0:
                    pairs += 1
            indices[num].append(curr)
        return pairs

    def test(self):
        for nums, k, expected in [
            ([3,1,2,2,2,1,3], 2, 4),
            ([1,2,3,4], 1, 0),
        ]:
            output = self.countPairs(nums, k)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
