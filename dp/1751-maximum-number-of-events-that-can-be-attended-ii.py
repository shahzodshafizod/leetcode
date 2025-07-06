from bisect import bisect_right
from typing import List
import unittest

# https://leetcode.com/problems/maximum-number-of-events-that-can-be-attended-ii/

# python3 -m unittest dp/1751-maximum-number-of-events-that-can-be-attended-ii.py


class Solution(unittest.TestCase):
    # # Approach: Top-Down Dynamic Programming + Binary Search
    # # Time: O(n⋅logn + n⋅k)
    # # Space: O(n⋅k)
    # def maxValue(self, events: List[List[int]], k: int) -> int:
    #     events.sort()
    #     n = len(events)
    #     next = [n] * n
    #     for idx in range(n):
    #         next[idx] = bisect_right(events, [events[idx][1]+1], lo=idx+1, hi=n)
    #     memo = [[None] * k for _ in range(n)]
    #     def dp(idx: int, count: int) -> int:
    #         if idx == n or count == k:
    #             return 0
    #         if memo[idx][count] != None:
    #             return memo[idx][count]
    #         memo[idx][count] = max(
    #             # skip this event
    #             dp(idx+1, count),
    #             # attend this event
    #             events[idx][2] + dp(next[idx], count+1),
    #         )
    #         return memo[idx][count]
    #     return dp(0, 0)

    # Approach: Bottom-Up Dynamic Programming + Binary Search
    # Time: O(n⋅logn + n⋅k)
    # Space: O(n⋅k)
    def maxValue(self, events: List[List[int]], k: int) -> int:
        events.sort()  # O(nlogn)
        n = len(events)
        nxt = [n] * n
        for idx in range(n):  # O(nlogn)
            nxt[idx] = bisect_right(events, [events[idx][1] + 1], lo=idx + 1, hi=n)
        dp = [[0] * (k + 1) for _ in range(n + 1)]
        for idx in range(n - 1, -1, -1):  # O(n)
            for count in range(k - 1, -1, -1):  # O(k)
                dp[idx][count] = max(
                    # skip this event
                    dp[idx + 1][count],
                    # attend this event
                    events[idx][2] + dp[nxt[idx]][count + 1],
                )
        return dp[0][0]

    def test(self):
        for events, k, expected in [
            ([[1, 2, 4], [3, 4, 3], [2, 3, 1]], 2, 7),
            ([[1, 2, 4], [3, 4, 3], [2, 3, 10]], 2, 10),
            ([[1, 1, 1], [2, 2, 2], [3, 3, 3], [4, 4, 4]], 3, 9),
            ([[1, 3, 4], [2, 4, 1], [1, 1, 4], [3, 5, 1], [2, 5, 5]], 3, 9),
            # ([[1,100,1],[1,100,1],[1,100,1],[1,100,1],[1,100,1],[1,100,1],[1,100,1],[1,100,1],[1,100,1],[1,100,1],[1,100,1],[1,100,1],[1,100,1],[1,100,1],[1,100,1],[1,100,1],[1,100,1],[1,100,1],[1,100,1],[1,100,1],[1,100,1],[1,100,1],[1,100,1],[1,100,1],[1,100,1],[1,100,1],[1,100,1],[1,100,1],[1,100,1],[1,100,1],[1,100,1],[1,100,1],[1,100,1],[1,100,1],[1,100,1],[1,100,1],[1,100,1],[1,100,1],[1,100,1],[1,100,1],[1,100,1],[1,100,1],[1,100,1],[1,100,1],[1,100,1],[1,100,1],[1,100,1],[1,100,1],[1,100,1],[1,100,1],[1,100,1],[1,100,1],[1,100,1],[1,100,1],[1,100,1],[1,100,1],[1,100,1],[1,100,1],[1,100,1],[1,100,1],[1,100,1],[1,100,1],[1,100,1],[1,100,1],[1,100,1],[1,100,1],[1,100,1],[1,100,1],[1,100,1],[1,100,1],[1,100,1],[1,100,1],[1,100,1],[1,100,1],[1,100,1],[1,100,1],[1,100,1],[1,100,1],[1,100,1],[1,100,1],[1,100,1],[1,100,1],[1,100,1],[1,100,1],[1,100,1],[1,100,1],[1,100,1],[1,100,1],[1,100,1],[1,100,1],[1,100,1],[1,100,1],[1,100,1],[1,100,1],[1,100,1],[1,100,1],[1,100,1],[1,100,1],[1,100,1],[1,100,1]], 2, 1),
        ]:
            output = self.maxValue(events, k)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
