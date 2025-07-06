from typing import List
import unittest

# https://leetcode.com/problems/solving-questions-with-brainpower/

# python3 -m unittest dp/2140-solving-questions-with-brainpower.py


class Solution(unittest.TestCase):
    # # Approach #1: Top-Down Dynamic Programming
    # # Time: O(n)
    # # Space: O(n)
    # def mostPoints(self, questions: List[List[int]]) -> int:
    #     n = len(questions)
    #     memo = [None] * n
    #     def dp(idx: int) -> int:
    #         if idx >= n:
    #             return 0
    #         if memo[idx] != None:
    #             return memo[idx]
    #         memo[idx] = max(
    #             dp(idx+1),
    #             questions[idx][0] + dp(idx+questions[idx][1]+1),
    #         )
    #         return memo[idx]
    #     return dp(0)

    # # Approach #2: Bottom-Up Dynamic Programming
    # # Time: O(n)
    # # Space: O(n)
    # def mostPoints(self, questions: List[List[int]]) -> int:
    #     dp = [q[0] for q in questions]
    #     n = len(questions)
    #     for idx in range(n-2,-1,-1):
    #         nextidx = idx + questions[idx][1] + 1
    #         if nextidx < n:
    #             dp[idx] += dp[nextidx]
    #         dp[idx] = max(dp[idx], dp[idx+1])
    #     return dp[0]

    # Approach #3: Bottom-Up Dynamic Programming (Space-Optimized)
    # Time: O(n)
    # Space: O(1)
    def mostPoints(self, questions: List[List[int]]) -> int:
        n = len(questions)
        for idx in range(n - 2, -1, -1):
            nextidx = idx + questions[idx][1] + 1
            if nextidx < n:
                questions[idx][0] += questions[nextidx][0]
            questions[idx][0] = max(questions[idx][0], questions[idx + 1][0])
        return questions[0][0]

    def test(self):
        for questions, expected in [
            ([[3, 2], [4, 3], [4, 4], [2, 5]], 5),
            ([[1, 1], [2, 2], [3, 3], [4, 4], [5, 5]], 7),
        ]:
            output = self.mostPoints(questions)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
