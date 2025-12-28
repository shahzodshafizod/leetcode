import unittest

# https://leetcode.com/problems/stone-game-iv/

# python3 -m unittest dp/1510-stone-game-iv.py


class Solution(unittest.TestCase):
    # # Approach: Top-Down Dynamic Programming
    # # Time: O(n*sqrt(n))
    # # Space: O(n*sqrt(n))
    # def winnerSquareGame(self, n: int) -> bool:
    #     cache = [None] * (n+1)
    #     def dp(curr: int) -> bool:
    #         if not curr: return False
    #         if cache[curr] != None: return cache[curr]
    #         for x in range(int(math.sqrt(curr)),0,-1):
    #             # if next player loses, the current one wins
    #             if not dp(curr - x*x):
    #                 cache[curr] = True
    #                 return True
    #         cache[curr] = False
    #         return False
    #     return dp(n)

    # Approach: Bottom-Up Dynamic Programming
    # Time: O(n*sqrt(n))
    # Space: O(n)
    def winnerSquareGame(self, n: int) -> bool:
        dp = [False] * (n + 1)
        dp[0] = False
        for curr in range(1, n + 1):
            x = 1
            while x * x <= curr:
                if not dp[curr - x * x]:
                    dp[curr] = True
                    break
                x += 1
        return dp[n]

    def test(self):
        for n, expected in [
            (1, True),
            (2, False),
            (4, True),
        ]:
            output = self.winnerSquareGame(n)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
