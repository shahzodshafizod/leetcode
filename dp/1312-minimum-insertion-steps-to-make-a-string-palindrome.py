import unittest

# https://leetcode.com/problems/minimum-insertion-steps-to-make-a-string-palindrome/

# python3 -m unittest dp/1312-minimum-insertion-steps-to-make-a-string-palindrome.py


class Solution(unittest.TestCase):
    # Approach #1 (faster): Top-Down Dynamic Programming
    # Time: O(nn)
    # Space: O(nn)
    def minInsertions(self, s: str) -> int:
        n = len(s)
        memo = [[-1] * n for _ in range(n)]

        def dp(left: int, right: int) -> int:
            if left >= right:
                return 0
            if memo[left][right] != -1:
                return memo[left][right]
            if s[left] != s[right]:
                memo[left][right] = 1 + min(dp(left + 1, right), dp(left, right - 1))
            else:
                memo[left][right] = dp(left + 1, right - 1)
            return memo[left][right]

        return dp(0, len(s) - 1)

    # # Approach #2: Bottom-Up Dynamic Programming: Longest Common Sequence
    # # Time: O(nn)
    # # Space: O(nn)
    # def minInsertions(self, s: str) -> int:
    #     n = len(s)
    #     dp = [[0] * (n + 1) for _ in range(n + 1)]
    #     for i in range(n):
    #         for j in range(n):
    #             dp[i + 1][j + 1] = dp[i][j] + 1 if s[i] == s[n - j - 1] else max(dp[i][j + 1], dp[i + 1][j])
    #     return n - dp[n][n]

    def test(self):
        for s, expected in [
            ("zzazz", 0),
            ("mbadm", 2),
            ("leetcode", 5),
        ]:
            output = self.minInsertions(s)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
