import unittest

# https://leetcode.com/problems/number-of-people-aware-of-a-secret/

# python3 -m unittest dp/2327-number-of-people-aware-of-a-secret.py


class Solution(unittest.TestCase):
    # # Approach #1: Top-Down Dynamic Programming
    # # Time: O(nn)
    # # Space: O(n)
    # def peopleAwareOfSecret(self, n: int, delay: int, forget: int) -> int:
    #     MOD = 10**9 + 7
    #     memo = [-1] * (n + 1)

    #     def dp(day: int) -> int:
    #         if memo[day] != -1:
    #             return memo[day]
    #         # if this person contributes to the last day
    #         people = 1 if day + forget - 1 >= n else 0
    #         for nxt in range(delay, forget):
    #             if day + nxt <= n:
    #                 people = (people + dp(day + nxt)) % MOD
    #         memo[day] = people
    #         return people

    #     return dp(1)

    # Approach #2: Bottom-Up Dynamic Programming
    # Time: O(n)
    # Space: O(n)
    def peopleAwareOfSecret(self, n: int, delay: int, forget: int) -> int:
        MOD = 10**9 + 7
        dp = [0] * (n + 1)
        dp[1] = 1
        people = 0
        for day in range(2, n + 1):
            if day - delay > 0:
                people = (people + dp[day - delay]) % MOD
            if day - forget > 0:
                people = (people - dp[day - forget] + MOD) % MOD
            dp[day] = people
        return sum(dp[n - forget + 1 :]) % MOD

    def test(self):
        for n, delay, forget, expected in [
            (6, 2, 4, 5),
            (4, 1, 3, 6),
        ]:
            output = self.peopleAwareOfSecret(n, delay, forget)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
