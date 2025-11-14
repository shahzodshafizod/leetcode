from typing import List
import unittest

# https://leetcode.com/problems/tallest-billboard/

# python3 -m unittest dp/0956-tallest-billboard.py


class Solution(unittest.TestCase):
    # # Approach #1: Brute-Force (Top-Down Dynamic Programming)
    # # Time: O(3^(n/2))
    # # Space: O(3^(n/2))
    # def tallestBillboard(self, rods: List[int]) -> int:
    #     n, total = len(rods), sum(rods)
    #     memo = [[[-1] * (total + 1) for _ in range(total + 1)] for _ in range(n)]

    #     def dp(i: int, left: int, right: int) -> int:
    #         if i == n:
    #             return left if left == right else 0
    #         if memo[i][left][right] != -1:
    #             return memo[i][left][right]
    #         memo[i][left][right] = max(
    #             dp(i + 1, left, right),  # skip
    #             dp(i + 1, left + rods[i], right),  # include in left
    #             dp(i + 1, left, right + rods[i]),  # include in right
    #         )
    #         return memo[i][left][right]

    #     return dp(0, 0, 0)

    # # Approach #2: Top-Down Dynamic Programming
    # # Time: O(n * total)
    # # Space: O(n * total)
    # def tallestBillboard(self, rods: List[int]) -> int:
    #     INF = int(1e9)
    #     n, total = len(rods), sum(rods)
    #     memo = [[-1] * (total + 1) for _ in range(n)]

    #     def dp(i: int, height: int) -> int:
    #         if i == n:
    #             if height == 0:
    #                 return 0
    #             return -INF
    #         if 2 * height > total:
    #             return -INF
    #         if memo[i][height] != -1:
    #             return memo[i][height]
    #         memo[i][height] = max(
    #             dp(i + 1, height),  # skip
    #             dp(i + 1, abs(height - rods[i])) + min(rods[i], height),  # include in left
    #             dp(i + 1, height + rods[i]),  # include in right
    #         )
    #         return memo[i][height]

    #     return max(0, dp(0, 0))

    # Approach #3: Bottom-Up Dynamic Programming
    # Time: O(n * total)
    # Space: O(n * total)
    def tallestBillboard(self, rods: List[int]) -> int:
        n, total = len(rods), sum(rods)

        dp = [[-int(1e9)] * (total + 1) for _ in range(n + 1)]
        dp[0][0] = 0

        for i in range(n):
            for height in range(total + 1):
                if dp[i][height] < 0:
                    continue

                # Skip current rod
                dp[i + 1][height] = max(dp[i + 1][height], dp[i][height])

                # Include in left support
                new_height = height + rods[i]
                if new_height <= total:
                    dp[i + 1][new_height] = max(dp[i + 1][new_height], dp[i][height])

                # Include in right support
                new_height = abs(height - rods[i])
                if new_height <= total:
                    dp[i + 1][new_height] = max(dp[i + 1][new_height], dp[i][height] + min(rods[i], height))

        return max(0, dp[n][0])

    def test(self):
        for rods, expected in [
            ([1, 2, 3, 6], 6),
            ([1, 2, 3, 4, 5, 6], 10),
            ([1, 2], 0),
            ([96, 112, 101, 100, 104, 93, 106, 99, 114, 81, 94], 503),
            ([142, 178, 178, 143, 133, 139, 117, 153, 144, 162, 160, 147, 136, 149, 163, 160, 130, 157, 159], 1425),
            ([33, 30, 41, 34, 37, 29, 26, 31, 42, 33, 38, 27, 33, 31, 35, 900, 900, 900, 900, 900], 2050),
        ]:
            output = self.tallestBillboard(rods)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
