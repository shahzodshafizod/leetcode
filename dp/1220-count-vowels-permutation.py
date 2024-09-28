from functools import lru_cache
import unittest

# https://leetcode.com/problems/count-vowels-permutation/

# python3 -m unittest dp/1220-count-vowels-permutation.py

class Solution(unittest.TestCase):
    # # Approach: Top-Down Dynamic Programming
    # # Time: O(n)
    # # Space: O(n)
    # def countVowelPermutation(self, n: int) -> int:
    #     adj_list = {
    #         '.': ['a','e','i','o','u'],
    #         'a': ['e','i','u'],
    #         'e': ['a','i'],
    #         'i': ['e','o'],
    #         'o': ['i'],
    #         'u': ['i','o'],
    #     }
    #     mod = int(1e9+7)
    #     @lru_cache(None)
    #     def dfs(idx: int, prev: str) -> int:
    #         if idx == n:
    #             return 1
    #         count = 0
    #         for next in adj_list[prev]:
    #             count = (count + dfs(idx+1, next)) % mod
    #         return count
    #     return dfs(0, '.')

    # Approach: Bottom-Up Dynamic Programming
    # Time: O(n)
    # Space: O(1)
    def countVowelPermutation(self, n: int) -> int:
        mod = int(1e9+7)
        a, e, i, o, u = 1, 1, 1, 1, 1
        for _ in range(n-1):
            next_a = e
            next_e = (a + i) % mod
            next_i = (a + e + o + u) % mod
            next_o = (i + u) % mod
            next_u = a
            a, e, i, o, u = next_a, next_e, next_i, next_o, next_u
        return (a + e + i + o + u) % mod

    def testCountVowelPermutation(self) -> None:
        for n, expected in [
            (1, 5),
            (2, 10),
            (5, 68),
            (144, 18208803),
        ]:
            output = self.countVowelPermutation(n)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
