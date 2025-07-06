import unittest

# https://leetcode.com/problems/number-of-ways-to-stay-in-the-same-place-after-some-steps/

# python3 -m unittest dp/1269-number-of-ways-to-stay-in-the-same-place-after-some-steps.py


class Solution(unittest.TestCase):
    # # Approach #1: Top-Down Dynamic Programming
    # # Time: O(S⋅min(S,A)), S=steps, A=arrLen
    # # Space: O(S⋅min(S,A))
    # def numWays(self, steps: int, arrLen: int) -> int:
    #     MOD = int(1e9+7)
    #     arrLen = min(arrLen, steps) # cannot move further than "steps" steps
    #     memo = [[None] * (steps+1) for _ in range(arrLen)]
    #     def dp(pos: int, steps: int) -> int:
    #         if pos < 0 or pos == arrLen: return 0
    #         if steps == 0:
    #             return 1 if pos == 0 else 0
    #         if memo[pos][steps] != None:
    #             return memo[pos][steps]
    #         memo[pos][steps] = (
    #             dp(pos-1, steps-1)
    #             + dp(pos, steps-1)
    #             + dp(pos+1, steps-1)
    #         ) % MOD
    #         return memo[pos][steps]
    #     return dp(0, steps)

    # # Approach #2: Bottom-Up Dynamic Programming
    # # Time: O(S⋅min(S,A)), S=steps, A=arrLen
    # # Space: O(S⋅min(S,A))
    # def numWays(self, steps: int, arrLen: int) -> int:
    #     MOD = int(1e9+7)
    #     arrLen = min(arrLen, steps) # cannot move further than "steps" steps
    #     dp = [[0] * (steps+1) for _ in range(arrLen)]
    #     dp[0][0] = 1
    #     for st in range(1, steps+1):
    #         for pos in range(arrLen-1,-1,-1):
    #             dp[pos][st] = dp[pos][st-1]
    #             if pos > 0:
    #                 dp[pos][st] = (dp[pos][st] + dp[pos-1][st-1]) % MOD
    #             if pos+1 < arrLen:
    #                 dp[pos][st] = (dp[pos][st] + dp[pos+1][st-1]) % MOD
    #     return dp[0][steps]

    # Approach #3: Bottom-Up Dynamic Programming (Space-Optimized)
    # Time: O(S⋅min(S,A)), S=steps, A=arrLen
    # Space: O(min(S,A))
    def numWays(self, steps: int, arrLen: int) -> int:
        MOD = int(1e9 + 7)
        arrLen = min(arrLen, steps)  # cannot move further than "steps" steps
        prev, curr = [0] * arrLen, [0] * arrLen
        curr[0] = 1
        for _ in range(steps):
            prev, curr = curr, prev
            for pos in range(arrLen - 1, -1, -1):
                curr[pos] = prev[pos]
                if pos > 0:
                    curr[pos] = (curr[pos] + prev[pos - 1]) % MOD
                if pos + 1 < arrLen:
                    curr[pos] = (curr[pos] + prev[pos + 1]) % MOD
        return curr[0]

    def test(self):
        for steps, arrLen, expected in [
            (3, 2, 4),
            (2, 4, 2),
            (4, 2, 8),
            # (500, 1000000, 374847123),
        ]:
            output = self.numWays(steps, arrLen)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
