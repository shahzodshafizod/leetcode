import bisect
from typing import List, Dict
import unittest

# https://leetcode.com/problems/make-array-strictly-increasing/

# python3 -m unittest dp/1187-make-array-strictly-increasing.py


class Solution(unittest.TestCase):
    # # Approach #1: Top-Down Dynamic Programming
    # # Time: O(n * m * log(m)) where n = len(arr1), m = len(arr2)
    # # Space: O(n * m)
    # def makeArrayIncreasing(self, arr1: List[int], arr2: List[int]) -> int:
    #     arr2 = sorted(list(set(arr2)))
    #     n, m = len(arr1), len(arr2)
    #     memo: Dict[Tuple[int, int], int] = {}

    #     def dp(idx: int, prev: int) -> int:
    #         if idx == n:
    #             return 0
    #         state = (idx, prev)
    #         if state in memo:
    #             return memo[state]

    #         cost = n + 1
    #         # keep
    #         if prev < arr1[idx]:
    #             cost = min(cost, dp(idx + 1, arr1[idx]))
    #         # replace
    #         j = bisect.bisect_right(arr2, prev)
    #         if j < m:
    #             cost = min(cost, 1 + dp(idx + 1, arr2[j]))

    #         memo[state] = cost
    #         return cost

    #     ops = dp(0, -1)
    #     return ops if ops <= n else -1

    # Approach #2: Bottom-Up Dynamic Programming
    # Time: O(n * m * log(m)) where n = len(arr1), m = len(arr2)
    # Space: O(m)
    def makeArrayIncreasing(self, arr1: List[int], arr2: List[int]) -> int:
        arr2 = sorted(set(arr2))
        n, m = len(arr1), len(arr2)
        dp: Dict[int, int] = {-1: 0}
        for num in arr1:
            new_dp: Dict[int, int] = {}
            for prev, ops in dp.items():
                # keep
                if num > prev:
                    new_dp[num] = min(new_dp.get(num, n + 1), ops)
                # replace
                idx = bisect.bisect_right(arr2, prev)
                if idx < m:
                    new_dp[arr2[idx]] = min(new_dp.get(arr2[idx], n + 1), ops + 1)
            dp = new_dp
            if not dp:  # No valid state found
                return -1
        return min(dp.values()) if dp else -1

    def test(self):
        for arr1, arr2, expected in [
            ([1, 5, 3, 6, 7], [1, 3, 2, 4], 1),
            ([1, 5, 3, 6, 7], [4, 3, 1], 2),
            ([1, 5, 3, 6, 7], [1, 6, 3, 3], -1),
            ([31, 18, 1, 12, 23, 14, 25, 4, 17, 18, 29, 28, 35, 34, 19, 8, 25, 6, 35], [13, 10, 25, 18, 3, 8, 37, 20, 23, 12, 9, 36, 17, 22, 29, 6, 1, 12, 37, 6, 3, 14, 37, 2], 19),
            ([0, 11, 6, 1, 4, 3], [5, 4, 11, 10, 1, 0], 5),
        ]:
            output = self.makeArrayIncreasing(arr1, arr2)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
