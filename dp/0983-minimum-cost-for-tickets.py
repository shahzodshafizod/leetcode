from typing import List
import unittest

# https://leetcode.com/problems/minimum-cost-for-tickets/

# python3 -m unittest dp/0983-minimum-cost-for-tickets.py

class Solution(unittest.TestCase):
    def mincostTickets(self, days: List[int], costs: List[int]) -> int:
        n = len(days)
        memo = [None] * n
        def dp(idx: int) -> int:
            if idx == n: return 0
            key = idx
            if memo[key] != None:
                return memo[key]
            idx += 1
            day = 1
            cost = float("inf")
            for c, limit in enumerate([1,7,30]):
                while idx < n and day + (days[idx] - days[idx-1]) <= limit:
                    day += days[idx] - days[idx-1]
                    idx += 1
                cost = min(cost, costs[c] + dp(idx))
            memo[key] = cost
            return cost
        return dp(0)

    def test(self):
        for days, costs, expected in [
            ([1,4,6,7,8,20], [2,7,15], 11),
            ([1,2,3,4,5,6,7,8,9,10,30,31], [2,7,15], 17),
        ]:
            output = self.mincostTickets(days, costs)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
