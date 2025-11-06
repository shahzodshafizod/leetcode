import unittest

# https://leetcode.com/problems/count-k-reducible-numbers-less-than-n/

# python3 -m unittest dp/3352-count-k-reducible-numbers-less-than-n.py


class Solution(unittest.TestCase):
    # # Approach 1: Brute Force (Only works for small values of n)
    # # Time Complexity: O(n * k * log(n)) - iterates n numbers, each takes k reductions
    # # Space Complexity: O(1)
    # def countKReducibleNumbers(self, s: str, k: int) -> int:
    #     def is_k_reducible(num: int, k: int) -> bool:
    #         steps = 0
    #         while num > 1 and steps < k:
    #             num = num.bit_count()
    #             steps += 1

    #         return num == 1

    #     n = int(s, 2)  # Convert binary string to integer
    #     count = 0
    #     for i in range(1, n):
    #         if is_k_reducible(i, k):
    #             count += 1
    #     return count % (10**9 + 7)

    # Approach 2: Top-Down DP (Memoization) with Digit DP
    # Time Complexity: O(2nn) where n is length of string: len(s)
    # Space Complexity: O(2nn) for memoization
    def countKReducibleNumbers(self, s: str, k: int) -> int:
        MOD = 10**9 + 7
        n = len(s)

        # Step 1: Precompute steps needed to reduce each bit count to 1
        steps = [0] * (n + 1)
        steps[1] = 0  # no steps needed to reduce 1 to 1 (1 is already 1)
        for i in range(2, n + 1):
            steps[i] = steps[i.bit_count()] + 1

        # Step 2: Use digit DP to count numbers with specific bit counts
        memo = [[[-1] * 2 for _ in range(n)] for _ in range(n)]

        def dp(pos: int, ones: int, bound: int) -> int:
            if pos == n:
                return 1 if not bound and ones > 0 and steps[ones] < k else 0

            if memo[pos][ones][bound] != -1:
                return memo[pos][ones][bound]

            # Determine the maximum digit we can place
            limit = int(s[pos]) if bound else 1

            result = 0
            for digit in range(limit + 1):
                new_bound = bound and (digit == limit)
                result = (result + dp(pos + 1, ones + digit, new_bound)) % MOD

            memo[pos][ones][bound] = result
            return result

        return dp(0, 0, 1)

    # # Approach 3: Bottom-Up DP (Tabulation)
    # # Time Complexity: O(2n²) where n is length of string
    # # Space Complexity: O(2n²)
    # def countKReducibleNumbers(self, s: str, k: int) -> int:
    #     MOD = 10**9 + 7
    #     n = len(s)

    #     steps = [0] * (n + 1)
    #     for i in range(2, n + 1):
    #         steps[i] = steps[i.bit_count()] + 1

    #     dp = [[[0] * 2 for _ in range(n + 1)] for _ in range(n + 1)]

    #     # Base case: at position n (end of string)
    #     for ones in range(n + 1):
    #         if ones > 0 and steps[ones] < k:
    #             dp[n][ones][0] = 1

    #     # Fill the DP table from position n-1 down to 0
    #     for pos in range(n - 1, -1, -1):
    #         for ones in range(n + 1):
    #             for bound in range(2):
    #                 # Determine the maximum digit we can place
    #                 limit = int(s[pos]) if bound else 1

    #                 result = 0
    #                 for digit in range(limit + 1):
    #                     new_ones = ones + digit
    #                     new_bound = bound and (digit == limit)
    #                     if new_ones <= n:  # Ensure we don't go out of bounds
    #                         result = (result + dp[pos + 1][new_ones][new_bound]) % MOD

    #                 dp[pos][ones][bound] = result

    #     return dp[0][0][1]

    def test(self):
        for s, k, expected in [
            ("111", 1, 3),
            ("1000", 2, 6),
            ("1", 3, 0),
            ("101", 2, 4),
        ]:
            output = self.countKReducibleNumbers(s, k)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
