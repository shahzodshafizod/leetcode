from math import ceil
import unittest

# https://leetcode.com/problems/soup-servings/

# python3 -m unittest dp/0808-soup-servings.py

# On avarage, we serve more of soup A than soup B:
# 	A: (100 + 75 + 50 + 25) / 4 = 62.5
# 	B: (  0 + 25 + 50 + 75) / 4 = 37.5
# As n grows, the probability that A empties first approaches 1.
# For n > 4800, the probability is so close to 1 that we can just return 1.0.


class Solution(unittest.TestCase):
    # Approach: Top-Down Dynamic Programming
    # Time: O(192 * 192), 192 = 4800 / 25
    # Space: O(192 * 192), 192 = 4800 / 25
    def soupServings(self, n: int) -> float:
        if n > 4800:
            return 1
        n = ceil(n / 25)
        memo = [[float(-1)] * (n + 1) for _ in range(n + 1)]

        def dp(a: int, b: int) -> float:
            if a <= 0:
                if b <= 0:
                    return 0.5
                return 1
            if b <= 0:
                return 0
            if memo[a][b] != -1:
                return memo[a][b]
            memo[a][b] = 0.25 * (dp(a - 4, b) + dp(a - 3, b - 1) + dp(a - 2, b - 2) + dp(a - 1, b - 3))
            return memo[a][b]

        return dp(n, n)

    def test(self):
        for n, expected in [
            (50, 0.62500),
            (100, 0.71875),
        ]:
            output = self.soupServings(n)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
