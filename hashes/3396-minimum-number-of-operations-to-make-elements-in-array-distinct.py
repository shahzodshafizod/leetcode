from typing import List
import unittest

# https://leetcode.com/problems/minimum-number-of-operations-to-make-elements-in-array-distinct/

# python3 -m unittest hashes/3396-minimum-number-of-operations-to-make-elements-in-array-distinct.py


class Solution(unittest.TestCase):
    # Approach: Reverse traversal
    # Time: O(n)
    # Space: O(n)
    def minimumOperations(self, nums: List[int]) -> int:
        visited = set()
        for idx in range(len(nums) - 1, -1, -1):
            if nums[idx] in visited:
                return idx // 3 + 1
            visited.add(nums[idx])
        return 0

    def test(self):
        for nums, expected in [
            ([1, 2, 3, 4, 2, 3, 3, 5, 7], 2),
            ([4, 5, 6, 4, 4], 2),
            ([6, 7, 8, 9], 0),
        ]:
            output = self.minimumOperations(nums)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
