import unittest

# https://leetcode.com/problems/k-inverse-pairs-array/

# python3 -m unittest dp/0629-k-inverse-pairs-array.py

"""
n=3, k=2
|   | 0 | 1 | 2 |
|---+---+---+---|
| 0 | 1 | 0 | 0 |
| 1 | 1 | 0 | 0 |
| 2 | 1 | 1 | 0 |
| 3 | 1 | 2 | 2 |
dp[n][k] - # of arrays with k inverse pairs, given n elements
"""

class Solution(unittest.TestCase):
    # # Approach #1: Top-Down Dynamic Programming (Time Limit Exceeded)
    # # Time: O(N^2 * K)
    # # Space: O(N * K)
    # def kInversePairs(self, n: int, k: int) -> int:
    #     MOD = int(1e9+7)
    #     dp = {} # (n, k) -> count
    #     def count(n: int, k: int) -> int: # O(N x K)
    #         if k <= 0:
    #             return int(k == 0)
    #         if (n, k) in dp:
    #             return dp[(n, k)]
    #         dp[(n, k)] = 0
    #         for i in range(n): # O(N)
    #             dp[(n, k)] = (
    #                 dp[(n, k)] + count(n-1, k-i)
    #             ) % MOD
    #         return dp[(n, k)]
    #     return count(n, k)

    # # Approach #2: Bottom-Up Dynamic Programming (Time Limit Exceeded)
    # # Time: O(N^2 * K)
    # # Space: O(N * K)
    # def kInversePairs(self, n: int, k: int) -> int:
    #     MOD = int(1e9+7)
    #     count = [[0] * (k+1) for _ in range(n+1)]
    #     count[0][0] = 1
    #     for ni in range(1, n+1):
    #         for ki in range(k+1):
    #             for pairs in range(min(ni, ki+1)):
    #                 count[ni][ki] = (count[ni][ki] + count[ni-1][ki-pairs]) % MOD
    #     return count[n][k]

    # Approach #3: Bottom-Up Dynamic Programming (Time & Memory Optimized)
    # Time: O(N * K)
    # Space: O(K)
    def kInversePairs(self, n: int, k: int) -> int:
        MOD = int(1e9+7)
        prev, curr = [0] * (k+1), [0] * (k+1)
        curr[0] = 1
        for ni in range(1, n+1):
            prev, curr = curr, prev
            total = 0 # sliding window
            for ki in range(k+1):
                if ki >= ni: # equal, because we're starting from 1
                    total -= prev[ki-ni]
                total = (total + prev[ki]) % MOD
                curr[ki] = total
        return curr[k]

    def test(self) -> None:
        for n, k, expected in [
            (3, 0, 1),
		    (3, 1, 2),
            (3, 3, 1),
            (4, 5, 3),
            (4, 2, 5),
            # (1000, 1000, 663677020),
        ]:
            output = self.kInversePairs(n, k)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")

"""
Sample test case: n=3:
1 2 3 -
1 3 2 - (3,2)
2 1 3 - (2,1)
2 3 1 - (2,1), (3,1)
3 1 2 - (3,1), (3,2)
3 2 1 - (3,1), (3,2), (2,1)

number of arrays with exactly k=0 inverse pairs = 1
number of arrays with exactly k=1 inverse pair = 2
number of arrays with exactly k=3 inverse pair = 1
"""
