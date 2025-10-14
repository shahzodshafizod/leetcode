from typing import List, Set, Tuple
import unittest

# https://leetcode.com/problems/frog-jump/

# python3 -m unittest dp/0403-frog-jump.py

# EARLY EXIT:
# if (n * (n - 1) // 2) < stones[-1]:
#     return False
# If the frog makes n jumps and always increases its steps by 1,
# the total distance is: 1+2+3+...+n = n(n+1)/2
# However, the frog needs n-1 steps to reach the last stone: (n-1)n/2
# Means: If the last stone is farther than the maximum possible distance the frog could travel


class Solution(unittest.TestCase):
    # Approach #1 (faster): Top-Down Dynamic Programming
    # Time: O(nn)
    # Space: O(nn)
    def canCross(self, stones: List[int]) -> bool:
        n = len(stones)
        if (n * (n - 1) // 2) < stones[-1]:
            return False
        visited: Set[Tuple[int, int]] = set()
        indices = {stone: idx for idx, stone in enumerate(stones)}

        def dp(i: int, k: int) -> bool:
            if i == n - 1:
                return True
            if (i, k) in visited:
                return False
            j = indices.get(stones[i] + k, -1)
            if j <= i:
                return False
            if dp(j, k - 1) or dp(j, k) or dp(j, k + 1):
                return True
            visited.add((i, k))
            return False

        return dp(0, 1)

    # # Approach #2: Bottom-Up Dynamic Programming
    # # Time: O(nn)
    # # Space: O(nn)
    # def canCross(self, stones: List[int]) -> bool:
    #     n = len(stones)
    #     if (n * (n - 1) // 2) < stones[-1]:
    #         return False
    #     dp: List[List[bool]] = [[False] * (n + 1) for _ in range(n + 1)]
    #     dp[0][1] = True
    #     for i in range(1, n):
    #         for j in range(i):
    #             k = stones[i] - stones[j]
    #             if k <= n and dp[j][k]:
    #                 if i == n - 1:
    #                     return True
    #                 dp[i][k - 1] = True
    #                 dp[i][k] = True
    #                 dp[i][k + 1] = True
    #     return False

    def test(self):
        for stones, expected in [
            ([0, 1, 3, 5, 6, 8, 12, 17], True),
            ([0, 1, 2, 3, 4, 8, 9, 11], False),
        ]:
            output = self.canCross(stones)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
