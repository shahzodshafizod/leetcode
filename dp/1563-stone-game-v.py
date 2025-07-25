from typing import List
import unittest

# https://leetcode.com/problems/stone-game-v/

# python3 -m unittest dp/1563-stone-game-v.py

# TODO: Solve it in O(nnlogn) time complexity # pylint: disable=fixme


class Solution(unittest.TestCase):
    # # Approach: Top-Down Dynamic Programming (Time Limit Exceeded)
    # # Time: O(nnn)
    # # Space: O(nn)
    # def stoneGameV(self, stoneValue: List[int]) -> int:
    #     n = len(stoneValue)
    #     presum = [0] * (n+1)
    #     for idx in range(n):
    #         presum[idx+1] = stoneValue[idx] + presum[idx]
    #     memo = [[None] * n for _ in range(n)]
    #     def dp(left: int, right: int) -> int:
    #         if left == right: return 0
    #         if memo[left][right] != None:
    #             return memo[left][right]
    #         result, lsum = 0, 0
    #         for idx in range(left, right):
    #             lsum += stoneValue[idx]
    #             rsum = presum[right+1] - presum[idx+1]
    #             if lsum <= rsum:
    #                 result = max(result, lsum + dp(left, idx))
    #             if rsum <= lsum:
    #                 result = max(result, rsum + dp(idx+1, right))
    #         memo[left][right] = result
    #         return result
    #     return dp(0, n-1)

    # Approach: Bottom-Up Dynamic Programming (Time Limit Exceeded)
    # Time: O(nnn)
    # Space: O(nn)
    def stoneGameV(self, stoneValue: List[int]) -> int:
        n = len(stoneValue)
        presum = [0] * (n + 1)
        for idx in range(n):
            presum[idx + 1] = stoneValue[idx] + presum[idx]
        dp = [[0] * n for _ in range(n)]
        for length in range(2, n + 1):
            for left in range(n - length + 1):
                right = left + length - 1
                lsum = 0
                for idx in range(left, right):
                    lsum += stoneValue[idx]
                    rsum = presum[right + 1] - presum[idx + 1]
                    if lsum <= rsum:
                        dp[left][right] = max(dp[left][right], lsum + dp[left][idx])
                    if rsum <= lsum:
                        dp[left][right] = max(dp[left][right], rsum + dp[idx + 1][right])
        return dp[0][n - 1]

    def test(self):
        for stoneValue, expected in [
            ([6, 2, 3, 4, 5, 5], 18),
            ([7, 7, 7, 7, 7, 7, 7], 28),
            ([4], 0),
            ([1, 1, 2], 3),
            ([68, 75, 25, 50, 34, 29, 77, 1, 2, 69], 304),
            ([10, 9, 8, 7, 6, 5, 4, 3, 2, 1], 37),
            # ([4,9000,6,2,5,5,6,8,3,7,3,4,5,2,1,5,1,6,10,10,3,3,9,3,8,5,5,1,6,6,1,3,7,3,7,8,1,9,5,2,3,9,2,1,4,10,2,10,4,5,6,1,8,5,10,10,9,1,2,5,1,1,10,2,6,8,3,5,8,3,9,3,4,8,1,6,8,5,3,7,3,5,1,10,10,4,6,6,8,5,7,4,1,5,10,2,6,7,5,7,8,4,9,5,2,9,3,7,9,10,1,1,4,3,5,8,9,2,6,3,9,8,9,4,4,9,10,3,7,5,3,4,2,9,7,2,3,1,1,4,10,5,1,2,3,2,7,7,1,5,6,2,4,9,6,10,9,7,8,9,3,3,7,7,3,2,10,9,3,4,6,10,10,2,8,6,10,10,6,1,10,5,1,9,3,4,3,7,5,6,9,2,6,2,4,9,1,9,1,4,3,10,3,6,10,6,10,6,3,7,7,2,5,6,9,10,7,6,7,3,3,5,2,9,5,4,10,6,1,9,3,6,3,10,2,6,3,4,1,10,1,4,9,5,10,2,2,4,8,3,3,8,10,2,6,3,8,9,6,6,7,3,7,3,2,1,3,4,3,9,10,7,4,6,7,8,3,3,5,9,8,2,10,4,6,7,2,10,10,2,5,1,7,2,9,9,5,1,10,5,1,1,1,7,8,10,3,1,6,3,7,9,1,10,5,5,2,5,10,8,10,6,6,8,3,6,4,3,6,7,8,1,3,2,1,4,7,7,8,1,1,4,3,3,7,7,7,6,8,8,1,10,6,4,4,9,9,9,2,3,9,2,10,2,2,6,9,9,1,7,8,1,2,7,8,8,10,10,4,10,8,4,1,6,4,3,8,6,1,7,3,2,7,4,3,6,4,3,10,6,10,10,7,5,10,1,6,8,6,6,3,7,8,7,5,6,5,3,1,4,4,8,8,10,7,4,10,10,8,5,9,6,2,7,10,5,7,1,3,5,3,5,7,5,2,10,3,10,4,6,5,6,2,2,4,1,7,1,1,9,7,8,7,5,4,7,4,8,8,1,2,10,6,6,1,6,6,5,8,5,3,5,9,10,6,9,4,10,10,5,4,1,6000,12500,25000,50000,100000,200000,400000,990000], 1205175),
        ]:
            output = self.stoneGameV(stoneValue)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
