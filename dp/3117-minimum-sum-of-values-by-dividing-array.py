from typing import List, Dict, Tuple
import unittest

# https://leetcode.com/problems/minimum-sum-of-values-by-dividing-array/

# python3 -m unittest dp/3117-minimum-sum-of-values-by-dividing-array.py


class Solution(unittest.TestCase):
    # Approach: Dynamic Programming with memoization
    # Divide array into segments where each segment's AND equals corresponding value in andValues
    # Minimize sum of last elements of each segment
    # Time: O(n * m * maxAnd) where maxAnd is number of unique AND values
    # Space: O(n * m * maxAnd) for memoization
    def minimumValueSum(self, nums: List[int], andValues: List[int]) -> int:
        n = len(nums)
        m = len(andValues)
        MAX = 1 << 31

        memo: Dict[Tuple[int, int, int], int] = {}

        def dp(i: int, j: int, current_and: int) -> int:
            # i: current position in nums
            # j: current position in andValues
            # current_and: current AND value of segment

            if i == n and j == m:
                return 0
            if i == n or j == m:
                return MAX

            state = (i, j, current_and)
            if state in memo:
                return memo[state]

            # Update AND with current element
            new_and = current_and & nums[i]

            result = MAX

            # Option 1: Continue current segment
            result = min(result, dp(i + 1, j, new_and))

            # Option 2: End current segment if AND matches
            if new_and == andValues[j]:
                result = min(result, nums[i] + dp(i + 1, j + 1, (1 << 20) - 1))

            memo[state] = result
            return result

        # Start with all bits set (AND identity)
        ans = dp(0, 0, (1 << 20) - 1)
        return ans if ans != MAX else -1

    def test(self):
        for nums, andValues, expected in [
            ([1, 4, 3, 3, 2], [0, 3, 3, 2], 12),
            ([2, 3, 5, 7, 7, 7, 5], [0, 7, 5], 17),
            ([1, 2, 3, 4], [2], -1),
        ]:
            output = self.minimumValueSum(nums, andValues)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
