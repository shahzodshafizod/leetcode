from typing import List
import unittest

# https://leetcode.com/problems/maximum-value-of-k-coins-from-piles/

# python3 -m unittest dp/2218-maximum-value-of-k-coins-from-piles.py

class Solution(unittest.TestCase):
    # # Approach: Top-Down Dynamic Programming
    # # Time: O(nkp), n=len(piles), p=len(piles[i])
    # # Space: O(nk)
    # def maxValueOfCoins(self, piles: List[List[int]], k: int) -> int:
    #     n = len(piles)
    #     memo = [[None] * (k+1) for _ in range(n)]
    #     def dp(i: int, coins: int) -> int:
    #         if i == n: return 0
    #         if memo[i][coins] != None:
    #             return memo[i][coins]
    #         # skip curr pile
    #         count = dp(i+1, coins)
    #         # take from curr pile
    #         total = 0
    #         for j in range(min(coins, len(piles[i]))): # O(len(piles[i]))
    #             total += piles[i][j]
    #             count = max(count, total + dp(i+1, coins-j-1))
    #         memo[i][coins] = count
    #         return count
    #     return dp(0, k)

    # # Approach: Bottom-Up Dynamic Programming
    # # Time: O(nkp), n=len(piles), p=len(piles[i])
    # # Space: O(nk)
    # def maxValueOfCoins(self, piles: List[List[int]], k: int) -> int:
    #     n = len(piles)
    #     memo = [[0] * (k+1) for _ in range(n+1)]
    #     memo[n][1] = 1
    #     for i in range(n-1,-1,-1):
    #         for coins in range(1, k+1):
    #             # skip curr pile
    #             memo[i][coins] = memo[i+1][coins]
    #             # take from curr pile
    #             total = 0
    #             for j in range(min(coins, len(piles[i]))):
    #                 total += piles[i][j]
    #                 memo[i][coins] = max(
    #                     memo[i][coins],
    #                     total + memo[i+1][coins-j-1]
    #                 )
    #     return memo[0][k]

    # Approach: Bottom-Up Dynamic Programming (Space-Optimized)
    # Time: O(nkp), n=len(piles), p=len(piles[i])
    # Space: O(k)
    def maxValueOfCoins(self, piles: List[List[int]], k: int) -> int:
        n = len(piles)
        curr, next = [0] * (k+1), [0] * (k+1)
        curr[1] = 1
        for i in range(n-1,-1,-1):
            curr, next = next, curr
            for coins in range(1, k+1):
                # skip curr pile
                curr[coins] = next[coins]
                # take from curr pile
                total = 0
                for j in range(min(coins, len(piles[i]))):
                    total += piles[i][j]
                    curr[coins] = max(
                        curr[coins],
                        total + next[coins-j-1]
                    )
        return curr[k]

    def test(self) -> None:
        for piles, k, expected in [
            ([[1,100,3],[7,8,9]], 2, 101),
            ([[1],[100],[100],[100],[100],[100],[10,1,1,1,1,1,700]], 5, 500),
		    ([[100],[100],[100],[100],[100],[100],[1,1,1,1,1,1,700]], 7, 706),
            ([[80,62,78,78,40,59,98,35],[79,19,100,15],[79,2,27,73,12,13,11,37,27,55,54,55,87,10,97,26,78,20,75,23,46,94,56,32,14,70,70,37,60,46,1,53]], 25, 1332),
        ]:
            output = self.maxValueOfCoins(piles, k)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
