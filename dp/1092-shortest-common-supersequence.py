import unittest

# https://leetcode.com/problems/shortest-common-supersequence/

# python3 -m unittest dp/1092-shortest-common-supersequence.py

class Solution(unittest.TestCase):
    # # Approach #1: Top-Down Dynamic Programming (Time Limit Exceeded)
    # # Time: O(2^(n+m)⋅(n+m)), n=len(str1), m=len(str2)
    # # Space: O((n+m)^2) for recursion stack
    # def shortestCommonSupersequence(self, str1: str, str2: str) -> str:
    #     n1, n2 = len(str1), len(str2)
    #     def backtrack(i1: int, i2: int) -> str:
    #         if i1 == n1:
    #             return str2[i2:]
    #         if i2 == n2:
    #             return str1[i1:]
    #         if str1[i1] == str2[i2]:
    #             return str1[i1] + backtrack(i1+1, i2+1)
    #         res1 = str1[i1] + backtrack(i1+1, i2)
    #         res2 = str2[i2] + backtrack(i1, i2+1)
    #         if len(res1) < len(res2):
    #             return res1
    #         return res2
    #     return backtrack(0, 0)

    # # Approach #2: Top-Down Dynamic Programming (Memory Limit Exceeded)
    # # Time: O(n⋅m⋅(n+m)), n=len(str1), m=len(str2)
    # # Space: O(n⋅m⋅(n+m))
    # def shortestCommonSupersequence(self, str1: str, str2: str) -> str:
    #     n1, n2 = len(str1), len(str2)
    #     memo = [[None] * n2 for _ in range(n1)]
    #     def backtrack(i1: int, i2: int) -> str:
    #         if i1 == n1:
    #             return str2[i2:]
    #         if i2 == n2:
    #             return str1[i1:]
    #         if memo[i1][i2] != None:
    #             return memo[i1][i2]
    #         if str1[i1] == str2[i2]:
    #             memo[i1][i2] = str1[i1] + backtrack(i1+1, i2+1)
    #             return memo[i1][i2]
    #         res1 = str1[i1] + backtrack(i1+1, i2)
    #         res2 = str2[i2] + backtrack(i1, i2+1)
    #         if len(res1) < len(res2):
    #             memo[i1][i2] = res1
    #         else:
    #             memo[i1][i2] = res2
    #         return memo[i1][i2]
    #     return backtrack(0, 0)

    # Approach #3: Bottom-Up Dynamic Programming (Memory Optimized)
    # Time: O(n⋅m)
    # Space: O(n⋅m)
    def shortestCommonSupersequence(self, str1: str, str2: str) -> str:
        n1, n2 = len(str1), len(str2)
        curr = [str2[i2:] for i2 in range(n2)] + [""]
        for i1 in range(n1-1,-1,-1):
            next = curr
            curr = [""] * (n2+1)
            curr[n2] = str1[i1:]
            for i2 in range(n2-1,-1,-1):
                if str1[i1] == str2[i2]:
                    curr[i2] = str1[i1] + next[i2+1]
                else:
                    substr1 = str1[i1] + next[i2]
                    substr2 = str2[i2] + curr[i2+1]
                    if len(substr1) < len(substr2):
                        curr[i2] = substr1
                    else:
                        curr[i2] = substr2
        return curr[0]

    def test(self):
        for str1, str2, expected in [
            ("abac", "cab", "cabac"),
            ("aaaaaaaa", "aaaaaaaa", "aaaaaaaa"),
        ]:
            output = self.shortestCommonSupersequence(str1, str2)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
