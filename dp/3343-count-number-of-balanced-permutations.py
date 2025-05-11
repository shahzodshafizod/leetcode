from functools import cache
from math import comb
import unittest

# https://leetcode.com/problems/count-number-of-balanced-permutations/

# python3 -m unittest dp/3343-count-number-of-balanced-permutations.py

class Solution(unittest.TestCase):
    def countBalancedPermutations(self, num: str) -> int:
        cnt = [0] * 10
        total = 0
        for digit in map(int, num):
            cnt[digit] += 1
            total += digit
        if total & 1:
            return 0

        post_cnt = [0] * 11
        for i in range(9,-1,-1):
            post_cnt[i] = cnt[i] + post_cnt[i+1]

        n = len(num)
        @cache
        def dfs(digit: int, odd: int, balance: int) -> int:
            even = n - odd - post_cnt[digit+1]
            if odd == 0 and even == 0 and balance == 0:
                return 1
            if digit < 0 or odd < 0 or even < 0 or balance < 0:
                return 0
            res = 0
            for freq in range(cnt[digit]+1):
                res += comb(odd, freq) * comb(even, cnt[digit]-freq) * dfs(
                    digit - 1,
                    odd - freq,
                    balance - digit*freq
                )
            return res % 1000000007

        return dfs(9, (n+1)//2, total >> 1)

    def test(self):
        for num, expected in [
            ("123", 2),
            ("112", 1),
            ("12345", 0),
            ("0101", 4),
            ("325419", 72),
            # ("96509861244547846", 851350267),
            # ("9650986124454784696509861244547846", 24468941),
            # ("07162132741696526558195913376589552512680680143682957878451885242156416378132712", 285771532),
            # ("01010101010101010101010101010101010101010101010101010101010101010101010101010101", 432969423),
            # ("00000000000000000000000000000000000000000000000000000000000000000000000000000000", 1),
        ]:
            output = self.countBalancedPermutations(num)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
