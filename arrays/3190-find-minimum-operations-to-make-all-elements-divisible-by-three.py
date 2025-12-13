from typing import List
import unittest

# https://leetcode.com/problems/find-minimum-operations-to-make-all-elements-divisible-by-three/

# python3 -m unittest arrays/3190-find-minimum-operations-to-make-all-elements-divisible-by-three.py


class Solution(unittest.TestCase):
    # Approach: Count Remainders
    # Time: O(n)
    # Space: O(1)
    def minimumOperations(self, nums: List[int]) -> int:
        return sum(1 for num in nums if num % 3 != 0)

    def test(self):
        for nums, expected in [
            ([1, 2, 3, 4], 3),  # 1->0, 2->3, 4->3: 3 operations
            ([3, 6, 9], 0),  # all divisible by 3
            ([1, 2, 4, 5], 4),  # all need 1 operation
            ([0], 0),  # 0 is divisible by 3
            ([5], 1),  # 5->6 or 5->3: 1 operation
        ]:
            output = self.minimumOperations(nums)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
