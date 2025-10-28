from typing import List

# from typing import Dict
# from typing import Tuple
import unittest

# from functools import cache
# from itertools import permutations

# https://leetcode.com/problems/count-the-number-of-inversions/

# python3 -m unittest dp/3193-count-the-number-of-inversions.py


class Solution(unittest.TestCase):
    # # Approach 1: Brute-force - Generate all permutations and count inversions
    # # Time: O(n! * n^2) - generate all permutations and count inversions for each
    # # Space: O(n) - for storing permutation
    # def numberOfPermutations(self, n: int, requirements: List[List[int]]) -> int:
    #     def count_inversions(perm: List[int], end: int) -> int:
    #         """Count inversions in perm[0..end]"""
    #         count = 0
    #         for i in range(end + 1):
    #             for j in range(i + 1, end + 1):
    #                 if perm[i] > perm[j]:
    #                     count += 1
    #         return count

    #     count = 0
    #     # Generate all permutations of [0, 1, 2, ..., n-1]
    #     for perm in permutations(range(n)):
    #         # Check if this permutation satisfies all requirements
    #         count += 1
    #         for end, cnt in requirements:
    #             if count_inversions(list(perm), end) != cnt:
    #                 count -= 1
    #                 break

    #     return count % (10**9 + 7)

    # # Approach 2: Top-down DP (Memoization)
    # # Time: O(n * max_cnt * (n + max_cnt))
    # # Space: O(n * max_cnt) for memoization
    # def numberOfPermutations(self, n: int, requirements: List[List[int]]) -> int:
    #     MOD = 10**9 + 7
    #     reqs = [-1] * n
    #     max_cnt = 0
    #     for end, cnt in requirements:
    #         reqs[end] = cnt
    #         max_cnt = max(max_cnt, cnt)

    #     if reqs[0] > 0:
    #         return 0

    #     memo: List[List[int]] = [[-1] * (max_cnt + 1) for _ in range(n)]

    #     def dp(i: int, cnt: int) -> int:
    #         if i == n:
    #             return 1
    #         if cnt > max_cnt:
    #             return 0
    #         if memo[i][cnt] != -1:
    #             return memo[i][cnt]
    #         total = 0
    #         for ncnt in range(cnt, cnt + i + 1):
    #             if reqs[i] == -1 or ncnt == reqs[i]:
    #                 total = (total + dp(i + 1, ncnt)) % MOD
    #         memo[i][cnt] = total
    #         return total

    #     return dp(0, 0)

    # Approach 3: Bottom-up DP (Tabulation)
    # Time: O(n*max_cnt)
    # Space: O(n+max_cnt)
    def numberOfPermutations(self, n: int, requirements: List[List[int]]) -> int:
        MOD = 10**9 + 7
        reqs = [-1] * n
        max_cnt = 0
        for end, cnt in requirements:
            reqs[end] = cnt
            max_cnt = max(max_cnt, cnt)
        if reqs[0] > 0:
            return 0
        dp = [0] * (max_cnt + 1)
        dp[0] = 1
        for i in range(1, n + 1):
            prefix = [0] * (max_cnt + 2)
            for cnt in range(max_cnt + 1):
                prefix[cnt + 1] = (prefix[cnt] + dp[cnt]) % MOD

            ndp = [0] * (max_cnt + 1)
            for cnt in range(max_cnt + 1):
                left = max(0, cnt - (i - 1))
                ndp[cnt] = (prefix[cnt + 1] - prefix[left]) % MOD

            if reqs[i - 1] != -1:
                for cnt in range(max_cnt + 1):
                    if cnt != reqs[i - 1]:
                        ndp[cnt] = 0
            dp = ndp
        return sum(dp) % MOD

    def test(self):
        for n, requirements, expected in [
            (3, [[2, 2], [0, 0]], 2),
            (3, [[2, 2], [1, 1], [0, 0]], 1),
            (2, [[0, 0], [1, 0]], 1),
            (3, [[2, 0]], 1),
            (15, [[14, 58], [0, 0], [10, 28]], 243296005),
        ]:
            output = self.numberOfPermutations(n, requirements)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
