import unittest

# https://leetcode.com/problems/find-the-original-typed-string-ii/

# python3 -m unittest dp/3333-find-the-original-typed-string-ii.py

class Solution(unittest.TestCase):
    # Approach: Bottom-Up Dynamic Programming (Tabulation) + Prefix Sum
    # Time: O(n + m*k), n=len(word), m=len(groups)
    # Space: O(k)
    def possibleStringCount(self, word: str, k: int) -> int:
        MOD = 10**9 + 7
        groups = []
        count, total = 1, 1
        for idx in range(1, len(word)):
            if word[idx-1] != word[idx]:
                groups.append(count)
                total = (total * count) % MOD
                count = 0
            count += 1
        groups.append(count)
        total = (total * count) % MOD
        if len(groups) >= k:
            return total
        dp = [0] * k
        dp[0] = 1
        for count in groups:
            new_dp = [0] * k
            presum = 0
            for i in range(1, k):
                presum = (presum + dp[i-1]) % MOD
                if i > count:
                    presum = (presum - dp[i-count-1] + MOD) % MOD
                new_dp[i] = presum
            dp = new_dp
        return (total - sum(dp) + MOD) % MOD

    def test(self):
        for word, k, expected in [
            ("aabbccdd", 7, 5),
            ("aabbccdd", 8, 1),
            ("aaabbb", 3, 8),
        ]:
            output = self.possibleStringCount(word, k)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
