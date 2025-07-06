from typing import List
import unittest

# https://leetcode.com/problems/minimum-operations-to-make-array-values-equal-to-k/

# python3 -m unittest hashes/3375-minimum-operations-to-make-array-values-equal-to-k.py


class Solution(unittest.TestCase):
    # Approach: Hash map
    # Time: O(n)
    # Space: O(n)
    def minOperations(self, nums: List[int], k: int) -> int:
        visited = set()
        for num in nums:
            if num < k:
                return -1
            elif num > k:
                visited.add(num)
        return len(visited)

    def test(self):
        for nums, k, expected in [
            ([5, 2, 5, 4, 5], 2, 2),
            ([2, 1, 2], 2, -1),
            ([9, 7, 5, 3], 1, 4),
        ]:
            output = self.minOperations(nums, k)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
