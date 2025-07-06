import itertools
import unittest

# https://leetcode.com/problems/strange-printer/

# python3 -m unittest dp/0664-strange-printer.py


class Solution(unittest.TestCase):
    # # Bottom Up Dynamic Programming (Tabulation)
    # # Time: O(N^3)
    # # Space: O(N^2)
    def strangePrinter(self, s: str) -> int:
        s = "".join(c for c, _ in itertools.groupby(s))
        n = len(s)
        indexes = {}
        nxt = [n] * n
        for idx in range(n - 1, -1, -1):
            if s[idx] in indexes:
                nxt[idx] = indexes[s[idx]]
            indexes[s[idx]] = idx
        dp = [[0] * n for _ in range(n + 1)]
        for idx in range(n):
            dp[idx][idx] = 1
        for start in range(n - 2, -1, -1):
            for end in range(start + 1, n):
                dp[start][end] = 1 + dp[start + 1][end]
                k = nxt[start]
                while k <= end:
                    dp[start][end] = min(dp[start][end], dp[start][k - 1] + dp[k + 1][end])
                    k = nxt[k]
        return dp[0][n - 1]

    # # Top Down Dynamic Programming (Memoization)
    # # Time: O(N^3)
    # # Space: O(N^2)
    # def strangePrinter(self, s: str) -> int:
    #     # remove consecutive duplicate characters
    #     s = "".join(c for c,_ in itertools.groupby(s))
    #     memo = {}
    #     def minTurns(start: int, end: int) -> int:
    #         if start > end:
    #             return 0
    #         if (start, end) in memo:
    #             return memo[(start, end)]
    #         turns = 1 + minTurns(start+1, end)
    #         for k in range(start+1, end+1):
    #             if s[k] == s[start]:
    #                 turns = min(turns, minTurns(start, k-1) + minTurns(k+1, end))
    #         memo[(start, end)] = turns
    #         return turns
    #     return minTurns(0, len(s)-1)

    def test(self):
        for s, expected in [
            ("aaabbb", 2),
            ("aba", 2),
            ("abcab", 4),
            ("abcabc", 5),
            ("zzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzz", 1),
            ("abcabccb", 5),
            ("abcdbca", 5),
            ("abbbccbbbbccccbbbbbccccc", 5),
            ("cabad", 4),
            ("ababa", 3),
        ]:
            output = self.strangePrinter(s)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
