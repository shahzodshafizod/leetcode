from typing import List
import unittest

# https://leetcode.com/problems/total-characters-in-string-after-transformations-ii/

# python3 -m unittest maths/3337-total-characters-in-string-after-transformations-ii.py


class Solution(unittest.TestCase):
    # Approach: Matrix Multiplication + Matrix Exponentiation By Squaring
    # Time: O(n+logt×26^3) = O(n+logt)
    # Space: O(26^2) = O(1)
    def lengthAfterTransformations(self, s: str, t: int, nums: List[int]) -> int:
        MOD = int(1e9 + 7)

        class Matrix:
            def __init__(self):
                self.m = [[0] * 26 for _ in range(26)]

            def __mul__(self, other: "Matrix") -> "Matrix":
                c = Matrix()
                for i in range(26):
                    for j in range(26):
                        for k in range(26):
                            c.m[i][j] = (c.m[i][j] + self.m[i][k] * other.m[k][j]) % MOD
                return c

        def pow_matrix(base: Matrix, exp: int) -> Matrix:
            # identity matrix
            res = Matrix()
            for i in range(26):
                res.m[i][i] = 1
            while exp > 0:
                if exp & 1:
                    res = res * base
                base = base * base
                exp >>= 1
            return res

        transformation = Matrix()
        for i in range(26):
            for j in range(1, nums[i] + 1):
                transformation.m[(i + j) % 26][i] = 1
        final = pow_matrix(transformation, t)
        freq = [0] * 26
        for c in s:
            freq[ord(c) - ord('a')] += 1
        res = 0
        for i in range(26):
            for j in range(26):
                res = (res + final.m[i][j] * freq[j]) % MOD
        return res

    def test(self):
        for s, t, nums, expected in [
            ("abcyy", 2, [1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2], 7),
            ("azbk", 1, [2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2], 8),
            # ("zcxjrwgcqgmzgxyinfmnexquaopmzjllxlwtgxnpqvsgqbqvnguwymblhzkjkytdejhgewxtqwhpgmgrmvrazebvjikqplfcgnnosmxrcakvfvvjyvmkgzpglcpnphjttgmowbbfeqmqkbfszohhtrrjixnqctndm", 8942, [21,10,18,20,22,15,19,16,1,18,17,1,22,18,3,7,13,5,21,6,10,20,22,2,25,21], 906321724),
        ]:
            output = self.lengthAfterTransformations(s, t, nums)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
