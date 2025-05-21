from collections import defaultdict
from typing import List
import unittest

# https://leetcode.com/problems/number-of-ways-to-wear-different-hats-to-each-other/

# python3 -m unittest dp/1434-number-of-ways-to-wear-different-hats-to-each-other.py

class Solution(unittest.TestCase):
    # Approach: Top-Down Dynamic Programming + Bitmasks
    # Time: O(k⋅n⋅2^n), k=# of hats, n=# of people
    # Space: O(k⋅2^n)
    def numberWays(self, hats: List[List[int]]) -> int:
        MOD = int(1e9+7)
        people = defaultdict(list)
        for person, hts in enumerate(hats):
            for hat in hts:
                people[hat-1].append(person)
        n = len(hats)
        DONE = (1 << n) - 1
        MAX_HAT = 40
        memo = [[None] * DONE for _ in range(MAX_HAT)]
        def dp(hat: int, mask: int) -> int:
            if mask == DONE: return 1
            if hat == MAX_HAT: return 0
            if memo[hat][mask] != None:
                return memo[hat][mask]
            # 1. skip
            ways = dp(hat+1, mask)
            # 2. include
            for person in people[hat]:
                person = 1 << person
                if mask&person == 0:
                    ways += dp(hat+1, mask|person)
            memo[hat][mask] = ways % MOD
            return memo[hat][mask]
        return dp(0, 0)

    def test(self):
        for hats, expected in [
            ([[3,4],[4,5],[5]], 1),
            ([[3,5,1],[3,5]], 4),
            ([[1,2,3,4],[1,2,3,4],[1,2,3,4],[1,2,3,4]], 24),
        ]:
            output = self.numberWays(hats)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
